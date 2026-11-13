package database

import (
	"api/internal/artwork/model"
	"api/internal/database"
)

// FindArtworkBySlug returns an artwork with given slug
// and loads its constituents and palettes
func FindArtworkBySlug(slug string) (*model.Artwork, *[]model.ArtworkConstituentDetail, *[]model.ArtworkPaletteDetails, error) {

	var artwork model.Artwork
	var constituents []model.ArtworkConstituentDetail
	paletteDetails := make([]model.ArtworkPaletteDetails, 0)

	if err := database.DB.Where("slug = ?", slug).First(&artwork).Error; err != nil {
		return nil, nil, nil, err
	}

	fields := `
		constituents.name AS name,
		artwork_constituents.is_primary AS is_primary
	`

	if err := database.DB.Table("artwork_constituents").
		Select(fields).
		Joins("INNER JOIN constituents ON constituents.id = artwork_constituents.constituent_id").
		Order("artwork_constituents.is_primary DESC, constituents.name ASC").
		Where("artwork_constituents.artwork_id = ?", artwork.ID).
		Find(&constituents).Error; err != nil {
		return &artwork, nil, nil, err
	}

	var palettes []model.ArtworkPalette
	if err := database.DB.Table("artwork_palettes").
		Where("artwork_id = ?", artwork.ID).
		Order("id ASC").
		Find(&palettes).Error; err != nil {
		return &artwork, &constituents, nil, err
	}

	for _, palette := range palettes {
		if !model.IsValidPaletteMood(palette.Mood) {
			continue
		}

		var colors []model.ArtworkPaletteColor
		if err := database.DB.Table("artwork_palette_colors").
			Select("artwork_palette_colors.palette_id, artwork_palette_colors.position, colors.hex").
			Joins("INNER JOIN colors ON colors.id = artwork_palette_colors.color_id").
			Where("artwork_palette_colors.palette_id = ?", palette.ID).
			Order("artwork_palette_colors.position ASC").
			Find(&colors).Error; err != nil {
			return &artwork, &constituents, nil, err
		}

		if len(colors) > 0 {
			paletteDetails = append(paletteDetails, model.ArtworkPaletteDetails{
				Mood:        palette.Mood,
				Composition: palette.Composition,
				Colors:      colors,
			})
		}
	}

	return &artwork, &constituents, &paletteDetails, nil
}

// FindArtworks returns a set of results for condensed artwork data
// with pagination info (hasMore, totalResults, etc.)
func FindArtworks(sort string, page, pageSize int, from, to *int64, objectType, medium, culture, search string) (*model.PreviewsData, error) {
	var artworks []model.Artwork
	var constituents []model.ConstituentPreview
	var totalResults int64

	if err := filterArtworks(database.DB.Model(&model.Artwork{}), from, to, objectType, medium, culture, search).Count(&totalResults).Error; err != nil {
		return nil, err
	}

	artworkQuery := filterArtworks(database.DB.Model(&model.Artwork{}), from, to, objectType, medium, culture, search).
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	if err := orderArtworks(artworkQuery, sort, search).Find(&artworks).Error; err != nil {
		return nil, err
	}

	var hasMore = int64(page*pageSize) < totalResults

	fields := `
		artwork_constituents.artwork_id AS artwork_id,
		constituents.name AS name,
		artwork_constituents.is_primary as is_primary
	`

	for _, artwork := range artworks {
		var currConstituents []model.ConstituentPreview

		if err := database.DB.Table("artwork_constituents").
			Select(fields).
			Joins("INNER JOIN constituents ON constituents.id = artwork_constituents.constituent_id").
			Where("artwork_constituents.artwork_id = ?", artwork.ID).
			Find(&currConstituents).Error; err != nil {
			return &model.PreviewsData{
				Artworks:     []model.Artwork{},
				TotalResults: totalResults,
				PageSize:     pageSize,
				Page:         page,
				HasMore:      hasMore,
			}, err
		}
		constituents = append(constituents, currConstituents...)
	}

	return &model.PreviewsData{
		Artworks:     artworks,
		Constituents: constituents,
		TotalResults: totalResults,
		PageSize:     pageSize,
		Page:         page,
		HasMore:      hasMore,
	}, nil
}

func FindRelatedArtworks(slug string, limit int) (*model.PreviewsData, error) {
	var artwork model.Artwork
	if err := database.DB.Where("slug = ?", slug).First(&artwork).Error; err != nil {
		return nil, err
	}

	var primaryConstituentIDs []uint
	if err := database.DB.Table("artwork_constituents").
		Where("artwork_id = ? AND is_primary = ?", artwork.ID, true).
		Pluck("constituent_id", &primaryConstituentIDs).Error; err != nil {
		return nil, err
	}

	var artworks []model.Artwork
	query := database.DB.Table("artworks AS related").
		Select(`
			related.*,
			CASE WHEN related.classification = ? THEN 4 ELSE 0 END +
			CASE WHEN related.department = ? THEN 2 ELSE 0 END +
			CASE WHEN related.culture = ? THEN 1 ELSE 0 END AS relevance`,
			artwork.Classification, artwork.Department, artwork.Culture,
		).
		Where("related.id <> ?", artwork.ID).
		Where(
			"related.classification = ? OR related.department = ? OR related.culture = ?",
			artwork.Classification, artwork.Department, artwork.Culture,
		)

	if len(primaryConstituentIDs) > 0 {
		query = query.Select(`
			related.*,
			CASE WHEN EXISTS (
				SELECT 1 FROM artwork_constituents related_ac
				WHERE related_ac.artwork_id = related.id
				AND related_ac.is_primary = true
				AND related_ac.constituent_id IN ?
			) THEN 8 ELSE 0 END +
			CASE WHEN related.classification = ? THEN 4 ELSE 0 END +
			CASE WHEN related.department = ? THEN 2 ELSE 0 END +
			CASE WHEN related.culture = ? THEN 1 ELSE 0 END AS relevance`,
			primaryConstituentIDs,
			artwork.Classification, artwork.Department, artwork.Culture,
		).Where(`
			related.classification = ? OR related.department = ? OR related.culture = ? OR EXISTS (
				SELECT 1 FROM artwork_constituents related_ac
				WHERE related_ac.artwork_id = related.id
				AND related_ac.is_primary = true
				AND related_ac.constituent_id IN ?
			)`,
			artwork.Classification, artwork.Department, artwork.Culture, primaryConstituentIDs,
		)
	}

	if err := query.Order("relevance DESC, related.id DESC").Limit(limit).Find(&artworks).Error; err != nil {
		return nil, err
	}

	constituents := make([]model.ConstituentPreview, 0)
	for _, related := range artworks {
		var current []model.ConstituentPreview
		if err := database.DB.Table("artwork_constituents").
			Select(`
				artwork_constituents.artwork_id AS artwork_id,
				constituents.name AS name,
				artwork_constituents.is_primary AS is_primary
			`).
			Joins("INNER JOIN constituents ON constituents.id = artwork_constituents.constituent_id").
			Where("artwork_constituents.artwork_id = ?", related.ID).
			Find(&current).Error; err != nil {
			return nil, err
		}
		constituents = append(constituents, current...)
	}

	return &model.PreviewsData{
		Artworks:     artworks,
		Constituents: constituents,
		TotalResults: int64(len(artworks)),
		PageSize:     limit,
		Page:         1,
		HasMore:      false,
	}, nil
}

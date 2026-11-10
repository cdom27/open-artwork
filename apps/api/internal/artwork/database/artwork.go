package database

import (
	"api/internal/artwork/model"
	"api/internal/database"
)

// FindArtworkBySlug returns an artwork with given slug
// and loads its constituents and palettes.
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

package database

import (
	"api/internal/artwork/model"
	"api/internal/database"
)

// FindArtworkBySlug returns an artwork with given slug
// and loads constituents
func FindArtworkBySlug(slug string) (*model.Artwork, *[]model.ConstituentWithRole, error) {
	// TODO: Load color palettes

	var artwork model.Artwork
	var constituents []model.ConstituentWithRole

	if err := database.DB.Where("slug = ?", slug).First(&artwork).Error; err != nil {
		return nil, nil, err
	}

	fields := `
		constituents.name AS name,
		constituents.bio AS bio,
		constituents.nationality AS nationality,
		constituents.date_begin AS date_begin,
		constituents.date_end AS date_end,
		artwork_constituents.role AS role
	`

	if err := database.DB.Table("artwork_constituents").
		Select(fields).
		Joins("INNER JOIN constituents ON constituents.id = artwork_constituents.constituent_id").
		Where("artwork_constituents.artwork_id = ?", artwork.ID).
		Find(&constituents).Error; err != nil {
		return &artwork, nil, err
	}

	return &artwork, &constituents, nil
}

// FindArtworks returns a set of results for condensed artwork data
// with pagination info (hasMore, totalResults, etc.)
func FindArtworks(sort string, page, pageSize int, from, to *int64, objectType, medium, culture string) (*model.PreviewsData, error) {
	var artworks []model.Artwork
	var constituents []model.ConstituentPreview
	var totalResults int64

	if err := filterArtworks(database.DB.Model(&model.Artwork{}), from, to, objectType, medium, culture).Count(&totalResults).Error; err != nil {
		return nil, err
	}

	if err := filterArtworks(database.DB.Model(&model.Artwork{}), from, to, objectType, medium, culture).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order(artworkOrder(sort)).
		Find(&artworks).Error; err != nil {
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

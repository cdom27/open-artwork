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
		artwork_constituents.role AS role,
		artwork_constituents.is_primary AS is_primary
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

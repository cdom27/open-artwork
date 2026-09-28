package database

import (
	"api/internal/artwork/model"
	"api/internal/database"
)

// FindArtworkBySlug returns an artwork with given slug
// joined with constituents (artists) and color palettes
func FindArtworkBySlug(slug string) (*model.Artwork, error) {
	var artwork model.Artwork

	err := database.DB.First(&artwork, "slug", slug).Error

	if err == nil {
		// todo: Load constituents and palettes
	}

	if err != nil {
		return nil, err
	}

	return &artwork, nil
}

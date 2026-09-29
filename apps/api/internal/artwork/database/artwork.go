package database

import (
	"api/internal/artwork/model"
	"api/internal/database"
)

// FindArtworkBySlug returns an artwork with given slug
// and loads constituents
func FindArtworkBySlug(slug string) (*model.Artwork, *[]model.Constituent, error) {
	// TODO: Load color palettes

	var artwork model.Artwork
	var awConstituents []model.ArtworkConstituent
	var constituents []model.Constituent

	if err := database.DB.Where("slug = ?", slug).First(&artwork).Error; err != nil {
		return nil, nil, err
	}

	if err := database.DB.Where("artwork_id = ?", artwork.ID).Find(&awConstituents).Error; err != nil {
		return &artwork, nil, err
	}

	var constituentIDs []uint
	for _, c := range awConstituents {
		constituentIDs = append(constituentIDs, c.ConstituentID)
	}

	if len(constituentIDs) > 0 {
		if err := database.DB.Where("id IN ?", constituentIDs).Find(&constituents).Error; err != nil {
			return &artwork, nil, err
		}
	}

	return &artwork, &constituents, nil
}

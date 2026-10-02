package database

import "gorm.io/gorm"

func filterArtworks(query *gorm.DB, from, to *int64, objectType, medium, culture string) *gorm.DB {
	if from != nil {
		query = query.Where("date_end >= ?", *from)
	}
	if to != nil {
		query = query.Where("date_begin <= ?", *to)
	}
	if objectType != "" {
		query = query.Where("object_type = ?", objectType)
	}
	if medium != "" {
		query = query.Where("medium = ?", medium)
	}
	if culture != "" {
		query = query.Where("culture = ?", culture)
	}

	return query
}

func artworkOrder(sort string) string {
	const primaryArtistName = `(
		SELECT constituents.name
		FROM artwork_constituents
		INNER JOIN constituents ON constituents.id = artwork_constituents.constituent_id
		WHERE artwork_constituents.artwork_id = artworks.id
		ORDER BY artwork_constituents.is_primary DESC, constituents.name ASC
		LIMIT 1
	)`

	switch sort {
	case "title":
		return "artworks.title ASC, artworks.id DESC"
	case "newest":
		return "artworks.date_begin DESC, artworks.id DESC"
	case "oldest":
		return "artworks.date_begin ASC, artworks.id DESC"
	case "artistAsc":
		return "COALESCE(" + primaryArtistName + ", '') ASC, artworks.id DESC"
	case "artistDesc":
		return "COALESCE(" + primaryArtistName + ", '') DESC, artworks.id DESC"
	case "relevance":
		fallthrough
	default:
		return "artworks.id DESC"
	}
}

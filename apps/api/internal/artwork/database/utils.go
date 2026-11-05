package database

import (
	"regexp"
	"strings"

	"gorm.io/gorm"
)

var searchTokenPattern = regexp.MustCompile(`[\p{L}\p{N}]+`)

const artworkSearchVector = `
	setweight(to_tsvector('simple', coalesce(artworks.title, '')), 'A') ||
	setweight(to_tsvector('simple', coalesce(array_to_string(artworks.tags, ' '), '')), 'B') ||
	setweight(to_tsvector('simple', coalesce(artworks.object_type, '')), 'B') ||
	setweight(to_tsvector('simple', concat_ws(' ', artworks.medium, artworks.culture, artworks.period, artworks.classification, artworks.department)), 'C')
`

const constituentSearchVector = `setweight(to_tsvector('simple', coalesce(constituents.name, '')), 'A')`

func filterArtworks(query *gorm.DB, from, to *int64, objectType, medium, culture, search string) *gorm.DB {
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

	for _, term := range artworkSearchTerms(search) {
		query = query.Where(`(
			`+artworkSearchVector+` @@ to_tsquery('simple', ?)
			OR EXISTS (
				SELECT 1
				FROM artwork_constituents
				INNER JOIN constituents ON constituents.id = artwork_constituents.constituent_id
				WHERE artwork_constituents.artwork_id = artworks.id
					AND `+constituentSearchVector+` @@ to_tsquery('simple', ?)
			)
		)`, term, term)
	}

	return query
}

func artworkSearchTerms(search string) []string {
	tokens := searchTokenPattern.FindAllString(strings.ToLower(search), -1)
	terms := make([]string, len(tokens))
	for i, token := range tokens {
		terms[i] = token + ":*"
	}

	return terms
}

func orderArtworks(query *gorm.DB, sort, search string) *gorm.DB {
	terms := artworkSearchTerms(search)
	if sort != "relevance" || len(terms) == 0 {
		return query.Order(artworkOrder(sort))
	}

	searchQuery := strings.Join(terms, " & ")
	return query.Select(`artworks.*, (
		ts_rank_cd((`+artworkSearchVector+`), to_tsquery('simple', ?)) +
		COALESCE((
			SELECT MAX(ts_rank_cd(`+constituentSearchVector+`, to_tsquery('simple', ?)))
			FROM artwork_constituents
			INNER JOIN constituents ON constituents.id = artwork_constituents.constituent_id
			WHERE artwork_constituents.artwork_id = artworks.id
		), 0)
	) AS search_rank`, searchQuery, searchQuery).
		Order("search_rank DESC, artworks.id DESC")
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

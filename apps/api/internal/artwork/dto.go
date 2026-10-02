package artwork

import "api/internal/artwork/model"

type ArtworkResponse struct {
	Title          string                      `json:"title"`
	Slug           string                      `json:"slug"`
	ObjectType     string                      `json:"objectType"`
	Department     string                      `json:"department"`
	Classification string                      `json:"classification"`
	Culture        string                      `json:"culture"`
	Period         string                      `json:"period"`
	DateDisplay    string                      `json:"dateDisplay"`
	DateBegin      int64                       `json:"dateBegin"`
	DateEnd        int64                       `json:"dateEnd"`
	Medium         string                      `json:"medium"`
	Dimensions     string                      `json:"dimensions"`
	Creditline     string                      `json:"creditLine"`
	Constituents   []model.ConstituentWithRole `json:"constituents"`
	Tags           []string                    `json:"tags"`
}

type ArtworkPreview struct {
	Title             string `json:"title"`
	Slug              string `json:"slug"`
	DateDisplay       string `json:"dateDisplay"`
	PrimaryArtistName string `json:"primaryArtistName"`
	ImageURL          string `json:"imageUrl"`
}
type PreviewResponse struct {
	Artworks     []ArtworkPreview `json:"artworks"`
	TotalResults int64            `json:"totalResults"`
	PageSize     int              `json:"pageSize"`
	Page         int              `json:"page"`
	HasMore      bool             `json:"hasMore"`
}

type Sort string

const (
	SortRelevance  Sort = "relevance"
	SortTitle      Sort = "title"
	SortNewest     Sort = "newest"
	SortOldest     Sort = "oldest"
	SortArtistAsc  Sort = "artistAsc"
	SortArtistDesc Sort = "artistDesc"
)

type SearchParams struct {
	Q          string
	Page       int
	PageSize   int
	Sort       Sort
	From       *int64
	To         *int64
	ObjectType string
	Medium     string
	Culture    string
}

// ToArtworkResponse assists in mapping DB struct into client response
func ToArtworkResponse(a *model.Artwork, b *[]model.ConstituentWithRole) ArtworkResponse {
	return ArtworkResponse{
		Title:          a.Title,
		ObjectType:     a.ObjectType,
		Department:     a.Department,
		Classification: a.Classification,
		Culture:        a.Culture,
		Period:         a.Period,
		DateDisplay:    a.DateDisplay,
		DateBegin:      a.DateBegin,
		DateEnd:        a.DateEnd,
		Medium:         a.Medium,
		Dimensions:     a.Dimensions,
		Creditline:     a.Creditline,
		Constituents:   *b,
		Tags:           a.Tags,
	}
}

// ToArtworkPreviews assists in mapping multiple DB tables (artworks, constituents) into a preview response
func ToArtworkPreviews(pd *model.PreviewsData) PreviewResponse {
	constituentsByArtwork := make(map[uint][]model.ConstituentPreview)
	for _, c := range pd.Constituents {
		constituentsByArtwork[c.ArtworkID] = append(constituentsByArtwork[c.ArtworkID], c)
	}

	previews := make([]ArtworkPreview, len(pd.Artworks))
	for i, artwork := range pd.Artworks {
		var primaryName string

		// prefer primary artist if found, fallbacks -> first constituent -> Unknown Artist
		if cs, ok := constituentsByArtwork[artwork.ID]; ok && len(cs) > 0 {
			primaryName = cs[0].Name
			for _, c := range cs {
				if c.IsPrimary {
					primaryName = c.Name
					break
				}
			}
		}

		previews[i] = ArtworkPreview{
			Title:             artwork.Title,
			Slug:              artwork.Slug,
			DateDisplay:       artwork.DateDisplay,
			PrimaryArtistName: primaryName,
		}
	}

	return PreviewResponse{
		Artworks:     previews,
		TotalResults: pd.TotalResults,
		PageSize:     pd.PageSize,
		Page:         pd.Page,
		HasMore:      pd.HasMore,
	}
}

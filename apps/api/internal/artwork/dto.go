package artwork

import (
	"api/internal/artwork/model"
	"fmt"
)

type ArtworkResponse struct {
	Title          string                   `json:"title"`
	ObjectType     string                   `json:"objectType"`
	Department     string                   `json:"department"`
	Classification string                   `json:"classification"`
	Culture        string                   `json:"culture"`
	Period         string                   `json:"period"`
	DateDisplay    string                   `json:"dateDisplay"`
	Medium         string                   `json:"medium"`
	Dimensions     string                   `json:"dimensions"`
	Creditline     string                   `json:"creditLine"`
	Tags           []string                 `json:"tags"`
	Constituents   []string                 `json:"constituents"`
	Palettes       []ArtworkPaletteResponse `json:"palettes"`
	FullURL        string                   `json:"fullUrl"`
}

type ArtworkPaletteResponse struct {
	Mood        model.PaletteMood      `json:"mood"`
	Composition string                 `json:"composition"`
	Colors      []ArtworkColorResponse `json:"colors"`
}

type ArtworkColorResponse struct {
	Hex string `json:"hex"`
}

type ArtworkPreview struct {
	Title             string `json:"title"`
	Slug              string `json:"slug"`
	DateDisplay       string `json:"dateDisplay"`
	PrimaryArtistName string `json:"primaryArtistName"`
	ThumbnailURL      string `json:"thumbnailUrl"`
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
func ToArtworkResponse(a *model.Artwork, constituents *[]model.ArtworkConstituentDetail, paletteDetails *[]model.ArtworkPaletteDetails) ArtworkResponse {
	names := make([]string, len(*constituents))
	for i, constituent := range *constituents {
		names[i] = constituent.Name
	}

	tags := []string(a.Tags)
	if tags == nil {
		tags = []string{}
	}

	palettes := make([]ArtworkPaletteResponse, len(*paletteDetails))
	for i, palette := range *paletteDetails {
		colors := make([]ArtworkColorResponse, len(palette.Colors))
		for j, color := range palette.Colors {
			colors[j] = ArtworkColorResponse{Hex: color.Hex}
		}
		palettes[i] = ArtworkPaletteResponse{
			Mood:        palette.Mood,
			Composition: palette.Composition,
			Colors:      colors,
		}
	}

	return ArtworkResponse{
		Title:          a.Title,
		ObjectType:     a.ObjectType,
		Department:     a.Department,
		Classification: a.Classification,
		Culture:        a.Culture,
		Period:         a.Period,
		DateDisplay:    a.DateDisplay,
		Medium:         a.Medium,
		Dimensions:     a.Dimensions,
		Creditline:     a.Creditline,
		Tags:           tags,
		Constituents:   names,
		Palettes:       palettes,
		FullURL:        fmt.Sprintf("https://images.openart.work/images/%d/primary-full.webp", a.ID),
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
			ThumbnailURL:      fmt.Sprintf("https://images.openart.work/images/%d/primary-small.webp", artwork.ID),
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

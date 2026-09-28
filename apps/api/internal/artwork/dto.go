package artwork

import "api/internal/artwork/model"

type ArtworkResponse struct {
	Title          string   `json:"title"`
	Slug           string   `json:"slug"`
	ObjectType     string   `json:"objectType"`
	Department     string   `json:"department"`
	Classification string   `json:"classification"`
	Culture        string   `json:"culture"`
	Period         string   `json:"period"`
	DateDisplay    string   `json:"dateDisplay"`
	DateBegin      uint     `json:"dateBegin"`
	DateEnd        uint     `json:"dateEnd"`
	Medium         string   `json:"medium"`
	Dimensions     string   `json:"dimensions"`
	Creditline     string   `json:"creditLine"`
	Tags           []string `json:"tags"`
}

// ToArtworkResponse assists in mapping DB struct into client response
func ToArtworkResponse(a *model.Artwork) ArtworkResponse {
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
		Tags:           a.Tags,
	}
}

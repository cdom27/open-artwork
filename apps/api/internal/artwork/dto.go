package artwork

import "api/internal/artwork/model"

type ArtworkResponse struct {
	Title          string                `json:"title"`
	Slug           string                `json:"slug"`
	ObjectType     string                `json:"objectType"`
	Department     string                `json:"department"`
	Classification string                `json:"classification"`
	Culture        string                `json:"culture"`
	Period         string                `json:"period"`
	DateDisplay    string                `json:"dateDisplay"`
	DateBegin      uint                  `json:"dateBegin"`
	DateEnd        uint                  `json:"dateEnd"`
	Medium         string                `json:"medium"`
	Dimensions     string                `json:"dimensions"`
	Creditline     string                `json:"creditLine"`
	Constituents   []ConstituentResponse `json:"constituents"`
	Tags           []string              `json:"tags"`
}

type ConstituentResponse struct {
	Name        string `gorm:"column:name;not null"`
	Bio         string `gorm:"column:bio"`
	Nationality string `gorm:"column:nationality"`
	DateBegin   uint   `gorm:"column:date_begin"`
	DateEnd     uint   `gorm:"column:date_end"`
}

// ToArtworkResponse assists in mapping DB struct into client response
func ToArtworkResponse(a *model.Artwork, b *[]model.Constituent) ArtworkResponse {
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
		Constituents:   toConstituentResponse(*b),
		Tags:           a.Tags,
	}
}

func toConstituentResponse(a []model.Constituent) []ConstituentResponse {
	var constituents []ConstituentResponse

	for _, c := range a {
		constituents = append(constituents, ConstituentResponse{
			Name:        c.Name,
			Bio:         c.Bio,
			Nationality: c.Nationality,
			DateBegin:   c.DateBegin,
			DateEnd:     c.DateEnd,
		})
	}

	return constituents
}

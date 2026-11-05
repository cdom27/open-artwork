package model

import (
	"time"

	"github.com/lib/pq"
)

type Artwork struct {
	ID             uint           `gorm:"primaryKey;column:id;autoIncrement;index:artwork_pkey,unique,type:btree"`
	Title          string         `gorm:"column:title;not null"`
	Slug           string         `gorm:"column:slug;not null"`
	ObjectType     string         `gorm:"column:object_type"`
	Department     string         `gorm:"column:department"`
	Classification string         `gorm:"column:classification"`
	Culture        string         `gorm:"column:culture"`
	Period         string         `gorm:"column:period"`
	DateDisplay    string         `gorm:"column:date_display"`
	DateBegin      int64          `gorm:"column:date_begin"`
	DateEnd        int64          `gorm:"column:date_end"`
	Medium         string         `gorm:"column:medium"`
	Dimensions     string         `gorm:"column:dimensions"`
	Creditline     string         `gorm:"column:credit_line"`
	Tags           pq.StringArray `gorm:"column:tags;type:text[]"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
}

type ConstituentWithRole struct {
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	Nationality string `json:"nationality"`
	DateBegin   int64  `json:"dateBegin"`
	DateEnd     int64  `json:"dateEnd"`
	Role        string `json:"role"`
}

type ConstituentPreview struct {
	ArtworkID uint   `json:"-"`
	Name      string `json:"name"`
	IsPrimary bool   `json:"isPrimaryArtist"`
}

type PreviewsData struct {
	Artworks     []Artwork            `json:"artworks"`
	Constituents []ConstituentPreview `json:"constituents"`
	TotalResults int64                `json:"totalResults"`
	PageSize     int                  `json:"pageSize"`
	Page         int                  `json:"page"`
	HasMore      bool                 `json:"hasMore"`
}

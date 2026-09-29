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
	DateBegin      uint           `gorm:"column:date_begin"`
	DateEnd        uint           `gorm:"column:date_end"`
	Medium         string         `gorm:"column:medium"`
	Dimensions     string         `gorm:"column:dimensions"`
	Creditline     string         `gorm:"column:credit_line"`
	Tags           pq.StringArray `gorm:"column:tags;type:text[]"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
}

type ArtworkConstituent struct {
	ArtworkID     uint      `gorm:"column:artwork_id;primaryKey"`
	ConstituentID uint      `gorm:"column:constituent_id;primaryKey"`
	Role          string    `gorm:"column:role"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

type Constituent struct {
	ID          uint      `gorm:"primaryKey;column:id;autoIncrement;index:artwork_pkey,unique,type:btree"`
	Name        string    `gorm:"column:name;not null"`
	Bio         string    `gorm:"column:bio"`
	Nationality string    `gorm:"column:nationality"`
	DateBegin   uint      `gorm:"column:date_begin"`
	DateEnd     uint      `gorm:"column:date_end"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type ConstituentWithRole struct {
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	Nationality string `json:"nationality"`
	DateBegin   uint   `json:"dateBegin"`
	DateEnd     uint   `json:"dateEnd"`
	Role        string `json:"role"`
	IsPrimary   bool   `json:"isPrimary"`
}

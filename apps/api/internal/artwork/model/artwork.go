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

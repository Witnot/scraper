package models

import "time"

type Product struct {
	ID           uint `gorm:"primaryKey"`
	Source       string
	ExternalID   string `gorm:"index"` // e.g. site-specific id
	Name         string
	URL          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	PriceRecords []PriceRecord
}

type PriceRecord struct {
	ID         uint `gorm:"primaryKey"`
	ProductID  uint `gorm:"index"`
	Price      float64
	Currency   string
	RecordedAt time.Time `gorm:"index"`
}

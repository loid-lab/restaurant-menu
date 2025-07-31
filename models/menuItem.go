package models

import "time"

type MenuItem struct {
	ID          uint   `gorm:"primaryKey"`
	CategoryID  uint   `gorm:"not null"`
	Name        string `gorm:"not null"`
	ImageURL    string
	Price       float64 `gorm:"not null"`
	Category    string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	CreatedBy   uint    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

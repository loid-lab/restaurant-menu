package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FullName  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"type:varchar(20);default:'customer'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Orders    []Order
	Addresses []Address
}

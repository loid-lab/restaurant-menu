package models

type OrderItem struct {
	ID         uint    `gorm:"primaryKey"`
	OrderID    uint    `gorm:"not null"`
	MenuItemID uint    `gorm:"not null"`
	Quantity   int     `gorm:"not null"`
	UnitPrice  float64 `gorm:"not null"`
	TotalPrice float64 `gorm:"not null"`
	MenuItem   MenuItem
}

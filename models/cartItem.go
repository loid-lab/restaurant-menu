package models

type CartItem struct {
	ID         uint `gorm:"primaryKey"`
	CartID     uint `gorm:"not null"`
	MenuItemID uint `gorm:"not null"`
	Quantity   int  `gorm:"not null"`
	MenuItem   MenuItem
	UserID     uint
}

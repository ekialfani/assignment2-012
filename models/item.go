package models

import "time"

type Item struct {
	ID uint `gorm:"primaryKey"`
	ItemCode string `gorm:"not null; unique;type:varchar(191)" json:"item_code"`
	Description string `json:"description"`
	Quantity uint `gorm:"not null; default:0" json:"quantity"`
	OrderID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
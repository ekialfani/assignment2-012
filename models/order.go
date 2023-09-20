package models

import "time"

type Order struct {
	ID uint `gorm:"primaryKey"`
	CustomerName string `gorm:"not null; type:varchar(191)" json:"customer_name"`
	Items []Item
	OrderedAt time.Time `json:"ordered_at"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
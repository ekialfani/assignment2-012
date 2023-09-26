package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID uint `gorm:"primaryKey"`
	ItemCode string `gorm:"not null; unique;type:char(3)" json:"item_code"`
	Description string `json:"description"`
	Quantity uint `gorm:"not null; default:0" json:"quantity"`
	OrderID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	if len(item.ItemCode) > 3 {
		err = errors.New("The item code must be unique and should not exceed 3 characters.")
	}
	return
}

func (item *Item) BeforeUpdate(tx *gorm.DB) (err error) {
	if len(item.ItemCode) > 3 {
		err = errors.New("The item code must be unique and should not exceed 3 characters.")
	}
	return
}
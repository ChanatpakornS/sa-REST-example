package model

import (
	"time"

	"gorm.io/gorm"
)

type Concert struct {
	// GORM automatically add id, created_at, updated_at, deleted_at fields
	gorm.Model
	Title       string    `gorm:"size:200;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Location    string    `gorm:"size:200;not null" json:"location"`
	Date        time.Time `gorm:"not null" json:"date"`
	Capacity    int       `gorm:"not null" json:"capacity"`
	Price       float64   `gorm:"not null" json:"price"`
}

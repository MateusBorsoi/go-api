package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	TargetPrice float64   `json:"target_price"`
	LastPrice   float64   `json:"last_price"`
	LastChecked time.Time `json:"last_checked"`
	Status      string    `json:"status" gorm:"default:'pending'"`
}

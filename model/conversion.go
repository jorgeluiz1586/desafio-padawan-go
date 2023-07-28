package model

import (
	"time"

	"gorm.io/gorm"
)

type Conversion struct {
	gorm.Model
	ID             uint `gorm:"primaryKey"`
	Amount         float64
	ConvertedValue float64
	From           string
	To             string
	Rate           float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

package models

import "time"

type Dasha struct {
	ID       uint      `gorm:"primaryKey"`
	FromDate time.Time `gorm:"not null"`
	ToDate   time.Time `gorm:"not null"`
	Summary  string    `gorm:"not null"`
	Planet   string    `gorm:"not null"`
}

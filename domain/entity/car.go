package entity

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:100"`
	Brand        string `gorm:"size:150;unique"`
	LicensePlate string `gorm:"size:10;unique"`
	Year         time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

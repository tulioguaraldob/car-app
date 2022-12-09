package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"size:125"`
	LastName  string `gorm:"size:150"`
	Email     string `gorm:"size:125;unique"`
	Password  string `gorm:"size:100"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

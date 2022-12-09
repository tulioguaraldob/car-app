package migration

import (
	"github.com/TulioGuaraldoB/car-app/domain/entity"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}

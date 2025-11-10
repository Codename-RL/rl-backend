package repository

import (
	"codename-rl/internal/entity"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{})
}

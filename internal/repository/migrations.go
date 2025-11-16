package repository

import (
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB, entities ...interface{}) error {
	return db.AutoMigrate(entities...)
}

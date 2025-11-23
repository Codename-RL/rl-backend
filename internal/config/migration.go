package config

import (
	"codename-rl/internal/entity"
	"codename-rl/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewDatabaseMigration(db *gorm.DB, log *logrus.Logger) error {
	log.Info("Running database migrations...")

	if err := repository.AutoMigrate(
		db,
		&entity.User{},
		&entity.Otp{},
		&entity.Tag{},
	); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
		return err
	}

	log.Info("Database migrations completed successfully.")
	return nil
}

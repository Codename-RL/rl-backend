package main

import (
	"codename-rl/internal/config"
	"fmt"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator(viperConfig)
	app := config.NewFiber(viperConfig)
	jwt := config.NewJwt(viperConfig)
	emailClient := config.NewEmail(viperConfig, log)

	migrationErr := config.NewDatabaseMigration(db, log)
	if migrationErr != nil {
		log.Errorf("Failed to start migration: %v", migrationErr)
	}

	config.Bootstrap(&config.BootstrapConfig{
		DB:          db,
		App:         app,
		Log:         log,
		Validate:    validate,
		Config:      viperConfig,
		EmailClient: emailClient,
		JWTService:  jwt,
	})

	appPort := viperConfig.GetInt("server.port")
	err := app.Listen(fmt.Sprintf(":%d", appPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

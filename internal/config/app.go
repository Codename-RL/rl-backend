package config

import (
	"codename-rl/internal/delivery/http/handler"
	"codename-rl/internal/delivery/http/middleware"
	"codename-rl/internal/delivery/http/route"
	"codename-rl/internal/pkg/auth"
	"codename-rl/internal/repository"
	"codename-rl/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB         *gorm.DB
	App        *fiber.App
	Log        *logrus.Logger
	Validate   *validator.Validate
	JWTService *auth.JwtService
	Config     *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	OtpRepository := repository.NewOtpRepository(config.Log)

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository, OtpRepository, config.JWTService)
	otpUseCase := usecase.NewOtpUseCase(config.DB, config.Log, config.Validate, OtpRepository, userRepository, config.JWTService)

	// setup controller
	userController := handler.NewUserController(userUseCase, config.Log)
	otpController := handler.NewOtpController(otpUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.Config{
		App:            config.App,
		UserController: userController,
		OtpController:  otpController,
		AuthMiddleware: authMiddleware,
	}
	routeConfig.Setup()
}

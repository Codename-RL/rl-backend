package config

import (
	"codename-rl/internal/delivery/http/handler"
	"codename-rl/internal/delivery/http/middleware"
	"codename-rl/internal/delivery/http/route"
	"codename-rl/internal/pkg/auth"
	"codename-rl/internal/pkg/email"
	"codename-rl/internal/repository"
	"codename-rl/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB          *gorm.DB
	App         *fiber.App
	Log         *logrus.Logger
	Validate    *validator.Validate
	JWTService  *auth.JwtService
	EmailClient *email.Client
	Config      *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	otpRepository := repository.NewOtpRepository(config.Log)
	tagRepository := repository.NewTagRepository(config.Log)
	personRepository := repository.NewPersonRepository(config.Log)

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository, otpRepository, config.JWTService)
	otpUseCase := usecase.NewOtpUseCase(config.DB, config.Log, config.Validate, otpRepository, userRepository, config.EmailClient, config.JWTService)
	tagUseCase := usecase.NewTagUseCase(config.DB, config.Log, config.Validate, tagRepository, config.JWTService)
	personUseCase := usecase.NewPersonUseCase(config.DB, config.Log, config.Validate, personRepository, config.JWTService)

	// setup controller
	userController := handler.NewUserController(userUseCase, config.Log)
	otpController := handler.NewOtpController(otpUseCase, config.Log)
	tagHandler := handler.NewTagHandler(tagUseCase, config.Log)
	personHandler := handler.NewPersonHandler(personUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.Config{
		App:              config.App,
		UserController:   userController,
		OtpController:    otpController,
		TagController:    tagHandler,
		PersonController: personHandler,
		AuthMiddleware:   authMiddleware,
	}
	routeConfig.Setup()
}

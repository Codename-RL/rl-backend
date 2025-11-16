package route

import (
	"codename-rl/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App            *fiber.App
	UserController *handler.UserController
	OtpController  *handler.OtpController
	AuthMiddleware fiber.Handler
}

func (c *Config) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *Config) SetupGuestRoute() {
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/_login", c.UserController.Login)
	c.App.Post("/api/users/_otp", c.OtpController.CreateOtp)
	c.App.Post("/api/users/_otp/forgot", c.OtpController.VerifyOtpForgotPassword)
	c.App.Patch("/api/users/_password", c.UserController.UpdatePassword)
}

func (c *Config) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Delete("/api/users", c.UserController.Logout)
	c.App.Patch("/api/users/_current", c.UserController.Update)
	c.App.Get("/api/users/_current", c.UserController.Current)
	c.App.Post("/api/users/_otp/verify", c.OtpController.VerifyOtpUser)
}

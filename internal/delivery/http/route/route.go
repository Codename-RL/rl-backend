package route

import (
	"codename-rl/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App            *fiber.App
	UserController *handler.UserController
	OtpController  *handler.OtpController
	TagController  *handler.TagHandler
	AuthMiddleware fiber.Handler
}

func (c *Config) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *Config) SetupGuestRoute() {
	// User
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/_login", c.UserController.Login)

	// OTP
	c.App.Post("/api/users/_otp", c.OtpController.CreateOtp)
	c.App.Post("/api/users/_otp/forgot", c.OtpController.VerifyOtpForgotPassword)
	c.App.Patch("/api/users/_password", c.UserController.UpdatePassword)
}

func (c *Config) SetupAuthRoute() {
	// Middleware
	c.App.Use(c.AuthMiddleware)
	c.App.Delete("/api/users", c.UserController.Logout)
	c.App.Patch("/api/users/_current", c.UserController.Update)
	c.App.Get("/api/users/_current", c.UserController.Current)

	// OTP
	c.App.Post("/api/users/_otp/verify", c.OtpController.VerifyOtpUser)

	//Tags
	c.App.Post("/api/tags", c.TagController.Create)
	c.App.Get("/api/tags", c.TagController.Get)
	c.App.Patch("/api/tags", c.TagController.Update)
	c.App.Delete("/api/tags", c.TagController.Delete)
}

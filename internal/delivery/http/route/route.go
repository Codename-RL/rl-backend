package route

import (
	"codename-rl/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App                     *fiber.App
	UserController          *handler.UserController
	OtpController           *handler.OtpController
	TagController           *handler.TagHandler
	PersonController        *handler.PersonHandler
	RelationshipController  *handler.RelationshipHandler
	PhoneController         *handler.PhoneHandler
	ImportantDateController *handler.ImportantDateHandler
	AuthMiddleware          fiber.Handler
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

	//Persons
	c.App.Post("/api/persons", c.PersonController.Create)
	c.App.Get("/api/persons", c.PersonController.Get)
	c.App.Patch("/api/persons", c.PersonController.Update)
	c.App.Delete("/api/persons", c.PersonController.Delete)

	//Relationships
	c.App.Post("/api/relationships", c.RelationshipController.Create)
	c.App.Get("/api/relationships", c.RelationshipController.Get)
	c.App.Patch("/api/relationships", c.RelationshipController.Update)
	c.App.Delete("/api/relationships", c.RelationshipController.Delete)

	//Phones
	c.App.Post("/api/phones", c.PhoneController.Create)
	c.App.Get("/api/phones", c.PhoneController.Get)
	c.App.Patch("/api/phones", c.PhoneController.Update)
	c.App.Delete("/api/phones", c.PhoneController.Delete)

	//Important Dates
	c.App.Post("/api/importantdates", c.ImportantDateController.Create)
	c.App.Get("/api/importantdates", c.ImportantDateController.Get)
	c.App.Patch("/api/importantdates", c.ImportantDateController.Update)
	c.App.Delete("/api/importantdates", c.ImportantDateController.Delete)
}

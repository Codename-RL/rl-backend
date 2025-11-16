package handler

import (
	"codename-rl/internal/delivery/http/middleware"
	"codename-rl/internal/delivery/http/response"
	"codename-rl/internal/model"
	"codename-rl/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	responseData, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		resp := response.NewErrorResponse("Failed to register user", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	resp := response.NewResponse("User registered successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	responseData, err := c.UseCase.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login user : %+v", err)
		resp := response.NewErrorResponse("Failed to login user", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	resp := response.NewResponse("Login successful", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *UserController) Current(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetUserRequest{
		ID: auth.ID,
	}

	responseData, err := c.UseCase.Current(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get current user")
		resp := response.NewErrorResponse("Failed to get current user", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	resp := response.NewResponse("Current user fetched successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *UserController) Logout(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.LogoutUserRequest{
		ID: auth.ID,
	}

	responseData, err := c.UseCase.Logout(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to logout user")
		resp := response.NewErrorResponse("Failed to logout user", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	resp := response.NewResponse("Logout successful", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.ID = auth.ID

	responseData, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update user")
		resp := response.NewErrorResponse("Failed to update user", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	resp := response.NewResponse("User updated successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *UserController) UpdatePassword(ctx *fiber.Ctx) error {
	request := new(model.UpdateUserPasswordRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.ID = auth.ID

	responseData, err := c.UseCase.UpdatePassword(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update user")
		resp := response.NewErrorResponse("Failed to update user", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	resp := response.NewResponse("User updated successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

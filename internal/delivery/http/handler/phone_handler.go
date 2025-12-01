package handler

import (
	"codename-rl/internal/delivery/http/middleware"
	"codename-rl/internal/delivery/http/response"
	"codename-rl/internal/model"
	"codename-rl/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PhoneHandler struct {
	Log     *logrus.Logger
	UseCase *usecase.PhoneUseCase
}

func NewPhoneHandler(useCase *usecase.PhoneUseCase, logger *logrus.Logger) *PhoneHandler {
	return &PhoneHandler{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *PhoneHandler) Create(ctx *fiber.Ctx) error {
	request := new(model.CreatePhoneRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create phone : %+v", err)
		resp := response.NewErrorResponse("Failed to create phone", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("Phone createed successfully", responseData)
	return ctx.Status(fiber.StatusCreated).JSON(resp)
}

func (c *PhoneHandler) Get(ctx *fiber.Ctx) error {
	request := new(model.GetPhoneRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, total, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get phones")
		resp := response.NewErrorResponse("Failed to get phones", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	pageSize := request.Query.Limit
	if pageSize == 0 {
		pageSize = 10
	}
	page := (request.Query.Offset / pageSize) + 1

	resp := response.NewPaginatedResponse("Get phones fetched successfully", *responseData, total, page, pageSize)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *PhoneHandler) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdatePhoneRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update phone")
		resp := response.NewErrorResponse("Failed to update phone", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("User updated successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *PhoneHandler) Delete(ctx *fiber.Ctx) error {
	request := new(model.DeletePhoneRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to delete phone")
		resp := response.NewErrorResponse("Failed to delete phone", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("Phone deleted successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

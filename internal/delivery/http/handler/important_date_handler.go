package handler

import (
	"codename-rl/internal/delivery/http/middleware"
	"codename-rl/internal/delivery/http/response"
	"codename-rl/internal/model"
	"codename-rl/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ImportantDateHandler struct {
	Log     *logrus.Logger
	UseCase *usecase.ImportantDateUseCase
}

func NewImportantDateHandler(useCase *usecase.ImportantDateUseCase, logger *logrus.Logger) *ImportantDateHandler {
	return &ImportantDateHandler{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *ImportantDateHandler) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateImportantDateRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create important date : %+v", err)
		resp := response.NewErrorResponse("Failed to create important date", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("ImportantDate createed successfully", responseData)
	return ctx.Status(fiber.StatusCreated).JSON(resp)
}

func (c *ImportantDateHandler) Get(ctx *fiber.Ctx) error {
	request := new(model.GetImportantDateRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, total, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get important dates")
		resp := response.NewErrorResponse("Failed to get important dates", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	pageSize := request.Query.Limit
	if pageSize == 0 {
		pageSize = 10
	}
	page := (request.Query.Offset / pageSize) + 1

	resp := response.NewPaginatedResponse("Get important dates fetched successfully", *responseData, total, page, pageSize)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *ImportantDateHandler) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateImportantDateRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update important date")
		resp := response.NewErrorResponse("Failed to update important date", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("User updated successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *ImportantDateHandler) Delete(ctx *fiber.Ctx) error {
	request := new(model.DeleteImportantDateRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to delete important date")
		resp := response.NewErrorResponse("Failed to delete important date", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("Important Dates deleted successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

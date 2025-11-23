package handler

import (
	"codename-rl/internal/delivery/http/middleware"
	"codename-rl/internal/delivery/http/response"
	"codename-rl/internal/model"
	"codename-rl/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PersonHandler struct {
	Log     *logrus.Logger
	UseCase *usecase.PersonUseCase
}

func NewPersonHandler(useCase *usecase.PersonUseCase, logger *logrus.Logger) *PersonHandler {
	return &PersonHandler{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *PersonHandler) Create(ctx *fiber.Ctx) error {
	request := new(model.CreatePersonRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create person : %+v", err)
		resp := response.NewErrorResponse("Failed to create person", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("Person created successfully", responseData)
	return ctx.Status(fiber.StatusCreated).JSON(resp)
}

func (c *PersonHandler) Get(ctx *fiber.Ctx) error {
	request := new(model.GetPersonRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, total, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get persons")
		resp := response.NewErrorResponse("Failed to get persons", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	pageSize := request.Query.Limit
	if pageSize == 0 {
		pageSize = 10
	}
	page := (request.Query.Offset / pageSize) + 1

	resp := response.NewPaginatedResponse("Get persons fetched successfully", *responseData, total, page, pageSize)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *PersonHandler) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdatePersonRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update person")
		resp := response.NewErrorResponse("Failed to update person", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("Person updated successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *PersonHandler) Delete(ctx *fiber.Ctx) error {
	request := new(model.DeletePersonRequest)

	request.ID = ctx.Params("id")

	responseData, err := c.UseCase.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to delete person")
		resp := response.NewErrorResponse("Failed to delete person", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("Person deleted successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

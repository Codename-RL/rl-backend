package handler

import (
	"codename-rl/internal/delivery/http/middleware"
	"codename-rl/internal/delivery/http/response"
	"codename-rl/internal/model"
	"codename-rl/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TagHandler struct {
	Log     *logrus.Logger
	UseCase *usecase.TagUseCase
}

func NewTagHandler(useCase *usecase.TagUseCase, logger *logrus.Logger) *TagHandler {
	return &TagHandler{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *TagHandler) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateTagRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create tag : %+v", err)
		resp := response.NewErrorResponse("Failed to create tag", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("Tag createed successfully", responseData)
	return ctx.Status(fiber.StatusCreated).JSON(resp)
}

func (c *TagHandler) Get(ctx *fiber.Ctx) error {
	request := new(model.GetTagRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, total, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get tags")
		resp := response.NewErrorResponse("Failed to get tags", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	pageSize := request.Query.Limit
	if pageSize == 0 {
		pageSize = 10
	}
	page := (request.Query.Offset / pageSize) + 1

	resp := response.NewPaginatedResponse("Get tags fetched successfully", *responseData, total, page, pageSize)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *TagHandler) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateTagRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update user")
		resp := response.NewErrorResponse("Failed to update user", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("User updated successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *TagHandler) Delete(ctx *fiber.Ctx) error {
	request := new(model.DeleteTagRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update user")
		resp := response.NewErrorResponse("Failed to update user", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("User updated successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

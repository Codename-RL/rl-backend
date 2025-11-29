package handler

import (
	"codename-rl/internal/delivery/http/middleware"
	"codename-rl/internal/delivery/http/response"
	"codename-rl/internal/model"
	"codename-rl/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RelationshipHandler struct {
	Log     *logrus.Logger
	UseCase *usecase.RelationshipUseCase
}

func NewRelationshipHandler(useCase *usecase.RelationshipUseCase, logger *logrus.Logger) *RelationshipHandler {
	return &RelationshipHandler{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *RelationshipHandler) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateRelationshipRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create relationship : %+v", err)
		resp := response.NewErrorResponse("Failed to create relationship", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("Relationship created successfully", responseData)
	return ctx.Status(fiber.StatusCreated).JSON(resp)
}

func (c *RelationshipHandler) Get(ctx *fiber.Ctx) error {
	request := new(model.GetRelationshipRequest)

	if err := ctx.QueryParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, total, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get relationships")
		resp := response.NewErrorResponse("Failed to get relationships", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	pageSize := request.Query.Limit
	if pageSize == 0 {
		pageSize = 10
	}
	page := (request.Query.Offset / pageSize) + 1

	resp := response.NewPaginatedResponse("Get relationships fetched successfully", *responseData, total, page, pageSize)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *RelationshipHandler) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateRelationshipRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	request.ID = ctx.Params("id")

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update relationship")
		resp := response.NewErrorResponse("Failed to update relationship", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("Relationship updated successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *RelationshipHandler) Delete(ctx *fiber.Ctx) error {
	request := new(model.DeleteRelationshipRequest)

	request.ID = ctx.Params("id")

	auth := middleware.GetUser(ctx)
	request.UserID = auth.ID

	responseData, err := c.UseCase.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to delete relationship")
		resp := response.NewErrorResponse("Failed to delete relationship", err)
		return ctx.Status(err.Code).JSON(resp)
	}

	resp := response.NewResponse("Relationship deleted successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

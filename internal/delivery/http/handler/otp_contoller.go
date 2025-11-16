package handler

import (
	"codename-rl/internal/delivery/http/response"
	"codename-rl/internal/model"
	"codename-rl/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type OtpController struct {
	Log     *logrus.Logger
	UseCase *usecase.OtpUseCase
}

func NewOtpController(useCase *usecase.OtpUseCase, logger *logrus.Logger) *OtpController {
	return &OtpController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *OtpController) CreateOtp(ctx *fiber.Ctx) error {
	request := new(model.CreateOtpRequest)
	err := ctx.BodyParser(request)

	if err != nil {
		c.Log.Warnf("Failed to parse request body: %v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	responseData, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create new OTP : %+v", err)
		resp := response.NewErrorResponse("Failed to create new OTP", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	resp := response.NewResponse("OTP Created Successfully", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *OtpController) VerifyOtpUser(ctx *fiber.Ctx) error {
	request := new(model.VerifyOtpRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	responseData, err := c.UseCase.VerifyUser(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to verify OTP : %+v", err)
		resp := response.NewErrorResponse("Failed to verify OTP", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	resp := response.NewResponse("OTP verification successful", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *OtpController) VerifyOtpForgotPassword(ctx *fiber.Ctx) error {
	request := new(model.VerifyOtpRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		resp := response.NewErrorResponse("Invalid request body", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	responseData, err := c.UseCase.VerifyUser(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to verify OTP : %+v", err)
		resp := response.NewErrorResponse("Failed to verify OTP", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	resp := response.NewResponse("OTP verification successful", responseData)
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

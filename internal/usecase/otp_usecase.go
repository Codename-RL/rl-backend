package usecase

import (
	"codename-rl/internal/entity"
	"codename-rl/internal/model"
	"codename-rl/internal/model/converter"
	"codename-rl/internal/pkg/auth"
	"codename-rl/internal/pkg/email"
	"codename-rl/internal/pkg/utils"
	"codename-rl/internal/repository"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type OtpUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	OtpRepository  *repository.OtpRepository
	UserRepository *repository.UserRepository
	EmailClient    *email.Client
	JWTService     *auth.JwtService
}

func NewOtpUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	otpRepository *repository.OtpRepository, UserRepository *repository.UserRepository, emailClient *email.Client, JWTService *auth.JwtService) *OtpUseCase {
	return &OtpUseCase{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		OtpRepository:  otpRepository,
		UserRepository: UserRepository,
		EmailClient:    emailClient,
		JWTService:     JWTService,
	}
}

func (c *OtpUseCase) Create(ctx context.Context, request *model.CreateOtpRequest) (*model.OtpResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	if err = c.UserRepository.FindByEmail(tx, user, request.Email); err != nil {
		c.Log.Warnf("Failed to find user by email : %+v", err)
		return nil, fiber.ErrNotFound
	}

	token, err := c.JWTService.GenerateToken(user, 5*time.Minute)
	if err != nil {
		c.Log.Errorf("Failed to generate token : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	otpNumeric, err := utils.GenerateNumericOTP(6)
	if err != nil {
		c.Log.Warnf("Failed to generate OTP : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	otpHashed, err := utils.HashPassword(otpNumeric)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrypt hash : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	otp := &entity.Otp{
		ID:        uuid.New().String(),
		Otp:       string(otpHashed),
		Token:     token,
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		UserID:    user.ID,
	}

	if err := c.OtpRepository.Create(tx, otp); err != nil {
		c.Log.Warnf("Failed create otp to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// send email here
	if err := c.EmailClient.SendOTP(request.Email, otpNumeric); err != nil {
		c.Log.Warnf("Failed to send OTP to email : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}
	return converter.OtpToResponse(otp), nil
}

func (c *OtpUseCase) VerifyUser(ctx context.Context, request *model.VerifyOtpRequest) (*model.OtpResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	otp := new(entity.Otp)
	if err = c.OtpRepository.FindByToken(ctx, tx, otp, request.Token); err != nil {
		c.Log.Warnf("Failed to find OTP by token : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if otp.VerifiedAt != 0 {
		c.Log.Warnf("OTP already verified")
		return nil, fiber.ErrBadRequest
	}

	if otp.ExpiresAt < time.Now().Unix() {
		c.Log.Warnf("OTP expired")
		return nil, fiber.ErrBadRequest
	}

	if err = utils.ComparePassword(otp.Otp, request.Otp); err != nil {
		c.Log.Warnf("Failed to compare OTP hash : %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	if err = c.UserRepository.VerifyUser(ctx, tx, otp.UserID); err != nil {
		c.Log.Errorf("Failed to verify user: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	otp.VerifiedAt = time.Now().Unix()
	if err = c.OtpRepository.Update(tx, otp); err != nil {
		c.Log.Errorf("Failed to update OTP: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err = c.OtpRepository.DeleteByToken(ctx, tx, otp.Token); err != nil {
		c.Log.Warnf("Failed delete OTP by id : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err = tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.OtpToResponse(otp), nil

}

func (c *OtpUseCase) VerifyForgotPassword(ctx context.Context, request *model.VerifyOtpRequest) (*model.OtpResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}
	otp := new(entity.Otp)
	if err = c.OtpRepository.FindByToken(ctx, tx, otp, request.Token); err != nil {
		c.Log.Warnf("Failed to find OTP by token : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if otp.VerifiedAt != 0 {
		c.Log.Warnf("OTP already verified")
		return nil, fiber.ErrBadRequest
	}

	if otp.ExpiresAt < time.Now().Unix() {
		c.Log.Warnf("OTP expired")
		return nil, fiber.ErrBadRequest
	}

	if err = bcrypt.CompareHashAndPassword([]byte(otp.Otp), []byte(request.Otp)); err != nil {
		c.Log.Warnf("Failed to compare OTP hash : %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	otp.VerifiedAt = time.Now().Unix()
	if err = c.OtpRepository.Update(tx, otp); err != nil {
		c.Log.Errorf("Failed to update OTP: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err = tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.OtpToResponse(otp), nil

}

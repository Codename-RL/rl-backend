package usecase

import (
	"codename-rl/internal/entity"
	"codename-rl/internal/model"
	"codename-rl/internal/model/converter"
	"codename-rl/internal/pkg/auth"
	"codename-rl/internal/repository"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PhoneUseCase struct {
	DB              *gorm.DB
	Log             *logrus.Logger
	Validate        *validator.Validate
	PhoneRepository *repository.PhoneRepository
	JWTService      *auth.JwtService
}

func NewPhoneUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	phoneRepository *repository.PhoneRepository, JWTService *auth.JwtService) *PhoneUseCase {
	return &PhoneUseCase{
		DB:              db,
		Log:             logger,
		Validate:        validate,
		PhoneRepository: phoneRepository,
		JWTService:      JWTService,
	}
}

func (c *PhoneUseCase) Create(ctx context.Context, request *model.CreatePhoneRequest) (*model.PhoneResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	phone := new(entity.Phone)

	exists, err := c.PhoneRepository.ExistsByNumber(tx, phone, request.Number)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to check phone existence by number")
		return nil, fiber.ErrInternalServerError
	}

	if exists {
		c.Log.Warnf("Phone already exists with number: %s", request.Number)
		return nil, fiber.ErrConflict
	}

	*phone = entity.Phone{
		ID:        uuid.New().String(),
		Number:    request.Number,
		PersonID:  request.PersonID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Person:    nil,
	}

	if err := c.PhoneRepository.Create(tx, phone, request.PersonID); err != nil {
		c.Log.Warnf("Failed create phone to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.PhoneToResponse(phone), nil
}

func (c *PhoneUseCase) Get(ctx context.Context, request *model.GetPhoneRequest) (*[]model.PhoneResponse, int64, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, 0, fiber.ErrBadRequest
	}

	var phones []entity.Phone
	total, err := c.PhoneRepository.FindAll(tx, &phones, &request.Query)
	if err != nil {
		c.Log.Warnf("Failed find phones : %+v", err)
		return nil, 0, fiber.ErrNotFound
	}

	if len(phones) == 0 {
		return nil, 0, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, 0, fiber.ErrInternalServerError
	}

	return converter.PhonesToResponses(&phones), total, nil
}

func (c *PhoneUseCase) Update(ctx context.Context, request *model.UpdatePhoneRequest) (*model.PhoneResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	phone := new(entity.Phone)
	if err := c.PhoneRepository.FindById(tx, phone, request.ID); err != nil {
		c.Log.Warnf("Failed find phone by id : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if request.Number != "" {
		phone.Number = request.Number
	}

	if err := c.PhoneRepository.Update(tx, phone); err != nil {
		c.Log.Warnf("Failed save phone : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.PhoneToResponse(phone), nil
}

func (c *PhoneUseCase) Delete(ctx context.Context, request *model.DeletePhoneRequest) (*model.PhoneResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	phone := new(entity.Phone)
	if err := c.PhoneRepository.FindById(tx, phone, request.ID); err != nil {
		c.Log.Warnf("Failed find phone by id : %+v", err)
		return nil, fiber.ErrNotFound
	}
	phone.ID = request.ID

	if err := c.PhoneRepository.Delete(tx, phone); err != nil {
		c.Log.Warnf("Failed save phone : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.PhoneToResponse(phone), nil
}

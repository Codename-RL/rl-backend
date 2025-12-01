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

type ImportantDateUseCase struct {
	DB                      *gorm.DB
	Log                     *logrus.Logger
	Validate                *validator.Validate
	ImportantDateRepository *repository.ImportantDateRepository
	JWTService              *auth.JwtService
}

func NewImportantDateUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	importantDateRepository *repository.ImportantDateRepository, JWTService *auth.JwtService) *ImportantDateUseCase {
	return &ImportantDateUseCase{
		DB:                      db,
		Log:                     logger,
		Validate:                validate,
		ImportantDateRepository: importantDateRepository,
		JWTService:              JWTService,
	}
}

func (c *ImportantDateUseCase) Create(ctx context.Context, request *model.CreateImportantDateRequest) (*model.ImportantDateResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	importantDate := new(entity.ImportantDate)

	exists, err := c.ImportantDateRepository.ExistsByName(tx, importantDate, request.Name)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to check important date existence by name")
		return nil, fiber.ErrInternalServerError
	}

	if exists {
		c.Log.Warnf("Important date already exists with name: %s", request.Name)
		return nil, fiber.ErrConflict
	}

	*importantDate = entity.ImportantDate{
		ID:        uuid.New().String(),
		Name:      request.Name,
		PersonID:  request.PersonID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Person:    nil,
	}

	if err := c.ImportantDateRepository.Create(tx, importantDate, request.PersonID); err != nil {
		c.Log.Warnf("Failed create important date to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ImportantDateToResponse(importantDate), nil
}

func (c *ImportantDateUseCase) Get(ctx context.Context, request *model.GetImportantDateRequest) (*[]model.ImportantDateResponse, int64, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, 0, fiber.ErrBadRequest
	}

	var importantDates []entity.ImportantDate
	total, err := c.ImportantDateRepository.FindAll(tx, &importantDates, &request.Query)
	if err != nil {
		c.Log.Warnf("Failed find important dates : %+v", err)
		return nil, 0, fiber.ErrNotFound
	}

	if len(importantDates) == 0 {
		return nil, 0, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, 0, fiber.ErrInternalServerError
	}

	return converter.ImportantDatesToResponses(&importantDates), total, nil
}

func (c *ImportantDateUseCase) Update(ctx context.Context, request *model.UpdateImportantDateRequest) (*model.ImportantDateResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	importantDate := new(entity.ImportantDate)
	if err := c.ImportantDateRepository.FindById(tx, importantDate, request.ID); err != nil {
		c.Log.Warnf("Failed find important date by id : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if request.Name != "" {
		importantDate.Name = request.Name
	}

	if err := c.ImportantDateRepository.Update(tx, importantDate); err != nil {
		c.Log.Warnf("Failed save important date : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ImportantDateToResponse(importantDate), nil
}

func (c *ImportantDateUseCase) Delete(ctx context.Context, request *model.DeleteImportantDateRequest) (*model.ImportantDateResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	importantDate := new(entity.ImportantDate)
	if err := c.ImportantDateRepository.FindById(tx, importantDate, request.ID); err != nil {
		c.Log.Warnf("Failed find important date by id : %+v", err)
		return nil, fiber.ErrNotFound
	}
	importantDate.ID = request.ID

	if err := c.ImportantDateRepository.Delete(tx, importantDate); err != nil {
		c.Log.Warnf("Failed save important date : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ImportantDateToResponse(importantDate), nil
}

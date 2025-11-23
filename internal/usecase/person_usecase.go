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

type PersonUseCase struct {
	DB               *gorm.DB
	Log              *logrus.Logger
	Validate         *validator.Validate
	PersonRepository *repository.PersonRepository
	JWTService       *auth.JwtService
}

func NewPersonUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	personRepository *repository.PersonRepository,
	JWTService *auth.JwtService,
) *PersonUseCase {
	return &PersonUseCase{
		DB:               db,
		Log:              logger,
		Validate:         validate,
		PersonRepository: personRepository,
		JWTService:       JWTService,
	}
}

func (c *PersonUseCase) Create(ctx context.Context, request *model.CreatePersonRequest) (*model.PersonResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	person := new(entity.Person)

	*person = entity.Person{
		ID:          uuid.New().String(),
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Nickname:    request.Nickname,
		Avatar:      request.Avatar,
		Description: request.Description,
		UserID:      request.UserID,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	if err := c.PersonRepository.Create(tx, person, request.TagIDs); err != nil {
		c.Log.Warnf("Failed create person to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.PersonToResponse(person), nil
}

func (c *PersonUseCase) Get(ctx context.Context, request *model.GetPersonRequest) (*[]model.PersonResponse, int64, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, 0, fiber.ErrBadRequest
	}

	var persons []entity.Person
	total, err := c.PersonRepository.FindAll(tx, &persons, &request.Query)
	if err != nil {
		c.Log.Warnf("Failed find persons : %+v", err)
		return nil, 0, fiber.ErrNotFound
	}

	if len(persons) == 0 {
		return nil, 0, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, 0, fiber.ErrInternalServerError
	}

	return converter.PersonsToResponses(&persons), total, nil
}

func (c *PersonUseCase) Update(ctx context.Context, request *model.UpdatePersonRequest) (*model.PersonResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	person := new(entity.Person)
	if err := c.PersonRepository.FindById(tx, person, request.ID); err != nil {
		c.Log.Warnf("Failed find person by id : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if request.FirstName != "" {
		person.FirstName = request.FirstName
	}
	if request.LastName != "" {
		person.LastName = request.LastName
	}
	if request.Nickname != "" {
		person.Nickname = request.Nickname
	}
	if request.Avatar != "" {
		person.Avatar = request.Avatar
	}
	if request.Description != "" {
		person.Description = request.Description
	}

	if err := c.PersonRepository.Update(tx, person, request.TagIDs); err != nil {
		c.Log.Warnf("Failed save person : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.PersonToResponse(person), nil
}

func (c *PersonUseCase) Delete(ctx context.Context, request *model.DeletePersonRequest) (*model.PersonResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	person := new(entity.Person)
	if err := c.PersonRepository.FindById(tx, person, request.ID); err != nil {
		c.Log.Warnf("Failed find person by id : %+v", err)
		return nil, fiber.ErrNotFound
	}
	person.ID = request.ID

	if err := c.PersonRepository.Delete(tx, person); err != nil {
		c.Log.Warnf("Failed save person : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.PersonToResponse(person), nil
}

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

type TagUseCase struct {
	DB            *gorm.DB
	Log           *logrus.Logger
	Validate      *validator.Validate
	TagRepository *repository.TagRepository
	JWTService    *auth.JwtService
}

func NewTagUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	tagRepository *repository.TagRepository, JWTService *auth.JwtService) *TagUseCase {
	return &TagUseCase{
		DB:            db,
		Log:           logger,
		Validate:      validate,
		TagRepository: tagRepository,
		JWTService:    JWTService,
	}
}

func (c *TagUseCase) Create(ctx context.Context, request *model.CreateTagRequest) (*model.TagResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	tag := new(entity.Tag)

	exists, err := c.TagRepository.ExistsByName(tx, tag, request.Name)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to check tag existence by email")
		return nil, fiber.ErrInternalServerError
	}

	if exists {
		c.Log.Warnf("Tag already exists with email: %s", request.Name)
		return nil, fiber.ErrConflict
	}

	*tag = entity.Tag{
		ID:        uuid.New().String(),
		Name:      request.Name,
		UserID:    request.UserID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Persons:   nil,
		User:      nil,
	}

	if err := c.TagRepository.Create(tx, tag, request.PersonIDs); err != nil {
		c.Log.Warnf("Failed create tag to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.TagToResponse(tag), nil
}

func (c *TagUseCase) Get(ctx context.Context, request *model.GetTagRequest) (*[]model.TagResponse, int64, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, 0, fiber.ErrBadRequest
	}

	var tags []entity.Tag
	total, err := c.TagRepository.FindAll(tx, &tags, &request.Query)
	if err != nil {
		c.Log.Warnf("Failed find tags : %+v", err)
		return nil, 0, fiber.ErrNotFound
	}

	if len(tags) == 0 {
		return nil, 0, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, 0, fiber.ErrInternalServerError
	}

	return converter.TagsToResponses(&tags), total, nil
}

func (c *TagUseCase) Update(ctx context.Context, request *model.UpdateTagRequest) (*model.TagResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	tag := new(entity.Tag)
	if err := c.TagRepository.FindById(tx, tag, request.ID); err != nil {
		c.Log.Warnf("Failed find tag by id : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if request.Name != "" {
		tag.Name = request.Name
	}

	if err := c.TagRepository.Update(tx, tag); err != nil {
		c.Log.Warnf("Failed save tag : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.TagToResponse(tag), nil
}

func (c *TagUseCase) Delete(ctx context.Context, request *model.DeleteTagRequest) (*model.TagResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	tag := new(entity.Tag)
	if err := c.TagRepository.FindById(tx, tag, request.ID); err != nil {
		c.Log.Warnf("Failed find tag by id : %+v", err)
		return nil, fiber.ErrNotFound
	}
	tag.ID = request.ID

	if err := c.TagRepository.Delete(tx, tag); err != nil {
		c.Log.Warnf("Failed save tag : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.TagToResponse(tag), nil
}

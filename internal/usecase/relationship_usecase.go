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

type RelationshipUseCase struct {
	DB                     *gorm.DB
	Log                    *logrus.Logger
	Validate               *validator.Validate
	RelationshipRepository *repository.RelationshipRepository
	JWTService             *auth.JwtService
}

func NewRelationshipUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	relationshipRepository *repository.RelationshipRepository,
	JWTService *auth.JwtService,
) *RelationshipUseCase {
	return &RelationshipUseCase{
		DB:                     db,
		Log:                    logger,
		Validate:               validate,
		RelationshipRepository: relationshipRepository,
		JWTService:             JWTService,
	}
}

func (c *RelationshipUseCase) Create(ctx context.Context, request *model.CreateRelationshipRequest) (*model.RelationshipResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	relationship := new(entity.Relationship)

	exists, err := c.RelationshipRepository.ExistsByName(tx, relationship, request.Name)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to check relationship existence by name")
		return nil, fiber.ErrInternalServerError
	}

	if exists {
		c.Log.Warnf("Relationship already exists with name: %s", request.Name)
		return nil, fiber.ErrConflict
	}

	*relationship = entity.Relationship{
		ID:        uuid.New().String(),
		Name:      request.Name,
		UserID:    request.UserID,
		Color:     request.Color,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	if err := c.RelationshipRepository.Create(tx, relationship, request.PersonIDs); err != nil {
		c.Log.Warnf("Failed create relationship to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.RelationshipToResponse(relationship), nil
}

func (c *RelationshipUseCase) Get(ctx context.Context, request *model.GetRelationshipRequest) (*[]model.RelationshipResponse, int64, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, 0, fiber.ErrBadRequest
	}

	var relationships []entity.Relationship
	total, err := c.RelationshipRepository.FindAll(tx, &relationships, &request.Query)
	if err != nil {
		c.Log.Warnf("Failed find relationships : %+v", err)
		return nil, 0, fiber.ErrNotFound
	}

	if len(relationships) == 0 {
		return nil, 0, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, 0, fiber.ErrInternalServerError
	}

	return converter.RelationshipsToResponses(&relationships), total, nil
}

func (c *RelationshipUseCase) Update(ctx context.Context, request *model.UpdateRelationshipRequest) (*model.RelationshipResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	relationship := new(entity.Relationship)
	if err := c.RelationshipRepository.FindById(tx, relationship, request.ID); err != nil {
		c.Log.Warnf("Failed find relationship by id : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if request.Name != "" {
		relationship.Name = request.Name
	}

	if err := c.RelationshipRepository.Update(tx, relationship, request.PersonIDs); err != nil {
		c.Log.Warnf("Failed save relationship : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.RelationshipToResponse(relationship), nil
}

func (c *RelationshipUseCase) Delete(ctx context.Context, request *model.DeleteRelationshipRequest) (*model.RelationshipResponse, *fiber.Error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	relationship := new(entity.Relationship)
	if err := c.RelationshipRepository.FindById(tx, relationship, request.ID); err != nil {
		c.Log.Warnf("Failed find relationship by id : %+v", err)
		return nil, fiber.ErrNotFound
	}
	relationship.ID = request.ID

	if err := c.RelationshipRepository.Delete(tx, relationship); err != nil {
		c.Log.Warnf("Failed save relationship : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.RelationshipToResponse(relationship), nil
}

package repository

import (
	"codename-rl/internal/entity"
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) ExistsByEmail(tx *gorm.DB, user *entity.User, email string) (bool, error) {
	var exists bool
	err := tx.Model(user).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).Error
	return exists, err
}

func (r *UserRepository) FindByEmail(tx *gorm.DB, user *entity.User, email string) error {
	return tx.Where("email = ?", email).First(user).Error
}

func (r *UserRepository) VerifyUser(ctx context.Context, db *gorm.DB, userID string) error {
	return db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", userID).
		Update("verified_at", time.Now().UnixMilli()).
		Error
}

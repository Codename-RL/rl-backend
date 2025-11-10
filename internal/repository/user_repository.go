package repository

import (
	"codename-rl/internal/entity"

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

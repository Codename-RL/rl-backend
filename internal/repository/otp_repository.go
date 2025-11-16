package repository

import (
	"codename-rl/internal/entity"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OtpRepository struct {
	Repository[entity.Otp]
	Log *logrus.Logger
}

func NewOtpRepository(log *logrus.Logger) *OtpRepository {
	return &OtpRepository{
		Log: log,
	}
}

func (r *OtpRepository) FindByToken(ctx context.Context, db *gorm.DB, otp *entity.Otp, token string) error {
	return db.WithContext(ctx).Where("token = ?", token).First(otp).Error
}

func (r *OtpRepository) FindByUserID(ctx context.Context, db *gorm.DB, userID string) (*entity.Otp, error) {
	var otp entity.Otp
	if err := db.WithContext(ctx).Where("user_id = ?", userID).First(&otp).Error; err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *OtpRepository) FindWithUser(ctx context.Context, db *gorm.DB, otpID string) (*entity.Otp, error) {
	var otp entity.Otp
	if err := db.WithContext(ctx).Preload("User").First(&otp, "id = ?", otpID).Error; err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *OtpRepository) DeleteByUserID(ctx context.Context, db *gorm.DB, userID string) error {
	return db.WithContext(ctx).Where("user_id = ?", userID).Delete(&entity.Otp{}).Error
}

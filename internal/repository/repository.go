package repository

import (
	"codename-rl/internal/entity"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

func (r *Repository[T]) CountById(db *gorm.DB, id any) (int64, error) {
	var total int64
	err := db.Model(new(T)).Where("id = ?", id).Count(&total).Error
	return total, err
}

func (r *UserRepository) ExistsById(tx *gorm.DB, id string) (bool, error) {
	var exists bool
	err := tx.Model(&entity.User{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).Error
	return exists, err
}

func (r *Repository[T]) FindById(db *gorm.DB, entity *T, id any) error {
	return db.Where("id = ?", id).Take(entity).Error
}

func (r *Repository[T]) ExistsById(db *gorm.DB, id any) (bool, error) {
	var exists bool
	err := db.Model(new(T)).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).Error
	return exists, err
}

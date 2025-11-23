package repository

import (
	"codename-rl/internal/entity"
	"codename-rl/internal/model"
	"codename-rl/internal/pkg/utils"
	"strings"

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

func (r *Repository[T]) FindAll(db *gorm.DB, result *[]T, q *model.Query) (int64, error) {
	tx := db.Model(new(T))

	// -----------------------------------
	// 0. PRELOAD RELATIONS
	// -----------------------------------
	for _, rel := range q.Preload {
		tx = tx.Preload(rel)
	}

	// -----------------------------------
	// 1. SEARCH MAP (field → value)
	// -----------------------------------
	if len(q.Search) > 0 {
		for field, value := range q.Search {
			var condition string
			var arg interface{}

			if field == "id" {
				// Exact match
				condition = field + " = ?"
				arg = value
			} else if strings.HasSuffix(field, "_id") {
				// Exact match for any *_id field
				condition = field + " = ?"
				arg = value
			} else {
				// LIKE search
				condition = field + " LIKE ?"
				arg = "%" + value + "%"
			}

			if q.Or {
				tx = tx.Or(condition, arg)
			} else {
				tx = tx.Where(condition, arg)
			}
		}
	}

	// -----------------------------------
	// 2. DATE RANGE
	// -----------------------------------
	for field, dr := range q.DateRanges {

		// ---- FROM ----
		if dr.From != "" {
			fromTime, err := utils.ParseFlexibleTime(dr.From)
			if err == nil {
				tx = tx.Where(field+" >= ?", fromTime)
			}
		}

		// ---- TO ----
		if dr.To != "" {
			toTime, err := utils.ParseFlexibleTime(dr.To)
			if err == nil {
				tx = tx.Where(field+" <= ?", toTime)
			}
		}
	}

	// -----------------------------------
	// 3. COUNT TOTAL
	// -----------------------------------
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		return 0, err
	}

	// -----------------------------------
	// 4. SORTING
	// -----------------------------------
	if q.SortBy != "" {
		order := "ASC"
		if strings.ToUpper(q.Order) == "DESC" {
			order = "DESC"
		}
		tx = tx.Order(q.SortBy + " " + order)
	}

	// -----------------------------------
	// 5. PAGINATION
	// -----------------------------------
	if q.Limit > 0 {
		tx = tx.Limit(q.Limit)
	}
	if q.Offset > 0 {
		tx = tx.Offset(q.Offset)
	}

	return total, tx.Find(result).Error
}

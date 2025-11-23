package repository

import (
	"codename-rl/internal/entity"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PersonRepository struct {
	Repository[entity.Person]
	Log *logrus.Logger
}

func NewPersonRepository(log *logrus.Logger) *PersonRepository {
	return &PersonRepository{
		Log: log,
	}
}

func (r *PersonRepository) Create(tx *gorm.DB, person *entity.Person, tagIDs []string) error {
	if err := tx.Create(person).Error; err != nil {
		return err
	}

	if len(tagIDs) == 0 {
		return nil
	}

	var count int64
	if err := tx.Model(&entity.Tag{}).
		Where("id IN ?", tagIDs).
		Count(&count).Error; err != nil {
		return err
	}

	if count != int64(len(tagIDs)) {
		return fmt.Errorf("one or more tag IDs do not exist")
	}

	tags := make([]entity.Tag, 0, len(tagIDs))
	for _, id := range tagIDs {
		tags = append(tags, entity.Tag{ID: id})
	}

	if err := tx.Model(person).Association("Tags").Append(tags); err != nil {
		return err
	}

	return nil
}

func (r *PersonRepository) Update(tx *gorm.DB, person *entity.Person, tagIDs []string) error {
	if err := tx.Save(person).Error; err != nil {
		return err
	}

	if len(tagIDs) > 0 {
		var count int64
		if err := tx.Model(&entity.Tag{}).
			Where("id IN ?", tagIDs).
			Count(&count).Error; err != nil {
			return err
		}

		if count != int64(len(tagIDs)) {
			return fmt.Errorf("one or more tag IDs do not exist")
		}

		tags := make([]entity.Tag, 0, len(tagIDs))
		for _, id := range tagIDs {
			tags = append(tags, entity.Tag{ID: id})
		}

		if err := tx.Model(person).Association("Tags").Replace(tags); err != nil {
			return err
		}
	}

	return nil
}

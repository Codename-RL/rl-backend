package repository

import (
	"codename-rl/internal/entity"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TagRepository struct {
	Repository[entity.Tag]
	Log *logrus.Logger
}

func NewTagRepository(log *logrus.Logger) *TagRepository {
	return &TagRepository{
		Log: log,
	}
}

func (r *TagRepository) ExistsByName(tx *gorm.DB, tag *entity.Tag, name string) (bool, error) {
	var exists bool
	err := tx.Model(tag).
		Select("count(*) > 0").
		Where("name = ?", name).
		Find(&exists).Error
	return exists, err
}

func (r *TagRepository) Create(tx *gorm.DB, tag *entity.Tag, personIDs []string) error {

	// 1. Create Tag
	if err := tx.Create(tag).Error; err != nil {
		return err
	}

	// If no persons → skip
	if len(personIDs) == 0 {
		return nil
	}

	// 2. Validate that all PersonIDs actually exist
	var count int64
	if err := tx.Model(&entity.Person{}).
		Where("id IN ?", personIDs).
		Count(&count).Error; err != nil {
		return err
	}

	if count != int64(len(personIDs)) {
		return fmt.Errorf("one or more person IDs do not exist")
	}

	// 3. Prepare dummy Person objects (only IDs needed)
	persons := make([]entity.Person, 0, len(personIDs))
	for _, id := range personIDs {
		persons = append(persons, entity.Person{ID: id})
	}

	// 4. Save relations
	if err := tx.Model(tag).Association("Persons").Append(persons); err != nil {
		return err
	}

	return nil
}

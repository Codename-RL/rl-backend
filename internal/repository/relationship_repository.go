package repository

import (
	"codename-rl/internal/entity"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RelationshipRepository struct {
	Repository[entity.Relationship]
	Log *logrus.Logger
}

func NewRelationshipRepository(log *logrus.Logger) *RelationshipRepository {
	return &RelationshipRepository{
		Log: log,
	}
}

func (r *RelationshipRepository) ExistsByName(tx *gorm.DB, relationship *entity.Relationship, name string) (bool, error) {
	var exists bool
	err := tx.Model(relationship).
		Select("count(*) > 0").
		Where("name = ?", name).
		Find(&exists).Error
	return exists, err
}

func (r *RelationshipRepository) Create(tx *gorm.DB, relationship *entity.Relationship, personIDs []string) error {
	if err := tx.Create(relationship).Error; err != nil {
		return err
	}

	if len(personIDs) == 0 {
		return nil
	}

	var count int64
	if err := tx.Model(&entity.Person{}).
		Where("id IN ?", personIDs).
		Count(&count).Error; err != nil {
		return err
	}

	if count != int64(len(personIDs)) {
		return fmt.Errorf("one or more person IDs do not exist")
	}

	persons := make([]entity.Person, 0, len(personIDs))
	for _, id := range personIDs {
		persons = append(persons, entity.Person{ID: id})
	}

	if err := tx.Model(relationship).Association("Persons").Append(persons); err != nil {
		return err
	}

	return nil
}

func (r *RelationshipRepository) Update(tx *gorm.DB, relationship *entity.Relationship, personIDs []string) error {
	if err := tx.Save(relationship).Error; err != nil {
		return err
	}

	if len(personIDs) > 0 {
		var count int64
		if err := tx.Model(&entity.Person{}).
			Where("id IN ?", personIDs).
			Count(&count).Error; err != nil {
			return err
		}

		if count != int64(len(personIDs)) {
			return fmt.Errorf("one or more person IDs do not exist")
		}

		persons := make([]entity.Person, 0, len(personIDs))
		for _, id := range personIDs {
			persons = append(persons, entity.Person{ID: id})
		}

		if err := tx.Model(relationship).Association("Persons").Replace(persons); err != nil {
			return err
		}
	}

	return nil
}

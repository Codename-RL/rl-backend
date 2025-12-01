package repository

import (
	"codename-rl/internal/entity"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ImportantDateRepository struct {
	Repository[entity.ImportantDate]
	Log *logrus.Logger
}

func NewImportantDateRepository(log *logrus.Logger) *ImportantDateRepository {
	return &ImportantDateRepository{
		Log: log,
	}
}

func (r *ImportantDateRepository) ExistsByName(tx *gorm.DB, phone *entity.ImportantDate, name string) (bool, error) {
	var exists bool
	err := tx.Model(phone).
		Select("count(*) > 0").
		Where("name = ?", name).
		Find(&exists).Error
	return exists, err
}

func (r *ImportantDateRepository) Create(tx *gorm.DB, phone *entity.ImportantDate, personID string) error {

	// 1. Create ImportantDate
	if err := tx.Create(phone).Error; err != nil {
		return err
	}

	// If no persons → skip
	if len(personID) == 0 {
		return nil
	}

	// 2. Validate that all PersonIDs actually exist
	var count int64
	if err := tx.Model(&entity.Person{}).
		Where("id IN ?", personID).
		Count(&count).Error; err != nil {
		return err
	}

	if count != int64(len(personID)) {
		return fmt.Errorf("one or more person IDs do not exist")
	}

	// 3. Prepare dummy Person objects (only IDs needed)
	persons := &entity.Person{ID: personID}

	// 4. Save relations
	if err := tx.Model(phone).Association("Persons").Append(persons); err != nil {
		return err
	}

	return nil
}

package entity

import (
	"time"
)

type ImportantDate struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name;unique"`
	Date      string    `gorm:"column:date;type:timestamptz"`
	PersonID  string    `gorm:"column:person_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:timestamptz"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;type:timestamptz"`

	Person *Person `gorm:"foreignKey:PersonID;references:ID"`
}

func (u *ImportantDate) TableName() string {
	return "important_dates"
}

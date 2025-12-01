package entity

import (
	"time"
)

type Phone struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	Number    string    `gorm:"column:number;unique"`
	PersonID  string    `gorm:"column:person_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:timestamptz"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;type:timestamptz"`

	Person *Person `gorm:"foreignKey:PersonID;references:ID"`
}

func (u *Phone) TableName() string {
	return "phones"
}

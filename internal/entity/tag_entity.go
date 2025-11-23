package entity

import (
	"time"
)

type Tag struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name;unique"`
	UserID    string    `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:timestamptz"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;type:timestamptz"`

	Persons []Person `gorm:"many2many:persons_tags"`
	User    *User    `gorm:"foreignKey:UserID;references:ID"`
}

func (u *Tag) TableName() string {
	return "tags"
}

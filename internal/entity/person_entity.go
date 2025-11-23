package entity

import (
	"time"
)

type Person struct {
	ID          string    `gorm:"column:id;primaryKey"`
	FirstName   string    `gorm:"column:first_name"`
	LastName    string    `gorm:"column:last_name"`
	Nickname    string    `gorm:"column:nickname"`
	Avatar      string    `gorm:"column:avatar"`
	Description string    `gorm:"column:description"`
	UserID      string    `gorm:"column:user_id"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;type:timestamptz"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime;type:timestamptz"`

	Tags          []Tag          `gorm:"many2many:persons_tags"`
	Relationships []Relationship `gorm:"many2many:persons_relationships"`
	User          *User          `gorm:"foreignKey:UserID;references:ID"`
}

func (u *Person) TableName() string {
	return "persons"
}

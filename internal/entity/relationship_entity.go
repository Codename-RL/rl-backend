package entity

import "time"

type Relationship struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name;unique"`
	Color     string    `gorm:"column:color"`
	UserID    string    `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:timestamptz"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;type:timestamptz"`

	Persons []Person `gorm:"many2many:person_relationships"`
	User    *User    `gorm:"foreignKey:UserID;references:ID"`
}

func (u *Relationship) TableName() string {
	return "relationships"
}

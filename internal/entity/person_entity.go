package entity

type Person struct {
	ID          string `gorm:"column:id;primaryKey"`
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	Nickname    string `gorm:"column:nickname"`
	Avatar      string `gorm:"column:avatar"`
	Description string `gorm:"column:description"`
	UserID      string `gorm:"column:user_id"`
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`

	User *User `gorm:"foreignKey:UserID;references:ID"`
}

func (u *Person) TableName() string {
	return "person"
}

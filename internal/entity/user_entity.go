package entity

// User is a struct that represents a user entity
type User struct {
	ID         string `gorm:"column:id;primaryKey"`
	Email      string `gorm:"column:email;unique;not null"`
	Password   string `gorm:"column:password;not null"`
	Name       string `gorm:"column:name"`
	Avatar     string `gorm:"column:avatar"`
	VerifiedAt int64  `gorm:"column:verified_at"`
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt  int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	Token      string `gorm:"-"`

	Otps    []Otp    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Persons []Person `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (u *User) TableName() string {
	return "users"
}

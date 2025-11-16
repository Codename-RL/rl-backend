package entity

// Otp is a struct that represents a otp entity
type Otp struct {
	ID         string `gorm:"column:id;primaryKey"`
	Otp        string `gorm:"column:otp;not null"`
	Token      string `gorm:"column:token"`
	UserID     string `gorm:"column:user_id"`
	VerifiedAt int64  `gorm:"column:verified_at"`
	ExpiresAt  int64  `gorm:"column:expires_at"`
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime:milli"`

	User *User `gorm:"foreignKey:UserID;references:ID"`
}

func (u *Otp) TableName() string {
	return "otps"
}

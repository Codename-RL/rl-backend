package entity

// Otp is a struct that represents a otp entity
type Otp struct {
	ID         string `gorm:"column:id;primaryKey"`
	Otp        string `gorm:"column:otp;not null"`
	Token      string `gorm:"column:token"`
	UserID     string `gorm:"foreignKey:UserID;references:ID"`
	VerifiedAt int64  `gorm:"column:verified_at"`
	ExpiresAt  int64  `gorm:"column:expires_at"`
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime:milli"`
}

func (u *Otp) TableName() string {
	return "otps"
}

package model

type OtpResponse struct {
	ID        string `json:"id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	Token     string `json:"token,omitempty"`
	VerfiedAt int64  `json:"verified_at,omitempty"`
	ExpiresAt int64  `json:"expires_at,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

type VerifyOtpRequest struct {
	Otp   string `json:"otp" validate:"required"`
	Token string `json:"token" validate:"required"`
}

type CreateOtpRequest struct {
	Email string `json:"email" validate:"required"`
}

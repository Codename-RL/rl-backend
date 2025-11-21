package converter

import (
	"codename-rl/internal/entity"
	"codename-rl/internal/model"
)

func OtpToResponse(otp *entity.Otp) *model.OtpResponse {
	if otp == nil {
		return nil
	}

	return &model.OtpResponse{
		ID:        otp.ID,
		UserID:    otp.UserID,
		Token:     otp.Token,
		VerfiedAt: otp.VerifiedAt,
		ExpiresAt: otp.ExpiresAt,
		CreatedAt: otp.CreatedAt,
	}
}

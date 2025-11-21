package converter

import (
	"codename-rl/internal/entity"
	"codename-rl/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	if user == nil {
		return nil
	}

	return &model.UserResponse{
		ID:         user.ID,
		Email:      user.Email,
		Name:       user.Name,
		Avatar:     user.Avatar,
		VerifiedAt: user.VerifiedAt,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

func UserToTokenResponse(user *entity.User) *model.UserResponse {
	if user == nil {
		return nil
	}

	return &model.UserResponse{
		Token: user.Token,
	}
}

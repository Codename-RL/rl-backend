package model

type UserResponse struct {
	ID         string `json:"id,omitempty"`
	Email      string `json:"email,omitempty"`
	Name       string `json:"name,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	Token      string `json:"token,omitempty"`
	VerifiedAt int64  `json:"verified_at,omitempty"`
	CreatedAt  int64  `json:"created_at,omitempty"`
	UpdatedAt  int64  `json:"updated_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required"`
}

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type UpdateUserRequest struct {
	ID         string `json:"-"`
	Email      string `json:"email,omitempty"`
	Name       string `json:"name,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	Password   string `json:"password,omitempty"`
	Token      string `json:"token,omitempty"`
	VerifiedAt string `json:"verified_at,omitempty"`
}
type UpdateUserPasswordRequest struct {
	ID       string `json:"-"`
	Password string `json:"password" validate:"required"`
}

type LoginUserRequest struct {
	ID       string `json:"-"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password" validate:"required"`
}

type LogoutUserRequest struct {
	ID string `json:"-"`
}

type GetUserRequest struct {
	ID string `json:"-"`
}

package model

import "time"

type PersonResponse struct {
	ID          string    `json:"id,omitempty"`
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	Nickname    string    `json:"nickname,omitempty"`
	Avatar      string    `json:"avatar,omitempty"`
	Description string    `json:"description,omitempty"`
	UserID      string    `json:"user_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`

	User *UserResponse `json:"user,omitempty"`
}

type CreatePersonRequest struct {
	FirstName   string   `json:"first_name" validate:"required"`
	LastName    string   `json:"last_name,omitempty"`
	Nickname    string   `json:"nickname,omitempty"`
	Avatar      string   `json:"avatar,omitempty"`
	Description string   `json:"description,omitempty"`
	UserID      string   `json:"-"`
	TagIDs      []string `json:"tag_ids,omitempty"`
}

type UpdatePersonRequest struct {
	ID          string   `json:"id" validate:"required"`
	FirstName   string   `json:"first_name,omitempty"`
	LastName    string   `json:"last_name,omitempty"`
	Nickname    string   `json:"nickname,omitempty"`
	Avatar      string   `json:"avatar,omitempty"`
	Description string   `json:"description,omitempty"`
	UserID      string   `json:"-"`
	TagIDs      []string `json:"tag_ids,omitempty"`
}
type GetPersonRequest struct {
	Query
	UserID string `json:"-"`
}

type DeletePersonRequest struct {
	ID string `json:"id" validate:"required"`
}

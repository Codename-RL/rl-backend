package model

import (
	"time"
)

type TagResponse struct {
	ID        string           `json:"id,omitempty"`
	Name      string           `json:"name,omitempty"`
	UserID    string           `json:"user_id,omitempty"`
	CreatedAt time.Time        `json:"created_at,omitempty"`
	UpdatedAt time.Time        `json:"updated_at,omitempty"`
	Persons   []PersonResponse `json:"persons,omitempty"`
	User      *UserResponse    `json:"user,omitempty"`
}

type CreateTagRequest struct {
	Name      string   `json:"name" validate:"required"`
	PersonIDs []string `json:"person_ids"`
	UserID    string   `json:"-"`
}

type UpdateTagRequest struct {
	ID        string   `json:"id" validate:"required"`
	Name      string   `json:"name" validate:"required"`
	PersonIDs []string `json:"person_ids"`
	UserID    string   `json:"-"`
}
type GetTagRequest struct {
	Query
	UserID string `json:"-"`
}

type DeleteTagRequest struct {
	ID     string `json:"id" validate:"required"`
	UserID string `json:"-"`
}

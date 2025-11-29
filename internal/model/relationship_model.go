package model

import (
	"time"
)

type RelationshipResponse struct {
	ID        string            `json:"id,omitempty"`
	Name      string            `json:"name,omitempty"`
	Color     string            `json:"color,omitempty"`
	UserID    string            `json:"user_id,omitempty"`
	CreatedAt time.Time         `json:"created_at,omitempty"`
	UpdatedAt time.Time         `json:"updated_at,omitempty"`
	Persons   *[]PersonResponse `json:"persons,omitempty"`
	User      *UserResponse     `json:"user,omitempty"`
}

type CreateRelationshipRequest struct {
	Name      string   `json:"name" validate:"required"`
	Color     string   `json:"color"`
	PersonIDs []string `json:"person_ids"`
	UserID    string   `json:"-"`
}

type UpdateRelationshipRequest struct {
	ID        string   `json:"id" validate:"required"`
	Name      string   `json:"name"`
	Color     string   `json:"color"`
	PersonIDs []string `json:"person_ids"`
	UserID    string   `json:"-"`
}
type GetRelationshipRequest struct {
	Query
	UserID string `json:"-"`
}

type DeleteRelationshipRequest struct {
	ID     string `json:"id" validate:"required"`
	UserID string `json:"-"`
}

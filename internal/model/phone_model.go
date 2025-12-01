package model

import (
	"time"
)

type PhoneResponse struct {
	ID        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Number    string          `json:"number,omitempty"`
	PersonID  string          `json:"person_id,omitempty"`
	CreatedAt time.Time       `json:"created_at,omitempty"`
	UpdatedAt time.Time       `json:"updated_at,omitempty"`
	Person    *PersonResponse `json:"person,omitempty"`
}

type CreatePhoneRequest struct {
	Name     string `json:"name" validate:"required"`
	Number   string `json:"number" validate:"required"`
	PersonID string `json:"person_id"`
	UserID   string `json:"-"`
}

type UpdatePhoneRequest struct {
	ID       string `json:"id" validate:"required"`
	Number   string `json:"number"`
	Name     string `json:"name"`
	PersonID string `json:"person_id"`
	UserID   string `json:"-"`
}
type GetPhoneRequest struct {
	Query
	UserID string `json:"-"`
}

type DeletePhoneRequest struct {
	ID     string `json:"id" validate:"required"`
	UserID string `json:"-"`
}

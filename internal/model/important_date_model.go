package model

import (
	"time"
)

type ImportantDateResponse struct {
	ID        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Date      string          `json:"date,omitempty"`
	PersonID  string          `json:"person_id,omitempty"`
	CreatedAt time.Time       `json:"created_at,omitempty"`
	UpdatedAt time.Time       `json:"updated_at,omitempty"`
	Person    *PersonResponse `json:"person,omitempty"`
}

type CreateImportantDateRequest struct {
	Name     string `json:"name" validate:"required"`
	Date     string `json:"date" validate:"required"`
	PersonID string `json:"person_id"`
	UserID   string `json:"-"`
}

type UpdateImportantDateRequest struct {
	ID       string `json:"id" validate:"required"`
	Date     string `json:"date"`
	Name     string `json:"name"`
	PersonID string `json:"person_id"`
	UserID   string `json:"-"`
}
type GetImportantDateRequest struct {
	Query
	UserID string `json:"-"`
}

type DeleteImportantDateRequest struct {
	ID     string `json:"id" validate:"required"`
	UserID string `json:"-"`
}

package model

type BaseResponse[T any] struct {
	Message      string          `json:"message"`
	Data         T               `json:"data,omitempty"`
	PageMetadata *PaginationMeta `json:"meta,omitempty"`
	Errors       string          `json:"errors,omitempty"`
	Timestamp    string          `json:"timestamp"`
}

type PageResponse[T any] struct {
	Message      string         `json:"message"`
	Data         []T            `json:"data,omitempty"`
	PageMetadata PaginationMeta `json:"meta,omitempty"`
	Timestamp    string         `json:"timestamp"`
}

type PaginationMeta struct {
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalPages int64 `json:"total_pages"`
}

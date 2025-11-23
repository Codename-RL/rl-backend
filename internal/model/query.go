package model

type DateRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Query struct {
	Search     map[string]string    `json:"search"` // field → value
	Or         bool                 `json:"or"`     // optional: OR instead of AND
	SortBy     string               `json:"sort_by"`
	Order      string               `json:"order"`
	Limit      int                  `json:"limit"`
	Offset     int                  `json:"offset"`
	DateRanges map[string]DateRange `json:"date_ranges"`
	Preload    []string             `json:"preload"`
}

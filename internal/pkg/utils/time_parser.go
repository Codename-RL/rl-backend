package utils

import (
	"codename-rl/internal/model"
	"fmt"
	"time"
)

// Acceptable time layouts
var timeLayouts = []string{
	time.RFC3339,          // 2025-11-19T10:04:58+08:00
	"2006-01-02T15:04:05", // 2025-11-19T10:04:58
	"2006-01-02 15:04:05", // 2025-11-19 10:04:58
	"2006-01-02",          // 2025-11-19
}

// ParseFlexibleTime returns a real time.Time
// Adds day start/end automatically depending on field
func ParseFlexibleTime(s string) (time.Time, error) {
	// Try known formats
	for _, layout := range timeLayouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			// If only date is provided (YYYY-MM-DD)
			if len(s) == 10 {
				// Mark as full day start or end via caller
				return t, nil

			}
			return t, nil
		}
	}
	t, err := time.Parse(time.RFC3339Nano, s)

	// Final fallback: try RFC3339Nano
	return t, err
}

func ValidateDateRange(q *model.Query) error {
	for field, dr := range q.DateRanges {
		if dr.From != "" {
			if _, err := ParseFlexibleTime(dr.From); err != nil {
				return fmt.Errorf("invalid '%s_from' format: %v", field, err)
			}
		}
		if dr.To != "" {
			if _, err := ParseFlexibleTime(dr.To); err != nil {
				return fmt.Errorf("invalid '%s_to' format: %v", field, err)
			}
		}
	}
	return nil
}

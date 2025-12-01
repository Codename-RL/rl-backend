package converter

import (
	"codename-rl/internal/entity"
	"codename-rl/internal/model"
)

func ImportantDateToResponse(importantDate *entity.ImportantDate) *model.ImportantDateResponse {
	if importantDate == nil {
		return nil
	}

	return &model.ImportantDateResponse{
		ID:        importantDate.ID,
		Name:      importantDate.Name,
		Date:      importantDate.Date,
		PersonID:  importantDate.PersonID,
		CreatedAt: importantDate.CreatedAt,
		UpdatedAt: importantDate.UpdatedAt,
		Person:    PersonToResponse(importantDate.Person),
	}
}

func ImportantDatesToResponses(importantDates *[]entity.ImportantDate) *[]model.ImportantDateResponse {
	if importantDates == nil {
		return nil
	}

	responses := make([]model.ImportantDateResponse, 0, len(*importantDates))

	for _, importantDate := range *importantDates {
		responses = append(responses, model.ImportantDateResponse{
			ID:        importantDate.ID,
			Name:      importantDate.Name,
			Date:      importantDate.Date,
			PersonID:  importantDate.PersonID,
			CreatedAt: importantDate.CreatedAt,
			UpdatedAt: importantDate.UpdatedAt,
			Person:    PersonToResponse(importantDate.Person),
		})
	}

	return &responses
}

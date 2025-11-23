package converter

import (
	"codename-rl/internal/entity"
	"codename-rl/internal/model"
)

func PersonToResponse(person *entity.Person) *model.PersonResponse {
	if person == nil {
		return nil
	}

	return &model.PersonResponse{
		ID:          person.ID,
		FirstName:   person.FirstName,
		LastName:    person.LastName,
		Nickname:    person.Nickname,
		Avatar:      person.Avatar,
		Description: person.Description,
		UserID:      person.UserID,
		CreatedAt:   person.CreatedAt,
		UpdatedAt:   person.UpdatedAt,
		User:        UserToResponse(person.User),
	}
}

func PersonsToResponses(persons *[]entity.Person) *[]model.PersonResponse {
	if persons == nil {
		return nil
	}

	responses := make([]model.PersonResponse, 0, len(*persons))

	for _, person := range *persons {
		responses = append(responses, model.PersonResponse{
			ID:          person.ID,
			FirstName:   person.FirstName,
			LastName:    person.LastName,
			Nickname:    person.Nickname,
			Avatar:      person.Avatar,
			Description: person.Description,
			UserID:      person.UserID,
			CreatedAt:   person.CreatedAt,
			UpdatedAt:   person.UpdatedAt,
			User:        UserToResponse(person.User),
		})
	}

	return &responses
}

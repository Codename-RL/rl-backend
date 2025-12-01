package converter

import (
	"codename-rl/internal/entity"
	"codename-rl/internal/model"
)

func PhoneToResponse(phone *entity.Phone) *model.PhoneResponse {
	if phone == nil {
		return nil
	}

	return &model.PhoneResponse{
		ID:        phone.ID,
		Name:      phone.Name,
		Number:    phone.Number,
		PersonID:  phone.PersonID,
		CreatedAt: phone.CreatedAt,
		UpdatedAt: phone.UpdatedAt,
		Person:    PersonToResponse(phone.Person),
	}
}

func PhonesToResponses(phones *[]entity.Phone) *[]model.PhoneResponse {
	if phones == nil {
		return nil
	}

	responses := make([]model.PhoneResponse, 0, len(*phones))

	for _, phone := range *phones {
		responses = append(responses, model.PhoneResponse{
			ID:        phone.ID,
			Name:      phone.Name,
			Number:    phone.Number,
			PersonID:  phone.PersonID,
			CreatedAt: phone.CreatedAt,
			UpdatedAt: phone.UpdatedAt,
			Person:    PersonToResponse(phone.Person),
		})
	}

	return &responses
}

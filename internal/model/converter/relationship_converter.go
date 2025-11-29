package converter

import (
	"codename-rl/internal/entity"
	"codename-rl/internal/model"
)

func RelationshipToResponse(relationship *entity.Relationship) *model.RelationshipResponse {
	if relationship == nil {
		return nil
	}

	return &model.RelationshipResponse{
		ID:        relationship.ID,
		Name:      relationship.Name,
		Color:     relationship.Color,
		UserID:    relationship.UserID,
		CreatedAt: relationship.CreatedAt,
		UpdatedAt: relationship.UpdatedAt,
		User:      UserToResponse(relationship.User),
	}
}

func RelationshipsToResponses(relationships *[]entity.Relationship) *[]model.RelationshipResponse {
	if relationships == nil {
		return nil
	}

	responses := make([]model.RelationshipResponse, 0, len(*relationships))

	for _, relationship := range *relationships {
		responses = append(responses, model.RelationshipResponse{
			ID:        relationship.ID,
			Name:      relationship.Name,
			Color:     relationship.Color,
			UserID:    relationship.UserID,
			CreatedAt: relationship.CreatedAt,
			UpdatedAt: relationship.UpdatedAt,
			Persons:   PersonsToResponses(&relationship.Persons),
			User:      UserToResponse(relationship.User),
		})
	}

	return &responses
}

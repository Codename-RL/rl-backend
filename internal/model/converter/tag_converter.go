package converter

import (
	"codename-rl/internal/entity"
	"codename-rl/internal/model"
)

func TagToResponse(tag *entity.Tag) *model.TagResponse {
	if tag == nil {
		return nil
	}

	return &model.TagResponse{
		ID:        tag.ID,
		Name:      tag.Name,
		UserID:    tag.UserID,
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdatedAt,
		//Persons:   tag.Persons,
		User: UserToResponse(tag.User),
	}
}

func TagsToResponses(tags *[]entity.Tag) *[]model.TagResponse {
	if tags == nil {
		return nil
	}

	responses := make([]model.TagResponse, 0, len(*tags))

	for _, tag := range *tags {
		responses = append(responses, model.TagResponse{
			ID:        tag.ID,
			Name:      tag.Name,
			UserID:    tag.UserID,
			CreatedAt: tag.CreatedAt,
			UpdatedAt: tag.UpdatedAt,
			//Persons:   tag.Persons,
			User: UserToResponse(tag.User),
		})
	}

	return &responses
}

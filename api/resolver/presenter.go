package resolver

import (
	"atlas/api/usecase"
	"atlas/graphql/model"
)

func toUserModel(input *usecase.User) *model.User {
	return &model.User{
		ID:    input.Id,
		Name:  input.Name,
		Email: input.Email,
	}
}

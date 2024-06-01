package resolver

import (
	"github.com/spark-tokyo/atlas/api/usecase"
	"github.com/spark-tokyo/atlas/graphql/model"
)

func toUserModel(input *usecase.User) *model.User {
	return &model.User{
		ID:    input.Id,
		Name:  input.Name,
		Email: input.Email,
	}
}

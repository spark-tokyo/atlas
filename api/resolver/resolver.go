package resolver

import "github.com/spark-tokyo/atlas/api/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userUsecase usecase.IFUserUsecase
}

func NewResolver(
	userUsecase *usecase.UserUsecase,
) *Resolver {
	return &Resolver{
		userUsecase: userUsecase,
	}
}

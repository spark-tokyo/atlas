package provider

import (
	"github.com/google/wire"

	"atlas/api/usecase"
)

var usecaseSet = wire.NewSet(
	usecase.NewUserUsecase,
)

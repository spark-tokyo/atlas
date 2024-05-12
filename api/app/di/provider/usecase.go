package provider

import (
	"github.com/google/wire"

	"atlas/api/usecase"
)

var newUsecase = wire.NewSet(
	usecase.NewUserUsecase,
)

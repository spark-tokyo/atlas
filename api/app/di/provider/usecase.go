package provider

import (
	"github.com/google/wire"

	"github.com/spark-tokyo/atlas/api/usecase"
)

var usecaseSet = wire.NewSet(
	usecase.NewUserUsecase,
)

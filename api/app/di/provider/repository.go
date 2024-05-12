package provider

import (
	"github.com/google/wire"

	"atlas/api/repository"
)

var newRepository = wire.NewSet(
	repository.NewUserRepository,
)

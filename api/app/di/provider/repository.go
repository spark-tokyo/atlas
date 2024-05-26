package provider

import (
	"github.com/google/wire"

	"atlas/api/repository"
)

var repositorySet = wire.NewSet(
	repository.NewUserRepository,
)

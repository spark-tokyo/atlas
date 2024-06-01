package provider

import (
	"github.com/google/wire"

	"github.com/spark-tokyo/atlas/api/repository"
)

var repositorySet = wire.NewSet(
	repository.NewUserRepository,
)

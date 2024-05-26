package provider

import (
	"github.com/google/wire"

	"atlas/api/infra"
)

var infraSet = wire.NewSet(
	infra.NewEnt,
)

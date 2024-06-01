package provider

import (
	"github.com/google/wire"

	"github.com/spark-tokyo/atlas/api/infra"
)

var infraSet = wire.NewSet(
	infra.NewEnt,
)

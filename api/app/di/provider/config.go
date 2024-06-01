package provider

import (
	"github.com/google/wire"

	"github.com/spark-tokyo/atlas/config"
)

var configSet = wire.NewSet(
	config.NewLoadConfig,
)

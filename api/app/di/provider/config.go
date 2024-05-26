package provider

import (
	"github.com/google/wire"

	"atlas/config"
)

var configSet = wire.NewSet(
	config.NewLoadConfig,
)

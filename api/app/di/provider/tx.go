package provider

import (
	"github.com/google/wire"

	"github.com/spark-tokyo/atlas/tx"
)

var txSet = wire.NewSet(
	tx.NewTxManager,
)

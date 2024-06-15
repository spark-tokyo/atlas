package repository

import (
	"github.com/spark-tokyo/atlas/ent"
	"github.com/spark-tokyo/atlas/utils/exception"
)

func FetchTx(tx *ent.Tx) (*ent.Tx, error) {
	if tx == nil {
		return nil, exception.New(exception.Internal, "no transaction")
	}
	return tx, nil
}

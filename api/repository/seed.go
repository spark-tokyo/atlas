package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spark-tokyo/atlas/api/repository/fixtures"
	"github.com/spark-tokyo/atlas/ent"
)

func InitSeed(ctx context.Context, tx *ent.Tx, t *testing.T) {
	t.Helper()
	err := tx.User.CreateBulk(fixtures.SetUsers(tx)...).Exec(ctx)
	assert.NoError(t, err)
	err = tx.Pet.CreateBulk().Exec(ctx)
	assert.NoError(t, err)
}

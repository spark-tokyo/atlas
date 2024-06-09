package test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/zeebo/assert"

	"github.com/spark-tokyo/atlas/api/infra"
	"github.com/spark-tokyo/atlas/ent"
	"github.com/spark-tokyo/atlas/ent/enttest"
)

func NewTestClient(ctx context.Context, t *testing.T) *infra.Ent {
	t.Helper()
	client := enttest.Open(t, "sqlite3", "file::memory:?_fk=1") // nolint
	client.Use(
	// wrapCreatedAt(),
	// wrapUpdatedAt(),
	// wrapCheckPredicateSize(),
	)
	return &infra.Ent{Client: client}
}

// FetchTestReadWriteTransaction はテスト用の読み書きトランザクションを返す。
// NOTE: 単体テストで、テストデータを用意する際にWriteが必要なためReadOnlyは用意しない
func FetchTestReadWriteTransaction(ctx context.Context, t *testing.T) (*ent.Tx, func()) {
	t.Helper()
	client := NewTestClient(ctx, t)
	tx, err := client.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	assert.NoError(t, err)
	return tx, func() {
		err := tx.Commit()
		assert.NoError(t, err)
		client.Close()
	}
}

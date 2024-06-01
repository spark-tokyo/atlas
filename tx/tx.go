package tx

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/spark-tokyo/atlas/api/infra"
	"github.com/spark-tokyo/atlas/ent"
)

// type IFTx interface {
// 	Begin(ctx context.Context) (*ent.Tx, error)
// 	Commit(ctx context.Context, tx *ent.Tx) error
// 	Rollback(ctx context.Context, tx *ent.Tx) error
// }

// トランザクションを貼るメソッドを定義
type IFTxManager interface {
	WitTx(ctx context.Context, fn func(ctx context.Context, tx *ent.Tx) error) error
}

type TxManager struct {
	ent *infra.Ent
}

func NewTxManager(
	ent *infra.Ent,
) *TxManager {
	return &TxManager{
		ent: ent,
	}
}

// トランジションを貼る repositoryは基本トランザクションを貼って処理するため、包括的にするためにusecaseで使用
func (m *TxManager) WitTx(ctx context.Context, fn func(ctx context.Context, tx *ent.Tx) error) error {
	// トランザクションの開始
	tx, err := m.ent.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return errors.New(err.Error())
	}

	// エラーが発生していた場合はロールバックする
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Println(rollbackErr.Error())
			}
			return
		}
	}()

	// トランザクション処理を実行
	err = fn(ctx, tx)
	if err != nil {
		return errors.New(err.Error())
	}

	// 処理を保存する
	if err = tx.Commit(); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

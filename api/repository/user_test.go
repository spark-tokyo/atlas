package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/spark-tokyo/atlas/api/entity"
	"github.com/spark-tokyo/atlas/api/test"
	"github.com/spark-tokyo/atlas/ent"
	"github.com/spark-tokyo/atlas/utils/helpertest"
)

func TestUser_Get(t *testing.T) {
	t.Parallel()

	type Args struct {
		ctx context.Context
		tx  *ent.Tx
		id  string
	}

	type Returns struct {
		user *entity.User
		err  error
	}

	type testContext struct {
		args      Args
		returns   Returns
		closeFunc func()
	}

	tests := []struct {
		name        string
		testContext func() *testContext
	}{
		{
			name: "正常: 取得",
			testContext: func() *testContext {
				ctx := context.Background()
				tx, closeFunc := test.FetchTestReadWriteTransaction(ctx, t)
				tx, err := FetchTx(tx)
				assert.NoError(t, err)
				InitSeed(ctx, tx, t)

				args := Args{
					tx:  tx,
					ctx: ctx,
					id:  "1",
				}

				entity := &entity.User{
					Id:    "1",
					Name:  "Name",
					Email: "Email",
				}
				returns := Returns{
					user: entity,
					err:  nil,
				}

				return &testContext{
					args:      args,
					returns:   returns,
					closeFunc: closeFunc,
				}
			},
		},
		{
			name: "異常: 存在しないID",
			testContext: func() *testContext {
				ctx := context.Background()
				tx, closeFunc := test.FetchTestReadWriteTransaction(ctx, t)
				tx, err := FetchTx(tx)
				assert.NoError(t, err)
				InitSeed(ctx, tx, t)

				args := Args{
					tx:  tx,
					ctx: ctx,
					id:  "4",
				}

				returns := Returns{
					user: nil,
					err:  errors.New("ent: user not found"),
				}

				return &testContext{
					args:      args,
					returns:   returns,
					closeFunc: closeFunc,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helpertest.RunWithMock(t, func(ctrl *gomock.Controller) {
				tc := tt.testContext()
				repo := &UserRepository{}
				t.Cleanup(tc.closeFunc)
				user, err := repo.Get(tc.args.ctx, tc.args.tx, tc.args.id)
				if tc.returns.err != nil {
					assert.EqualError(t, tc.returns.err, err.Error())
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tc.returns.user, user)
				}
			})
		})
	}
}

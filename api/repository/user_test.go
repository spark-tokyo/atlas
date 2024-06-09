package repository

import (
	"context"
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
				if tx == nil {
					panic("tx is nil")
				}
				InitSeed(ctx, tx, t)

				args := Args{
					ctx: ctx,
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helpertest.RunWithMock(t, func(ctrl *gomock.Controller) {
				print("testing")
				tc := tt.testContext()
				repo := &UserRepository{}
				t.Cleanup(tc.closeFunc)
				print("cleaning up repo")
				user, err := repo.Get(tc.args.ctx, tc.args.tx)
				print("user")
				if tc.returns.err != nil {
					panic(err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tc.returns.user, user)
				}
			})
		})
	}
}

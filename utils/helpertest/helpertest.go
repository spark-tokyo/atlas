package helpertest

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"
)

func RunWithMock(t *testing.T, f func(ctrl *gomock.Controller)) {
	ctx, cancel := context.WithCancel(context.Background())
	ctrl, ctx := gomock.WithContext(ctx, t)
	defer ctrl.Finish()

	go func() {
		f(ctrl)
		cancel()
	}()

	<-ctx.Done()
}

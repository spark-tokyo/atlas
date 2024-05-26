//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"atlas/api/app/di/provider"
)

var NewSet = wire.NewSet(
	provider.NewApp,
)

func NewApp() (*provider.App, func(), error) {
	wire.Build(
		provider.NewSet,
	)

	return nil, nil, nil
}

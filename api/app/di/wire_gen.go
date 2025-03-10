// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire gen -tags "wireinject"
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"github.com/spark-tokyo/atlas/api/app/di/provider"
	"github.com/spark-tokyo/atlas/api/infra"
	"github.com/spark-tokyo/atlas/api/repository"
	"github.com/spark-tokyo/atlas/api/resolver"
	"github.com/spark-tokyo/atlas/api/usecase"
	"github.com/spark-tokyo/atlas/config"
	"github.com/spark-tokyo/atlas/router"
	"github.com/spark-tokyo/atlas/tx"
)

// Injectors from wire.go:

func NewApp() (*provider.App, func(), error) {
	configConfig, err := config.NewLoadConfig()
	if err != nil {
		return nil, nil, err
	}
	userRepository := repository.NewUserRepository(configConfig)
	ent, err := infra.NewEnt(configConfig)
	if err != nil {
		return nil, nil, err
	}
	txManager := tx.NewTxManager(ent)
	userUsecase := usecase.NewUserUsecase(userRepository, txManager, ent)
	resolverResolver := resolver.NewResolver(userUsecase)
	routerRouter := router.NewRouter(resolverResolver, configConfig)
	app := provider.NewApp(routerRouter, configConfig)
	return app, func() {
	}, nil
}

// wire.go:

var NewSet = wire.NewSet(provider.NewApp)

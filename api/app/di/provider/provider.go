package provider

import (
	"github.com/google/wire"

	"github.com/spark-tokyo/atlas/api/resolver"
	"github.com/spark-tokyo/atlas/config"
	"github.com/spark-tokyo/atlas/router"
)

var NewSet = wire.NewSet(
	NewApp,
	resolver.NewResolver,
	router.NewRouter,

	txSet,
	middlewareSet,
	contlloerSet,
	gatewaySet,
	infraSet,
	usecaseSet,
	repositorySet,
	configSet,
)

type App struct {
	Router *router.Router
	Config *config.Config
}

func NewApp(
	router *router.Router,
	cfg *config.Config,
) *App {
	return &App{
		Router: router,
		Config: cfg,
	}
}

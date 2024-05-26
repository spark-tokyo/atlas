package provider

import (
	"github.com/google/wire"

	"atlas/api/resolver"
	"atlas/config"
	"atlas/router"
)

var NewSet = wire.NewSet(
	NewApp,
	resolver.NewResolver,
	router.NewRouter,

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

package provider

import (
	"github.com/google/wire"

	"atlas/api/resolver"
	"atlas/router"
)

var NewSet = wire.NewSet(
	NewApp,
	resolver.NewResolver,
	router.NewRouter,

	middleware,
	contlloer,
	gateway,
	infra,
	newRepository,
	newUsecase,
)

type App struct {
	Router *router.Router
}

func NewApp(
	router *router.Router,
) *App {
	return &App{
		Router: router,
	}
}

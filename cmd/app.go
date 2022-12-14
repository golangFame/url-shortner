package main

import (
	"github.com/goferHiro/url-shortner/config"
	"github.com/goferHiro/url-shortner/handlers"
	"github.com/goferHiro/url-shortner/internal"
	"github.com/goferHiro/url-shortner/middlewares"
	"github.com/goferHiro/url-shortner/router"
	"github.com/goferHiro/url-shortner/server"
	"github.com/goferHiro/url-shortner/services"
	"go.uber.org/fx"
)

func serverRun() {
	app := fx.New(
		config.Module,
		internal.Module,
		services.Module,
		middlewares.Module,
		handlers.Module,
		router.Module,
		server.Module,
	)
	app.Run()
}

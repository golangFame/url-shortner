package services

import (
	"github.com/goferHiro/url-shortner/services/app"
	"github.com/goferHiro/url-shortner/services/auth"
	"github.com/goferHiro/url-shortner/services/dummy"
	"go.uber.org/fx"
)

var Module = fx.Options(
	dummy.Module,
	app.Module,
	auth.Module,
)

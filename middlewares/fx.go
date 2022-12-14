package middlewares

import (
	"github.com/goferHiro/url-shortner/services/auth"
	"github.com/hiroBzinga/bun"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"os"
)

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

type in struct {
	fx.In
	Conf         *viper.Viper
	DB           *bun.DB `name:"db"`
	AuthServices auth.Services
}

type out struct {
	fx.Out
	Middlewares *Middleware
}

func New(i in) (o out) {
	m := &Middleware{
		&logrus.Logger{
			Out:       os.Stderr,
			Formatter: new(logrus.TextFormatter),
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		},
		i.DB,
		i.Conf,
		i.AuthServices,
	}
	o = out{
		Middlewares: m,
	}
	return
}

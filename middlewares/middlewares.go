package middlewares

import (
	"github.com/goferHiro/url-shortner/services/auth"
	"github.com/hiroBzinga/bun"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Middleware struct {
	log          *logrus.Logger //not yet used
	db           *bun.DB
	conf         *viper.Viper
	authServices auth.Services
}

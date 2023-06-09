package router

import (
	"github.com/goferHiro/url-shortner/handlers"
	"github.com/goferHiro/url-shortner/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//type services struct {
//	Root NoAuthRouting
//}

type service struct {
	conf         *viper.Viper
	middlewares  *middlewares.Middleware
	homeHandler  *handlers.HomeHandler
	dummyHandler *handlers.DummyHandler
	appHandler   *handlers.AppHandler
}

type Middleware func() gin.HandlerFunc

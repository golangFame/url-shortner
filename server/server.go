package server

import (
	"github.com/gin-gonic/gin"
	"github.com/goferHiro/url-shortner/enums"
	"github.com/goferHiro/url-shortner/middlewares"
	"github.com/goferHiro/url-shortner/router"
	"github.com/hiroBzinga/bun"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Server struct {
	middlewares    *middlewares.Middleware
	routerServices router.Services
	conf           *viper.Viper
	log            *logrus.Logger
	db             *bun.DB
}

func (s *Server) setupRouter() (engine *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)

	engine = gin.New()
	engine.Use(gin.Recovery(), gin.Logger())
	//engine.Use(gindump.Dump())
	engine.Use(s.middlewares.Logger())

	if s.conf.GetString(enums.MODE) == enums.DEVELOPMENT {

	}

	root := engine.Group("/")

	s.routerServices.RoutesWithNoAuth(root)
	return
}

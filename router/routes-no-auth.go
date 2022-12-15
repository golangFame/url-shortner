package router

import (
	"github.com/gin-gonic/gin"
)

func (s *service) RoutesWithNoAuth(r *gin.RouterGroup, mws ...Middleware) {

	for _, i := range mws {
		r.Use(i())
	}

	r.GET("/", s.homeHandler.Home)

	r.POST("/dummy", s.dummyHandler.Dummy)

	appRouter := r.Group("/app")

	appRouter.GET("/", s.appHandler.Home)

	appRouter.POST("/url", s.appHandler.ShortenUrl)
	appRouter.GET("/url/:surl", s.appHandler.RetrieveUrl)

}

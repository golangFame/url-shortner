package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/goferHiro/url-shortner/entities"
	"net/http"
)

func (h *AppHandler) Home(c *gin.Context) {
	var (
		res = entities.GenericResponse{}
		//err error
	)

	//defer h.handleResponse(c, &res, &err)

	res.Success = true
	res.Message = "Home service is up and running"
	c.JSON(http.StatusOK, res)
}

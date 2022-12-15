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
	res.Message = "App service is up and running"
	c.JSON(http.StatusOK, res)
}

type urlReq struct {
	url string `json:"url" binding:"required"`
}

func (h *AppHandler) ShortenUrl(c *gin.Context) {

	var u1 urlReq

	if err := c.ShouldBind(&u1); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request does not match the accepted format",
		})
		return

	}

	shortUrl := h.appServices.ShortenUrl(u1.url)

	c.JSON(http.StatusOK, gin.H{
		"short-url": shortUrl,
		"url":       u1.url,
	})

}

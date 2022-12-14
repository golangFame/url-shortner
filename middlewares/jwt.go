package middlewares

import (
	"github.com/gin-gonic/gin"
	models_enums "github.com/goferHiro/url-shortner/models/models-enums"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (m *Middleware) AuthorizeJWT() gin.HandlerFunc {

	return func(c *gin.Context) {

		//goland:noinspection GoSnakeCaseUsage
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := m.authServices.ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)

			c.Set(models_enums.AppUserID, claims["ID"])

			m.log.Info("Claims:", claims)
			c.Next()
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

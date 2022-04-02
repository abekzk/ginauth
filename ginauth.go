package ginauth

import "github.com/gin-gonic/gin"

type Auth interface {
	apply(*gin.Context)
}

func NewAuthorizer(auth Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.apply(c)
		c.Next()
	}
}

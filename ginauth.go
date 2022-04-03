package ginauth

import "github.com/gin-gonic/gin"

type Provider interface {
	apply(*gin.Context)
}

func NewAuthorizer(p Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		p.apply(c)
		c.Next()
	}
}

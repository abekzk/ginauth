package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzuabe/ginauth"
)

func main() {
	router := gin.Default()

	auth := ginauth.NewFirebaseAuth()
	router.Use(ginauth.NewAuthorizer(auth))

	router.GET("/", func(c *gin.Context) {
		token := c.MustGet(ginauth.FirebaseAuthTokenKey).(ginauth.FirebaseAuthToken)
		c.String(http.StatusOK, "Your ID is %s", token.UID)
	})

	router.Run()
}

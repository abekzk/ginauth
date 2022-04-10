package main

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/kzuabe/ginauth"
)

func main() {
	router := gin.Default()

	provider := ginauth.NewFirebaseAuthProvider()
	router.Use(ginauth.NewAuthorizer(provider))

	router.GET("/", func(c *gin.Context) {
		token := c.MustGet(ginauth.FirebaseAuthTokenKey).(*auth.Token)
		c.String(http.StatusOK, "Your ID is %s", token.UID)
	})

	router.Run()
}

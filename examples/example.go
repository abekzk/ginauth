package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzuabe/ginauth"
)

func main() {
	router := gin.Default()
	client, err := ginauth.NewFirebaseClient()
	if err != nil {
		log.Fatalf("error initializing: %v\n", err)
	}

	router.Use(ginauth.NewFirebaseAuthorizer(client))

	router.GET("/", func(c *gin.Context) {
		token := c.MustGet(ginauth.FirebaseAuthTokenKey).(ginauth.FirebaseAuthToken)
		c.String(http.StatusOK, "Your ID is %s", token.UID)
	})

	router.Run()
}

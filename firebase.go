package ginauth

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

const FirebaseAuthTokenKey = "firebaseToken"

type FirebaseClient *auth.Client

// ref: https://firebase.google.com/docs/admin/setup#go
func NewFirebaseClient() (FirebaseClient, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	return client, nil
}

// ref: https://firebase.google.com/docs/auth/admin/verify-id-tokens#go
func NewFirebaseAuthorizer(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken, ok := extractTokenFromAuthHeader(c.Request.Header.Get("Authorization"))
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set(FirebaseAuthTokenKey, token)
		c.Next()
	}
}

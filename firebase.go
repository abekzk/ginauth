package ginauth

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

const FirebaseAuthTokenKey = "firebaseToken"

type FirebaseAuthToken *auth.Token
type FirebaseClient interface {
	VerifyIDToken(context.Context, string) (*auth.Token, error)
}

// ref: https://firebase.google.com/docs/admin/setup#go
func NewFirebaseClient() (*auth.Client, error) {
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
func NewFirebaseAuthorizer(client FirebaseClient) gin.HandlerFunc {
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
		c.Set(FirebaseAuthTokenKey, FirebaseAuthToken(token))
		c.Next()
	}
}

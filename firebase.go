package ginauth

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

const FirebaseAuthTokenKey = "firebaseToken"

type FirebaseAuthToken *auth.Token

type firebaseClient interface {
	VerifyIDToken(context.Context, string) (*auth.Token, error)
}

type FirebaseAuth struct {
	Client firebaseClient
}

// ref: https://firebase.google.com/docs/admin/setup#go
func NewFirebaseAuth() *FirebaseAuth {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	return &FirebaseAuth{
		Client: client,
	}
}

// ref: https://firebase.google.com/docs/auth/admin/verify-id-tokens#go
func (firebaseAuth *FirebaseAuth) apply(c *gin.Context) {
	idToken, ok := extractTokenFromAuthHeader(c.Request.Header.Get("Authorization"))
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := firebaseAuth.Client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set(FirebaseAuthTokenKey, FirebaseAuthToken(token))
}

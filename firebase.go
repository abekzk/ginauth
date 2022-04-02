package ginauth

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type FirebaseClient *auth.Client

// ref: https://github.com/firebase/firebase-admin-go/blob/e60757f9b29711f19fa1f44ce9b5a6fae3baf3a5/snippets/init.go#L82-L98
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

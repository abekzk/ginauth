package ginauth

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockfirebaseClient struct {
	mock.Mock
}

func (m *mockfirebaseClient) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	args := m.Called(ctx, idToken)
	return args.Get(0).(*auth.Token), args.Error(1)
}

func newTestRouterFirebaseAuth(p Provider) *gin.Engine {
	router := gin.New()

	router.Use(NewAuthorizer(p))

	router.GET("/", func(c *gin.Context) {
		token := c.MustGet(FirebaseAuthTokenKey).(*auth.Token)
		c.String(http.StatusOK, token.UID)
	})

	return router
}

func newTestFirebaseAuthProvider(client firebaseClient) *FirebaseAuthProvider {
	return &FirebaseAuthProvider{Client: client}
}

func performRequest(router http.Handler, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestFirebaseAuthorize(t *testing.T) {
	mockClient := new(mockfirebaseClient)
	mockClient.On("VerifyIDToken", mock.Anything, "token1").Return(&auth.Token{UID: "user1"}, nil)

	provider := newTestFirebaseAuthProvider(mockClient)
	router := newTestRouterFirebaseAuth(provider)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "bearer token1")

	w := performRequest(router, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "user1", w.Body.String())
	mockClient.AssertExpectations(t)
}

func TestFirebaseAuthorizeError(t *testing.T) {
	mockClient := new(mockfirebaseClient)
	mockClient.On("VerifyIDToken", mock.Anything, "token1").Return(&auth.Token{}, errors.New("error1"))

	provider := newTestFirebaseAuthProvider(mockClient)
	router := newTestRouterFirebaseAuth(provider)

	// No header
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := performRequest(router, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Invalid token
	req, _ = http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "bearer token1")
	w = performRequest(router, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	mockClient.AssertExpectations(t)
}

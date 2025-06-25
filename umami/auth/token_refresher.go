package auth

import (
	"context"
	"github.com/AdamShannag/umami-client/umami/token"
	"time"
)

type TokenRefresherAuth struct {
	token *token.Refresher
}

func NewTokenRefresherAuth(ctx context.Context, GetTokenFunc func() (string, time.Duration, error)) Auth {
	return &TokenRefresherAuth{token: token.NewRefresher(ctx, GetTokenFunc)}
}

func (a *TokenRefresherAuth) Get() (string, error) {
	get, err := a.token.Get()
	return "Bearer " + get, err
}

func (a *TokenRefresherAuth) Header() string {
	return "Authorization"
}

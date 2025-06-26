package auth_test

import (
	"context"
	"errors"
	"github.com/AdamShannag/umami-client/umami/auth"
	"testing"
	"time"
)

func TestSingleTokenAuth(t *testing.T) {
	a := auth.NewSingleTokenAuth("abc123")

	header := a.Header()
	if header != "Authorization" {
		t.Errorf("expected header Authorization, got %s", header)
	}

	token, err := a.Get()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if token != "Bearer abc123" {
		t.Errorf("expected token 'Bearer abc123', got %s", token)
	}
}

func TestApiKeyAuth(t *testing.T) {
	a := auth.NewApiKeyAuth("key-123")

	header := a.Header()
	if header != "x-umami-api-key" {
		t.Errorf("expected header x-umami-api-key, got %s", header)
	}

	token, err := a.Get()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if token != "key-123" {
		t.Errorf("expected token 'key-123', got %s", token)
	}
}

func TestDefaultAuth(t *testing.T) {
	a := auth.NewDefaultAuth()

	header := a.Header()
	if header != "" {
		t.Errorf("expected empty header, got %s", header)
	}

	token, err := a.Get()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if token != "" {
		t.Errorf("expected empty token, got %s", token)
	}
}

func TestTokenRefresherAuth_Success(t *testing.T) {
	callCount := 0
	getToken := func() (string, time.Duration, error) {
		callCount++
		return "dynamic-token", 1 * time.Second, nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a := auth.NewTokenRefresherAuth(ctx, getToken)

	header := a.Header()
	if header != "Authorization" {
		t.Errorf("expected header Authorization, got %s", header)
	}

	token, err := a.Get()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if token != "Bearer dynamic-token" {
		t.Errorf("expected token 'Bearer dynamic-token', got %s", token)
	}

	if callCount == 0 {
		t.Errorf("expected getToken to be called at least once")
	}
}

func TestTokenRefresherAuth_Error(t *testing.T) {
	getToken := func() (string, time.Duration, error) {
		return "", 0, errors.New("failed to fetch token")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a := auth.NewTokenRefresherAuth(ctx, getToken)

	_, err := a.Get()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

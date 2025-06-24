package token_test

import (
	"context"
	"github.com/AdamShannag/umami-client/umami/token"
	"testing"
	"time"
)

func TestRefresher_GetAndRefresh(t *testing.T) {
	authorize := func() (string, time.Duration, error) {
		return "token1", 30 * time.Second, nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	refresher := token.NewRefresher(ctx, authorize)

	token1, err := refresher.Get()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token1 != "token1" {
		t.Errorf("expected token1, got %s", token1)
	}

	cancel()
}

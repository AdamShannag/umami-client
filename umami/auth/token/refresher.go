package token

import (
	"context"
	"time"
)

type tokenResponse struct {
	Token string
	Err   error
}

type Refresher struct {
	accessToken chan tokenResponse
	authorize   func() (string, time.Duration, error)
}

func NewRefresher(ctx context.Context, auth func() (string, time.Duration, error)) *Refresher {
	t := &Refresher{
		accessToken: make(chan tokenResponse),
		authorize:   auth,
	}
	go t.refresh(ctx)
	return t
}

func (t *Refresher) refresh(ctx context.Context) {
	token, expiresIn, err := t.authorize()
	expired := time.After(expiresIn - 10*time.Second)

	for {
		select {
		case t.accessToken <- tokenResponse{Token: token, Err: err}:
		case <-expired:
			token, expiresIn, err = t.authorize()
			if err != nil {
				expiresIn = 5 * time.Second
			}
			expired = time.After(expiresIn - 10*time.Second)
		case <-ctx.Done():
			return
		}
	}
}

func (t *Refresher) Get() (string, error) {
	res := <-t.accessToken
	return res.Token, res.Err
}

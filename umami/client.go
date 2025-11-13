package umami

import (
	"context"
	"net/http"
	"time"

	"github.com/AdamShannag/umami-client/umami/api"
	"github.com/AdamShannag/umami-client/umami/auth"
	"github.com/AdamShannag/umami-client/umami/request"
)

const defaultTokenExpiry = 24 * time.Hour

// Option represents a functional client option used during initialization.
type Option func(*client) error

// Client Umami API client.
type Client interface {
	// GetToken returns the current auth token and its remaining TTL.
	GetToken(string, string) (string, time.Duration, error)

	// User returns the User API interface.
	User() api.User

	// Team returns the Team API interface.
	Team() api.Team

	// Event returns the Event API interface.
	Event() api.Event

	// Session returns the Session API interface.
	Session() api.Session

	// Website returns the Website API interface.
	Website() api.Website

	// WebsiteStats returns the WebsiteStats API interface.
	WebsiteStats() api.WebsiteStats

	// Public returns the Public API interface.
	Public() api.Public

	// Report returns the Report API interface.
	Report() api.Report

	// Close shuts down background token refreshes.
	Close()
}

// client is the internal implementation for the Client interface.
type client struct {
	hostURL     string
	tokenExpiry time.Duration

	cancel     context.CancelFunc
	httpClient request.Client
}

func NewClient(hostURL string, opts ...Option) (Client, error) {
	c := &client{
		hostURL:     hostURL,
		tokenExpiry: defaultTokenExpiry,
		httpClient:  request.NewClient(),
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func WithApiKey(apiKey string) Option {
	return func(c *client) error {
		c.httpClient.WithAuth(auth.NewApiKeyAuth(apiKey))
		return nil
	}
}

func WithSingleToken(username, password string) Option {
	return func(c *client) error {
		getToken, _, err := c.GetToken(username, password)
		if err != nil {
			return err
		}

		c.httpClient.WithAuth(auth.NewSingleTokenAuth(getToken))
		return nil
	}
}

func WithToken(token string) Option {
	return func(c *client) error {
		c.httpClient.WithAuth(auth.NewSingleTokenAuth(token))
		return nil
	}
}

func WithTokenRefresh(username, password string) Option {
	return func(c *client) error {
		ctx, cancel := context.WithCancel(context.Background())
		c.cancel = cancel
		c.httpClient.WithAuth(auth.NewTokenRefresherAuth(ctx, func() (string, time.Duration, error) {
			return c.GetToken(username, password)
		}))
		return nil
	}
}

func WithTokenExpiry(d time.Duration) Option {
	return func(c *client) error {
		c.tokenExpiry = d
		return nil
	}
}

func WithHttpClient(httpClient *http.Client) Option {
	return func(c *client) error {
		c.httpClient.WithHttpClient(httpClient)
		return nil
	}
}

func (c *client) Close() {
	if c.cancel != nil {
		c.cancel()
	}
}

func (c *client) User() api.User {
	return c
}
func (c *client) Team() api.Team {
	return c
}
func (c *client) Event() api.Event {
	return c
}
func (c *client) Session() api.Session {
	return c
}
func (c *client) Website() api.Website {
	return c
}
func (c *client) WebsiteStats() api.WebsiteStats {
	return c
}
func (c *client) Public() api.Public { return c }
func (c *client) Report() api.Report { return c }

func (c *client) getRequest(ctx context.Context, endpoint string, query map[string]string, v any) error {
	return c.httpClient.Send(ctx, request.Request{
		Method:   http.MethodGet,
		Endpoint: endpoint,
		Headers:  nil,
		Query:    query,
		Payload:  nil,
	}, v)
}

func (c *client) postRequest(ctx context.Context, endpoint string, payload any, v any) error {
	return c.httpClient.Send(ctx, request.Request{
		Method:   http.MethodPost,
		Endpoint: endpoint,
		Headers:  nil,
		Query:    nil,
		Payload:  payload,
	}, v)
}

func (c *client) deleteRequest(ctx context.Context, endpoint string) error {
	return c.httpClient.Send(ctx, request.Request{
		Method:   http.MethodDelete,
		Endpoint: endpoint,
		Headers:  nil,
		Query:    nil,
		Payload:  nil,
	}, nil)
}

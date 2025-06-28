package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/AdamShannag/umami-client/umami/auth"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Method   string
	Endpoint string
	Headers  map[string]string
	Query    map[string]string
	Payload  any
	Public   bool
}

type Client interface {
	Send(ctx context.Context, req Request, res any) error
	WithHttpClient(*http.Client)
	WithAuth(auth.Auth)
}

type client struct {
	auth       auth.Auth
	httpClient *http.Client
}

func NewClient() Client {
	return &client{
		auth:       auth.NewDefaultAuth(),
		httpClient: http.DefaultClient,
	}
}

func (c *client) Send(ctx context.Context, req Request, res any) error {
	if req.Headers == nil {
		req.Headers = map[string]string{}
	}

	reqURL := c.withQuery(req.Endpoint, req.Query)
	body, err := c.marshalBody(req.Payload)
	if err != nil {
		return err
	}

	r, err := http.NewRequestWithContext(ctx, req.Method, reqURL, body)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	if req.Payload != nil {
		req.Headers["Content-Type"] = "application/json"
	}
	c.setHeaders(r, req.Headers)

	if !req.Public {
		key, authErr := c.auth.Get()
		if authErr != nil {
			return fmt.Errorf("auth: %w", authErr)
		}
		r.Header.Set(c.auth.Header(), key)
	}

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return c.decodeResponse(resp, res)
}

func (c *client) WithHttpClient(h *http.Client) {
	c.httpClient = h
}

func (c *client) WithAuth(a auth.Auth) {
	c.auth = a
}

func (c *client) setHeaders(req *http.Request, headers map[string]string) {
	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func (c *client) marshalBody(payload any) (io.Reader, error) {
	if payload == nil {
		return nil, nil
	}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal payload: %w", err)
	}
	return bytes.NewReader(bodyBytes), nil
}

func (c *client) decodeResponse(resp *http.Response, v any) error {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed [%d]: %s", resp.StatusCode, string(bodyBytes))
	}

	if v == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *client) withQuery(base string, params map[string]string) string {
	u, _ := url.Parse(base)
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

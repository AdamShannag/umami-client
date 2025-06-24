package umami

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func (c *client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.Clone(ctx)

	if c.apiKey != "" {
		req.Header.Set("x-umami-api-key", c.apiKey)
	} else {
		tkn, err := c.token.Get()
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+tkn)
	}

	httpClient := c.httpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return httpClient.Do(req)
}

func (c *client) getRequest(ctx context.Context, endpoint string, query map[string]string, v any) error {
	return c.sendRequest(ctx, http.MethodGet, endpoint, query, nil, v)
}

func (c *client) postRequest(ctx context.Context, endpoint string, payload any, v any) error {
	return c.sendRequest(ctx, http.MethodPost, endpoint, nil, payload, v)
}

func (c *client) deleteRequest(ctx context.Context, endpoint string) error {
	return c.sendRequest(ctx, http.MethodDelete, endpoint, nil, nil, nil)
}

func (c *client) sendRequest(ctx context.Context, method, endpoint string, query map[string]string, payload any, v any) error {
	reqURL := withQuery(endpoint, query)

	var body io.Reader
	if payload != nil {
		bodyBytes, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		body = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}

	return nil
}

func buildTimeQuery(start, end time.Time) map[string]string {
	return map[string]string{
		"startAt": fmt.Sprintf("%d", start.UnixMilli()),
		"endAt":   fmt.Sprintf("%d", end.UnixMilli()),
	}
}

func appendOptionalParams(dst map[string]string, src map[string]string) {
	for k, v := range src {
		if v != "" {
			dst[k] = v
		}
	}
}

func withQuery(base string, params map[string]string) string {
	u, _ := url.Parse(base)
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

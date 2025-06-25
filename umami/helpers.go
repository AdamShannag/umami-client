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

// do sends an HTTP request, optionally using public authentication
func (c *client) do(ctx context.Context, req *http.Request, isPublic bool) (*http.Response, error) {
	req = req.Clone(ctx)

	// Auth header
	if !isPublic {
		if c.apiKey != "" {
			req.Header.Set("x-umami-api-key", c.apiKey)
		} else {
			token, err := c.token.Get()
			if err != nil {
				return nil, fmt.Errorf("get token: %w", err)
			}
			req.Header.Set("Authorization", "Bearer "+token)
		}
	}

	httpClient := c.httpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return httpClient.Do(req)
}

// getRequest issues a GET request and decodes the response into v.
func (c *client) getRequest(ctx context.Context, endpoint string, query map[string]string, v any) error {
	return c.sendRequest(ctx, http.MethodGet, endpoint, query, nil, v)
}

// postRequest issues a POST request with payload and decodes the response into v.
func (c *client) postRequest(ctx context.Context, endpoint string, payload any, v any) error {
	return c.sendRequest(ctx, http.MethodPost, endpoint, nil, payload, v)
}

// deleteRequest sends a DELETE request without a response body.
func (c *client) deleteRequest(ctx context.Context, endpoint string) error {
	return c.sendRequest(ctx, http.MethodDelete, endpoint, nil, nil, nil)
}

// sendRequest handles all internal requests with proper encoding and decoding
func (c *client) sendRequest(ctx context.Context, method, endpoint string, query map[string]string, payload any, v any) error {
	reqURL := withQuery(endpoint, query)
	body, err := marshalBody(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.do(ctx, req, false)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return decodeResponse(resp, v)
}

// sendPublicRequest sends an unauthenticated request (e.g. for tracking)
func (c *client) sendPublicRequest(ctx context.Context, method, endpoint, userAgent string, query map[string]string, payload any, v any) error {
	reqURL := withQuery(endpoint, query)
	body, err := marshalBody(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return fmt.Errorf("create public request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.do(ctx, req, true)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return decodeResponse(resp, v)
}

// marshalBody encodes payload to JSON if not nil
func marshalBody(payload any) (io.Reader, error) {
	if payload == nil {
		return nil, nil
	}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal payload: %w", err)
	}
	return bytes.NewReader(bodyBytes), nil
}

// decodeResponse decodes JSON response or returns error on bad status
func decodeResponse(resp *http.Response, v any) error {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed [%d]: %s", resp.StatusCode, string(bodyBytes))
	}

	if v == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

// buildTimeQuery formats start and end times into a query map
func buildTimeQuery(start, end time.Time) map[string]string {
	return map[string]string{
		"startAt": fmt.Sprintf("%d", start.UnixMilli()),
		"endAt":   fmt.Sprintf("%d", end.UnixMilli()),
	}
}

// appendOptionalParams adds non-empty parameters from src to dst
func appendOptionalParams(dst, src map[string]string) {
	for k, v := range src {
		if v != "" {
			dst[k] = v
		}
	}
}

// withQuery adds query parameters to a base URL
func withQuery(base string, params map[string]string) string {
	u, _ := url.Parse(base)
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

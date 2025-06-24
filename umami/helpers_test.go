package umami

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestClient_do(t *testing.T) {
	c := &client{
		apiKey: "test-key",
		httpClient: &http.Client{Transport: &mockRoundTripper{
			fn: func(r *http.Request) *http.Response {
				assertEqual(t, r.Header.Get("x-umami-api-key"), "test-key")
				return mockJSONResp([]byte(`{}`))
			},
		}},
	}

	req, _ := http.NewRequest("GET", "https://example.com", nil)
	_, err := c.do(context.Background(), req)
	assertNil(t, err)
}

func TestClient_sendRequest_success(t *testing.T) {
	type response struct {
		Msg string `json:"msg"`
	}
	want := response{Msg: "ok"}
	body, _ := json.Marshal(want)

	c := &client{
		apiKey: "x",
		httpClient: &http.Client{Transport: &mockRoundTripper{
			fn: func(r *http.Request) *http.Response {
				assertEqual(t, r.Method, http.MethodGet)
				assertEqual(t, r.URL.Path, "/test")
				return mockJSONResp(body)
			},
		}},
	}

	var got response
	err := c.sendRequest(context.Background(), http.MethodGet, "https://example.com/test", nil, nil, &got)
	assertNil(t, err)
	assertEqual(t, got.Msg, "ok")
}

func TestClient_sendRequest_post_withPayload(t *testing.T) {
	type payload struct {
		Name string `json:"name"`
	}
	type response struct {
		Ok bool `json:"ok"`
	}
	p := payload{Name: "gopher"}
	want := response{Ok: true}
	body, _ := json.Marshal(want)

	c := &client{
		apiKey: "x",
		httpClient: &http.Client{Transport: &mockRoundTripper{
			fn: func(r *http.Request) *http.Response {
				assertEqual(t, r.Method, http.MethodPost)
				assertEqual(t, r.Header.Get("Content-Type"), "application/json")

				var in payload
				_ = json.NewDecoder(r.Body).Decode(&in)
				assertEqual(t, in.Name, "gopher")

				return mockJSONResp(body)
			},
		}},
	}

	var got response
	err := c.sendRequest(context.Background(), http.MethodPost, "https://example.com/test", nil, p, &got)
	assertNil(t, err)
	assertEqual(t, got.Ok, true)
}

func TestClient_sendRequest_errorStatus(t *testing.T) {
	c := &client{
		apiKey: "x",
		httpClient: &http.Client{Transport: &mockRoundTripper{
			fn: func(r *http.Request) *http.Response {
				return &http.Response{
					StatusCode: 500,
					Body:       io.NopCloser(bytes.NewReader([]byte("server error"))),
					Header:     make(http.Header),
				}
			},
		}},
	}

	err := c.sendRequest(context.Background(), http.MethodGet, "https://example.com/bad", nil, nil, nil)
	if err == nil || err.Error() != "unexpected status 500: server error" {
		t.Fatalf("expected error status, got %v", err)
	}
}

func TestClient_getRequest_delegatesToSendRequest(t *testing.T) {
	called := false
	c := &client{
		apiKey: "x",
		httpClient: &http.Client{Transport: &mockRoundTripper{
			fn: func(r *http.Request) *http.Response {
				called = true
				return mockJSONResp([]byte(`{}`))
			},
		}},
	}

	err := c.getRequest(context.Background(), "https://example.com", nil, nil)
	assertNil(t, err)
	assertEqual(t, called, true)
}

func TestWithQuery_mergesQueryParams(t *testing.T) {
	base := "https://example.com/track"
	params := map[string]string{
		"event": "click",
		"id":    "123",
	}

	full := withQuery(base, params)
	u, _ := url.Parse(full)
	assertEqual(t, u.Query().Get("event"), "click")
	assertEqual(t, u.Query().Get("id"), "123")
}

func TestBuildTimeQuery(t *testing.T) {
	start := time.Unix(1000, 0)
	end := time.Unix(2000, 0)
	q := buildTimeQuery(start, end)

	assertEqual(t, q["startAt"], "1000000")
	assertEqual(t, q["endAt"], "2000000")
}

func TestAppendOptionalParams(t *testing.T) {
	dst := map[string]string{"a": "1"}
	src := map[string]string{"b": "2", "c": ""}
	appendOptionalParams(dst, src)

	if _, ok := dst["b"]; !ok {
		t.Error("expected b to be set")
	}
	if _, ok := dst["c"]; ok {
		t.Error("expected c to be omitted")
	}
}

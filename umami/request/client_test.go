package request_test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/AdamShannag/umami-client/umami/request"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockAuth struct {
	header string
	token  string
	err    error
}

func (m mockAuth) Header() string       { return m.header }
func (m mockAuth) Get() (string, error) { return m.token, m.err }

type mockResponse struct {
	Message string `json:"message"`
}

type mockPayload struct {
	Data string `json:"data"`
}

func TestSend_Success(t *testing.T) {
	expected := mockResponse{Message: "hello"}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer token" {
			t.Errorf("expected Authorization header not found")
		}
		_ = r.ParseForm()
		if r.URL.Query().Get("foo") != "bar" {
			t.Errorf("missing or incorrect query param")
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := request.NewClient()
	client.WithAuth(mockAuth{header: "Authorization", token: "Bearer token", err: nil})
	client.WithHttpClient(server.Client())

	var res mockResponse
	err := client.Send(context.Background(), request.Request{
		Method:   http.MethodGet,
		Endpoint: server.URL + "?existing=1",
		Query:    map[string]string{"foo": "bar"},
		Public:   false,
	}, &res)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res.Message != expected.Message {
		t.Errorf("unexpected response: %+v", res)
	}
}

func TestSend_WithPayload(t *testing.T) {
	expected := mockResponse{Message: "created"}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body mockPayload
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("error decoding body: %v", err)
		}
		if body.Data != "test" {
			t.Errorf("unexpected payload: %+v", body)
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := request.NewClient()
	client.WithAuth(mockAuth{header: "X-Auth", token: "value", err: nil})
	client.WithHttpClient(server.Client())

	var res mockResponse
	err := client.Send(context.Background(), request.Request{
		Method:   http.MethodPost,
		Endpoint: server.URL,
		Payload:  mockPayload{Data: "test"},
	}, &res)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Message != expected.Message {
		t.Errorf("unexpected response: %+v", res)
	}
}

func TestSend_AuthError(t *testing.T) {
	client := request.NewClient()
	client.WithAuth(mockAuth{header: "Authorization", err: errors.New("auth failed")})

	err := client.Send(context.Background(), request.Request{
		Method:   http.MethodGet,
		Endpoint: "http://fake",
	}, nil)

	if err == nil || err.Error() != "auth: auth failed" {
		t.Errorf("expected auth error, got %v", err)
	}
}

func TestSend_FailureResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad request", http.StatusBadRequest)
	}))
	defer server.Close()

	client := request.NewClient()
	client.WithHttpClient(server.Client())
	client.WithAuth(mockAuth{header: "Authorization", token: "abc"})

	err := client.Send(context.Background(), request.Request{
		Method:   http.MethodGet,
		Endpoint: server.URL,
	}, nil)

	if err == nil || err.Error() != "request failed [400]: bad request\n" {
		t.Errorf("expected 400 error, got %v", err)
	}
}

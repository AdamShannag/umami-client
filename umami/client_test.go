package umami

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/AdamShannag/umami-client/umami/types"
	"io"
	"net/http"
	"testing"
	"time"
)

type mockRoundTripper struct {
	fn func(*http.Request) *http.Response
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.fn(req), nil
}

func newMockClient(fn func(*http.Request) *http.Response) Client {
	return NewClient("https://example.com",
		WithApiKey("test"),
		WithHttpClient(&http.Client{Transport: &mockRoundTripper{fn: fn}}),
	)
}

func mockJSONResp(body []byte) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(body)),
		Header:     make(http.Header),
	}
}

func assertNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestUser_CreateUser(t *testing.T) {
	expected := types.User{ID: "u1", Username: "user1"}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodPost || req.URL.Path != "/api/users" {
			t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.User().CreateUser(context.Background(), types.CreateUserRequest{Username: "user1"})
	assertNil(t, err)
	assertEqual(t, got.ID, expected.ID)
}

func TestUser_ListUsers(t *testing.T) {
	expected := types.Users{Data: []types.UserInfo{{ID: "u2"}}}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodGet || req.URL.Path != "/api/admin/users" {
			t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.User().ListUsers(context.Background())
	assertNil(t, err)
	assertEqual(t, got.Data[0].ID, "u2")
}

func TestUser_GetUser(t *testing.T) {
	expected := types.User{ID: "u3", Username: "user3"}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodGet || req.URL.Path != "/api/users/u3" {
			t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.User().GetUser(context.Background(), "u3")
	assertNil(t, err)
	assertEqual(t, got.Username, expected.Username)
}

func TestUser_UpdateUser(t *testing.T) {
	expected := types.User{ID: "u4", Username: "updated"}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodPost || req.URL.Path != "/api/users/u4" {
			t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.User().UpdateUser(context.Background(), "u4", types.UpdateUserRequest{Username: "updated"})
	assertNil(t, err)
	assertEqual(t, got.Username, expected.Username)
}

func TestUser_DeleteUser(t *testing.T) {
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodDelete || req.URL.Path != "/api/users/u5" {
			t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		return &http.Response{StatusCode: http.StatusNoContent, Body: io.NopCloser(bytes.NewReader(nil))}
	})

	err := mock.User().DeleteUser(context.Background(), "u5")
	assertNil(t, err)
}

func TestUser_GetUserWebsites(t *testing.T) {
	expected := types.UserWebsites{Data: []types.UserWebsite{{ID: "w1"}}}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.URL.Path != "/api/users/u1/websites" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.User().GetUserWebsites(context.Background(), "u1", types.ListQueryParams{})
	assertNil(t, err)
	assertEqual(t, got.Data[0].ID, "w1")
}

func TestUser_ListUserTeams(t *testing.T) {
	expected := types.UserTeams{Data: []types.UserTeam{{ID: "t1"}}}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.URL.Path != "/api/users/u1/teams" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.User().ListUserTeams(context.Background(), "u1", types.ListQueryParams{})
	assertNil(t, err)
	assertEqual(t, got.Data[0].ID, "t1")
}

func TestTeam_CreateTeam(t *testing.T) {
	expected := []types.Team{{ID: "t1", Name: "Team 1"}}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodPost || req.URL.Path != "/api/teams" {
			t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.Team().CreateTeam(context.Background(), types.CreateTeamRequest{Name: "Team 1"})
	assertNil(t, err)
	assertEqual(t, got[0].ID, expected[0].ID)
}

func TestTeam_JoinTeam(t *testing.T) {
	expected := []types.Team{{ID: "t2", Name: "Joined"}}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodPost || req.URL.Path != "/api/teams/join" {
			t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.Team().JoinTeam(context.Background(), types.JoinTeamRequest{AccessCode: "xyz"})
	assertNil(t, err)
	assertEqual(t, got[0].Name, expected[0].Name)
}

func TestTeam_GetTeam(t *testing.T) {
	expected := types.Team{ID: "t3", Name: "Team 3"}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodGet || req.URL.Path != "/api/teams/t3" {
			t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.Team().GetTeam(context.Background(), "t3")
	assertNil(t, err)
	assertEqual(t, got.Name, expected.Name)
}

func TestTeam_UpdateTeam(t *testing.T) {
	expected := types.Team{ID: "t4", Name: "Updated"}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodPost || req.URL.Path != "/api/teams/t4" {
			t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.Team().UpdateTeam(context.Background(), "t4", types.UpdateTeamRequest{})
	assertNil(t, err)
	assertEqual(t, got.Name, expected.Name)
}

func TestTeam_DeleteTeam(t *testing.T) {
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodDelete || req.URL.Path != "/api/teams/t5" {
			t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		return &http.Response{StatusCode: http.StatusNoContent, Body: io.NopCloser(bytes.NewReader(nil))}
	})

	err := mock.Team().DeleteTeam(context.Background(), "t5")
	assertNil(t, err)
}

func TestTeam_ListTeamUsers(t *testing.T) {
	expected := types.Users{Data: []types.UserInfo{{ID: "u1"}}}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.URL.Path != "/api/teams/t1/users" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.Team().ListTeamUsers(context.Background(), "t1", types.ListQueryParams{})
	assertNil(t, err)
	assertEqual(t, got.Data[0].ID, "u1")
}

func TestTeam_AddUser(t *testing.T) {
	expected := types.TeamUserInfo{UserID: "u2", Role: "member"}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodPost || req.URL.Path != "/api/teams/t1/users" {
			t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.Team().AddUser(context.Background(), "t1", types.AddUserRequest{UserID: "u2", Role: "member"})
	assertNil(t, err)
	assertEqual(t, got.UserID, expected.UserID)
}

func TestTeam_GetTeamUser(t *testing.T) {
	expected := types.TeamUserInfo{UserID: "u3", Role: "admin"}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.URL.Path != "/api/teams/t1/users/u3" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.Team().GetTeamUser(context.Background(), "t1", "u3")
	assertNil(t, err)
	assertEqual(t, got.Role, expected.Role)
}

func TestTeam_UpdateUserRole(t *testing.T) {
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodPost || req.URL.Path != "/api/teams/t1/users/u4" {
			t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(nil))}
	})

	err := mock.Team().UpdateUserRole(context.Background(), "t1", "u4", "viewer")
	assertNil(t, err)
}

func TestTeam_RemoveUser(t *testing.T) {
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodDelete || req.URL.Path != "/api/teams/t1/users/u5" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return &http.Response{StatusCode: http.StatusNoContent, Body: io.NopCloser(bytes.NewReader(nil))}
	})

	err := mock.Team().RemoveUser(context.Background(), "t1", "u5")
	assertNil(t, err)
}

func TestTeam_ListTeamWebsites(t *testing.T) {
	expected := types.TeamWebsites{Data: []types.TeamWebsiteInfo{{ID: "w1"}}}
	mock := newMockClient(func(req *http.Request) *http.Response {
		if req.URL.Path != "/api/teams/t1/websites" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		b, _ := json.Marshal(expected)
		return mockJSONResp(b)
	})

	got, err := mock.Team().ListTeamWebsites(context.Background(), "t1", types.ListQueryParams{})
	assertNil(t, err)
	assertEqual(t, got.Data[0].ID, "w1")
}

func TestEvent_TestListEvents(t *testing.T) {
	now := time.Now()
	expected := types.ListEventsResponse{
		Data: []types.EventDetail{
			{
				ID:        "event1",
				WebsiteID: "website1",
				CreatedAt: now,
				EventName: "click",
			},
		},
		Count:    1,
		Page:     1,
		PageSize: 10,
	}

	mockClient := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", req.Method)
		}
		if req.URL.Path != "/api/websites/website1/events" {
			t.Errorf("unexpected URL path: %s", req.URL.Path)
		}

		body, _ := json.Marshal(expected)
		return mockJSONResp(body)
	})

	got, err := mockClient.Event().ListEvents(context.Background(), "website1", types.ListEventsParams{
		StartAt:  now.Add(-1 * time.Hour),
		EndAt:    now,
		Page:     "1",
		PageSize: "10",
	})
	assertNil(t, err)
	if len(got.Data) != 1 {
		t.Errorf("expected 1 event, got %d", len(got.Data))
	}
	assertEqual(t, got.Data[0].ID, "event1")
}

func TestEvent_TestGetEventProperties(t *testing.T) {
	expected := []types.EventPropertyCount{
		{
			EventName:    "purchase",
			PropertyName: "price",
			DataType:     2,
			Total:        5,
		},
	}

	mockClient := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", req.Method)
		}
		if req.URL.Path != "/api/websites/website2/event-data/events" {
			t.Errorf("unexpected URL path: %s", req.URL.Path)
		}

		body, _ := json.Marshal(expected)
		return mockJSONResp(body)
	})

	got, err := mockClient.Event().GetEventProperties(context.Background(), "website2", types.EventDataQueryParams{
		StartAt: time.Now().Add(-24 * time.Hour),
		EndAt:   time.Now(),
	})
	assertNil(t, err)
	if len(got) != 1 {
		t.Errorf("expected 1 property count, got %d", len(got))
	}
	assertEqual(t, got[0].EventName, "purchase")
}

func TestEvent_TestGetEventFields(t *testing.T) {
	expected := []types.EventFieldsStat{
		{
			PropertyName: "category",
			DataType:     1,
			Value:        "books",
			Total:        10,
		},
	}

	mockClient := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", req.Method)
		}
		if req.URL.Path != "/api/websites/website3/event-data/fields" {
			t.Errorf("unexpected URL path: %s", req.URL.Path)
		}

		body, _ := json.Marshal(expected)
		return mockJSONResp(body)
	})

	got, err := mockClient.Event().GetEventFields(context.Background(), "website3", types.EventDataQueryParams{
		StartAt: time.Now().Add(-7 * 24 * time.Hour),
		EndAt:   time.Now(),
	})
	assertNil(t, err)
	if len(got) != 1 {
		t.Errorf("expected 1 field stat, got %d", len(got))
	}
	assertEqual(t, got[0].PropertyName, "category")
}

func TestEvent_TestGetEventValues(t *testing.T) {
	expected := []types.EventValueCount{
		{
			Value: "blue",
			Total: 7,
		},
	}

	mockClient := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", req.Method)
		}
		if req.URL.Path != "/api/websites/website4/event-data/values" {
			t.Errorf("unexpected URL path: %s", req.URL.Path)
		}

		body, _ := json.Marshal(expected)
		return mockJSONResp(body)
	})

	got, err := mockClient.Event().GetEventValues(context.Background(), "website4", types.EventDataQueryParams{
		StartAt:      time.Now().Add(-3 * 24 * time.Hour),
		EndAt:        time.Now(),
		EventName:    "color_change",
		PropertyName: "color",
	})
	assertNil(t, err)
	if len(got) != 1 {
		t.Errorf("expected 1 value count, got %d", len(got))
	}
	assertEqual(t, got[0].Value, "blue")
}

func TestEvent_TestGetEventDataStats(t *testing.T) {
	expected := types.EventDataStats{
		Events:     100,
		Properties: 10,
	}

	mockClient := newMockClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", req.Method)
		}
		if req.URL.Path != "/api/websites/website5/event-data/stats" {
			t.Errorf("unexpected URL path: %s", req.URL.Path)
		}

		body, _ := json.Marshal(expected)
		return mockJSONResp(body)
	})

	got, err := mockClient.Event().GetEventDataStats(context.Background(), "website5", types.EventDataQueryParams{
		StartAt: time.Now().Add(-30 * 24 * time.Hour),
		EndAt:   time.Now(),
	})
	assertNil(t, err)
	assertEqual(t, got.Events, 100)
	assertEqual(t, got.Properties, 10)
}

func TestEvent_TestListSessions(t *testing.T) {
	expected := types.ListSessionsResponse{
		Data: []types.Session{
			{ID: "abc123"},
		},
		Count:    1,
		Page:     1,
		PageSize: 50,
	}

	mockClient := newMockClient(func(req *http.Request) *http.Response {
		assertEqual(t, req.URL.Path, "/api/websites/site123/sessions")
		body, _ := json.Marshal(expected)
		return mockJSONResp(body)
	})

	params := types.ListSessionsParams{
		StartAt: time.Now().Add(-1 * time.Hour),
		EndAt:   time.Now(),
	}

	resp, err := mockClient.Session().ListSessions(context.Background(), "site123", params)
	assertNil(t, err)
	assertEqual(t, resp.Count, expected.Count)
	assertEqual(t, resp.Data[0].ID, "abc123")
}

func TestSession_TestGetSessionStats(t *testing.T) {
	expected := types.SessionStats{
		Pageviews: types.SessionMetric{Value: 123},
		Visitors:  types.SessionMetric{Value: 50},
		Visits:    types.SessionMetric{Value: 70},
		Countries: types.SessionMetric{Value: 10},
		Events:    types.SessionMetric{Value: 33},
	}

	mockClient := newMockClient(func(req *http.Request) *http.Response {
		assertEqual(t, req.URL.Path, "/api/websites/site123/sessions/stats")
		body, _ := json.Marshal(expected)
		return mockJSONResp(body)
	})

	params := types.SessionStatsParams{
		StartAt: time.Now().Add(-24 * time.Hour),
		EndAt:   time.Now(),
	}

	stats, err := mockClient.Session().ListSessionStats(context.Background(), "site123", params)
	assertNil(t, err)
	assertEqual(t, stats.Visitors.Value, 50)
	assertEqual(t, stats.Events.Value, 33)
	assertEqual(t, stats.Pageviews.Value, 123)
}

func TestSession_GetSessionDetails(t *testing.T) {
	want := types.SessionDetails{
		ID:        "abc123",
		WebsiteID: "site1",
		Country:   "US",
	}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodGet)
		assertEqual(t, r.URL.Path, "/api/websites/site1/sessions/abc123")
		return mockJSONResp(b)
	})

	got, err := c.Session().GetSessionDetails(context.Background(), "site1", "abc123")
	assertNil(t, err)
	assertEqual(t, got.ID, want.ID)
	assertEqual(t, got.Country, want.Country)
}

func TestSession_GetSessionActivity(t *testing.T) {
	want := []types.SessionActivityItem{{
		ID: "evt123",
	}}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodGet)
		assertEqual(t, r.URL.Path, "/api/websites/site1/sessions/abc123/activity")
		return mockJSONResp(b)
	})

	params := types.SessionDataValuesParams{
		StartAt:      time.Date(2023, 7, 1, 12, 0, 0, 0, time.UTC),
		EndAt:        time.Date(2023, 7, 1, 13, 0, 0, 0, time.UTC),
		PropertyName: "device",
	}
	got, err := c.Session().ListSessionActivities(context.Background(), "site1", "abc123", params)
	assertNil(t, err)
	assertEqual(t, len(got), 1)
}

func TestSession_GetSessionProperties(t *testing.T) {
	want := []types.SessionProperty{{
		SessionID:   "abc123",
		DataKey:     "device",
		StringValue: "mobile",
	}}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodGet)
		assertEqual(t, r.URL.Path, "/api/websites/site1/sessions/abc123/properties")
		return mockJSONResp(b)
	})

	got, err := c.Session().GetSessionProperties(context.Background(), "site1", "abc123")
	assertNil(t, err)
	assertEqual(t, got[0].DataKey, "device")
	assertEqual(t, got[0].StringValue, "mobile")
}

func TestSession_GetSessionDataProperties(t *testing.T) {
	want := []types.SessionDataPropertyCount{{
		PropertyName: "os",
		Total:        42,
	}}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodGet)
		assertEqual(t, r.URL.Path, "/api/websites/site1/session-data/properties")
		return mockJSONResp(b)
	})

	params := types.SessionDataPropertiesParams{
		StartAt: time.Date(2023, 7, 1, 12, 0, 0, 0, time.UTC),
		EndAt:   time.Date(2023, 7, 1, 13, 0, 0, 0, time.UTC),
	}
	got, err := c.Session().GetSessionDataProperties(context.Background(), "site1", params)
	assertNil(t, err)
	assertEqual(t, got[0].PropertyName, "os")
}

func TestSession_GetSessionDataValues(t *testing.T) {
	want := []types.SessionDataValueCount{{
		Value: "iOS",
		Total: 20,
	}}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodGet)
		assertEqual(t, r.URL.Path, "/api/websites/site1/session-data/values")
		return mockJSONResp(b)
	})

	params := types.SessionDataValuesParams{
		StartAt:      time.Date(2023, 7, 1, 12, 0, 0, 0, time.UTC),
		EndAt:        time.Date(2023, 7, 1, 13, 0, 0, 0, time.UTC),
		PropertyName: "os",
	}
	got, err := c.Session().GetSessionDataValues(context.Background(), "site1", params)
	assertNil(t, err)
	assertEqual(t, got[0].Value, "iOS")
}

func TestWebsite_ListWebsites(t *testing.T) {
	want := types.Websites{Data: []types.Website{{
		ID:   "site123",
		Name: "Example",
	}}}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodGet)
		assertEqual(t, r.URL.Path, "/api/websites")
		return mockJSONResp(b)
	})

	got, err := c.Website().ListWebsites(context.Background(), types.ListQueryParams{})
	assertNil(t, err)
	assertEqual(t, len(got.Data), 1)
	assertEqual(t, got.Data[0].ID, "site123")
}

func TestWebsite_CreateWebsite(t *testing.T) {
	want := types.TeamWebsiteInfo{
		ID:   "newsite",
		Name: "My Site",
	}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodPost)
		assertEqual(t, r.URL.Path, "/api/websites")
		return mockJSONResp(b)
	})

	req := types.CreateWebsiteRequest{
		Name:   "My Site",
		Domain: "example.com",
	}
	got, err := c.Website().CreateWebsite(context.Background(), req)
	assertNil(t, err)
	assertEqual(t, got.ID, "newsite")
	assertEqual(t, got.Name, "My Site")
}

func TestWebsite_GetWebsite(t *testing.T) {
	want := types.TeamWebsiteInfo{
		ID:   "site456",
		Name: "Another Site",
	}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodGet)
		assertEqual(t, r.URL.Path, "/api/websites/site456")
		return mockJSONResp(b)
	})

	got, err := c.Website().GetWebsite(context.Background(), "site456")
	assertNil(t, err)
	assertEqual(t, got.ID, "site456")
}

func TestWebsite_UpdateWebsite(t *testing.T) {
	want := types.TeamWebsiteInfo{
		ID:   "site789",
		Name: "Updated Site",
	}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodPost)
		assertEqual(t, r.URL.Path, "/api/websites/site789")
		return mockJSONResp(b)
	})

	req := types.UpdateWebsiteRequest{}
	got, err := c.Website().UpdateWebsite(context.Background(), "site789", req)
	assertNil(t, err)
	assertEqual(t, got.Name, "Updated Site")
}

func TestWebsite_DeleteWebsite(t *testing.T) {
	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodDelete)
		assertEqual(t, r.URL.Path, "/api/websites/site999")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       http.NoBody,
			Header:     make(http.Header),
		}
	})

	err := c.Website().DeleteWebsite(context.Background(), "site999")
	assertNil(t, err)
}

func TestWebsite_ResetWebsite(t *testing.T) {
	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodPost)
		assertEqual(t, r.URL.Path, "/api/websites/site999/reset")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       http.NoBody,
			Header:     make(http.Header),
		}
	})

	err := c.Website().ResetWebsite(context.Background(), "site999")
	assertNil(t, err)
}
func TestWebsiteStats_GetWebsiteActiveUsers(t *testing.T) {
	want := types.WebsiteActiveUsers{Visitors: 23}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodGet)
		assertEqual(t, r.URL.Path, "/api/websites/site123/active")
		return mockJSONResp(b)
	})

	got, err := c.WebsiteStats().GetWebsiteActiveUsers(context.Background(), "site123")
	assertNil(t, err)
	assertEqual(t, got.Visitors, int64(23))
}

func TestWebsiteStats_GetWebsiteEvents(t *testing.T) {
	want := types.WebsiteEvents{Data: []types.WebsiteEvent{{
		CreatedAt: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		EventType: 42,
	}}}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodGet)
		assertEqual(t, r.URL.Path, "/api/websites/site123/events")
		return mockJSONResp(b)
	})

	got, err := c.WebsiteStats().GetWebsiteEvents(context.Background(), "site123", types.WebsiteEventsQueryParams{})
	assertNil(t, err)
	assertEqual(t, len(got.Data), 1)
	assertEqual(t, got.Data[0].EventType, 42)
}

func TestWebsiteStats_GetWebsitePageViews(t *testing.T) {
	want := types.WebsitePageViews{
		Pageviews: []types.TimeSeriesDataPoint{{
			Timestamp:        types.CustomTime{Time: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)},
			NumberOfVisitors: 100,
		}},
		Sessions: []types.TimeSeriesDataPoint{{
			Timestamp:        types.CustomTime{Time: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)},
			NumberOfVisitors: 50,
		}},
	}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodGet)
		assertEqual(t, r.URL.Path, "/api/websites/site123/pageviews")
		return mockJSONResp(b)
	})

	got, err := c.WebsiteStats().GetWebsitePageViews(context.Background(), "site123", types.WebsitePageViewsQueryParams{})
	assertNil(t, err)
	assertEqual(t, len(got.Pageviews), 1)
	assertEqual(t, got.Pageviews[0].NumberOfVisitors, 100)
	assertEqual(t, len(got.Sessions), 1)
	assertEqual(t, got.Sessions[0].NumberOfVisitors, 50)
}

func TestWebsiteStats_GetWebsiteStats(t *testing.T) {
	want := types.WebsiteStats{
		Pageviews: types.Metric{Value: 300, Prev: 280},
		Visitors:  types.Metric{Value: 120, Prev: 100},
		Visits:    types.Metric{Value: 90, Prev: 70},
		Bounces:   types.Metric{Value: 30, Prev: 25},
		TotalTime: types.Metric{Value: 4500, Prev: 4000},
	}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodGet)
		assertEqual(t, r.URL.Path, "/api/websites/site123/stats")
		return mockJSONResp(b)
	})

	got, err := c.WebsiteStats().GetWebsiteStats(context.Background(), "site123", types.WebsiteStatsQueryParams{})
	assertNil(t, err)
	assertEqual(t, got.Pageviews.Value, int64(300))
	assertEqual(t, got.Bounces.Value, int64(30))
}

func TestWebsiteStats_GetWebsiteMetrics(t *testing.T) {
	want := []types.WebsiteMetric{{
		Value:            "Chrome",
		NumberOfVisitors: 999,
	}}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodGet)
		assertEqual(t, r.URL.Path, "/api/websites/site123/metrics")
		return mockJSONResp(b)
	})

	got, err := c.WebsiteStats().GetWebsiteMetrics(context.Background(), "site123", types.WebsiteMetricsQueryParams{Type: "browser"})
	assertNil(t, err)
	assertEqual(t, len(got), 1)
	assertEqual(t, got[0].Value, "Chrome")
	assertEqual(t, got[0].NumberOfVisitors, 999)
}

func TestClient_Send(t *testing.T) {
	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodPost)
		assertEqual(t, r.URL.Path, "/api/send")
		assertEqual(t, r.Header.Get("User-Agent"), "TestAgent")
		return mockJSONResp([]byte(`{}`))
	})

	err := c.Public().Send(context.Background(), "TestAgent", types.SendEventRequest{})
	assertNil(t, err)
}

func TestClient_GetInsights(t *testing.T) {
	want := []types.ReportInsight{
		{
			Views:     "100",
			Visitors:  80,
			Visits:    90,
			Bounces:   10,
			Totaltime: "300",
			URL:       "/home",
		},
	}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodPost)
		assertEqual(t, r.URL.Path, "/api/reports/insights")
		return mockJSONResp(b)
	})

	got, err := c.Report().GetInsights(context.Background(), types.ReportInsightsRequest{})
	assertNil(t, err)
	assertEqual(t, got[0].URL, "/home")
	assertEqual(t, got[0].Visitors, 80)
}

func TestClient_GetFunnel(t *testing.T) {
	want := []types.ReportFunnel{
		{
			Type:     "event",
			Value:    "signup",
			Visitors: 50,
			Dropped:  10,
		},
	}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodPost)
		assertEqual(t, r.URL.Path, "/api/reports/funnel")
		return mockJSONResp(b)
	})

	got, err := c.Report().GetFunnel(context.Background(), types.ReportFunnelRequest{})
	assertNil(t, err)
	assertEqual(t, got[0].Value, "signup")
	assertEqual(t, got[0].Visitors, 50)
}

func TestClient_GetRetention(t *testing.T) {
	want := []types.ReportRetention{
		{
			Date:           "2024-06-01",
			Day:            0,
			Visitors:       100,
			ReturnVisitors: 50,
			Percentage:     50,
		},
	}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodPost)
		assertEqual(t, r.URL.Path, "/api/reports/retention")
		return mockJSONResp(b)
	})

	got, err := c.Report().GetRetention(context.Background(), types.ReportRetentionRequest{})
	assertNil(t, err)
	assertEqual(t, got[0].ReturnVisitors, int64(50))
}

func TestClient_GetUTM(t *testing.T) {
	want := types.ReportUTM{
		"utm_source": {"google": 10},
	}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodPost)
		assertEqual(t, r.URL.Path, "/api/reports/utm")
		return mockJSONResp(b)
	})

	got, err := c.Report().GetUTM(context.Background(), types.ReportUTMRequest{})
	assertNil(t, err)
	assertEqual(t, got["utm_source"]["google"], 10)
}

func TestClient_GetGoals(t *testing.T) {
	want := []types.ReportGoal{{
		Type:   "signup",
		Goal:   100,
		Result: 80,
	}}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodPost)
		assertEqual(t, r.URL.Path, "/api/reports/goals")
		return mockJSONResp(b)
	})

	got, err := c.Report().GetGoals(context.Background(), types.ReportGoalsRequest{})
	assertNil(t, err)
	assertEqual(t, got[0].Type, "signup")
	assertEqual(t, got[0].Goal, 100)
}

func TestClient_GetJourney(t *testing.T) {
	item := "page1"
	want := []types.ReportJourney{{
		Items: []*string{&item},
		Count: 5,
	}}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodPost)
		assertEqual(t, r.URL.Path, "/api/reports/journey")
		return mockJSONResp(b)
	})

	got, err := c.Report().GetJourney(context.Background(), types.ReportJourneyRequest{})
	assertNil(t, err)
	assertEqual(t, *got[0].Items[0], "page1")
	assertEqual(t, got[0].Count, 5)
}

func TestClient_GetRevenue(t *testing.T) {
	want := types.ReportRevenue{
		Total: types.RevenueTotal{
			Sum:   1000.5,
			Count: 10,
		},
	}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodPost)
		assertEqual(t, r.URL.Path, "/api/reports/attribution")
		return mockJSONResp(b)
	})

	got, err := c.Report().GetRevenue(context.Background(), types.ReportRevenueRequest{})
	assertNil(t, err)
	assertEqual(t, got.Total.Sum, 1000.5)
	assertEqual(t, got.Total.Count, 10)
}

func TestClient_GetAttribution(t *testing.T) {
	want := types.ReportAttribution{
		Total: types.AttributionTotal{
			Visitors:  100,
			Visits:    200,
			Pageviews: 300,
		},
	}
	b, _ := json.Marshal(want)

	c := newMockClient(func(r *http.Request) *http.Response {
		assertEqual(t, r.Method, http.MethodPost)
		assertEqual(t, r.URL.Path, "/api/reports/attribution")
		return mockJSONResp(b)
	})

	got, err := c.Report().GetAttribution(context.Background(), types.ReportAttributionRequest{})
	assertNil(t, err)
	assertEqual(t, got.Total.Visitors, 100)
	assertEqual(t, got.Total.Visits, 200)
}

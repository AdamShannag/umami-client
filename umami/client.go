package umami

import (
	"context"
	"github.com/AdamShannag/umami-client/umami/auth"
	"github.com/AdamShannag/umami-client/umami/types"
	"log"
	"net/http"
	"time"
)

const defaultTokenExpiry = 24 * time.Hour

// User defines user-related operations.
//
// Covers:
//   - Creating, reading, updating, and deleting users
//   - Listing user-owned websites and teams
type User interface {
	// CreateUser creates a new user.
	//
	// POST /api/users
	CreateUser(ctx context.Context, req types.CreateUserRequest) (types.User, error)

	// ListUsers returns all users. Admin access is required.
	//
	// GET /api/admin/users
	ListUsers(ctx context.Context) (types.Users, error)

	// GetUser gets a user by ID.
	//
	// GET /api/users/:userId
	GetUser(ctx context.Context, userId string) (types.User, error)

	// UpdateUser updates a user
	//
	// POST /api/users/:userId
	UpdateUser(ctx context.Context, userId string, req types.UpdateUserRequest) (types.User, error)

	// DeleteUser deletes a user.
	//
	// DELETE /api/users/:userId
	DeleteUser(ctx context.Context, userId string) error

	// GetUserWebsites gets all websites that belong to a user.
	//
	// GET /api/users/:userId/websites
	GetUserWebsites(ctx context.Context, userId string, params types.ListQueryParams) (types.UserWebsites, error)

	// ListUserTeams gets all teams that belong to a user.
	//
	// GET /api/users/:userId/teams
	ListUserTeams(ctx context.Context, userId string, params types.ListQueryParams) (types.UserTeams, error)
}

// Team defines team-related operations.
//
// Covers:
//   - Creating and joining teams
//   - Managing users in a team
//   - Listing team users and websites
type Team interface {
	// CreateTeam creates a new team.
	//
	// POST /api/teams
	CreateTeam(ctx context.Context, req types.CreateTeamRequest) ([]types.Team, error)

	// JoinTeam joins a team using an access code.
	//
	// POST /api/teams/join
	JoinTeam(ctx context.Context, req types.JoinTeamRequest) ([]types.Team, error)

	// GetTeam fetches a specific team by ID.
	//
	// GET /api/teams/:teamId
	GetTeam(ctx context.Context, teamID string) (types.Team, error)

	// UpdateTeam updates team properties (name, accessCode).
	//
	// POST /api/teams/:teamId
	UpdateTeam(ctx context.Context, teamID string, req types.UpdateTeamRequest) (types.Team, error)

	// DeleteTeam deletes a team.
	//
	// DELETE /api/teams/:teamId
	DeleteTeam(ctx context.Context, teamID string) error

	// ListTeamUsers retrieves users in a team.
	//
	// GET /api/teams/:teamId/users
	ListTeamUsers(ctx context.Context, teamID string, params types.ListQueryParams) (types.TeamUsers, error)

	// AddUser adds a user to a team with a specific role.
	//
	// POST /api/teams/:teamId/users
	AddUser(ctx context.Context, teamID string, req types.AddUserRequest) (types.TeamUserInfo, error)

	// GetTeamUser gets a specific user's role on a team.
	//
	// GET /api/teams/:teamId/users/:userId
	GetTeamUser(ctx context.Context, teamID, userID string) (types.TeamUserInfo, error)

	// UpdateUserRole updates a user's role within a team.
	//
	// POST /api/teams/:teamId/users/:userId
	UpdateUserRole(ctx context.Context, teamID, userID string, role string) error

	// RemoveUser removes a user from a team.
	//
	// DELETE /api/teams/:teamId/users/:userId
	RemoveUser(ctx context.Context, teamID, userID string) error

	// ListTeamWebsites lists websites linked to a team.
	//
	// GET /api/teams/:teamId/websites
	ListTeamWebsites(ctx context.Context, teamID string, params types.ListQueryParams) (types.TeamWebsites, error)
}

// Event defines endpoints for working with event data.
//
// Covers:
//   - Listing events
//   - Aggregating and inspecting event properties and counts
type Event interface {
	// ListEvents Gets website event details within a given time range.
	//
	// GET /api/websites/:websiteId/events
	ListEvents(ctx context.Context, websiteId string, params types.ListEventsParams) (types.ListEventsResponse, error)

	// GetEventProperties Gets event data names, properties, and counts
	//
	// GET /api/websites/:websiteId/event-data/events
	GetEventProperties(ctx context.Context, websiteId string, params types.EventDataQueryParams) ([]types.EventPropertyCount, error)

	// GetEventFields Gets event data property and value counts within a given time range.
	//
	// GET /api/websites/:websiteId/events/fields
	GetEventFields(ctx context.Context, websiteId string, params types.EventDataQueryParams) ([]types.EventFieldsStat, error)

	// GetEventValues Gets event data counts for a given event and property
	//
	// GET /api/websites/:websiteId/events/values
	GetEventValues(ctx context.Context, websiteId string, params types.EventDataQueryParams) ([]types.EventValueCount, error)

	// GetEventDataStats Gets summarized website events, fields, and records within a given time range.
	//
	// GET /api/websites/:websiteId/events/stats
	GetEventDataStats(ctx context.Context, websiteId string, params types.EventDataQueryParams) (types.EventDataStats, error)
}

// Session defines session tracking endpoints.
//
// Covers:
//   - Listing sessions
//   - Session details and metadata
type Session interface {
	// ListSessions returns all sessions for a website.
	//
	// GET /api/websites/:websiteId/sessions
	ListSessions(ctx context.Context, websiteId string, params types.ListSessionsParams) (types.ListSessionsResponse, error)

	// ListSessionStats returns aggregate stats over sessions.
	//
	// GET /api/websites/:websiteId/sessions/stats
	ListSessionStats(ctx context.Context, websiteId string, params types.SessionStatsParams) (types.SessionStats, error)

	// GetSessionDetails fetches a specific session's metadata.
	//
	// GET /api/websites/:websiteId/sessions/:sessionId
	GetSessionDetails(ctx context.Context, websiteId, sessionId string) (types.SessionDetails, error)

	// ListSessionActivities returns all events and activity for a session.
	//
	// GET /api/websites/:websiteId/sessions/:sessionId/activity
	ListSessionActivities(ctx context.Context, websiteId, sessionId string, params types.SessionDataValuesParams) ([]types.SessionActivityItem, error)

	// GetSessionProperties returns session metadata key-value pairs.
	//
	// GET /api/websites/:websiteId/sessions/:sessionId/properties
	GetSessionProperties(ctx context.Context, websiteId, sessionId string) ([]types.SessionProperty, error)

	// GetSessionDataProperties returns global data property distributions.
	//
	// GET /api/websites/:websiteId/sessions/data/properties
	GetSessionDataProperties(ctx context.Context, websiteId string, params types.SessionDataPropertiesParams) ([]types.SessionDataPropertyCount, error)

	// GetSessionDataValues returns distribution of a session data property.
	//
	// GET /api/websites/:websiteId/sessions/data/values
	GetSessionDataValues(ctx context.Context, websiteId string, params types.SessionDataValuesParams) ([]types.SessionDataValueCount, error)
}

// Website defines website CRUD operations.
//
// Covers:
//   - Listing, creating, updating, deleting, resetting websites
type Website interface {
	// ListWebsites returns all accessible websites.
	//
	// GET /api/websites
	ListWebsites(ctx context.Context, query types.ListQueryParams) (types.Websites, error)

	// CreateWebsite creates a new website.
	//
	// POST /api/websites
	CreateWebsite(ctx context.Context, req types.CreateWebsiteRequest) (types.Website, error)

	// GetWebsite fetches a single website by ID.
	//
	// GET /api/websites/:websiteId
	GetWebsite(ctx context.Context, websiteId string) (types.Website, error)

	// UpdateWebsite updates a website.
	//
	// POST /api/websites/:websiteId
	UpdateWebsite(ctx context.Context, websiteId string, req types.UpdateWebsiteRequest) (types.Website, error)

	// DeleteWebsite deletes a website.
	//
	// DELETE /api/websites/:websiteId
	DeleteWebsite(ctx context.Context, websiteId string) error

	// ResetWebsite resets website analytics.
	//
	// POST /api/websites/:websiteId/reset
	ResetWebsite(ctx context.Context, websiteId string) error
}

// WebsiteStats provides metrics and data visualizations per website.
//
// Covers:
//   - Visitors, pageviews, sessions, metrics, and event trends
type WebsiteStats interface {
	// GetWebsiteActiveUsers returns number of active users in last 5 minutes.
	//
	// GET /api/websites/:websiteId/active
	GetWebsiteActiveUsers(ctx context.Context, websiteId string) (types.WebsiteActiveUsers, error)

	// GetWebsiteEvents returns event time series data.
	//
	// GET /api/websites/:websiteId/events/series
	GetWebsiteEvents(ctx context.Context, websiteId string, params types.WebsiteEventsQueryParams) (types.WebsiteEvents, error)

	// GetWebsitePageViews returns time series of pageviews/sessions.
	//
	// GET /api/websites/:websiteId/pageviews
	GetWebsitePageViews(ctx context.Context, websiteId string, params types.WebsitePageViewsQueryParams) (types.WebsitePageViews, error)

	// GetWebsiteStats returns summarized website statistics.
	//
	// GET /api/websites/:websiteId/stats
	GetWebsiteStats(ctx context.Context, websiteId string, params types.WebsiteStatsQueryParams) (types.WebsiteStats, error)

	// GetWebsiteMetrics returns grouped data for metric types.
	//
	// GET /api/websites/:websiteId/metrics
	GetWebsiteMetrics(ctx context.Context, websiteId string, params types.WebsiteMetricsQueryParams) ([]types.WebsiteMetric, error)
}

type Public interface {
	// Send register event in umami
	//
	// POST /api/send
	Send(ctx context.Context, userAgent string, payload types.SendEventRequest) error
}

// Option represents a functional client option used during initialization.
type Option func(*client)

// Client aggregates all functional interface groups and session lifecycle.
//
// Provides:
//   - Resource clients (User, Team, Event, etc.)
//   - Session lifecycle control and token access
type Client interface {
	// Close shuts down the client and any background token refreshes.
	Close()

	// GetToken returns the current auth token and its remaining TTL.
	GetToken(string, string) (string, time.Duration, error)

	// User returns the User API interface.
	User() User

	// Team returns the Team API interface.
	Team() Team

	// Event returns the Event API interface.
	Event() Event

	// Session returns the Session API interface.
	Session() Session

	// Website returns the Website API interface.
	Website() Website

	// WebsiteStats returns the WebsiteStats API interface.
	WebsiteStats() WebsiteStats

	// Public returns the Public API interface.
	Public() Public
}

// client is the internal implementation for the Client interface.
type client struct {
	hostURL     string
	tokenExpiry time.Duration

	auth       auth.Auth
	cancel     context.CancelFunc
	httpClient *http.Client
}

func NewClient(hostURL string, opts ...Option) Client {
	c := &client{
		hostURL:     hostURL,
		tokenExpiry: defaultTokenExpiry,
		auth:        auth.NewDefaultAuth(),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithApiKey(apiKey string) Option {
	return func(c *client) {
		c.auth = auth.NewApiKeyAuth(apiKey)
	}
}

func WithSingleToken(username, password string) Option {
	return func(c *client) {
		getToken, _, err := c.GetToken(username, password)
		if err != nil {
			log.Fatal(err)
		}

		c.auth = auth.NewSingleTokenAuth(getToken)
	}
}

func WithTokenRefresh(username, password string) Option {
	return func(c *client) {
		ctx, cancel := context.WithCancel(context.Background())
		c.cancel = cancel
		c.auth = auth.NewTokenRefresherAuth(ctx,
			func() (string, time.Duration, error) {
				return c.GetToken(username, password)
			})
	}
}

func WithTokenExpiry(d time.Duration) Option {
	return func(c *client) {
		c.tokenExpiry = d
	}
}

func WithHttpClient(httpClient *http.Client) Option {
	return func(c *client) {
		c.httpClient = httpClient
	}
}

func (c *client) Close() {
	if c.cancel != nil {
		c.cancel()
	}
}

func (c *client) User() User {
	return c
}

func (c *client) Team() Team {
	return c
}

func (c *client) Event() Event {
	return c
}

func (c *client) Session() Session {
	return c
}

func (c *client) Website() Website {
	return c
}

func (c *client) WebsiteStats() WebsiteStats {
	return c
}

func (c *client) Public() Public {
	return c
}

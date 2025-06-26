package api

import (
	"context"
	"github.com/AdamShannag/umami-client/umami/types"
)

// Public provides an endpoint to send events to Umami server.
type Public interface {
	// Send register event in umami
	//
	// POST /api/send
	Send(ctx context.Context, userAgent string, payload types.SendEventRequest) error
}

// User defines user-related operations.
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

// Report provides structured access to Umami's reporting endpoints.
type Report interface {
	// GetInsights dive deeper into your data by using segments and filters.
	//
	// POST /api/reports/insights
	GetInsights(ctx context.Context, payload types.ReportInsightsRequest) ([]types.ReportInsight, error)

	// GetFunnel understand the conversion and drop-off rate of users
	//
	// POST /api/reports/funnel
	GetFunnel(ctx context.Context, payload types.ReportFunnelRequest) ([]types.ReportFunnel, error)

	// GetRetention measure your website stickiness by tracking how often users return.
	//
	// POST /api/reports/retention
	GetRetention(ctx context.Context, payload types.ReportRetentionRequest) ([]types.ReportRetention, error)

	// GetUTM track your campaigns through UTM parameters.
	//
	// POST /api/reports/utm
	GetUTM(ctx context.Context, payload types.ReportUTMRequest) (types.ReportUTM, error)

	// GetGoals track your goals for pageviews and events.
	//
	// POST /api/reports/goals
	GetGoals(ctx context.Context, payload types.ReportGoalsRequest) ([]types.ReportGoal, error)

	// GetJourney understand how users navigate through your website.
	//
	// POST /api/reports/journey
	GetJourney(ctx context.Context, payload types.ReportJourneyRequest) ([]types.ReportJourney, error)

	// GetRevenue look into your revenue data and how users are spending.
	//
	// POST /api/reports/revenue
	GetRevenue(ctx context.Context, payload types.ReportRevenueRequest) (types.ReportRevenue, error)

	// GetAttribution see how users engage with your marketing and what drives conversions.
	//
	// POST /api/reports/attribution
	GetAttribution(ctx context.Context, payload types.ReportAttributionRequest) (types.ReportAttribution, error)
}

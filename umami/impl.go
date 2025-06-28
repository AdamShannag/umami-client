package umami

import (
	"context"
	"fmt"
	"github.com/AdamShannag/umami-client/umami/request"
	"github.com/AdamShannag/umami-client/umami/types"
	"net/http"
	"time"
)

func (c *client) GetToken(username, password string) (string, time.Duration, error) {
	var resp struct {
		Token string `json:"token"`
	}

	if err := c.httpClient.Send(context.Background(), request.Request{
		Method:   http.MethodPost,
		Endpoint: fmt.Sprintf("%s/api/auth/login", c.hostURL),
		Headers:  nil,
		Query:    nil,
		Payload: map[string]string{
			"username": username,
			"password": password,
		},
		Public: true,
	}, &resp); err != nil {
		return "", 0, err
	}

	return resp.Token, c.tokenExpiry, nil
}

func (c *client) CreateUser(ctx context.Context, req types.CreateUserRequest) (types.User, error) {
	var result types.User
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/users", c.hostURL), req, &result)
}

func (c *client) ListUsers(ctx context.Context) (types.Users, error) {
	var result types.Users
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/admin/users", c.hostURL), nil, &result)
}

func (c *client) GetUser(ctx context.Context, userId string) (types.User, error) {
	var result types.User
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/users/%s", c.hostURL, userId), nil, &result)
}

func (c *client) UpdateUser(ctx context.Context, userId string, req types.UpdateUserRequest) (types.User, error) {
	var result types.User
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/users/%s", c.hostURL, userId), req, &result)
}

func (c *client) DeleteUser(ctx context.Context, userId string) error {
	return c.deleteRequest(ctx, fmt.Sprintf("%s/api/users/%s", c.hostURL, userId))
}

func (c *client) GetUserWebsites(ctx context.Context, userId string, params types.ListQueryParams) (types.UserWebsites, error) {
	var result types.UserWebsites
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/users/%s/websites", c.hostURL, userId), params.ToQueryMap(), &result)
}

func (c *client) ListUserTeams(ctx context.Context, userId string, params types.ListQueryParams) (types.UserTeams, error) {
	var result types.UserTeams
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/users/%s/teams", c.hostURL, userId), params.ToQueryMap(), &result)
}

func (c *client) CreateTeam(ctx context.Context, req types.CreateTeamRequest) ([]types.Team, error) {
	var result []types.Team
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/teams", c.hostURL), req, &result)
}

func (c *client) JoinTeam(ctx context.Context, req types.JoinTeamRequest) ([]types.Team, error) {
	var result []types.Team
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/teams/join", c.hostURL), req, &result)
}

func (c *client) GetTeam(ctx context.Context, teamID string) (types.Team, error) {
	var result types.Team
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/teams/%s", c.hostURL, teamID), nil, &result)
}

func (c *client) UpdateTeam(ctx context.Context, teamID string, req types.UpdateTeamRequest) (types.Team, error) {
	var result types.Team
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/teams/%s", c.hostURL, teamID), req, &result)
}

func (c *client) DeleteTeam(ctx context.Context, teamID string) error {
	return c.deleteRequest(ctx, fmt.Sprintf("%s/api/teams/%s", c.hostURL, teamID))
}

func (c *client) ListTeamUsers(ctx context.Context, teamID string, params types.ListQueryParams) (types.TeamUsers, error) {
	var result types.TeamUsers
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/teams/%s/users", c.hostURL, teamID), params.ToQueryMap(), &result)
}

func (c *client) AddUser(ctx context.Context, teamID string, req types.AddUserRequest) (types.TeamUserInfo, error) {
	var result types.TeamUserInfo
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/teams/%s/users", c.hostURL, teamID), req, &result)
}

func (c *client) GetTeamUser(ctx context.Context, teamID, userID string) (types.TeamUserInfo, error) {
	var result types.TeamUserInfo
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/teams/%s/users/%s", c.hostURL, teamID, userID), nil, &result)
}

func (c *client) UpdateUserRole(ctx context.Context, teamID, userID string, role string) error {
	return c.postRequest(ctx, fmt.Sprintf("%s/api/teams/%s/users/%s", c.hostURL, teamID, userID), map[string]string{"role": role}, nil)
}

func (c *client) RemoveUser(ctx context.Context, teamID, userID string) error {
	return c.deleteRequest(ctx, fmt.Sprintf("%s/api/teams/%s/users/%s", c.hostURL, teamID, userID))
}

func (c *client) ListTeamWebsites(ctx context.Context, teamID string, params types.ListQueryParams) (types.TeamWebsites, error) {
	var result types.TeamWebsites
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/teams/%s/websites", c.hostURL, teamID), params.ToQueryMap(), &result)
}

func (c *client) ListEvents(ctx context.Context, websiteId string, params types.ListEventsParams) (types.ListEventsResponse, error) {
	var result types.ListEventsResponse
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/events", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) GetEventProperties(ctx context.Context, websiteId string, params types.EventDataQueryParams) ([]types.EventPropertyCount, error) {
	var result []types.EventPropertyCount
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/event-data/events", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) GetEventFields(ctx context.Context, websiteId string, params types.EventDataQueryParams) ([]types.EventFieldsStat, error) {
	var result []types.EventFieldsStat
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/event-data/fields", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) GetEventValues(ctx context.Context, websiteId string, params types.EventDataQueryParams) ([]types.EventValueCount, error) {
	var result []types.EventValueCount
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/event-data/values", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) GetEventDataStats(ctx context.Context, websiteId string, params types.EventDataQueryParams) (types.EventDataStats, error) {
	var result types.EventDataStats
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/event-data/stats", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) ListSessions(ctx context.Context, websiteId string, params types.ListSessionsParams) (types.ListSessionsResponse, error) {
	var result types.ListSessionsResponse
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/sessions", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) ListSessionStats(ctx context.Context, websiteId string, params types.SessionStatsParams) (types.SessionStats, error) {
	var result types.SessionStats
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/sessions/stats", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) GetSessionDetails(ctx context.Context, websiteId, sessionId string) (types.SessionDetails, error) {
	var result types.SessionDetails
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/sessions/%s", c.hostURL, websiteId, sessionId), nil, &result)
}

func (c *client) ListSessionActivities(ctx context.Context, websiteId, sessionId string, params types.SessionDataValuesParams) ([]types.SessionActivityItem, error) {
	var result []types.SessionActivityItem
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/sessions/%s/activity", c.hostURL, websiteId, sessionId), params.ToQueryMap(), &result)
}

func (c *client) GetSessionProperties(ctx context.Context, websiteId, sessionId string) ([]types.SessionProperty, error) {
	var result []types.SessionProperty
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/sessions/%s/properties", c.hostURL, websiteId, sessionId), nil, &result)
}

func (c *client) GetSessionDataProperties(ctx context.Context, websiteId string, params types.SessionDataPropertiesParams) ([]types.SessionDataPropertyCount, error) {
	var result []types.SessionDataPropertyCount
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/session-data/properties", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) GetSessionDataValues(ctx context.Context, websiteId string, params types.SessionDataValuesParams) ([]types.SessionDataValueCount, error) {
	var result []types.SessionDataValueCount
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/session-data/values", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) ListWebsites(ctx context.Context, params types.ListQueryParams) (types.Websites, error) {
	var result types.Websites
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites", c.hostURL), params.ToQueryMap(), &result)
}

func (c *client) CreateWebsite(ctx context.Context, reqBody types.CreateWebsiteRequest) (types.Website, error) {
	var result types.Website
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/websites", c.hostURL), reqBody, &result)
}

func (c *client) GetWebsite(ctx context.Context, websiteId string) (types.Website, error) {
	var result types.Website
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s", c.hostURL, websiteId), nil, &result)
}

func (c *client) UpdateWebsite(ctx context.Context, websiteId string, reqBody types.UpdateWebsiteRequest) (types.Website, error) {
	var result types.Website
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/websites/%s", c.hostURL, websiteId), reqBody, &result)
}

func (c *client) DeleteWebsite(ctx context.Context, websiteId string) error {
	return c.deleteRequest(ctx, fmt.Sprintf("%s/api/websites/%s", c.hostURL, websiteId))
}

func (c *client) ResetWebsite(ctx context.Context, websiteId string) error {
	return c.postRequest(ctx, fmt.Sprintf("%s/api/websites/%s/reset", c.hostURL, websiteId), nil, nil)
}

func (c *client) GetWebsiteActiveUsers(ctx context.Context, websiteId string) (types.WebsiteActiveUsers, error) {
	var result types.WebsiteActiveUsers
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/active", c.hostURL, websiteId), nil, &result)
}

func (c *client) GetWebsiteEvents(ctx context.Context, websiteId string, params types.WebsiteEventsQueryParams) (types.WebsiteEvents, error) {
	var result types.WebsiteEvents
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/events", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) GetWebsitePageViews(ctx context.Context, websiteId string, params types.WebsitePageViewsQueryParams) (types.WebsitePageViews, error) {
	var result types.WebsitePageViews
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/pageviews", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) GetWebsiteMetrics(ctx context.Context, websiteId string, params types.WebsiteMetricsQueryParams) ([]types.WebsiteMetric, error) {
	var result []types.WebsiteMetric
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/metrics", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) GetWebsiteStats(ctx context.Context, websiteId string, params types.WebsiteStatsQueryParams) (types.WebsiteStats, error) {
	var result types.WebsiteStats
	return result, c.getRequest(ctx, fmt.Sprintf("%s/api/websites/%s/stats", c.hostURL, websiteId), params.ToQueryMap(), &result)
}

func (c *client) Send(ctx context.Context, userAgent string, payload types.SendEventRequest) error {
	return c.httpClient.Send(ctx, request.Request{
		Method:   http.MethodPost,
		Endpoint: fmt.Sprintf("%s/api/send", c.hostURL),
		Headers: map[string]string{
			"User-Agent": userAgent,
		},
		Query:   nil,
		Payload: payload,
		Public:  true,
	}, nil)
}

func (c *client) GetInsights(ctx context.Context, payload types.ReportInsightsRequest) ([]types.ReportInsight, error) {
	var result []types.ReportInsight
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/reports/insights", c.hostURL), payload, &result)
}

func (c *client) GetFunnel(ctx context.Context, payload types.ReportFunnelRequest) ([]types.ReportFunnel, error) {
	var result []types.ReportFunnel
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/reports/funnel", c.hostURL), payload, &result)
}

func (c *client) GetRetention(ctx context.Context, payload types.ReportRetentionRequest) ([]types.ReportRetention, error) {
	var result []types.ReportRetention
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/reports/retention", c.hostURL), payload, &result)
}

func (c *client) GetUTM(ctx context.Context, payload types.ReportUTMRequest) (types.ReportUTM, error) {
	var result types.ReportUTM
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/reports/utm", c.hostURL), payload, &result)
}

func (c *client) GetGoals(ctx context.Context, payload types.ReportGoalsRequest) ([]types.ReportGoal, error) {
	var result []types.ReportGoal
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/reports/goals", c.hostURL), payload, &result)
}

func (c *client) GetJourney(ctx context.Context, payload types.ReportJourneyRequest) ([]types.ReportJourney, error) {
	var result []types.ReportJourney
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/reports/journey", c.hostURL), payload, &result)
}

func (c *client) GetRevenue(ctx context.Context, payload types.ReportRevenueRequest) (types.ReportRevenue, error) {
	var result types.ReportRevenue
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/reports/attribution", c.hostURL), payload, &result)
}

func (c *client) GetAttribution(ctx context.Context, payload types.ReportAttributionRequest) (types.ReportAttribution, error) {
	var result types.ReportAttribution
	return result, c.postRequest(ctx, fmt.Sprintf("%s/api/reports/attribution", c.hostURL), payload, &result)
}

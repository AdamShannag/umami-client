package umami

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/AdamShannag/umami-client/umami/types"
	"io"
	"net/http"
	"time"
)

func (c *client) GetToken() (string, time.Duration, error) {
	body, _ := json.Marshal(map[string]string{
		"username": c.username,
		"password": c.password,
	})

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/auth/login", c.hostURL), bytes.NewReader(body))
	if err != nil {
		return "", 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(resp.Body)
		return "", 0, fmt.Errorf("login failed: %d %s", resp.StatusCode, msg)
	}

	var authResp struct {
		Token string `json:"token"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return "", 0, err
	}

	return authResp.Token, c.tokenExpiry, nil
}

func (c *client) CreateUser(ctx context.Context, req types.CreateUserRequest) (types.User, error) {
	var result types.User
	endpoint := fmt.Sprintf("%s/api/users", c.hostURL)
	err := c.postRequest(ctx, endpoint, req, &result)
	return result, err
}

func (c *client) ListUsers(ctx context.Context) (types.Users, error) {
	var result types.Users
	endpoint := fmt.Sprintf("%s/api/admin/users", c.hostURL)
	err := c.getRequest(ctx, endpoint, nil, &result)
	return result, err
}

func (c *client) GetUser(ctx context.Context, userId string) (types.User, error) {
	var result types.User
	endpoint := fmt.Sprintf("%s/api/users/%s", c.hostURL, userId)
	err := c.getRequest(ctx, endpoint, nil, &result)
	return result, err
}

func (c *client) UpdateUser(ctx context.Context, userId string, req types.UpdateUserRequest) (types.User, error) {
	var result types.User
	endpoint := fmt.Sprintf("%s/api/users/%s", c.hostURL, userId)
	err := c.postRequest(ctx, endpoint, req, &result)
	return result, err
}

func (c *client) DeleteUser(ctx context.Context, userId string) error {
	endpoint := fmt.Sprintf("%s/api/users/%s", c.hostURL, userId)
	return c.deleteRequest(ctx, endpoint)
}

func (c *client) GetUserWebsites(ctx context.Context, userId string, params types.ListQueryParams) (types.UserWebsites, error) {
	var result types.UserWebsites
	endpoint := fmt.Sprintf("%s/api/users/%s/websites", c.hostURL, userId)

	q := map[string]string{}
	appendOptionalParams(q, map[string]string{
		"query":    params.Query,
		"page":     params.Page,
		"pageSize": params.PageSize,
		"orderBy":  params.OrderBy,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) ListUserTeams(ctx context.Context, userId string, params types.ListQueryParams) (types.UserTeams, error) {
	var result types.UserTeams
	endpoint := fmt.Sprintf("%s/api/users/%s/teams", c.hostURL, userId)

	q := map[string]string{}
	appendOptionalParams(q, map[string]string{
		"query":    params.Query,
		"page":     params.Page,
		"pageSize": params.PageSize,
		"orderBy":  params.OrderBy,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) CreateTeam(ctx context.Context, req types.CreateTeamRequest) ([]types.Team, error) {
	var result []types.Team
	endpoint := fmt.Sprintf("%s/api/teams", c.hostURL)

	err := c.postRequest(ctx, endpoint, req, &result)
	return result, err
}

func (c *client) JoinTeam(ctx context.Context, req types.JoinTeamRequest) ([]types.Team, error) {
	var result []types.Team
	endpoint := fmt.Sprintf("%s/api/teams/join", c.hostURL)

	err := c.postRequest(ctx, endpoint, req, &result)
	return result, err
}

func (c *client) GetTeam(ctx context.Context, teamID string) (types.Team, error) {
	var result types.Team
	endpoint := fmt.Sprintf("%s/api/teams/%s", c.hostURL, teamID)

	err := c.getRequest(ctx, endpoint, nil, &result)
	return result, err
}

func (c *client) UpdateTeam(ctx context.Context, teamID string, req types.UpdateTeamRequest) (types.Team, error) {
	var result types.Team
	endpoint := fmt.Sprintf("%s/api/teams/%s", c.hostURL, teamID)

	err := c.postRequest(ctx, endpoint, req, &result)
	return result, err
}

func (c *client) DeleteTeam(ctx context.Context, teamID string) error {
	endpoint := fmt.Sprintf("%s/api/teams/%s", c.hostURL, teamID)

	err := c.deleteRequest(ctx, endpoint)
	return err
}

func (c *client) ListTeamUsers(ctx context.Context, teamID string, params types.ListQueryParams) (types.TeamUsers, error) {
	var result types.TeamUsers
	endpoint := fmt.Sprintf("%s/api/teams/%s/users", c.hostURL, teamID)

	q := map[string]string{}
	appendOptionalParams(q, map[string]string{
		"query":    params.Query,
		"page":     params.Page,
		"pageSize": params.PageSize,
		"orderBy":  params.OrderBy,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) AddUser(ctx context.Context, teamID string, req types.AddUserRequest) (types.TeamUserInfo, error) {
	var result types.TeamUserInfo
	endpoint := fmt.Sprintf("%s/api/teams/%s/users", c.hostURL, teamID)

	err := c.postRequest(ctx, endpoint, req, &result)
	return result, err
}

func (c *client) GetTeamUser(ctx context.Context, teamID, userID string) (types.TeamUserInfo, error) {
	var result types.TeamUserInfo
	endpoint := fmt.Sprintf("%s/api/teams/%s/users/%s", c.hostURL, teamID, userID)

	err := c.getRequest(ctx, endpoint, nil, &result)
	return result, err
}

func (c *client) UpdateUserRole(ctx context.Context, teamID, userID string, role string) error {
	endpoint := fmt.Sprintf("%s/api/teams/%s/users/%s", c.hostURL, teamID, userID)

	payload := map[string]string{"role": role}

	err := c.postRequest(ctx, endpoint, payload, nil)
	return err
}

func (c *client) RemoveUser(ctx context.Context, teamID, userID string) error {
	endpoint := fmt.Sprintf("%s/api/teams/%s/users/%s", c.hostURL, teamID, userID)

	err := c.deleteRequest(ctx, endpoint)
	return err
}

func (c *client) ListTeamWebsites(ctx context.Context, teamID string, params types.ListQueryParams) (types.TeamWebsites, error) {
	var result types.TeamWebsites
	endpoint := fmt.Sprintf("%s/api/teams/%s/websites", c.hostURL, teamID)

	q := map[string]string{}
	appendOptionalParams(q, map[string]string{
		"query":    params.Query,
		"page":     params.Page,
		"pageSize": params.PageSize,
		"orderBy":  params.OrderBy,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) ListEvents(ctx context.Context, websiteId string, params types.ListEventsParams) (types.ListEventsResponse, error) {
	var result types.ListEventsResponse
	endpoint := fmt.Sprintf("%s/api/websites/%s/events", c.hostURL, websiteId)

	q := buildTimeQuery(params.StartAt, params.EndAt)
	appendOptionalParams(q, map[string]string{
		"query":    params.Query,
		"page":     params.Page,
		"pageSize": params.PageSize,
		"orderBy":  params.OrderBy,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) GetEventProperties(ctx context.Context, websiteId string, params types.EventDataQueryParams) ([]types.EventPropertyCount, error) {
	var result []types.EventPropertyCount
	endpoint := fmt.Sprintf("%s/api/websites/%s/event-data/events", c.hostURL, websiteId)

	q := buildTimeQuery(params.StartAt, params.EndAt)
	appendOptionalParams(q, map[string]string{
		"event": params.EventName,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) GetEventFields(ctx context.Context, websiteId string, params types.EventDataQueryParams) ([]types.EventFieldsStat, error) {
	var result []types.EventFieldsStat
	endpoint := fmt.Sprintf("%s/api/websites/%s/event-data/fields", c.hostURL, websiteId)

	q := buildTimeQuery(params.StartAt, params.EndAt)

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) GetEventValues(ctx context.Context, websiteId string, params types.EventDataQueryParams) ([]types.EventValueCount, error) {
	var result []types.EventValueCount
	endpoint := fmt.Sprintf("%s/api/websites/%s/event-data/values", c.hostURL, websiteId)

	q := buildTimeQuery(params.StartAt, params.EndAt)
	appendOptionalParams(q, map[string]string{
		"eventName":    params.EventName,
		"propertyName": params.PropertyName,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) GetEventDataStats(ctx context.Context, websiteId string, params types.EventDataQueryParams) (types.EventDataStats, error) {
	var result types.EventDataStats
	endpoint := fmt.Sprintf("%s/api/websites/%s/event-data/stats", c.hostURL, websiteId)

	q := buildTimeQuery(params.StartAt, params.EndAt)

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) ListSessions(ctx context.Context, websiteId string, params types.ListSessionsParams) (types.ListSessionsResponse, error) {
	var result types.ListSessionsResponse
	endpoint := fmt.Sprintf("%s/api/websites/%s/sessions", c.hostURL, websiteId)

	q := buildTimeQuery(params.StartAt, params.EndAt)
	appendOptionalParams(q, map[string]string{
		"page":     params.Page,
		"pageSize": params.PageSize,
		"query":    params.Query,
		"orderBy":  params.OrderBy,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) ListSessionStats(ctx context.Context, websiteId string, params types.SessionStatsParams) (types.SessionStats, error) {
	var result types.SessionStats
	endpoint := fmt.Sprintf("%s/api/websites/%s/sessions/stats", c.hostURL, websiteId)

	q := buildTimeQuery(params.StartAt, params.EndAt)
	appendOptionalParams(q, map[string]string{
		"url":      params.URL,
		"referrer": params.Referrer,
		"title":    params.Title,
		"query":    params.Query,
		"event":    params.Event,
		"host":     params.Host,
		"os":       params.OS,
		"browser":  params.Browser,
		"device":   params.Device,
		"country":  params.Country,
		"region":   params.Region,
		"city":     params.City,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) GetSessionDetails(ctx context.Context, websiteId, sessionId string) (types.SessionDetails, error) {
	var result types.SessionDetails
	endpoint := fmt.Sprintf("%s/api/websites/%s/sessions/%s", c.hostURL, websiteId, sessionId)

	err := c.getRequest(ctx, endpoint, nil, &result)
	return result, err
}

func (c *client) ListSessionActivities(ctx context.Context, websiteId, sessionId string, params types.SessionDataValuesParams) ([]types.SessionActivityItem, error) {
	var result []types.SessionActivityItem
	endpoint := fmt.Sprintf("%s/api/websites/%s/sessions/%s/activity", c.hostURL, websiteId, sessionId)

	q := buildTimeQuery(params.StartAt, params.EndAt)

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) GetSessionProperties(ctx context.Context, websiteId, sessionId string) ([]types.SessionProperty, error) {
	var result []types.SessionProperty
	endpoint := fmt.Sprintf("%s/api/websites/%s/sessions/%s/properties", c.hostURL, websiteId, sessionId)

	err := c.getRequest(ctx, endpoint, nil, &result)
	return result, err
}

func (c *client) GetSessionDataProperties(ctx context.Context, websiteId string, params types.SessionDataPropertiesParams) ([]types.SessionDataPropertyCount, error) {
	var result []types.SessionDataPropertyCount
	endpoint := fmt.Sprintf("%s/api/websites/%s/session-data/properties", c.hostURL, websiteId)

	q := buildTimeQuery(params.StartAt, params.EndAt)

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) GetSessionDataValues(ctx context.Context, websiteId string, params types.SessionDataValuesParams) ([]types.SessionDataValueCount, error) {
	var result []types.SessionDataValueCount
	endpoint := fmt.Sprintf("%s/api/websites/%s/session-data/values", c.hostURL, websiteId)

	q := buildTimeQuery(params.StartAt, params.EndAt)
	appendOptionalParams(q, map[string]string{
		"propertyName": params.PropertyName,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) ListWebsites(ctx context.Context, params types.ListQueryParams) (types.Websites, error) {
	var result types.Websites
	endpoint := fmt.Sprintf("%s/api/websites", c.hostURL)

	q := map[string]string{}
	appendOptionalParams(q, map[string]string{
		"query":    params.Query,
		"page":     params.Page,
		"pageSize": params.PageSize,
		"orderBy":  params.OrderBy,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) CreateWebsite(ctx context.Context, reqBody types.CreateWebsiteRequest) (types.Website, error) {
	var result types.Website
	endpoint := fmt.Sprintf("%s/api/websites", c.hostURL)

	err := c.postRequest(ctx, endpoint, reqBody, &result)
	return result, err
}

func (c *client) GetWebsite(ctx context.Context, websiteId string) (types.Website, error) {
	var result types.Website
	endpoint := fmt.Sprintf("%s/api/websites/%s", c.hostURL, websiteId)

	err := c.getRequest(ctx, endpoint, nil, &result)
	return result, err
}

func (c *client) UpdateWebsite(ctx context.Context, websiteId string, reqBody types.UpdateWebsiteRequest) (types.Website, error) {
	var result types.Website
	endpoint := fmt.Sprintf("%s/api/websites/%s", c.hostURL, websiteId)

	err := c.postRequest(ctx, endpoint, reqBody, &result)
	return result, err
}

func (c *client) DeleteWebsite(ctx context.Context, websiteId string) error {
	endpoint := fmt.Sprintf("%s/api/websites/%s", c.hostURL, websiteId)

	err := c.deleteRequest(ctx, endpoint)
	return err
}

func (c *client) ResetWebsite(ctx context.Context, websiteId string) error {
	endpoint := fmt.Sprintf("%s/api/websites/%s/reset", c.hostURL, websiteId)

	err := c.postRequest(ctx, endpoint, nil, nil)
	return err
}

func (c *client) GetWebsiteActiveUsers(ctx context.Context, websiteId string) (types.WebsiteActiveUsers, error) {
	var result types.WebsiteActiveUsers
	endpoint := fmt.Sprintf("%s/api/websites/%s/active", c.hostURL, websiteId)
	err := c.getRequest(ctx, endpoint, nil, &result)
	return result, err
}

func (c *client) GetWebsiteEvents(ctx context.Context, websiteId string, params types.WebsiteEventsQueryParams) (types.WebsiteEvents, error) {
	var result types.WebsiteEvents
	endpoint := fmt.Sprintf("%s/api/websites/%s/events", c.hostURL, websiteId)

	q := buildTimeQuery(params.StartAt, params.EndAt)
	q["unit"] = params.Unit
	q["timezone"] = params.Timezone
	appendOptionalParams(q, map[string]string{
		"url":      params.URL,
		"referrer": params.Referrer,
		"title":    params.Title,
		"host":     params.Host,
		"os":       params.OS,
		"browser":  params.Browser,
		"device":   params.Device,
		"country":  params.Country,
		"region":   params.Region,
		"city":     params.City,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) GetWebsitePageViews(ctx context.Context, websiteId string, params types.WebsitePageViewsQueryParams) (types.WebsitePageViews, error) {
	var result types.WebsitePageViews
	endpoint := fmt.Sprintf("%s/api/websites/%s/pageviews", c.hostURL, websiteId)

	q := buildTimeQuery(params.StartAt, params.EndAt)
	q["unit"] = params.Unit
	q["timezone"] = params.Timezone
	appendOptionalParams(q, map[string]string{
		"url":      params.URL,
		"referrer": params.Referrer,
		"title":    params.Title,
		"host":     params.Host,
		"os":       params.OS,
		"browser":  params.Browser,
		"device":   params.Device,
		"country":  params.Country,
		"region":   params.Region,
		"city":     params.City,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) GetWebsiteMetrics(ctx context.Context, websiteId string, params types.WebsiteMetricsQueryParams) ([]types.WebsiteMetric, error) {
	var result []types.WebsiteMetric
	endpoint := fmt.Sprintf("%s/api/websites/%s/metrics", c.hostURL, websiteId)

	q := buildTimeQuery(params.StartAt, params.EndAt)
	q["type"] = params.Type
	if params.Limit > 0 {
		q["limit"] = fmt.Sprintf("%d", params.Limit)
	}
	appendOptionalParams(q, map[string]string{
		"url":      params.URL,
		"referrer": params.Referrer,
		"title":    params.Title,
		"query":    params.Query,
		"host":     params.Host,
		"os":       params.OS,
		"browser":  params.Browser,
		"device":   params.Device,
		"country":  params.Country,
		"region":   params.Region,
		"city":     params.City,
		"language": params.Language,
		"event":    params.Event,
	})

	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) GetWebsiteStats(ctx context.Context, websiteId string, params types.WebsiteStatsQueryParams) (types.WebsiteStats, error) {
	var result types.WebsiteStats

	q := buildTimeQuery(params.StartAt, params.EndAt)
	appendOptionalParams(q, map[string]string{
		"url":      params.URL,
		"referrer": params.Referrer,
		"title":    params.Title,
		"query":    params.Query,
		"event":    params.Event,
		"host":     params.Host,
		"os":       params.OS,
		"browser":  params.Browser,
		"device":   params.Device,
		"country":  params.Country,
		"region":   params.Region,
		"city":     params.City,
	})

	endpoint := fmt.Sprintf("%s/api/websites/%s/stats", c.hostURL, websiteId)
	err := c.getRequest(ctx, endpoint, q, &result)
	return result, err
}

func (c *client) Send(ctx context.Context, userAgent string, payload types.SendEventRequest) error {
	endpoint := fmt.Sprintf("%s/api/send", c.hostURL)
	return c.sendPublicRequest(ctx, http.MethodPost, endpoint, userAgent, nil, payload, nil)
}

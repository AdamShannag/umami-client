package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AdamShannag/umami-client/umami"
	"github.com/AdamShannag/umami-client/umami/types"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	client := umami.NewClient("https://umami.host.net", umami.WithTokenAuth("admin", "umami"))
	defer client.Close()

	ctx := context.Background()
	websiteID := "website-uuid"

	go users(ctx, client)
	go teams(ctx, client)
	go events(ctx, client, websiteID)
	go sessions(ctx, client, websiteID)
	go websites(ctx, client, websiteID)
	go websiteStats(ctx, client, websiteID)

	fmt.Println("Running examples... Press Ctrl+C to exit.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("\nShutting down.")
}

/*
Users

POST /api/users
GET /api/admin/users
POST /api/users/:userId
GET /api/users/:userId
DELETE /api/users/:userId
GET /api/users/:userId/websites
GET /api/users/:userId/teams
*/
func users(ctx context.Context, client umami.Client) {
	// POST /api/users
	createUser, err := client.User().CreateUser(ctx, types.CreateUserRequest{
		Username: "test-user",
		Password: "test-password",
		Role:     "User",
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Created User", createUser)

	// GET /api/admin/users
	usrs, err := client.User().ListUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Users", usrs)

	// GET /api/users/:userId
	user, err := client.User().GetUser(ctx, createUser.ID)
	if err != nil {
		log.Fatal(err)
	}
	logStruct("User", user)

	// POST /api/users/:userId
	updateUser, err := client.User().UpdateUser(ctx, user.ID, types.UpdateUserRequest{
		Username: "new-username",
		Password: "new-password",
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Updated User", updateUser)

	// GET /api/users/:userId/websites
	userWebsites, err := client.User().GetUserWebsites(ctx, updateUser.ID, types.ListQueryParams{})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("User Websites", userWebsites)

	// GET /api/users/:userId/teams
	userTeams, err := client.User().ListUserTeams(ctx, updateUser.ID, types.ListQueryParams{})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("User Teams", userTeams)

	// DELETE /api/users/:userId
	err = client.User().DeleteUser(ctx, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("User Deleted: %+v", updateUser.ID)
}

/*
Teams

POST /api/teams
POST /api/teams/join
POST /api/teams/:teamId
GET /api/teams/:teamId
DELETE /api/teams/:teamId
POST /api/teams/:teamId/users
GET /api/teams/:teamId/users
GET /api/teams/:teamId/users/:userId
POST /api/teams/:teamId/users/:userId
DELETE /api/teams/:teamId/users/:userId
GET /api/teams/:teamId/websites
*/
func teams(ctx context.Context, client umami.Client) {
	// POST /api/teams
	createTeam, err := client.Team().CreateTeam(ctx, types.CreateTeamRequest{Name: "Test Team"})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Created Team", createTeam)

	// POST /api/websites
	createWebsite, err := client.Website().CreateWebsite(ctx, types.CreateWebsiteRequest{
		Domain: "test.com",
		Name:   "Test",
		TeamID: &createTeam[0].ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Created Website", createWebsite)

	// GET /api/teams/:teamId
	team, err := client.Team().GetTeam(ctx, createTeam[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Get Team", team)

	// POST /api/teams/:teamId
	name := "new team"
	updateTeam, err := client.Team().UpdateTeam(ctx, team.ID, types.UpdateTeamRequest{
		Name: &name,
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Updated Team", updateTeam)

	// GET /api/teams/:teamId/users
	teamUsers, err := client.Team().ListTeamUsers(ctx, updateTeam.ID, types.ListQueryParams{})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Team Users", teamUsers)

	// GET /api/teams/:teamId/websites
	teamWebsites, err := client.Team().ListTeamWebsites(ctx, updateTeam.ID, types.ListQueryParams{})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Team Websites", teamWebsites)

	// DELETE /api/websites/:websiteId
	err = client.Website().DeleteWebsite(ctx, createWebsite.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Website Deleted: %s", createWebsite.ID)

	// DELETE /api/teams/:teamId
	err = client.Team().DeleteTeam(ctx, updateTeam.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Team Deleted: %s", updateTeam.ID)
}

/*
Events

GET /api/websites/:websiteId/events
GET /api/websites/:websiteId/event-data/events
GET /api/websites/:websiteId/event-data/fields
GET /api/websites/:websiteId/event-data/values
GET /api/websites/:websiteId/event-data/stats
*/
func events(ctx context.Context, client umami.Client, websiteID string) {
	start := time.Now().Add(-24 * time.Hour)
	end := time.Now()

	// GET /api/websites/:websiteId/events
	websiteEvents, err := client.Event().ListEvents(ctx, websiteID, types.ListEventsParams{
		StartAt: start,
		EndAt:   end,
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Website Events", websiteEvents)

	// GET /api/websites/:websiteId/event-data/events
	websiteEventProperties, err := client.Event().GetEventProperties(ctx, websiteID, types.EventDataQueryParams{
		StartAt: start,
		EndAt:   end,
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Website Event Properties", websiteEventProperties)

	// GET /api/websites/:websiteId/event-data/fields
	websiteEventFields, err := client.Event().GetEventFields(ctx, websiteID, types.EventDataQueryParams{
		StartAt: start,
		EndAt:   end,
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Website Event Fields", websiteEventFields)

	// GET /api/websites/:websiteId/event-data/values
	websiteEventValues, err := client.Event().GetEventValues(ctx, websiteID, types.EventDataQueryParams{
		StartAt:      start,
		EndAt:        end,
		EventName:    "test",
		PropertyName: "test",
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Website Event Values", websiteEventValues)

	// GET /api/websites/:websiteId/event-data/stats
	websiteEventStats, err := client.Event().GetEventDataStats(ctx, websiteID, types.EventDataQueryParams{
		StartAt: start,
		EndAt:   end,
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Website Event Stats", websiteEventStats)
}

/*
Sessions

GET /api/websites/:websiteId/sessions
GET /api/websites/:websiteId/sessions/stats
GET /api/websites/:websiteId/sessions/:sessionId
GET /api/websites/:websiteId/sessions/:sessionId/activity
GET /api/websites/:websiteId/sessions/:sessionId/properties
GET /api/websites/:websiteId/session-data/properties
GET /api/websites/:websiteId/session-data/values
*/
func sessions(ctx context.Context, client umami.Client, websiteID string) {
	start := time.Now().Add(-24 * time.Hour)
	end := time.Now()

	// GET /api/websites/:websiteId/sessions
	sessionDetails, err := client.Session().ListSessions(ctx, websiteID, types.ListSessionsParams{
		StartAt: start,
		EndAt:   end,
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Session Details", sessionDetails)

	// GET /api/websites/:websiteId/sessions/stats
	sessionStats, err := client.Session().ListSessionStats(ctx, websiteID, types.SessionStatsParams{
		StartAt: start,
		EndAt:   end,
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Session Stats", sessionStats)

	if len(sessionDetails.Data) == 0 {
		log.Println("No sessions available to inspect")
		return
	}
	sessionID := sessionDetails.Data[0].ID

	// GET /api/websites/:websiteId/sessions/:sessionId
	sessionStat, err := client.Session().GetSessionDetails(ctx, websiteID, sessionID)
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Session Details by ID", sessionStat)

	// GET /api/websites/:websiteId/sessions/:sessionId/activity
	sessionActivities, err := client.Session().ListSessionActivities(ctx, websiteID, sessionID, types.SessionDataValuesParams{
		StartAt: start,
		EndAt:   end,
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Session Activities", sessionActivities)

	// GET /api/websites/:websiteId/sessions/:sessionId/properties
	sessionProperties, err := client.Session().GetSessionProperties(ctx, websiteID, sessionID)
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Session Properties", sessionProperties)

	// GET /api/websites/:websiteId/session-data/properties
	sessionDataProperties, err := client.Session().GetSessionDataProperties(ctx, websiteID, types.SessionDataPropertiesParams{
		StartAt: start,
		EndAt:   end,
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Session Data Properties", sessionDataProperties)

	// GET /api/websites/:websiteId/session-data/values
	sessionDataValues, err := client.Session().GetSessionDataValues(ctx, websiteID, types.SessionDataValuesParams{
		StartAt:      start,
		EndAt:        end,
		PropertyName: "city",
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Session Data Values", sessionDataValues)
}

/*
Websites

GET /api/websites
POST /api/websites
GET /api/websites/:websiteId
POST /api/websites/:websiteId
POST /api/websites/:websiteId/reset
DELETE /api/websites/:websiteId
*/
func websites(ctx context.Context, client umami.Client, websiteID string) {
	// GET /api/websites
	sites, err := client.Website().ListWebsites(ctx, types.ListQueryParams{})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Websites", sites)

	// POST /api/websites
	createWebsite, err := client.Website().CreateWebsite(ctx, types.CreateWebsiteRequest{
		Domain: "test.com",
		Name:   "Test",
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Created Website", createWebsite)

	// GET /api/websites/:websiteId
	website, err := client.Website().GetWebsite(ctx, createWebsite.ID)
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Website", website)

	// POST /api/websites/:websiteId
	updatedWebsite, err := client.Website().UpdateWebsite(ctx, website.ID, types.UpdateWebsiteRequest{
		Name:   "Updated Test",
		Domain: "updated-test.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Updated Website", updatedWebsite)

	// POST /api/websites/:websiteId/reset
	err = client.Website().ResetWebsite(ctx, updatedWebsite.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Website Reset: %s", updatedWebsite.ID)

	// DELETE /api/websites/:websiteId
	err = client.Website().DeleteWebsite(ctx, updatedWebsite.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Website Deleted: %s", updatedWebsite.ID)
}

/*
Website stats

GET /api/websites/:websiteId/active
GET /api/websites/:websiteId/events
GET /api/websites/:websiteId/pageviews
GET /api/websites/:websiteId/metrics
GET /api/websites/:websiteId/stats
*/
func websiteStats(ctx context.Context, client umami.Client, websiteID string) {
	// GET /api/websites/:websiteId/active
	activeUsers, err := client.WebsiteStats().GetWebsiteActiveUsers(ctx, websiteID)
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Active Users", activeUsers)

	// GET /api/websites/:websiteId/events/series
	evts, err := client.WebsiteStats().GetWebsiteEvents(ctx, websiteID, types.WebsiteEventsQueryParams{
		StartAt:  time.Now().Add(-24 * time.Hour),
		EndAt:    time.Now(),
		Unit:     "month",
		Timezone: "America/Los_Angeles",
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Events", evts)

	// GET /api/websites/:websiteId/pageviews
	pageViews, err := client.WebsiteStats().GetWebsitePageViews(ctx, websiteID, types.WebsitePageViewsQueryParams{
		StartAt:  time.Now().Add(-24 * time.Hour),
		EndAt:    time.Now(),
		Unit:     "month",
		Timezone: "America/Los_Angeles",
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Page Views", pageViews)

	// GET /api/websites/:websiteId/stats
	stats, err := client.WebsiteStats().GetWebsiteStats(ctx, websiteID, types.WebsiteStatsQueryParams{
		StartAt: time.Now().Add(-24 * time.Hour),
		EndAt:   time.Now(),
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Website Stats", stats)

	// GET /api/websites/:websiteId/metrics
	websiteMetrics, err := client.WebsiteStats().GetWebsiteMetrics(ctx, websiteID, types.WebsiteMetricsQueryParams{
		StartAt: time.Now().Add(-24 * time.Hour),
		EndAt:   time.Now(),
		Type:    "url",
	})
	if err != nil {
		log.Fatal(err)
	}
	logStruct("Website Metrics", websiteMetrics)
}

func logStruct(label string, v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("%s: (error marshaling) %v", label, err)
		return
	}
	log.Printf("%s:\n%s", label, data)
}

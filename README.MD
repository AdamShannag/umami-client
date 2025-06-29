# umami-client

**A fully typed Go api client for the [Umami](https://umami.is) API.**

## Features

- Strongly typed request and response models
- Works with API key or Token Auth (with single and refresh token)
- Simple, flexible configuration
- Zero dependencies

## Installation

```bash
go get -u github.com/AdamShannag/umami-client@v1.2.1
````

## Example Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/AdamShannag/umami-client/umami"
	"github.com/AdamShannag/umami-client/umami/types"
	"log"
)

func main() {
	ctx := context.Background()

	client := umami.NewClient("https://umami.instance.com",
		umami.WithSingleToken("admin", "password123"),
	)

	websites, err := client.Website().ListWebsites(ctx, types.ListQueryParams{})
	if err != nil {
		log.Fatal(err)
	}

	for _, site := range websites.Data {
		fmt.Println(site.Name)
	}
}

```

### More Examples

Check out the [`example`](./example) directory for comprehensive usage examples of this SDK.

* The example covers all major Umami API endpoints.
* the example runs multiple goroutines which interact with the client concurrently.

## Authentication & Configuration

You can authenticate using either an **API key** or **Token** (username/password).

| Option                         | Description                                           |
|--------------------------------|-------------------------------------------------------|
| `WithApiKey(apiKey)`           | Authenticate using a static API key                   |
| `WithSingleToken(user, pass)`  | Authenticate once using a single token (no refresh)   |
| `WithTokenRefresh(user, pass)` | Automatically refresh access token in background      |
| `WithTokenExpiry(d)`           | Override token expiry duration (default: 24 hours)    |
| `WithHttpClient(client)`       | Provide a custom `*http.Client` instance for requests |

Example using API key:

```go
client := umami.NewClient("https://umami.example.com",
umami.WithAPIKey("your-api-key"),
)
```

Example using Single Token:

```go
client := umami.NewClient("https://umami.example.com",
umami.WithSingleToken("admin", "password123"),
)
```

To clean up a client that uses token refresh:

```go
client.Close()
```

## Date Ranges

The client provides helper methods for commonly used date ranges. These are useful when making report requests.

**Example:**

```go
params := types.ReportInsightsRequest{
    WebsiteID: "your-website-id",
    DateRange: daterange.Last7Days(),
    // ...
}
```

**Available helpers:**

| Helper                               | Description                          |
| ------------------------------------ |--------------------------------------|
| `daterange.Today()`                  | From start to end of the current day |
| `daterange.Last24Hours()`            | Last 24 hours from now               |
| `daterange.ThisWeek()`               | From start of current week to now    |
| `daterange.Last7Days()`              | Last 7 days from now                 |
| `daterange.ThisMonth()`              | From start of current month to now   |
| `daterange.Last30Days()`             | Last 30 days from now                |
| `daterange.Last90Days()`             | Last 90 days from now                |
| `daterange.ThisYear()`               | From start of current year to now    |
| `daterange.Last6Months()`            | Last 6 months from now               |
| `daterange.Last12Months()`           | Last 12 months from now              |
| `daterange.Custom(start, end, unit)` | Create a custom date range           |

## Interfaces

| Interface      | Description                          |
|----------------|--------------------------------------|
| `Public`       | Public API for sending events        |
| `User`         | User management                      |
| `Team`         | Team membership and management       |
| `Event`        | Event tracking and property analysis |
| `Session`      | Visitor sessions and activity        |
| `Website`      | Website CRUD                         |
| `WebsiteStats` | Analytics: metrics, trends           |
| `Reports`      | Reports                              |

## API Reference

### `Public` Interface

| Method | Endpoint         |
|--------|------------------|
| `Send` | `POST /api/send` |

### `User` Interface

| Method            | Endpoint                          |
|-------------------|-----------------------------------|
| `CreateUser`      | `POST /api/users`                 |
| `ListUsers`       | `GET /api/admin/users`            |
| `GetUser`         | `GET /api/users/:userId`          |
| `UpdateUser`      | `POST /api/users/:userId`         |
| `DeleteUser`      | `DELETE /api/users/:userId`       |
| `GetUserWebsites` | `GET /api/users/:userId/websites` |
| `ListUserTeams`   | `GET /api/users/:userId/teams`    |

### `Team` Interface

| Method              | Endpoint                                        |
|---------------------|-------------------------------------------------|
| `ListTeams`         | `GET /api/teams`                                |
| `CreateTeam`        | `POST /api/teams`                               |
| `GetTeam`           | `GET /api/teams/:teamId`                        |
| `UpdateTeam`        | `POST /api/teams/:teamId`                       |
| `DeleteTeam`        | `DELETE /api/teams/:teamId`                     |
| `GetTeamWebsites`   | `GET /api/teams/:teamId/websites`               |
| `AddTeamWebsite`    | `POST /api/teams/:teamId/websites`              |
| `RemoveTeamWebsite` | `DELETE /api/teams/:teamId/websites/:websiteId` |
| `GetTeamUsers`      | `GET /api/teams/:teamId/users`                  |
| `AddTeamUser`       | `POST /api/teams/:teamId/users`                 |
| `UpdateTeamUser`    | `POST /api/teams/:teamId/users/:userId`         |
| `RemoveTeamUser`    | `DELETE /api/teams/:teamId/users/:userId`       |

### `Event` Interface

| Method               | Endpoint                                         |
|----------------------|--------------------------------------------------|
| `ListEvents`         | `GET /api/websites/:websiteId/events`            |
| `GetEventProperties` | `GET /api/websites/:websiteId/events/properties` |
| `GetEventFields`     | `GET /api/websites/:websiteId/events/fields`     |
| `GetEventValues`     | `GET /api/websites/:websiteId/events/values`     |
| `GetEventDataStats`  | `GET /api/websites/:websiteId/events/stats`      |

### `Session` Interface

| Method                     | Endpoint                                                      |
|----------------------------|---------------------------------------------------------------|
| `ListSessions`             | `GET /api/websites/:websiteId/sessions`                       |
| `GetSessionStats`          | `GET /api/websites/:websiteId/sessions/stats`                 |
| `GetSessionDetails`        | `GET /api/websites/:websiteId/sessions/:sessionId`            |
| `GetSessionActivity`       | `GET /api/websites/:websiteId/sessions/:sessionId/activity`   |
| `GetSessionProperties`     | `GET /api/websites/:websiteId/sessions/:sessionId/properties` |
| `GetSessionDataProperties` | `GET /api/websites/:websiteId/sessions/data/properties`       |
| `GetSessionDataValues`     | `GET /api/websites/:websiteId/sessions/data/values`           |

### `Website` Interface

| Method          | Endpoint                              |
|-----------------|---------------------------------------|
| `ListWebsites`  | `GET /api/websites`                   |
| `CreateWebsite` | `POST /api/websites`                  |
| `GetWebsite`    | `GET /api/websites/:websiteId`        |
| `UpdateWebsite` | `POST /api/websites/:websiteId`       |
| `DeleteWebsite` | `DELETE /api/websites/:websiteId`     |
| `ResetWebsite`  | `POST /api/websites/:websiteId/reset` |

### `WebsiteStats` Interface

| Method                  | Endpoint                                     |
|-------------------------|----------------------------------------------|
| `GetWebsiteActiveUsers` | `GET /api/websites/:websiteId/active`        |
| `GetWebsiteEvents`      | `GET /api/websites/:websiteId/events/series` |
| `GetWebsitePageViews`   | `GET /api/websites/:websiteId/pageviews`     |
| `GetWebsiteStats`       | `GET /api/websites/:websiteId/stats`         |
| `GetWebsiteMetrics`     | `GET /api/websites/:websiteId/metrics`       |

### `Report` Interface

| Method           | Endpoint                        |
| ---------------- | ------------------------------- |
| `GetInsights`    | `POST /api/reports/insights`    |
| `GetFunnel`      | `POST /api/reports/funnel`      |
| `GetRetention`   | `POST /api/reports/retention`   |
| `GetUTM`         | `POST /api/reports/utm`         |
| `GetGoals`       | `POST /api/reports/goals`       |
| `GetJourney`     | `POST /api/reports/journey`     |
| `GetRevenue`     | `POST /api/reports/revenue`     |
| `GetAttribution` | `POST /api/reports/attribution` |

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

## Credits

This is an unofficial SDK for the [Umami Analytics](https://github.com/umami-software/umami) platform.
Maintained independently and not affiliated with the Umami team.

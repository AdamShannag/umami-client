package types

import "time"

type Auth struct {
	Token string `json:"token"`
}

type ListQueryParams struct {
	Query    string `json:"query,omitempty"`    // Optional search string
	Page     string `json:"page,omitempty"`     // Optional page number (default: 1)
	PageSize string `json:"pageSize,omitempty"` // Optional results per page
	OrderBy  string `json:"orderBy,omitempty"`  // Optional order column (default: name)
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"` // Role (e.g., admin, user, view-only)
	CreatedAt time.Time `json:"createdAt"`
}

type UserInfo struct {
	ID          string     `json:"id"`
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	Role        string     `json:"role"`
	LogoURL     *string    `json:"logoUrl"`
	DisplayName *string    `json:"displayName"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
	Count       struct {
		WebsiteUser int64 `json:"websiteUser"`
	} `json:"_count"`
}

type Users struct {
	Data     []UserInfo `json:"data"`
	Count    int64      `json:"count"`
	Page     int64      `json:"page"`
	PageSize int64      `json:"pageSize"`
	OrderBy  string     `json:"orderBy"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

type UserWebsite struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Domain    string     `json:"domain"`
	ShareID   *string    `json:"shareId"`
	ResetAt   *string    `json:"resetAt"`
	UserID    string     `json:"userId"`
	TeamID    *string    `json:"teamId"`
	CreatedBy string     `json:"createdBy"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	User      struct {
		Username string `json:"username"`
		ID       string `json:"id"`
	} `json:"user"`
}

type UserWebsites struct {
	Data     []UserWebsite `json:"data"`
	Count    int64         `json:"count"`
	Page     int64         `json:"page"`
	PageSize int64         `json:"pageSize"`
	OrderBy  string        `json:"orderBy"`
}

type UserTeam struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	AccessCode string     `json:"accessCode"`
	LogoURL    *string    `json:"logoUrl"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
	TeamUser   []struct {
		ID        string    `json:"id"`
		TeamID    string    `json:"teamId"`
		UserID    string    `json:"userId"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		User      struct {
			ID       string `json:"id"`
			Username string `json:"username"`
		} `json:"user"`
	} `json:"teamUser"`
	Count struct {
		Website  int64 `json:"website"`
		TeamUser int64 `json:"teamUser"`
	} `json:"_count"`
}

type UserTeams struct {
	Data     []UserTeam `json:"data"`
	Count    int64      `json:"count"`
	Page     int64      `json:"page"`
	PageSize int64      `json:"pageSize"`
	OrderBy  string     `json:"orderBy"`
}

type Team struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	AccessCode string         `json:"accessCode"`
	LogoURL    *string        `json:"logoUrl"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  *time.Time     `json:"updatedAt"`
	DeletedAt  *time.Time     `json:"deletedAt"`
	TeamUsers  []TeamUserInfo `json:"teamUser"`
}

type TeamUserInfo struct {
	ID        string     `json:"id"`
	TeamID    string     `json:"teamId"`
	UserID    string     `json:"userId"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	User      *TeamUser  `json:"user"`
}

type TeamUsers struct {
	Data     []TeamUserInfo `json:"data"`
	Count    int            `json:"count"`
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
}

type TeamUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type CreateTeamRequest struct {
	Name string `json:"name"`
}

type JoinTeamRequest struct {
	AccessCode string `json:"accessCode"`
}

type UpdateTeamRequest struct {
	Name       *string `json:"name,omitempty"`
	AccessCode *string `json:"accessCode,omitempty"`
}

type AddUserRequest struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
}

type ListTeamUsersResponse struct {
	Data     []TeamUserInfo `json:"data"`
	Count    int            `json:"count"`
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
}

type TeamWebsiteInfo struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Domain    string     `json:"domain"`
	ShareID   *string    `json:"shareId"`
	ResetAt   *time.Time `json:"resetAt"`
	UserID    *string    `json:"userId"`
	TeamID    string     `json:"teamId"`
	CreatedBy string     `json:"CreatedBy"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	User      TeamUser   `json:"createUser"`
}

type TeamWebsites struct {
	Data     []TeamWebsiteInfo `json:"data"`
	Count    int               `json:"count"`
	Page     int               `json:"page"`
	PageSize int               `json:"pageSize"`
}

type EventDetail struct {
	ID             string    `json:"id"`
	WebsiteID      string    `json:"websiteId"`
	SessionID      string    `json:"sessionId"`
	CreatedAt      time.Time `json:"createdAt"`
	URLPath        string    `json:"urlPath"`
	URLQuery       string    `json:"urlQuery"`
	ReferrerPath   string    `json:"referrerPath"`
	ReferrerQuery  string    `json:"referrerQuery"`
	ReferrerDomain string    `json:"referrerDomain"`
	PageTitle      string    `json:"pageTitle"`
	EventType      int       `json:"eventType"`
	EventName      string    `json:"eventName"`
}

type ListEventsResponse struct {
	Data     []EventDetail `json:"data"`
	Count    int           `json:"count"`
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
}

type ListEventsParams struct {
	StartAt  time.Time `json:"startAt"`
	EndAt    time.Time `json:"endAt"`
	Query    string    `json:"query,omitempty"`
	Page     string    `json:"page,omitempty"`
	PageSize string    `json:"pageSize,omitempty"`
	OrderBy  string    `json:"orderBy,omitempty"`
}

type EventPropertyCount struct {
	EventName    string `json:"eventName"`
	PropertyName string `json:"propertyName"`
	DataType     int    `json:"dataType"`
	Total        int    `json:"total"`
}

type EventFieldsStat struct {
	PropertyName string `json:"propertyName"`
	DataType     int    `json:"dataType"`
	Value        string `json:"value"`
	Total        int    `json:"total"`
}

type EventValueCount struct {
	Value string `json:"value"`
	Total int    `json:"total"`
}

type EventDataStats struct {
	Events     int64  `json:"events"`
	Properties int64  `json:"properties"`
	Records    *int64 `json:"records"`
}

type EventDataQueryParams struct {
	StartAt      time.Time `json:"startAt"`
	EndAt        time.Time `json:"endAt"`
	EventName    string    `json:"event,omitempty"`
	PropertyName string    `json:"propertyName,omitempty"`
}

type Session struct {
	ID          string    `json:"id"`
	WebsiteID   string    `json:"websiteId"`
	Hostname    string    `json:"hostname"`
	Browser     string    `json:"browser"`
	OS          string    `json:"os"`
	Device      string    `json:"device"`
	Screen      string    `json:"screen"`
	Language    string    `json:"language"`
	Country     string    `json:"country"`
	Subdivision string    `json:"subdivision1"`
	City        string    `json:"city"`
	FirstAt     time.Time `json:"firstAt"`
	LastAt      time.Time `json:"lastAt"`
	Visits      int       `json:"visits"`
	Views       int       `json:"views"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ListSessionsResponse struct {
	Data     []Session `json:"data"`
	Count    int       `json:"count"`
	Page     int       `json:"page"`
	PageSize int       `json:"pageSize"`
}

type ListSessionsParams struct {
	StartAt  time.Time `json:"startAt"`
	EndAt    time.Time `json:"endAt"`
	Query    string    `json:"query,omitempty"`
	Page     string    `json:"page,omitempty"`
	PageSize string    `json:"pageSize,omitempty"`
	OrderBy  string    `json:"orderBy,omitempty"`
}

type SessionStats struct {
	Pageviews SessionMetric `json:"pageviews"`
	Visitors  SessionMetric `json:"visitors"`
	Visits    SessionMetric `json:"visits"`
	Countries SessionMetric `json:"countries"`
	Events    SessionMetric `json:"events"`
}

type SessionMetric struct {
	Value int `json:"value"`
}

type SessionStatsParams struct {
	StartAt  time.Time `json:"startAt"`
	EndAt    time.Time `json:"endAt"`
	URL      string    `json:"url,omitempty"`
	Referrer string    `json:"referrer,omitempty"`
	Title    string    `json:"title,omitempty"`
	Query    string    `json:"query,omitempty"`
	Event    string    `json:"event,omitempty"`
	Host     string    `json:"host,omitempty"`
	OS       string    `json:"os,omitempty"`
	Browser  string    `json:"browser,omitempty"`
	Device   string    `json:"device,omitempty"`
	Country  string    `json:"country,omitempty"`
	Region   string    `json:"region,omitempty"`
	City     string    `json:"city,omitempty"`
}

type SessionDetails struct {
	ID          string    `json:"id"`
	DistinctID  *string   `json:"distinctId"`
	WebsiteID   string    `json:"websiteId"`
	Hostname    string    `json:"hostname"`
	Browser     string    `json:"browser"`
	OS          string    `json:"os"`
	Device      string    `json:"device"`
	Screen      string    `json:"screen"`
	Language    string    `json:"language"`
	Country     string    `json:"country"`
	Subdivision string    `json:"subdivision1"`
	City        *string   `json:"city"`
	Region      *string   `json:"region"`
	FirstAt     time.Time `json:"firstAt"`
	LastAt      time.Time `json:"lastAt"`
	Visits      int       `json:"visits"`
	Views       string    `json:"views"`
	Events      string    `json:"events"`
	TotalTime   string    `json:"totaltime"`
}

type SessionActivityItem struct {
	ID             string    `json:"id"`
	WebsiteID      string    `json:"websiteId"`
	SessionID      string    `json:"sessionId"`
	VisitID        string    `json:"visitId"`
	CreatedAt      time.Time `json:"createdAt"`
	URLPath        string    `json:"urlPath"`
	URLQuery       string    `json:"urlQuery"`
	UtmSource      *string   `json:"utmSource"`
	UtmMedium      *string   `json:"utmMedium"`
	UtmCampaign    *string   `json:"utmCampaign"`
	UtmContent     *string   `json:"utmContent"`
	UtmTerm        *string   `json:"utmTerm"`
	ReferrerPath   *string   `json:"referrerPath"`
	ReferrerQuery  *string   `json:"referrerQuery"`
	ReferrerDomain *string   `json:"referrerDomain"`
	PageTitle      string    `json:"pageTitle"`
	Gclid          *string   `json:"gclid"`
	Fbclid         *string   `json:"fbclid"`
	Msclkid        *string   `json:"msclkid"`
	Ttclid         *string   `json:"ttclid"`
	Lifatid        *string   `json:"lifatid"`
	Twclid         *string   `json:"twclid"`
	EventType      int       `json:"eventType"`
	EventName      *string   `json:"eventName"`
	Tag            *string   `json:"tag"`
	Hostname       string    `json:"hostname"`
}

type SessionProperty struct {
	WebsiteID   string     `json:"websiteId"`
	SessionID   string     `json:"sessionId"`
	DataKey     string     `json:"dataKey"`
	DataType    int        `json:"dataType"`
	StringValue string     `json:"stringValue,omitempty"`
	NumberValue *float64   `json:"numberValue,omitempty"`
	DateValue   *time.Time `json:"dateValue,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
}

type SessionDataPropertyCount struct {
	PropertyName string `json:"propertyName"`
	Total        int    `json:"total"`
}

type SessionDataValueCount struct {
	Value string `json:"value"`
	Total int    `json:"total"`
}

type SessionDataPropertiesParams struct {
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
}

type SessionDataValuesParams struct {
	StartAt      time.Time `json:"startAt"`
	EndAt        time.Time `json:"endAt"`
	PropertyName string    `json:"propertyName"`
}

type Website struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Domain    string     `json:"domain"`
	ShareID   *string    `json:"shareId"`
	ResetAt   *time.Time `json:"resetAt"`
	WebsiteID *string    `json:"websiteId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	UserID    string     `json:"userId"`
	TeamID    *string    `json:"teamId"`
	CreatedBy string     `json:"createdBy"`
	User      struct {
		Username string `json:"username"`
		ID       string `json:"id"`
	} `json:"user"`
}

type Websites struct {
	Data     []Website `json:"data"`
	Count    int64     `json:"count"`
	Page     int64     `json:"page"`
	PageSize int64     `json:"pageSize"`
	OrderBy  string    `json:"orderBy"`
}

type CreateWebsiteRequest struct {
	Domain  string  `json:"domain"`
	Name    string  `json:"name"`
	ShareID *string `json:"shareId,omitempty"`
	TeamID  *string `json:"teamId,omitempty"`
}

type UpdateWebsiteRequest struct {
	Name    string  `json:"name"`
	Domain  string  `json:"domain"`
	ShareID *string `json:"shareId"`
}

type WebsiteActiveUsers struct {
	Visitors int64 `json:"visitors"`
}

type WebsiteEvent struct {
	ID             string    `json:"id"`
	WebsiteID      string    `json:"websiteId"`
	SessionID      string    `json:"sessionId"`
	CreatedAt      time.Time `json:"createdAt"`
	URLPath        string    `json:"urlPath"`
	URLQuery       string    `json:"urlQuery"`
	ReferrerPath   *string   `json:"referrerPath"`
	ReferrerQuery  *string   `json:"referrerQuery"`
	ReferrerDomain *string   `json:"referrerDomain"`
	PageTitle      string    `json:"pageTitle"`
	EventType      int       `json:"eventType"`
	EventName      *string   `json:"eventName"`
}

type WebsiteEvents struct {
	Data     []WebsiteEvent `json:"data"`
	Count    int64          `json:"count"`
	Page     int64          `json:"page"`
	PageSize int64          `json:"pageSize"`
}

type WebsiteEventsQueryParams struct {
	StartAt  time.Time
	EndAt    time.Time
	Unit     string
	Timezone string
	URL      string
	Referrer string
	Title    string
	Host     string
	OS       string
	Browser  string
	Device   string
	Country  string
	Region   string
	City     string
}

type Metric struct {
	Value int64 `json:"value"`
	Prev  int64 `json:"prev"`
}

type WebsiteStats struct {
	Pageviews Metric `json:"pageviews"`
	Visitors  Metric `json:"visitors"`
	Visits    Metric `json:"visits"`
	Bounces   Metric `json:"bounces"`
	TotalTime Metric `json:"totaltime"`
}

type WebsiteStatsQueryParams struct {
	StartAt  time.Time
	EndAt    time.Time
	URL      string
	Referrer string
	Title    string
	Query    string
	Event    string
	Host     string
	OS       string
	Browser  string
	Device   string
	Country  string
	Region   string
	City     string
}

type WebsitePageViewsQueryParams struct {
	StartAt  time.Time
	EndAt    time.Time
	Unit     string
	Timezone string
	URL      string
	Referrer string
	Title    string
	Host     string
	OS       string
	Browser  string
	Device   string
	Country  string
	Region   string
	City     string
}

type TimeSeriesDataPoint struct {
	Timestamp        CustomTime `json:"x"`
	NumberOfVisitors int        `json:"y"`
}

type WebsitePageViews struct {
	Pageviews []TimeSeriesDataPoint `json:"pageviews"`
	Sessions  []TimeSeriesDataPoint `json:"sessions"`
}

type WebsiteMetricsQueryParams struct {
	StartAt  time.Time
	EndAt    time.Time
	Type     string
	URL      string
	Referrer string
	Title    string
	Query    string
	Host     string
	OS       string
	Browser  string
	Device   string
	Country  string
	Region   string
	City     string
	Language string
	Event    string
	Limit    int
}

type WebsiteMetric struct {
	Value            string `json:"x"`
	NumberOfVisitors int    `json:"y"`
}

type SendEventRequest struct {
	Payload SendEventPayload `json:"payload"`
	Type    string           `json:"type"`
}
type SendEventPayload struct {
	Website  string         `json:"website"`
	Session  string         `json:"session,omitempty"`
	Hostname string         `json:"hostname,omitempty"`
	Language string         `json:"language,omitempty"`
	Referrer string         `json:"referrer,omitempty"`
	Screen   string         `json:"screen,omitempty"`
	Title    string         `json:"title,omitempty"`
	URL      string         `json:"url,omitempty"`
	Name     string         `json:"name,omitempty"`
	Data     map[string]any `json:"data,omitempty"`
}

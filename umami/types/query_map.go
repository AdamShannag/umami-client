package types

import (
	"fmt"
	"strconv"
	"time"
)

func (p ListQueryParams) ToQueryMap() map[string]string {
	q := make(map[string]string)

	if p.Query != "" {
		q["query"] = p.Query
	}
	if p.Page != "" {
		q["page"] = p.Page
	}
	if p.PageSize != "" {
		q["pageSize"] = p.PageSize
	}
	if p.OrderBy != "" {
		q["orderBy"] = p.OrderBy
	}

	return q
}

func (p ListEventsParams) ToQueryMap() map[string]string {
	q := map[string]string{
		"startAt": fmt.Sprintf("%d", p.StartAt.UnixMilli()),
		"endAt":   fmt.Sprintf("%d", p.EndAt.UnixMilli()),
	}

	if p.Query != "" {
		q["query"] = p.Query
	}
	if p.Page != "" {
		q["page"] = p.Page
	}
	if p.PageSize != "" {
		q["pageSize"] = p.PageSize
	}
	if p.OrderBy != "" {
		q["orderBy"] = p.OrderBy
	}

	return q
}

func (p ListSessionsParams) ToQueryMap() map[string]string {
	q := map[string]string{
		"startAt": fmt.Sprintf("%d", p.StartAt.UnixMilli()),
		"endAt":   fmt.Sprintf("%d", p.EndAt.UnixMilli()),
	}

	if p.Query != "" {
		q["query"] = p.Query
	}
	if p.Page != "" {
		q["page"] = p.Page
	}
	if p.PageSize != "" {
		q["pageSize"] = p.PageSize
	}
	if p.OrderBy != "" {
		q["orderBy"] = p.OrderBy
	}

	return q
}

func (p EventDataQueryParams) ToQueryMap() map[string]string {
	q := map[string]string{
		"startAt": fmt.Sprintf("%d", p.StartAt.UnixMilli()),
		"endAt":   fmt.Sprintf("%d", p.EndAt.UnixMilli()),
	}

	if p.EventName != "" {
		q["event"] = p.EventName
	}
	if p.PropertyName != "" {
		q["propertyName"] = p.PropertyName
	}

	return q
}

func (p WebsiteEventsQueryParams) ToQueryMap() map[string]string {
	q := buildBaseQueryParams(
		p.StartAt, p.EndAt,
		p.URL, p.Referrer, p.Title,
		p.Host, p.OS, p.Browser, p.Device,
		p.Country, p.Region, p.City,
	)

	if p.Unit != "" {
		q["unit"] = p.Unit
	}
	if p.Timezone != "" {
		q["timezone"] = p.Timezone
	}
	return q
}

func (p WebsiteStatsQueryParams) ToQueryMap() map[string]string {
	q := buildBaseQueryParams(
		p.StartAt, p.EndAt,
		p.URL, p.Referrer, p.Title,
		p.Host, p.OS, p.Browser, p.Device,
		p.Country, p.Region, p.City,
	)

	if p.Query != "" {
		q["query"] = p.Query
	}
	if p.Event != "" {
		q["event"] = p.Event
	}
	return q
}

func (p WebsiteMetricsQueryParams) ToQueryMap() map[string]string {
	q := buildBaseQueryParams(
		p.StartAt, p.EndAt,
		p.URL, p.Referrer, p.Title,
		p.Host, p.OS, p.Browser, p.Device,
		p.Country, p.Region, p.City,
	)

	if p.Type != "" {
		q["type"] = p.Type
	}
	if p.Query != "" {
		q["query"] = p.Query
	}
	if p.Language != "" {
		q["language"] = p.Language
	}
	if p.Event != "" {
		q["event"] = p.Event
	}
	if p.Limit > 0 {
		q["limit"] = strconv.Itoa(p.Limit)
	}

	return q
}

func (p WebsitePageViewsQueryParams) ToQueryMap() map[string]string {
	q := buildBaseQueryParams(
		p.StartAt, p.EndAt,
		p.URL, p.Referrer, p.Title,
		p.Host, p.OS, p.Browser, p.Device,
		p.Country, p.Region, p.City,
	)

	if p.Unit != "" {
		q["unit"] = p.Unit
	}
	if p.Timezone != "" {
		q["timezone"] = p.Timezone
	}
	return q
}

func (p SessionStatsParams) ToQueryMap() map[string]string {
	q := buildBaseQueryParams(
		p.StartAt, p.EndAt,
		p.URL, p.Referrer, p.Title,
		p.Host, p.OS, p.Browser, p.Device,
		p.Country, p.Region, p.City,
	)

	if p.Query != "" {
		q["query"] = p.Query
	}
	if p.Event != "" {
		q["event"] = p.Event
	}

	return q
}

func (p SessionDataValuesParams) ToQueryMap() map[string]string {
	q := map[string]string{
		"startAt":      fmt.Sprintf("%d", p.StartAt.UnixMilli()),
		"endAt":        fmt.Sprintf("%d", p.EndAt.UnixMilli()),
		"propertyName": p.PropertyName,
	}

	return q
}

func (p SessionDataPropertiesParams) ToQueryMap() map[string]string {
	return map[string]string{
		"startAt": fmt.Sprintf("%d", p.StartAt.UnixMilli()),
		"endAt":   fmt.Sprintf("%d", p.EndAt.UnixMilli()),
	}
}

func buildBaseQueryParams(
	startAt, endAt time.Time,
	url, referrer, title, host, os, browser, device, country, region, city string,
) map[string]string {
	q := map[string]string{
		"startAt": fmt.Sprintf("%d", startAt.UnixMilli()),
		"endAt":   fmt.Sprintf("%d", endAt.UnixMilli()),
	}

	if url != "" {
		q["url"] = url
	}
	if referrer != "" {
		q["referrer"] = referrer
	}
	if title != "" {
		q["title"] = title
	}
	if host != "" {
		q["host"] = host
	}
	if os != "" {
		q["os"] = os
	}
	if browser != "" {
		q["browser"] = browser
	}
	if device != "" {
		q["device"] = device
	}
	if country != "" {
		q["country"] = country
	}
	if region != "" {
		q["region"] = region
	}
	if city != "" {
		q["city"] = city
	}
	return q
}

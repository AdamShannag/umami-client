package daterange

import (
	"fmt"
	"github.com/AdamShannag/umami-client/umami/types"
	"time"
)

func Today() types.DateRange {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999000000, time.UTC)

	return types.DateRange{
		StartDate: start,
		EndDate:   end,
		Unit:      "hour",
		Offset:    0,
		Num:       1,
		Value:     "0day",
	}
}

func Last24Hours() types.DateRange {
	now := time.Now()
	start := now.Add(-24 * time.Hour).Truncate(time.Hour)
	end := now
	return types.DateRange{
		StartDate: start,
		EndDate:   end,
		Unit:      "hour",
		Offset:    0,
		Num:       24,
		Value:     "24hour",
	}
}

func ThisWeek() types.DateRange {
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, -weekday+1)
	end := start.AddDate(0, 0, 7).Add(-time.Millisecond)

	return types.DateRange{
		StartDate: start,
		EndDate:   end,
		Unit:      "day",
		Offset:    0,
		Num:       1,
		Value:     "0week",
	}
}

func Last7Days() types.DateRange {
	now := time.Now()
	start := now.AddDate(0, 0, -7).Truncate(24 * time.Hour)
	end := now
	return types.DateRange{
		StartDate: start,
		EndDate:   end,
		Unit:      "day",
		Offset:    0,
		Num:       7,
		Value:     "7day",
	}
}

func ThisMonth() types.DateRange {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0).Add(-time.Millisecond)

	return types.DateRange{
		StartDate: start,
		EndDate:   end,
		Unit:      "day",
		Offset:    0,
		Num:       1,
		Value:     "0month",
	}
}

func Last30Days() types.DateRange {
	now := time.Now()
	start := now.AddDate(0, 0, -30).Truncate(24 * time.Hour)
	end := now
	return types.DateRange{
		StartDate: start,
		EndDate:   end,
		Unit:      "day",
		Offset:    0,
		Num:       30,
		Value:     "30day",
	}
}

func Last90Days() types.DateRange {
	now := time.Now()
	start := now.AddDate(0, 0, -90).Truncate(24 * time.Hour)
	end := now
	return types.DateRange{
		StartDate: start,
		EndDate:   end,
		Unit:      "day",
		Offset:    0,
		Num:       90,
		Value:     "90day",
	}
}

func ThisYear() types.DateRange {
	now := time.Now()
	start := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(now.Year(), 12, 31, 23, 59, 59, 999000000, time.UTC)

	return types.DateRange{
		StartDate: start,
		EndDate:   end,
		Unit:      "month",
		Offset:    0,
		Num:       1,
		Value:     "0year",
	}
}

func Last6Months() types.DateRange {
	now := time.Now()
	start := now.AddDate(0, -6, 0).Truncate(24 * time.Hour)
	end := now
	return types.DateRange{
		StartDate: start,
		EndDate:   end,
		Unit:      "month",
		Offset:    0,
		Num:       6,
		Value:     "6month",
	}
}

func Last12Months() types.DateRange {
	now := time.Now()
	start := now.AddDate(0, -12, 0).Truncate(24 * time.Hour)
	end := now
	return types.DateRange{
		StartDate: start,
		EndDate:   end,
		Unit:      "month",
		Offset:    0,
		Num:       12,
		Value:     "12month",
	}
}

func Custom(start, end time.Time, unit string) types.DateRange {
	start = start.Truncate(time.Millisecond)
	end = end.Truncate(time.Millisecond)

	startMs := start.UnixMilli()
	endMs := end.UnixMilli()

	value := fmt.Sprintf("range:%d:%d", startMs, endMs)

	return types.DateRange{
		StartDate: start,
		EndDate:   end,
		Unit:      unit,
		Offset:    0,
		Num:       int(end.Sub(start).Hours() / 24),
		Value:     value,
	}
}

package daterange

import (
	"fmt"
	"github.com/AdamShannag/umami-client/umami/types"
	"testing"
	"time"
)

func TestDateRanges(t *testing.T) {
	tests := []struct {
		name    string
		got     types.DateRange
		expUnit string
		expVal  string
	}{
		{"Today", Today(), "hour", "0day"},
		{"Last24Hours", Last24Hours(), "hour", "24hour"},
		{"ThisWeek", ThisWeek(), "day", "0week"},
		{"Last7Days", Last7Days(), "day", "7day"},
		{"ThisMonth", ThisMonth(), "day", "0month"},
		{"Last30Days", Last30Days(), "day", "30day"},
		{"Last90Days", Last90Days(), "day", "90day"},
		{"ThisYear", ThisYear(), "month", "0year"},
		{"Last6Months", Last6Months(), "month", "6month"},
		{"Last12Months", Last12Months(), "month", "12month"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.Unit != tt.expUnit {
				t.Errorf("expected unit %s, got %s", tt.expUnit, tt.got.Unit)
			}
			if tt.got.Value != tt.expVal {
				t.Errorf("expected value %s, got %s", tt.expVal, tt.got.Value)
			}
			if tt.got.StartDate.After(tt.got.EndDate) {
				t.Errorf("StartDate should be before EndDate")
			}
			if tt.got.EndDate.Before(tt.got.StartDate) {
				t.Errorf("EndDate should not be after StartDate")
			}
		})
	}
}

func TestNewCustomDateRange(t *testing.T) {
	start := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 6, 10, 23, 59, 59, 999000000, time.UTC)
	r := Custom(start, end, "day")

	if r.StartDate != start.Truncate(time.Millisecond) {
		t.Errorf("Expected start date %v, got %v", start, r.StartDate)
	}
	if r.EndDate != end.Truncate(time.Millisecond) {
		t.Errorf("Expected end date %v, got %v", end, r.EndDate)
	}
	if r.Unit != "day" {
		t.Errorf("Expected unit 'day', got %s", r.Unit)
	}

	expVal := fmt.Sprintf("range:%d:%d", start.UnixMilli(), end.UnixMilli())
	if r.Value != expVal {
		t.Errorf("Expected value %s, got %s", expVal, r.Value)
	}
}

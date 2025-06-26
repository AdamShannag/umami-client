package types_test

import (
	"encoding/json"
	"github.com/AdamShannag/umami-client/umami/types"
	"testing"
	"time"
)

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	jsonData := `"2025-06-26 12:34:56"`
	var ct types.CustomTime

	err := json.Unmarshal([]byte(jsonData), &ct)
	if err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}

	expected := time.Date(2025, 6, 26, 12, 34, 56, 0, time.UTC)
	if !ct.Time.Equal(expected) {
		t.Errorf("Expected time %v, got %v", expected, ct.Time)
	}
}

func TestCustomTime_UnmarshalJSON_Invalid(t *testing.T) {
	invalidJSON := `"invalid-date-format"`
	var ct types.CustomTime

	err := json.Unmarshal([]byte(invalidJSON), &ct)
	if err == nil {
		t.Fatalf("Expected error for invalid date format but got none")
	}
}

func TestCustomTime_MarshalJSON(t *testing.T) {
	ct := types.CustomTime{
		Time: time.Date(2025, 6, 26, 12, 34, 56, 0, time.UTC),
	}

	b, err := json.Marshal(&ct)
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}

	expectedJSON := `"2025-06-26 12:34:56"`
	if string(b) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(b))
	}
}

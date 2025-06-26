package types

import (
	"encoding/json"
	"strings"
	"time"
)

const customTimeLayout = "2006-01-02 15:04:05"

type CustomTime struct {
	time.Time
}

// UnmarshalJSON parses a JSON string to CustomTime
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	t, err := time.Parse(customTimeLayout, str)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

// MarshalJSON formats CustomTime to JSON string
func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Time.Format(customTimeLayout))
}

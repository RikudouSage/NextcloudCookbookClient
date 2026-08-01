package types

import (
	"encoding/json"
	"time"
)

const compactOffsetLayout = "2006-01-02T15:04:05-0700"

type APITime struct {
	time.Time
}

func (receiver *APITime) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}

	if value == "" {
		receiver.Time = time.Time{}
		return nil
	}

	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		parsed, err = time.Parse(compactOffsetLayout, value)
		if err != nil {
			return err
		}
	}

	receiver.Time = parsed
	return nil
}

func (receiver APITime) MarshalJSON() ([]byte, error) {
	return receiver.Time.MarshalJSON()
}

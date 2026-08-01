package types

import (
	"encoding/json"
	"net/url"
)

type APIURL struct {
	url.URL
}

func (receiver *APIURL) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}

	parsed, err := url.Parse(value)
	if err != nil {
		return err
	}

	receiver.URL = *parsed
	return nil
}

func (receiver APIURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(receiver.String())
}

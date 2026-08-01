package types

import (
	"encoding/json"
	"testing"
)

func TestAPIURLUnmarshalJSONAcceptsString(t *testing.T) {
	var parsed APIURL
	if err := json.Unmarshal([]byte(`"https://example.com/image.jpg?size=large"`), &parsed); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if parsed.String() != "https://example.com/image.jpg?size=large" {
		t.Fatalf("unexpected URL: %s", parsed.String())
	}
}

func TestAPIURLPointerFieldUnmarshalJSONAcceptsString(t *testing.T) {
	var decoded struct {
		ImageURL *APIURL `json:"imageUrl"`
	}

	if err := json.Unmarshal([]byte(`{"imageUrl":"https://example.com/image.jpg"}`), &decoded); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if decoded.ImageURL == nil {
		t.Fatal("expected ImageURL to be set")
	}
	if decoded.ImageURL.String() != "https://example.com/image.jpg" {
		t.Fatalf("unexpected URL: %s", decoded.ImageURL.String())
	}
}

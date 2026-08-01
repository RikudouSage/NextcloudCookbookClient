package types

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAPITimeUnmarshalJSONAcceptsCompactOffset(t *testing.T) {
	var parsed APITime
	if err := json.Unmarshal([]byte(`"2026-07-29T22:08:02+0000"`), &parsed); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := time.Date(2026, time.July, 29, 22, 8, 2, 0, time.UTC)
	if !parsed.Equal(expected) {
		t.Fatalf("expected %s, got %s", expected, parsed.Time)
	}
}

func TestAPITimeUnmarshalJSONAcceptsRFC3339(t *testing.T) {
	var parsed APITime
	if err := json.Unmarshal([]byte(`"2026-07-29T22:08:02Z"`), &parsed); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := time.Date(2026, time.July, 29, 22, 8, 2, 0, time.UTC)
	if !parsed.Equal(expected) {
		t.Fatalf("expected %s, got %s", expected, parsed.Time)
	}
}

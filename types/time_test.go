package types

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAPITimeUnmarshalJSONAcceptsStandardFormats(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
	}{
		{
			name:     "RFC3339 UTC",
			input:    "2026-07-29T22:08:02Z",
			expected: time.Date(2026, time.July, 29, 22, 8, 2, 0, time.UTC),
		},
		{
			name:     "RFC3339 numeric offset",
			input:    "2026-07-29T22:08:02+02:30",
			expected: time.Date(2026, time.July, 29, 22, 8, 2, 0, time.FixedZone("", 2*60*60+30*60)),
		},
		{
			name:     "compact numeric offset",
			input:    "2026-07-29T22:08:02+0000",
			expected: time.Date(2026, time.July, 29, 22, 8, 2, 0, time.UTC),
		},
		{
			name:     "RFC1123Z",
			input:    "Wed, 29 Jul 2026 22:08:02 +0230",
			expected: time.Date(2026, time.July, 29, 22, 8, 2, 0, time.FixedZone("", 2*60*60+30*60)),
		},
		{
			name:     "ANSIC",
			input:    "Wed Jul 29 22:08:02 2026",
			expected: time.Date(2026, time.July, 29, 22, 8, 2, 0, time.Local),
		},
		{
			name:     "RubyDate",
			input:    "Wed Jul 29 22:08:02 +0230 2026",
			expected: time.Date(2026, time.July, 29, 22, 8, 2, 0, time.FixedZone("", 2*60*60+30*60)),
		},
		{
			name:     "DateTime",
			input:    "2026-07-29 22:08:02",
			expected: time.Date(2026, time.July, 29, 22, 8, 2, 0, time.Local),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var parsed APITime
			input, err := json.Marshal(test.input)
			if err != nil {
				t.Fatalf("marshal input: %v", err)
			}
			if err := json.Unmarshal(input, &parsed); err != nil {
				t.Fatalf("unmarshal %q: %v", test.input, err)
			}
			if !parsed.Equal(test.expected) {
				t.Errorf("expected %s, got %s", test.expected, parsed.Time)
			}
		})
	}
}

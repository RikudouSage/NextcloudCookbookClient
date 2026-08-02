package cookbook

import (
	"net/url"
	"testing"
)

func TestURLWithPathPreservesInstancePath(t *testing.T) {
	baseURL, err := url.Parse("https://example.com/nextcloud")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	actual := urlWithPath(baseURL, "/recipes")
	expected := "https://example.com/nextcloud/apps/cookbook/api/v1/recipes"
	if actual.String() != expected {
		t.Fatalf("expected %q, got %q", expected, actual.String())
	}
}

func TestURLWithPathDoesNotDuplicateAPIPrefix(t *testing.T) {
	baseURL, err := url.Parse("https://example.com/nextcloud/")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	actual := urlWithPath(baseURL, "/apps/cookbook/api/v1/recipes")
	expected := "https://example.com/nextcloud/apps/cookbook/api/v1/recipes"
	if actual.String() != expected {
		t.Fatalf("expected %q, got %q", expected, actual.String())
	}
}

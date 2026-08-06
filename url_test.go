package cookbook

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
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

func TestValidateAndNormalizeNextcloudURLInsertsIndexPHPBeforeAPIPrefix(t *testing.T) {
	var paths []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		if r.URL.Path == "/nextcloud/index.php/apps/cookbook/api/v1/config" {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	parsedURL, err := url.Parse(server.URL + "/nextcloud/apps/cookbook/api/v1/recipes")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := validateAndNormalizeNextcloudURL(context.Background(), server.Client(), parsedURL); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedPaths := []string{
		"/nextcloud/apps/cookbook/api/v1/config",
		"/nextcloud/index.php/apps/cookbook/api/v1/config",
	}
	if !reflect.DeepEqual(paths, expectedPaths) {
		t.Fatalf("expected probe paths %v, got %v", expectedPaths, paths)
	}

	expectedPath := "/nextcloud/index.php/apps/cookbook/api/v1/recipes"
	if parsedURL.Path != expectedPath {
		t.Fatalf("expected normalized path %q, got %q", expectedPath, parsedURL.Path)
	}
}

func TestValidateAndNormalizeNextcloudURLLeavesEndpointPathWhenIndexPHPIsNotNeeded(t *testing.T) {
	var paths []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	parsedURL, err := url.Parse(server.URL + "/nextcloud/apps/cookbook/api/v1/recipes")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := validateAndNormalizeNextcloudURL(context.Background(), server.Client(), parsedURL); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedPaths := []string{"/nextcloud/apps/cookbook/api/v1/config"}
	if !reflect.DeepEqual(paths, expectedPaths) {
		t.Fatalf("expected probe paths %v, got %v", expectedPaths, paths)
	}

	expectedPath := "/nextcloud/apps/cookbook/api/v1/recipes"
	if parsedURL.Path != expectedPath {
		t.Fatalf("expected path %q, got %q", expectedPath, parsedURL.Path)
	}
}

package model

import (
	"net/url"
	"testing"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/types"
)

func TestPopulateAbsoluteUrlsPreservesInstancePath(t *testing.T) {
	baseURL, err := url.Parse("https://example.com/nextcloud")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	recipe := &Recipe{
		ImageURL: &types.APIURL{URL: url.URL{Path: "/apps/cookbook/api/v1/recipes/1/image"}},
	}

	if err := recipe.SetBaseURL(baseURL); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "https://example.com/nextcloud/apps/cookbook/api/v1/recipes/1/image"
	if recipe.ImageURL.String() != expected {
		t.Fatalf("expected %q, got %q", expected, recipe.ImageURL.String())
	}
}

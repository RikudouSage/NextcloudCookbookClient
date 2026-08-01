package cookbook

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRequestReturnsRawBytes(t *testing.T) {
	expected := []byte{0x89, 0x50, 0x4e, 0x47}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(expected)
	}))
	defer server.Close()

	parsedURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	actual, err := request[[]byte](
		context.Background(),
		server.Client(),
		http.MethodGet,
		parsedURL,
		nil,
		"",
		"",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(actual) != string(expected) {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

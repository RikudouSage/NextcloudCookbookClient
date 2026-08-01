package cookbook

import (
	"context"
	"net/http"
	"net/url"
)

type Misc interface {
	Reindex(ctx context.Context) error
}

type misc struct {
	username, password string
	url                *url.URL
	httpClient         *http.Client
}

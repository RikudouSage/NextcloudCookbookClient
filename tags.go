package cookbook

import (
	"context"
	"net/http"
	"net/url"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

type Tags interface {
	List(ctx context.Context) ([]*model.Keyword, error)
	Recipes(ctx context.Context, keywords []string) ([]*model.RecipeStub, error)
}

type tags struct {
	username, password string
	url                *url.URL
	httpClient         *http.Client
}

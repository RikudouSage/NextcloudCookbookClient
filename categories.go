package cookbook

import (
	"context"
	"net/http"
	"net/url"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

type Categories interface {
	List(ctx context.Context) ([]*model.Category, error)
	Recipes(ctx context.Context, category *model.Category) ([]*model.RecipeStub, error)
	Rename(ctx context.Context, category *model.Category, newName string) (*model.Category, error)
}

type categories struct {
	username, password string
	url                *url.URL
	httpClient         *http.Client
}

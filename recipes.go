package cookbook

import (
	"context"
	"net/http"
	"net/url"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

type Recipes interface {
	Import(ctx context.Context, url string) (*model.Recipe, error)
	List(ctx context.Context) ([]*model.RecipeStub, error)
	Image(ctx context.Context, ID string, size ImageSize) ([]byte, error)
	Search(ctx context.Context, query string) ([]*model.RecipeStub, error)
	Create(ctx context.Context, recipe *model.Recipe) (*model.Recipe, error)
	Get(ctx context.Context, ID string) (*model.Recipe, error)
	Update(ctx context.Context, recipe *model.Recipe) (*model.Recipe, error)
	Delete(ctx context.Context, ID string) error
}

type recipes struct {
	username, password string
	url                *url.URL
	httpClient         *http.Client
}

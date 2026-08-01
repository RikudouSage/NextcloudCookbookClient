package cookbook

import (
	"context"
	"fmt"
	"net/http"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

func (receiver *categories) Recipes(ctx context.Context, category *model.Category) ([]*model.RecipeStub, error) {
	return request[[]*model.RecipeStub](
		ctx,
		receiver.httpClient,
		http.MethodGet,
		urlWithPath(receiver.url, fmt.Sprintf("/category/%s", category.Name)),
		nil,
		receiver.username, receiver.password,
	)
}

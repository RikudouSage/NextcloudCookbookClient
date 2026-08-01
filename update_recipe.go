package cookbook

import (
	"context"
	"fmt"
	"net/http"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

func (receiver *recipes) Update(ctx context.Context, recipe *model.Recipe) (*model.Recipe, error) {
	_, err := request[any](
		ctx,
		receiver.httpClient,
		http.MethodPut,
		urlWithPath(receiver.url, fmt.Sprintf("recipes/%s", recipe.ID)),
		recipe,
		receiver.username, receiver.password,
	)
	if err != nil {
		return nil, fmt.Errorf("failed updating recipe: %w", err)
	}

	return recipe, nil
}

package cookbook

import (
	"context"
	"fmt"
	"net/http"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

func (receiver *recipes) Get(ctx context.Context, ID string) (*model.Recipe, error) {
	recipe, err := request[*model.Recipe](
		ctx,
		receiver.httpClient,
		http.MethodGet,
		urlWithPath(receiver.url, fmt.Sprintf("recipes/%s", ID)),
		nil,
		receiver.username, receiver.password,
	)
	if err != nil {
		return nil, fmt.Errorf("failed getting recipe: %w", err)
	}

	if err = recipe.SetBaseURL(receiver.url); err != nil {
		return nil, fmt.Errorf("failed setting base url: %w", err)
	}

	return recipe, nil
}

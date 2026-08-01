package cookbook

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

func (receiver *recipes) Create(ctx context.Context, recipe *model.Recipe) (*model.Recipe, error) {
	id, err := request[int](
		ctx,
		receiver.httpClient,
		http.MethodPost,
		urlWithPath(receiver.url, "/recipes"),
		recipe,
		receiver.username, receiver.password,
	)
	if err != nil {
		return nil, fmt.Errorf("failed creating a recipe: %w", err)
	}

	return receiver.Get(ctx, strconv.Itoa(id))
}

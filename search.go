package cookbook

import (
	"context"
	"fmt"
	"net/http"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

func (receiver *recipes) Search(ctx context.Context, query string) ([]*model.RecipeStub, error) {
	return request[[]*model.RecipeStub](
		ctx,
		receiver.httpClient,
		http.MethodGet,
		urlWithPath(receiver.url, fmt.Sprintf("/search/%s", query)),
		nil,
		receiver.username, receiver.password,
	)
}

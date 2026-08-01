package cookbook

import (
	"context"
	"net/http"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

func (receiver *recipes) List(ctx context.Context) ([]*model.RecipeStub, error) {
	return request[[]*model.RecipeStub](
		ctx,
		receiver.httpClient,
		http.MethodGet,
		urlWithPath(receiver.url, "/recipes"),
		nil,
		receiver.username, receiver.password,
	)
}

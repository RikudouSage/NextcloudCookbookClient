package cookbook

import (
	"context"
	"net/http"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

func (receiver *categories) List(ctx context.Context) ([]*model.Category, error) {
	return request[[]*model.Category](
		ctx,
		receiver.httpClient,
		http.MethodGet,
		urlWithPath(receiver.url, "/categories"),
		nil,
		receiver.username, receiver.password,
	)
}

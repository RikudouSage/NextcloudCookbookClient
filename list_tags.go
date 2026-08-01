package cookbook

import (
	"context"
	"net/http"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

func (receiver *tags) List(ctx context.Context) ([]*model.Keyword, error) {
	return request[[]*model.Keyword](
		ctx,
		receiver.httpClient,
		http.MethodGet,
		urlWithPath(receiver.url, "/keywords"),
		nil,
		receiver.username, receiver.password,
	)
}

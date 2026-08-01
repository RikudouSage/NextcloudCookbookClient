package cookbook

import (
	"context"
	"fmt"
	"net/http"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

func (receiver *recipes) Get(ctx context.Context, ID string) (*model.Recipe, error) {
	return request[*model.Recipe](
		ctx,
		receiver.httpClient,
		http.MethodGet,
		urlWithPath(receiver.url, fmt.Sprintf("recipes/%s", ID)),
		nil,
		receiver.username, receiver.password,
	)
}

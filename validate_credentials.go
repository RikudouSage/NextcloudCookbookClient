package cookbook

import (
	"context"
	"net/http"
)

func (receiver *client) ValidateCredentials(ctx context.Context) (bool, error) {
	_, err := request[any](
		ctx,
		receiver.httpClient,
		http.MethodGet,
		urlWithPath(receiver.url, "/config"),
		nil,
		receiver.username, receiver.password,
	)

	return err == nil, err
}

package cookbook

import (
	"context"
	"fmt"
	"net/http"
)

func (receiver *misc) Reindex(ctx context.Context) error {
	_, err := request[any](
		ctx,
		receiver.httpClient,
		http.MethodPost,
		urlWithPath(receiver.url, "reindex"),
		nil,
		receiver.username, receiver.password,
	)

	if err != nil {
		return fmt.Errorf("failed triggering reindex: %w", err)
	}

	return nil
}

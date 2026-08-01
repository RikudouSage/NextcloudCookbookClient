package cookbook

import (
	"context"
	"fmt"
	"net/http"
)

func (receiver *recipes) Delete(ctx context.Context, ID string) error {
	_, err := request[any](
		ctx,
		receiver.httpClient,
		http.MethodDelete,
		urlWithPath(receiver.url, fmt.Sprintf("recipes/%s", ID)),
		nil,
		receiver.username, receiver.password,
	)

	if err != nil {
		return fmt.Errorf("failed deleting a recipe: %w", err)
	}

	return nil
}

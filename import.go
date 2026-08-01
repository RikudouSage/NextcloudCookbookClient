package cookbook

import (
	"context"
	"errors"
	"net/http"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

var ErrRecipeAlreadyExists = errors.New("the recipe is already in the database")

func (receiver *recipes) Import(ctx context.Context, url string) (*model.Recipe, error) {
	resp, err := request[*model.Recipe](
		ctx,
		receiver.httpClient,
		http.MethodPost,
		urlWithPath(receiver.url, "/import"),
		map[string]string{
			"url": url,
		},
		receiver.username,
		receiver.password,
	)

	if err != nil && err.Error() == "unexpected status code: 409" {
		return nil, ErrRecipeAlreadyExists
	}

	return resp, nil
}

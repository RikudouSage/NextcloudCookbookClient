package cookbook

import (
	"context"
	"errors"
	"fmt"
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
	} else if err != nil {
		return nil, fmt.Errorf("failed importing: %w", err)
	}

	if err = resp.SetBaseURL(receiver.url); err != nil {
		return nil, fmt.Errorf("failed setting base url: %w", err)
	}

	return resp, nil
}

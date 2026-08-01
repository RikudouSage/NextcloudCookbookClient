package cookbook

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

func (receiver *tags) Recipes(ctx context.Context, keywords []string) ([]*model.RecipeStub, error) {
	return request[[]*model.RecipeStub](
		ctx,
		receiver.httpClient,
		http.MethodGet,
		urlWithPath(receiver.url, fmt.Sprintf("/tags/%s", strings.Join(keywords, ","))),
		nil,
		receiver.username, receiver.password,
	)
}

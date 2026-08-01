package cookbook

import (
	"context"
	"fmt"
	"net/http"

	"github.com/samber/lo"
	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

func (receiver *categories) Recipes(ctx context.Context, category *model.Category) ([]*model.RecipeStub, error) {
	items, err := request[[]*model.RecipeStub](
		ctx,
		receiver.httpClient,
		http.MethodGet,
		urlWithPath(receiver.url, fmt.Sprintf("/category/%s", category.Name)),
		nil,
		receiver.username, receiver.password,
	)

	if err != nil {
		return nil, fmt.Errorf("failed getting category recipes: %w", err)
	}

	return lo.MapErr(items, func(item *model.RecipeStub, _ int) (*model.RecipeStub, error) {
		if err := item.SetBaseURL(receiver.url); err != nil {
			return nil, fmt.Errorf("failed setting base url: %w", err)
		}

		return item, nil
	})
}

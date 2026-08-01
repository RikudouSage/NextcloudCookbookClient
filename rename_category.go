package cookbook

import (
	"context"
	"fmt"
	"net/http"

	clone "github.com/huandu/go-clone/generic"
	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
)

func (receiver *categories) Rename(ctx context.Context, category *model.Category, newName string) (*model.Category, error) {
	fetchedName, err := request[string](
		ctx,
		receiver.httpClient,
		http.MethodPut,
		urlWithPath(receiver.url, fmt.Sprintf("category/%s", category)),
		map[string]string{
			"name": newName,
		},
		receiver.username, receiver.password,
	)
	if err != nil {
		return nil, fmt.Errorf("failed renaming category: %w", err)
	}

	cloned := clone.Clone(category)
	cloned.Name = fetchedName

	return cloned, nil
}

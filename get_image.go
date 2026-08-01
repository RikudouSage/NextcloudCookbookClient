package cookbook

import (
	"context"
	"fmt"
	"net/http"
)

type ImageSize string

const (
	ImageSizeFull       ImageSize = "full"
	ImageSizeThumb      ImageSize = "thumb"
	ImageSizeThumbSmall ImageSize = "thumb16"
)

func (receiver *recipes) Image(ctx context.Context, ID string, size ImageSize) ([]byte, error) {
	return request[[]byte](
		ctx,
		receiver.httpClient,
		http.MethodGet,
		urlWithPath(receiver.url, fmt.Sprintf("/recipes/%s/image", ID)),
		nil,
		receiver.username, receiver.password,
	)
}

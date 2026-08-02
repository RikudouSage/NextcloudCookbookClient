package cookbook

import (
	"fmt"
	"net/url"
	"strings"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/internal/helper"
)

func urlWithPath(uri *url.URL, path string) *url.URL {
	const prefix = "/apps/cookbook/api/v1"
	path = "/" + strings.TrimPrefix(path, "/")
	if !strings.HasPrefix(path, prefix) {
		path = fmt.Sprintf("%s/%s", prefix, strings.TrimPrefix(path, "/"))
	}

	clone := new(*uri)
	clone.Path = helper.JoinURLPath(uri.Path, path)
	clone.RawPath = ""

	return clone
}

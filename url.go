package cookbook

import (
	"fmt"
	"net/url"
	"strings"
)

func urlWithPath(uri *url.URL, path string) *url.URL {
	const prefix = "/apps/cookbook/api/v1"
	if !strings.HasPrefix(path, prefix) {
		if strings.HasPrefix(path, "/") {
			path = path[1:]
		}
		path = fmt.Sprintf("%s/%s", prefix, path)
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	clone := new(*uri)
	clone.Path = path

	return clone
}

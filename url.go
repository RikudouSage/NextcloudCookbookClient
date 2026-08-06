package cookbook

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	clone "github.com/huandu/go-clone/generic"
	"go.chrastecky.dev/nextcloud-cookbook/cookbook/internal/helper"
)

const cookbookAPIPrefix = "/apps/cookbook/api/v1"
const checkedMarker = "checked-normalization"

func urlWithPath(uri *url.URL, path string) *url.URL {
	path = "/" + strings.TrimPrefix(path, "/")
	if !strings.HasPrefix(path, cookbookAPIPrefix) {
		path = fmt.Sprintf("%s/%s", cookbookAPIPrefix, strings.TrimPrefix(path, "/"))
	}

	cloned := new(*uri)
	cloned.Path = helper.JoinURLPath(uri.Path, path)
	cloned.RawPath = ""

	return cloned
}

func validateAndNormalizeNextcloudURL(ctx context.Context, httpClient *http.Client, uri *url.URL) error {
	getStatusCode := func(url *url.URL) (int, error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
		if err != nil {
			return 0, err
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			return 0, err
		}
		defer func() {
			_, _ = io.Copy(io.Discard, resp.Body)
			_ = resp.Body.Close()
		}()

		return resp.StatusCode, nil
	}

	if strings.Contains("/"+strings.Trim(uri.Path, "/")+"/", "/index.php/") {
		return nil
	}

	if uri.Fragment == checkedMarker {
		return nil
	}

	uri.Fragment = checkedMarker
	apiPrefixIndex := strings.Index(uri.Path, cookbookAPIPrefix)
	basePath, requestPath := uri.Path, ""
	if apiPrefixIndex != -1 {
		basePath = strings.TrimRight(uri.Path[:apiPrefixIndex], "/")
		requestPath = uri.Path[apiPrefixIndex:]
	}

	urlClone := clone.Clone(uri)
	urlClone.Fragment = ""
	urlClone.RawQuery = ""
	urlClone.Path = helper.JoinURLPath(basePath, cookbookAPIPrefix+"/config")

	statusCode, err := getStatusCode(urlClone)
	if err != nil {
		return err
	}

	if statusCode != http.StatusNotFound {
		return nil
	}

	urlClone.Path = helper.JoinURLPath(basePath, "/index.php"+cookbookAPIPrefix+"/config")
	statusCode, err = getStatusCode(urlClone)
	if err != nil {
		return err
	}

	if statusCode != http.StatusNotFound {
		uri.Path = helper.JoinURLPath(basePath, "/index.php"+requestPath)
	}

	return nil
}

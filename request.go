package cookbook

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	clone "github.com/huandu/go-clone/generic"
)

func request[TResponse any](
	ctx context.Context,
	httpClient *http.Client,
	method string,
	url *url.URL,
	body any,
	username, password string,
) (TResponse, error) {
	isDebug := os.Getenv("REQUEST_DEBUG") == "true"

	var requestBody io.Reader
	var out TResponse

	if err := validateAndNormalizeNextcloudURL(ctx, httpClient, url); err != nil {
		return out, fmt.Errorf("failed normalizing request url: %w", err)
	}

	if body != nil {
		var processedBody []byte
		if strBody, ok := body.(string); ok {
			processedBody = []byte(strBody)
		} else if bytesBody, ok := body.([]byte); ok {
			processedBody = bytesBody
		} else {
			var err error
			processedBody, err = json.Marshal(body)
			if err != nil {
				return out, fmt.Errorf("failed marshalling body to json: %w", err)
			}
		}
		requestBody = bytes.NewReader(processedBody)
	}

	newUrl := clone.Clone(url)
	newUrl.Fragment = ""

	req, err := http.NewRequestWithContext(ctx, method, newUrl.String(), requestBody)
	if err != nil {
		return out, fmt.Errorf("failed creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(username, password)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return out, fmt.Errorf("failed sending request: %w", err)
	}
	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return out, fmt.Errorf("failed reading response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if isDebug {
			return out, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, responseBody)
		}
		return out, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if _, ok := any(out).([]byte); ok {
		return any(responseBody).(TResponse), nil
	}

	if err = json.Unmarshal(responseBody, &out); err != nil {
		if errors.Is(err, io.EOF) {
			return out, nil
		}
		return out, fmt.Errorf("failed decoding response: %w", err)
	}

	return out, nil
}

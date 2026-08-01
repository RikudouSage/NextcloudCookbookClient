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
)

func request[TResponse any](
	ctx context.Context,
	httpClient *http.Client,
	method string,
	url *url.URL,
	body any,
	username, password string,
) (TResponse, error) {
	var requestBody io.Reader
	var out TResponse

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

	req, err := http.NewRequestWithContext(ctx, method, url.String(), requestBody)
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

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return out, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return out, fmt.Errorf("failed reading response body: %w", err)
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

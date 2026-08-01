package cookbook

import (
	"net/http"
	"net/url"
)

type Option func(clientInstance *client) error

func WithURL(urlStr string) Option {
	return func(clientInstance *client) error {
		var err error
		clientInstance.url, err = url.Parse(urlStr)
		return err
	}
}

func WithUsername(username string) Option {
	return func(clientInstance *client) error {
		clientInstance.username = username
		return nil
	}
}

func WithPassword(password string) Option {
	return func(clientInstance *client) error {
		clientInstance.password = password
		return nil
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(clientInstance *client) error {
		clientInstance.httpClient = httpClient
		return nil
	}
}

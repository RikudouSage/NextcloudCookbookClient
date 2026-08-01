package cookbook

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Client interface {
	Recipes() Recipes
	Categories() Categories
	Tags() Tags
	Misc() Misc
}

type client struct {
	url      *url.URL
	username string
	password string

	recipes    *recipes
	categories *categories
	misc       *misc
	tags       *tags

	httpClient *http.Client
}

func NewClient(options ...Option) (Client, error) {
	instance := &client{}

	for _, option := range options {
		if err := option(instance); err != nil {
			return nil, fmt.Errorf("failed applying option: %w", err)
		}
	}

	if err := instance.defaultsAndValidate(); err != nil {
		return nil, fmt.Errorf("there were validation errors: %w", err)
	}

	return instance, nil
}

func (receiver *client) Recipes() Recipes {
	if receiver.recipes == nil {
		receiver.recipes = &recipes{
			username:   receiver.username,
			password:   receiver.password,
			url:        receiver.url,
			httpClient: receiver.httpClient,
		}
	}

	return receiver.recipes
}

func (receiver *client) Categories() Categories {
	if receiver.categories == nil {
		receiver.categories = &categories{
			username:   receiver.username,
			password:   receiver.password,
			url:        receiver.url,
			httpClient: receiver.httpClient,
		}
	}

	return receiver.categories
}

func (receiver *client) Tags() Tags {
	if receiver.tags == nil {
		receiver.tags = &tags{
			username:   receiver.username,
			password:   receiver.password,
			url:        receiver.url,
			httpClient: receiver.httpClient,
		}
	}

	return receiver.tags
}

func (receiver *client) Misc() Misc {
	if receiver.misc == nil {
		receiver.misc = &misc{
			username:   receiver.username,
			password:   receiver.password,
			url:        receiver.url,
			httpClient: receiver.httpClient,
		}
	}

	return receiver.misc
}

func (receiver *client) defaultsAndValidate() error {
	errs := make([]error, 0)

	if receiver.httpClient == nil {
		receiver.httpClient = http.DefaultClient
	}

	if receiver.url == nil {
		errs = append(errs, errors.New("the Nextcloud instance URL is required"))
	}
	if receiver.username == "" {
		errs = append(errs, errors.New("the username is required"))
	}
	if receiver.password == "" {
		errs = append(errs, errors.New("the password is required"))
	}

	return errors.Join(errs...)
}

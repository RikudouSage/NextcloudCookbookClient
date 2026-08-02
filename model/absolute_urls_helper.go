package model

import (
	"fmt"
	"net/url"
	"reflect"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook/internal/helper"
	"go.chrastecky.dev/nextcloud-cookbook/cookbook/types"
)

func populateAbsoluteUrls(in any, baseURL *url.URL) error {
	ref := reflect.TypeOf(in)
	if ref.Kind() != reflect.Pointer && ref.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("the value must be struct pointer, got %T", in)
	}

	for field := range ref.Elem().Fields() {
		if field.Type.Kind() != reflect.Pointer || field.Type.Elem().Kind() != reflect.Struct {
			continue
		}
		if field.Type.Elem().Name() != "APIURL" {
			continue
		}

		val := reflect.ValueOf(in).Elem().FieldByName(field.Name)
		if val.IsNil() {
			continue
		}

		apiUrl := val.Interface().(*types.APIURL)
		isRelativeURL := apiUrl.Scheme == "" && apiUrl.Host == ""
		if apiUrl.Scheme == "" {
			apiUrl.Scheme = baseURL.Scheme
		}
		if apiUrl.Host == "" {
			apiUrl.Host = baseURL.Host
		}
		if isRelativeURL && baseURL.Path != "" {
			apiUrl.Path = helper.JoinURLPath(baseURL.Path, apiUrl.Path)
			apiUrl.RawPath = ""
		}
	}

	return nil
}

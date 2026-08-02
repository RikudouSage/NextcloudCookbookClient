package main

/*
#include "cb_common.h"
#include <stdbool.h>

typedef struct {
	const char* url;
	const char* username;
	const char* password;
} NewClientOptions;
*/
import "C"
import "go.chrastecky.dev/nextcloud-cookbook/cookbook"

//export CookbookNewClient
func CookbookNewClient(outHandle *C.ClientHandle, options C.NewClientOptions) C.CookbookResult {
	if outHandle == nil {
		setLastError(nullPointerError("outHandle"))
		return CookbookError
	}

	goOptions := make([]cookbook.Option, 0)
	if options.url != nil {
		goOptions = append(goOptions, cookbook.WithURL(C.GoString(options.url)))
	}
	if options.username != nil {
		goOptions = append(goOptions, cookbook.WithUsername(C.GoString(options.username)))
	}
	if options.password != nil {
		goOptions = append(goOptions, cookbook.WithPassword(C.GoString(options.password)))
	}

	client, err := cookbook.NewClient(goOptions...)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	result := registerHandle(client)
	*outHandle = C.ClientHandle(result)

	clearLastError()
	return CookbookSuccess
}

//export CookbookValidateCredentials
func CookbookValidateCredentials(ctx C.ContextHandle, client C.ClientHandle, outValid *C.bool) C.CookbookResult {
	if outValid == nil {
		setLastError(nullPointerError("outValid"))
		return CookbookError
	}

	ctxGo, clientGo, err := getContextAndClient(ctx, client)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	valid, err := clientGo.ValidateCredentials(ctxGo)
	if err != nil {
		setLastError(err)
	}

	*outValid = C.bool(valid)
	return CookbookSuccess
}

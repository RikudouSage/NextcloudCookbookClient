package main

/*
#include "cb_common.h"
*/
import "C"
import (
	"context"
	"fmt"

	"go.chrastecky.dev/nextcloud-cookbook/cookbook"
)

func getClient(clientHandle C.ClientHandle) (cookbook.Client, error) {
	client, err := getHandleObj[cookbook.Client](handle(clientHandle))
	if err != nil {
		return nil, fmt.Errorf("failed getting client out of the handle: %w", err)
	}

	return client, nil
}

func getContext(ctxHandle C.ContextHandle) (context.Context, error) {
	ctx, err := getHandleObj[contextHandle](handle(ctxHandle))
	if err != nil {
		return nil, fmt.Errorf("failed getting context out of the handle: %w", err)
	}

	return ctx.ctx, nil
}

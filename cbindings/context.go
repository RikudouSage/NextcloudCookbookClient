package main

/*
#include "cb_common.h"
*/
import "C"
import (
	"context"
	"time"
)

type contextHandle struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (receiver *contextHandle) Close() error {
	receiver.cancel()
	return nil
}

//export CookbookNewContext
func CookbookNewContext(outContext *C.ContextHandle) C.CookbookResult {
	if outContext == nil {
		setLastError(nullPointerError("outContext"))
		return CookbookError
	}

	ctx, cancel := context.WithCancel(context.Background())
	ctxHandle := &contextHandle{ctx: ctx, cancel: cancel}
	outHandle := registerHandle(ctxHandle)

	*outContext = C.ContextHandle(outHandle)

	clearLastError()
	return CookbookSuccess
}

//export CookbookNewTimeoutContext
func CookbookNewTimeoutContext(outContext *C.ContextHandle, timeoutMS C.int64_t) C.CookbookResult {
	if outContext == nil {
		setLastError(nullPointerError("outContext"))
		return CookbookError
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutMS)*time.Millisecond)
	ctxHandle := &contextHandle{ctx: ctx, cancel: cancel}
	outHandle := registerHandle(ctxHandle)

	*outContext = C.ContextHandle(outHandle)

	clearLastError()
	return CookbookSuccess
}

//export CookbookCancelContext
func CookbookCancelContext(handleRef C.ContextHandle) C.CookbookResult {
	ctxHandle, err := getHandleObj[*contextHandle](handle(handleRef))
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	ctxHandle.cancel()
	clearLastError()

	return CookbookSuccess
}

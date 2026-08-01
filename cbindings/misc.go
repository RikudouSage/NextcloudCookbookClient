package main

/*
#include "cb_common.h"
*/
import "C"

//export CookbookReindex
func CookbookReindex(ctx C.ContextHandle, client C.ClientHandle) C.CookbookResult {
	ctxGo, clientGo, err := getContextAndClient(ctx, client)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	if err := clientGo.Misc().Reindex(ctxGo); err != nil {
		setLastError(err)
		return CookbookError
	}

	clearLastError()
	return CookbookSuccess
}

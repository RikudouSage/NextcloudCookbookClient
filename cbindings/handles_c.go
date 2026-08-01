package main

/*
#include "cb_common.h"
*/
import "C"

//export CookbookCloseHandle
func CookbookCloseHandle(handleID C.Handle) C.CookbookResult {
	err := unregisterHandle(handle(handleID))
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	clearLastError()
	return CookbookSuccess
}

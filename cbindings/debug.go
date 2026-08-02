package main

/*
#include "cb_common.h"
#include <stdbool.h>
*/
import "C"
import (
	"os"
	"strconv"
)

//export CookbookSetDebug
func CookbookSetDebug(enabled C.bool) C.CookbookResult {
	enabledGo := strconv.FormatBool(bool(enabled))
	err := os.Setenv("REQUEST_DEBUG", enabledGo)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	clearLastError()
	return CookbookSuccess
}

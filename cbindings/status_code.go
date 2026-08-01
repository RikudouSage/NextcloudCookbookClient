package main

/*
#include "cb_common.h"
*/
import "C"

const (
	// CookbookSuccess indicates that a C API call completed successfully.
	CookbookSuccess C.CookbookResult = iota
	// CookbookError indicates that a C API call failed and last error is available.
	CookbookError
)

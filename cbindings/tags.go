package main

/*
#include "cb_common.h"
#include "cb_recipes.h"
#include "cb_taxonomy.h"
*/
import "C"

//export CookbookListKeywords
func CookbookListKeywords(ctx C.ContextHandle, client C.ClientHandle, out *C.CookbookKeywordSlice) C.CookbookResult {
	if out == nil {
		setLastError(nullPointerError("out"))
		return CookbookError
	}

	ctxGo, clientGo, err := getContextAndClient(ctx, client)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	keywords, err := clientGo.Tags().List(ctxGo)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	keywordSliceIntoC(out, keywords)
	clearLastError()
	return CookbookSuccess
}

//export CookbookGetKeywordRecipes
func CookbookGetKeywordRecipes(ctx C.ContextHandle, client C.ClientHandle, keywords C.StringSlice, out *C.CookbookRecipeStubSlice) C.CookbookResult {
	if out == nil {
		setLastError(nullPointerError("out"))
		return CookbookError
	}

	ctxGo, clientGo, err := getContextAndClient(ctx, client)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	recipes, err := clientGo.Tags().Recipes(ctxGo, goStringSliceFromC(keywords))
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	recipeStubSliceIntoC(out, recipes)
	clearLastError()
	return CookbookSuccess
}

//export CookbookFreeKeywordSlice
func CookbookFreeKeywordSlice(keywords *C.CookbookKeywordSlice) {
	freeKeywordSlice(keywords)
}

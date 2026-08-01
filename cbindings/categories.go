package main

/*
#include "cb_common.h"
#include "cb_recipes.h"
#include "cb_taxonomy.h"
*/
import "C"

//export CookbookListCategories
func CookbookListCategories(ctx C.ContextHandle, client C.ClientHandle, out *C.CookbookCategorySlice) C.CookbookResult {
	if out == nil {
		setLastError(nullPointerError("out"))
		return CookbookError
	}

	ctxGo, clientGo, err := getContextAndClient(ctx, client)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	categories, err := clientGo.Categories().List(ctxGo)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	categorySliceIntoC(out, categories)
	clearLastError()
	return CookbookSuccess
}

//export CookbookGetCategoryRecipes
func CookbookGetCategoryRecipes(ctx C.ContextHandle, client C.ClientHandle, category *C.CookbookCategory, out *C.CookbookRecipeStubSlice) C.CookbookResult {
	if category == nil {
		setLastError(nullPointerError("category"))
		return CookbookError
	}
	if out == nil {
		setLastError(nullPointerError("out"))
		return CookbookError
	}

	ctxGo, clientGo, err := getContextAndClient(ctx, client)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	recipes, err := clientGo.Categories().Recipes(ctxGo, categoryFromC(category))
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	recipeStubSliceIntoC(out, recipes)
	clearLastError()
	return CookbookSuccess
}

//export CookbookRenameCategory
func CookbookRenameCategory(ctx C.ContextHandle, client C.ClientHandle, category *C.CookbookCategory, newName *C.char, out *C.CookbookCategory) C.CookbookResult {
	if category == nil {
		setLastError(nullPointerError("category"))
		return CookbookError
	}
	if newName == nil {
		setLastError(nullPointerError("newName"))
		return CookbookError
	}
	if out == nil {
		setLastError(nullPointerError("out"))
		return CookbookError
	}

	ctxGo, clientGo, err := getContextAndClient(ctx, client)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	renamed, err := clientGo.Categories().Rename(ctxGo, categoryFromC(category), C.GoString(newName))
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	categoryIntoC(out, renamed)
	clearLastError()
	return CookbookSuccess
}

//export CookbookFreeCategorySlice
func CookbookFreeCategorySlice(categories *C.CookbookCategorySlice) {
	freeCategorySlice(categories)
}

//export CookbookFreeCategory
func CookbookFreeCategory(category *C.CookbookCategory) {
	freeCategory(category)
}

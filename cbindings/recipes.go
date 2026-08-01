package main

/*
#include "cb_common.h"
#include "cb_recipes.h"
*/
import "C"
import "go.chrastecky.dev/nextcloud-cookbook/cookbook"

//export CookbookListRecipes
func CookbookListRecipes(ctx C.ContextHandle, client C.ClientHandle, out *C.CookbookRecipeStubSlice) C.CookbookResult {
	if out == nil {
		setLastError(nullPointerError("out"))
		return CookbookError
	}

	ctxGo, clientGo, err := getContextAndClient(ctx, client)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	recipes, err := clientGo.Recipes().List(ctxGo)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	recipeStubSliceIntoC(out, recipes)
	clearLastError()
	return CookbookSuccess
}

//export CookbookSearchRecipes
func CookbookSearchRecipes(ctx C.ContextHandle, client C.ClientHandle, query *C.char, out *C.CookbookRecipeStubSlice) C.CookbookResult {
	if query == nil {
		setLastError(nullPointerError("query"))
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

	recipes, err := clientGo.Recipes().Search(ctxGo, C.GoString(query))
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	recipeStubSliceIntoC(out, recipes)
	clearLastError()
	return CookbookSuccess
}

//export CookbookGetRecipe
func CookbookGetRecipe(ctx C.ContextHandle, client C.ClientHandle, id *C.char, out *C.CookbookRecipe) C.CookbookResult {
	if id == nil {
		setLastError(nullPointerError("id"))
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

	recipe, err := clientGo.Recipes().Get(ctxGo, C.GoString(id))
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	recipeIntoC(out, recipe)
	clearLastError()
	return CookbookSuccess
}

//export CookbookImportRecipe
func CookbookImportRecipe(ctx C.ContextHandle, client C.ClientHandle, url *C.char, out *C.CookbookRecipe) C.CookbookResult {
	if url == nil {
		setLastError(nullPointerError("url"))
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

	recipe, err := clientGo.Recipes().Import(ctxGo, C.GoString(url))
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	recipeIntoC(out, recipe)
	clearLastError()
	return CookbookSuccess
}

//export CookbookCreateRecipe
func CookbookCreateRecipe(ctx C.ContextHandle, client C.ClientHandle, recipe *C.CookbookRecipe, out *C.CookbookRecipe) C.CookbookResult {
	if recipe == nil {
		setLastError(nullPointerError("recipe"))
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

	created, err := clientGo.Recipes().Create(ctxGo, recipeFromC(recipe))
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	recipeIntoC(out, created)
	clearLastError()
	return CookbookSuccess
}

//export CookbookUpdateRecipe
func CookbookUpdateRecipe(ctx C.ContextHandle, client C.ClientHandle, recipe *C.CookbookRecipe, out *C.CookbookRecipe) C.CookbookResult {
	if recipe == nil {
		setLastError(nullPointerError("recipe"))
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

	updated, err := clientGo.Recipes().Update(ctxGo, recipeFromC(recipe))
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	recipeIntoC(out, updated)
	clearLastError()
	return CookbookSuccess
}

//export CookbookDeleteRecipe
func CookbookDeleteRecipe(ctx C.ContextHandle, client C.ClientHandle, id *C.char) C.CookbookResult {
	if id == nil {
		setLastError(nullPointerError("id"))
		return CookbookError
	}

	ctxGo, clientGo, err := getContextAndClient(ctx, client)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	if err := clientGo.Recipes().Delete(ctxGo, C.GoString(id)); err != nil {
		setLastError(err)
		return CookbookError
	}

	clearLastError()
	return CookbookSuccess
}

//export CookbookGetRecipeImage
func CookbookGetRecipeImage(ctx C.ContextHandle, client C.ClientHandle, id *C.char, size *C.char, out *C.ByteSlice) C.CookbookResult {
	if id == nil {
		setLastError(nullPointerError("id"))
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

	sizeGo := cookbook.ImageSizeFull
	if size != nil {
		sizeGo = cookbook.ImageSize(C.GoString(size))
	}

	image, err := clientGo.Recipes().Image(ctxGo, C.GoString(id), sizeGo)
	if err != nil {
		setLastError(err)
		return CookbookError
	}

	byteSliceIntoC(out, image)
	clearLastError()
	return CookbookSuccess
}

//export CookbookFreeRecipeStubSlice
func CookbookFreeRecipeStubSlice(recipes *C.CookbookRecipeStubSlice) {
	freeRecipeStubSlice(recipes)
}

//export CookbookFreeRecipe
func CookbookFreeRecipe(recipe *C.CookbookRecipe) {
	freeRecipe(recipe)
}

//export CookbookFreeBytes
func CookbookFreeBytes(bytes *C.ByteSlice) {
	freeByteSlice(bytes)
}

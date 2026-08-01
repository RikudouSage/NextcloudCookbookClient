package main

/*
#include "cb_common.h"
#include "cb_recipes.h"
#include "cb_taxonomy.h"
#include <stdlib.h>
*/
import "C"
import (
	"net/url"
	"time"
	"unsafe"

	"github.com/sosodev/duration"
	"go.chrastecky.dev/nextcloud-cookbook/cookbook/model"
	"go.chrastecky.dev/nextcloud-cookbook/cookbook/types"
)

func cUnixMillis(value time.Time) C.int64_t {
	if value.IsZero() {
		return 0
	}
	return C.int64_t(value.UnixMilli())
}

func goTimeFromCUnixMillis(value C.int64_t) time.Time {
	if value == 0 {
		return time.Time{}
	}
	return time.UnixMilli(int64(value))
}

func cURLPtr(value *types.APIURL) unsafe.Pointer {
	if value == nil {
		return nil
	}
	return cStringPtr(value.String())
}

func goURLFromCPtr(value *C.char) *types.APIURL {
	if value == nil {
		return nil
	}

	parsed, err := url.Parse(C.GoString(value))
	if err != nil {
		return nil
	}
	return &types.APIURL{URL: *parsed}
}

func cDurationPtr(value *duration.Duration) unsafe.Pointer {
	if value == nil {
		return nil
	}
	return cStringPtr(value.String())
}

func goDurationFromCPtr(value *C.char) *duration.Duration {
	if value == nil {
		return nil
	}

	parsed, err := duration.Parse(C.GoString(value))
	if err != nil {
		return nil
	}
	return parsed
}

func cStringSliceParts(values []string) (unsafe.Pointer, C.size_t) {
	if len(values) == 0 {
		return nil, 0
	}

	items := (**C.char)(C.malloc(C.size_t(len(values)) * C.size_t(unsafe.Sizeof((*C.char)(nil)))))
	out := unsafe.Slice(items, len(values))
	for i, value := range values {
		out[i] = C.CString(value)
	}

	return unsafe.Pointer(items), C.size_t(len(values))
}

func cStringSliceIntoC(out *C.StringSlice, values []string) {
	if len(values) == 0 {
		clearC(out)
		return
	}

	items, length := cStringSliceParts(values)
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.items), items)
	putCValue(unsafe.Pointer(out), unsafe.Offsetof(out.len), length)
}

func goStringSliceFromC(value C.StringSlice) []string {
	if value.items == nil || value.len == 0 {
		return nil
	}

	items := unsafe.Slice(value.items, int(value.len))
	out := make([]string, len(items))
	for i, item := range items {
		if item != nil {
			out[i] = C.GoString(item)
		}
	}

	return out
}

func freeStringSlice(value *C.StringSlice) {
	if value == nil {
		return
	}

	items := unsafe.Slice(value.items, int(value.len))
	for _, item := range items {
		C.free(unsafe.Pointer(item))
	}
	C.free(unsafe.Pointer(value.items))
	clearC(value)
}

func recipeStubIntoC(out *C.CookbookRecipeStub, recipe *model.RecipeStub) {
	if recipe == nil {
		clearC(out)
		return
	}

	keywordsItems, keywordsLen := cStringSliceParts([]string(recipe.Keywords))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.id), cStringPtr(recipe.ID))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.name), cStringPtr(recipe.Name))
	putCSlice(unsafe.Pointer(out), unsafe.Offsetof(out.keywords), unsafe.Offsetof(out.keywords.items), unsafe.Offsetof(out.keywords.len), keywordsItems, keywordsLen)
	putCValue(unsafe.Pointer(out), unsafe.Offsetof(out.createdDate), cUnixMillis(recipe.CreatedDate.Time))
	putCValue(unsafe.Pointer(out), unsafe.Offsetof(out.modifiedDate), cUnixMillis(recipe.ModifiedDate.Time))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.imageUrl), cURLPtr(recipe.ImageURL))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.imagePlaceholderUrl), cURLPtr(recipe.ImagePlaceholderURL))
}

func recipeStubSliceIntoC(out *C.CookbookRecipeStubSlice, recipes []*model.RecipeStub) {
	if len(recipes) == 0 {
		clearC(out)
		return
	}

	items := (*C.CookbookRecipeStub)(C.malloc(C.size_t(len(recipes)) * C.size_t(unsafe.Sizeof(C.CookbookRecipeStub{}))))
	slice := unsafe.Slice(items, len(recipes))
	for i, recipe := range recipes {
		recipeStubIntoC(&slice[i], recipe)
	}

	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.items), unsafe.Pointer(items))
	putCValue(unsafe.Pointer(out), unsafe.Offsetof(out.len), C.size_t(len(recipes)))
}

func freeRecipeStub(value *C.CookbookRecipeStub) {
	if value == nil {
		return
	}

	C.free(unsafe.Pointer(value.id))
	C.free(unsafe.Pointer(value.name))
	freeStringSlice(&value.keywords)
	C.free(unsafe.Pointer(value.imageUrl))
	C.free(unsafe.Pointer(value.imagePlaceholderUrl))
	clearC(value)
}

func freeRecipeStubSlice(value *C.CookbookRecipeStubSlice) {
	if value == nil {
		return
	}

	items := unsafe.Slice(value.items, int(value.len))
	for i := range items {
		freeRecipeStub(&items[i])
	}
	C.free(unsafe.Pointer(value.items))
	clearC(value)
}

func nutritionIntoC(out *C.CookbookNutrition, value model.Nutrition) {
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out._type), cStringPtr(value.Type))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.calories), cStringPtr(value.Calories))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.carbohydrateContent), cStringPtr(value.CarbohydrateContent))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.cholesterolContent), cStringPtr(value.CholesterolContent))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.fatContent), cStringPtr(value.FatContent))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.fiberContent), cStringPtr(value.FiberContent))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.proteinContent), cStringPtr(value.ProteinContent))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.saturatedFatContent), cStringPtr(value.SaturatedFatContent))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.servingSize), cStringPtr(value.ServingSize))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.sodiumContent), cStringPtr(value.SodiumContent))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.sugarContent), cStringPtr(value.SugarContent))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.transFatContent), cStringPtr(value.TransFatContent))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.unsaturatedFatContent), cStringPtr(value.UnsaturatedFatContent))
}

func nutritionFromC(value C.CookbookNutrition) model.Nutrition {
	return model.Nutrition{
		Type:                  goStringValueFromCPtr(value._type),
		Calories:              goStringValueFromCPtr(value.calories),
		CarbohydrateContent:   goStringValueFromCPtr(value.carbohydrateContent),
		CholesterolContent:    goStringValueFromCPtr(value.cholesterolContent),
		FatContent:            goStringValueFromCPtr(value.fatContent),
		FiberContent:          goStringValueFromCPtr(value.fiberContent),
		ProteinContent:        goStringValueFromCPtr(value.proteinContent),
		SaturatedFatContent:   goStringValueFromCPtr(value.saturatedFatContent),
		ServingSize:           goStringValueFromCPtr(value.servingSize),
		SodiumContent:         goStringValueFromCPtr(value.sodiumContent),
		SugarContent:          goStringValueFromCPtr(value.sugarContent),
		TransFatContent:       goStringValueFromCPtr(value.transFatContent),
		UnsaturatedFatContent: goStringValueFromCPtr(value.unsaturatedFatContent),
	}
}

func freeNutrition(value *C.CookbookNutrition) {
	if value == nil {
		return
	}

	C.free(unsafe.Pointer(value._type))
	C.free(unsafe.Pointer(value.calories))
	C.free(unsafe.Pointer(value.carbohydrateContent))
	C.free(unsafe.Pointer(value.cholesterolContent))
	C.free(unsafe.Pointer(value.fatContent))
	C.free(unsafe.Pointer(value.fiberContent))
	C.free(unsafe.Pointer(value.proteinContent))
	C.free(unsafe.Pointer(value.saturatedFatContent))
	C.free(unsafe.Pointer(value.servingSize))
	C.free(unsafe.Pointer(value.sodiumContent))
	C.free(unsafe.Pointer(value.sugarContent))
	C.free(unsafe.Pointer(value.transFatContent))
	C.free(unsafe.Pointer(value.unsaturatedFatContent))
	clearC(value)
}

func recipeIntoC(out *C.CookbookRecipe, recipe *model.Recipe) {
	if recipe == nil {
		clearC(out)
		return
	}

	keywordsItems, keywordsLen := cStringSliceParts([]string(recipe.Keywords))
	toolsItems, toolsLen := cStringSliceParts(recipe.Tools)
	ingredientsItems, ingredientsLen := cStringSliceParts(recipe.Ingredients)
	instructionsItems, instructionsLen := cStringSliceParts(recipe.Instructions)

	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.id), cStringPtr(recipe.ID))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.schemaType), cStringPtr(recipe.Type))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.name), cStringPtr(recipe.Name))
	putCSlice(unsafe.Pointer(out), unsafe.Offsetof(out.keywords), unsafe.Offsetof(out.keywords.items), unsafe.Offsetof(out.keywords.len), keywordsItems, keywordsLen)
	putCValue(unsafe.Pointer(out), unsafe.Offsetof(out.createdDate), cUnixMillis(recipe.CreatedDate.Time))
	putCValue(unsafe.Pointer(out), unsafe.Offsetof(out.modifiedDate), cUnixMillis(recipe.ModifiedDate.Time))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.imageUrl), cURLPtr(recipe.ImageURL))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.imagePlaceholderUrl), cURLPtr(recipe.ImagePlaceholderURL))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.preparationTime), cDurationPtr(recipe.PreparationTime))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.cookTime), cDurationPtr(recipe.CookTime))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.totalTime), cDurationPtr(recipe.TotalTime))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.description), cStringPtr(recipe.Description))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.url), cURLPtr(recipe.URL))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.image), cURLPtr(recipe.ImageURL2))
	putCValue(unsafe.Pointer(out), unsafe.Offsetof(out.servings), C.uint32_t(recipe.Servings))
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.category), cStringPtr(recipe.Category))
	putCSlice(unsafe.Pointer(out), unsafe.Offsetof(out.tools), unsafe.Offsetof(out.tools.items), unsafe.Offsetof(out.tools.len), toolsItems, toolsLen)
	putCSlice(unsafe.Pointer(out), unsafe.Offsetof(out.ingredients), unsafe.Offsetof(out.ingredients.items), unsafe.Offsetof(out.ingredients.len), ingredientsItems, ingredientsLen)
	putCSlice(unsafe.Pointer(out), unsafe.Offsetof(out.instructions), unsafe.Offsetof(out.instructions.items), unsafe.Offsetof(out.instructions.len), instructionsItems, instructionsLen)
	nutritionIntoC(&out.nutrition, recipe.Nutrition)
}

func recipeFromC(recipe *C.CookbookRecipe) *model.Recipe {
	if recipe == nil {
		return nil
	}

	return &model.Recipe{
		ID:                  goStringValueFromCPtr(recipe.id),
		Type:                goStringValueFromCPtr(recipe.schemaType),
		Name:                goStringValueFromCPtr(recipe.name),
		Keywords:            types.CSVSlice(goStringSliceFromC(recipe.keywords)),
		CreatedDate:         types.APITime{Time: goTimeFromCUnixMillis(recipe.createdDate)},
		ModifiedDate:        types.APITime{Time: goTimeFromCUnixMillis(recipe.modifiedDate)},
		ImageURL:            goURLFromCPtr(recipe.imageUrl),
		ImagePlaceholderURL: goURLFromCPtr(recipe.imagePlaceholderUrl),
		PreparationTime:     goDurationFromCPtr(recipe.preparationTime),
		CookTime:            goDurationFromCPtr(recipe.cookTime),
		TotalTime:           goDurationFromCPtr(recipe.totalTime),
		Description:         goStringValueFromCPtr(recipe.description),
		URL:                 goURLFromCPtr(recipe.url),
		ImageURL2:           goURLFromCPtr(recipe.image),
		Servings:            uint(recipe.servings),
		Category:            goStringValueFromCPtr(recipe.category),
		Tools:               goStringSliceFromC(recipe.tools),
		Ingredients:         goStringSliceFromC(recipe.ingredients),
		Instructions:        goStringSliceFromC(recipe.instructions),
		Nutrition:           nutritionFromC(recipe.nutrition),
	}
}

func freeRecipe(value *C.CookbookRecipe) {
	if value == nil {
		return
	}

	C.free(unsafe.Pointer(value.id))
	C.free(unsafe.Pointer(value.schemaType))
	C.free(unsafe.Pointer(value.name))
	freeStringSlice(&value.keywords)
	C.free(unsafe.Pointer(value.imageUrl))
	C.free(unsafe.Pointer(value.imagePlaceholderUrl))
	C.free(unsafe.Pointer(value.preparationTime))
	C.free(unsafe.Pointer(value.cookTime))
	C.free(unsafe.Pointer(value.totalTime))
	C.free(unsafe.Pointer(value.description))
	C.free(unsafe.Pointer(value.url))
	C.free(unsafe.Pointer(value.image))
	C.free(unsafe.Pointer(value.category))
	freeStringSlice(&value.tools)
	freeStringSlice(&value.ingredients)
	freeStringSlice(&value.instructions)
	freeNutrition(&value.nutrition)
	clearC(value)
}

func categoryIntoC(out *C.CookbookCategory, category *model.Category) {
	if category == nil {
		clearC(out)
		return
	}

	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.name), cStringPtr(category.Name))
	putCValue(unsafe.Pointer(out), unsafe.Offsetof(out.recipeCount), C.uint32_t(category.RecipeCount))
}

func categorySliceIntoC(out *C.CookbookCategorySlice, categories []*model.Category) {
	if len(categories) == 0 {
		clearC(out)
		return
	}

	items := (*C.CookbookCategory)(C.malloc(C.size_t(len(categories)) * C.size_t(unsafe.Sizeof(C.CookbookCategory{}))))
	slice := unsafe.Slice(items, len(categories))
	for i, category := range categories {
		categoryIntoC(&slice[i], category)
	}

	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.items), unsafe.Pointer(items))
	putCValue(unsafe.Pointer(out), unsafe.Offsetof(out.len), C.size_t(len(categories)))
}

func categoryFromC(category *C.CookbookCategory) *model.Category {
	if category == nil {
		return nil
	}

	return &model.Category{
		Name:        goStringValueFromCPtr(category.name),
		RecipeCount: uint(category.recipeCount),
	}
}

func freeCategory(value *C.CookbookCategory) {
	if value == nil {
		return
	}

	C.free(unsafe.Pointer(value.name))
	clearC(value)
}

func freeCategorySlice(value *C.CookbookCategorySlice) {
	if value == nil {
		return
	}

	items := unsafe.Slice(value.items, int(value.len))
	for i := range items {
		freeCategory(&items[i])
	}
	C.free(unsafe.Pointer(value.items))
	clearC(value)
}

func keywordIntoC(out *C.CookbookKeyword, keyword *model.Keyword) {
	if keyword == nil {
		clearC(out)
		return
	}

	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.name), cStringPtr(keyword.Name))
	putCValue(unsafe.Pointer(out), unsafe.Offsetof(out.recipeCount), C.uint32_t(keyword.RecipeCount))
}

func keywordSliceIntoC(out *C.CookbookKeywordSlice, keywords []*model.Keyword) {
	if len(keywords) == 0 {
		clearC(out)
		return
	}

	items := (*C.CookbookKeyword)(C.malloc(C.size_t(len(keywords)) * C.size_t(unsafe.Sizeof(C.CookbookKeyword{}))))
	slice := unsafe.Slice(items, len(keywords))
	for i, keyword := range keywords {
		keywordIntoC(&slice[i], keyword)
	}

	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.items), unsafe.Pointer(items))
	putCValue(unsafe.Pointer(out), unsafe.Offsetof(out.len), C.size_t(len(keywords)))
}

func freeKeyword(value *C.CookbookKeyword) {
	if value == nil {
		return
	}

	C.free(unsafe.Pointer(value.name))
	clearC(value)
}

func freeKeywordSlice(value *C.CookbookKeywordSlice) {
	if value == nil {
		return
	}

	items := unsafe.Slice(value.items, int(value.len))
	for i := range items {
		freeKeyword(&items[i])
	}
	C.free(unsafe.Pointer(value.items))
	clearC(value)
}

func byteSliceIntoC(out *C.ByteSlice, values []byte) {
	if len(values) == 0 {
		clearC(out)
		return
	}

	items := (*C.uint8_t)(C.malloc(C.size_t(len(values))))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(items)), len(values)), values)
	putCPtr(unsafe.Pointer(out), unsafe.Offsetof(out.items), unsafe.Pointer(items))
	putCValue(unsafe.Pointer(out), unsafe.Offsetof(out.len), C.size_t(len(values)))
}

func freeByteSlice(value *C.ByteSlice) {
	if value == nil {
		return
	}

	C.free(unsafe.Pointer(value.items))
	clearC(value)
}

func goStringValueFromCPtr(value *C.char) string {
	if value == nil {
		return ""
	}
	return C.GoString(value)
}

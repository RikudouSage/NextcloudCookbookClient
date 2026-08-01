#ifndef COOKBOOK_RECIPES
#define COOKBOOK_RECIPES

#include "cb_common.h"

typedef struct {
    char* type;
    char* calories;
    char* carbohydrateContent;
    char* cholesterolContent;
    char* fatContent;
    char* fiberContent;
    char* proteinContent;
    char* saturatedFatContent;
    char* servingSize;
    char* sodiumContent;
    char* sugarContent;
    char* transFatContent;
    char* unsaturatedFatContent;
} CookbookNutrition;

typedef struct {
    char* id;
    char* name;
    StringSlice keywords;
    int64_t createdDate;
    int64_t modifiedDate;
    char* imageUrl;
    char* imagePlaceholderUrl;
} CookbookRecipeStub;

typedef struct {
    CookbookRecipeStub* items;
    size_t len;
} CookbookRecipeStubSlice;

typedef struct {
    char* id;
    char* schemaType;
    char* name;
    StringSlice keywords;
    int64_t createdDate;
    int64_t modifiedDate;
    char* imageUrl;
    char* imagePlaceholderUrl;
    char* preparationTime;
    char* cookTime;
    char* totalTime;
    char* description;
    char* url;
    char* image;
    uint32_t servings;
    char* category;
    StringSlice tools;
    StringSlice ingredients;
    StringSlice instructions;
    CookbookNutrition nutrition;
} CookbookRecipe;

#endif

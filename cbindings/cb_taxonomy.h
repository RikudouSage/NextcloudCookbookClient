#ifndef COOKBOOK_TAXONOMY
#define COOKBOOK_TAXONOMY

#include "cb_common.h"

typedef struct {
    char* name;
    uint32_t recipeCount;
} CookbookCategory;

typedef struct {
    CookbookCategory* items;
    size_t len;
} CookbookCategorySlice;

typedef struct {
    char* name;
    uint32_t recipeCount;
} CookbookKeyword;

typedef struct {
    CookbookKeyword* items;
    size_t len;
} CookbookKeywordSlice;

#endif

#ifndef COOKBOOK_COMMON
#define COOKBOOK_COMMON

#include <stddef.h>
#include <stdint.h>

typedef int CookbookResult;
typedef uint64_t Handle;
typedef Handle ClientHandle;
typedef Handle ContextHandle;


typedef struct {
    char** items;
    size_t len;
} StringSlice;

enum {
    CookbookSuccess = 0,
    CookbookError = 1,
};

#endif

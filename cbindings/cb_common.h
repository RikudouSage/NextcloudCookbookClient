#ifndef COOKBOOK_COMMON
#define COOKBOOK_COMMON

#include <stddef.h>
#include <stdint.h>

typedef int CookbookResult;
typedef uint64_t Handle;
typedef Handle ClientHandle;

enum {
    CookbookSuccess = 0,
    CookbookError = 1,
};

#endif

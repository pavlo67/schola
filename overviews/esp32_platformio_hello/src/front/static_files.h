#pragma once

#include <cstdint>
#include <stddef.h>

typedef struct {
    const char*    url;
    const char*    contentType;
    const uint8_t* content;
    const size_t   size;
} static_file;
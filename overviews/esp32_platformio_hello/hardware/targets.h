#pragma once

// common constants -----------------------------------------------------

#define LED_MAX_BRIGHTNESS 50 //0..255 for max led brightness


// specific constants and definitions -----------------------------------

#ifdef SUPER_MINI_ESP32_C3
    #include "data/SUPER_MINI_ESP32_C3.h"
    
#else
    // TODO???

#endif


// common definitions gotten from Express LRS ---------------------

#define WORD_ALIGNED_ATTR __attribute__((aligned(4)))
#define WORD_PADDED(size) (((size)+3) & ~3)
#ifndef DMA_ATTR
    #define DMA_ATTR
#endif

#undef  ICACHE_RAM_ATTR //fix to allow both esp32 and esp8266 to use ICACHE_RAM_ATTR for mapping to IRAM
#define ICACHE_RAM_ATTR IRAM_ATTR


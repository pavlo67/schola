#pragma once

#include <Arduino.h>

void    getKeyPress(const char* prompt);
bool    confirm(const char* prompt);
int64_t getInteger(const char* prompt, uint8_t base = 10);
void    showBuffer(uint8_t* buffer, int size);

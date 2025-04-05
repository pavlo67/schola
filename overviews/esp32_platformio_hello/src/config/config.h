# pragma once

#include <Arduino.h>

extern const char* VALUE_KEY;

void        updateConfig(const String& key, const String& value);
const char* getConfigJSON();
void        showConfig();

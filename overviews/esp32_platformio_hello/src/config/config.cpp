#include "config.h"

const char* VALUE_KEY = "value";

static uint32_t configValue = 0;

void updateConfig(const String& key, const String& value) {
	Serial.printf("updateConfig[%s]: %s\n", key.c_str(), value.c_str());

	if (key == VALUE_KEY) {
		configValue = value.toInt();
	} else {
		/* TODO!!! warning */
	}
}

// TODO!!! share this string more accuracy
static String str;

const char* getConfigJSON() {
	str = "{";
	str.concat("\""); str.concat(VALUE_KEY); str.concat("\":\""); str.concat(configValue); str.concat("\"");
	str.concat("}");
	return str.c_str();
}

void showConfig() {
	Serial.printf("%s: %d\n", VALUE_KEY, configValue);
}
#include <Arduino.h>
#include <ESPAsyncWebServer.h>

#include "config/config.h"
#include "static.h"

int testValue = 0;

AsyncWebServer asyncWebServer(80);

static void HandlePage(AsyncWebServerRequest *request) {
	for (size_t i = 0 ; i < FILES_LENGTH; i++) {
		if (request->url().equals(files[i].url)) {

			if (std::string(request->methodToString()) == "POST") {
				int params = request->params();
				for (int i = 0; i < params; i++) {
					AsyncWebParameter* p = request->getParam(i);
					updateConfig(p->name(), p->value());
				}
			}
			AsyncWebServerResponse *response = request->beginResponse_P(200, files[i].contentType, files[i].content, files[i].size);
			request->send(response);
			Serial.printf("served %d bytes to %s\n", files[i].size, files[i].url);
			return;
		}
	}
	request->send(404, "text/plain", "File not found");
	Serial.printf("File not found %s%s\n", request->host().c_str(), request->url().c_str());
}

static_file xhrs[] = {
	{"/config", "application/json", nullptr, 0},
};

size_t XHRS_LENGTH = sizeof(xhrs) / sizeof(static_file);

constexpr int PARAM_LENGTH = 20;
char paramBuffer[PARAM_LENGTH + 1];

static void HandleXHR(AsyncWebServerRequest *request) {
	for (size_t i = 0 ; i < XHRS_LENGTH; i++) {
		if (request->url().equals(xhrs[i].url)) {
			const char* valueBuffer = nullptr;
			if (request->url() == "/config") {
				valueBuffer = getConfigJSON();

			} else {
				// TODO...
			}

			AsyncWebServerResponse *response = request->beginResponse_P(200, xhrs[i].contentType, (const uint8_t*)valueBuffer, strlen(valueBuffer));
			response->addHeader("Access-Control-Allow-Origin", "*");
			response->addHeader("Access-Control-Allow-Headers", "authorization,content-type");
			response->addHeader("Access-Control-Allow-Methods", "HEAD,GET,POST,PUT,DELETE,OPTIONS");
			response->addHeader("Access-Control-Allow-Credentials", "true");
			request->send(response);
			Serial.printf("served %d bytes to %s\n", strlen(valueBuffer), xhrs[i].url);
			return;
		}
	}
	request->send(404, "text/plain", "File not found");
	Serial.printf("File not found %s%s\n", request->host().c_str(), request->url().c_str());
}

void startServer() {
	for (size_t i = 0 ; i < FILES_LENGTH; i++) {
		asyncWebServer.on(files[i].url, HandlePage);
	}
	for (size_t i = 0 ; i < XHRS_LENGTH; i++) {
		asyncWebServer.on(xhrs[i].url, HandleXHR);
	}
	asyncWebServer.onNotFound(HandlePage);
	asyncWebServer.begin();

	Serial.println("HTTPUpdateServer ready!"); // Open http://%s.local in your browser , wifi_hostname
}
#include <Arduino.h>

#include "user_defines.h"

#include <WiFi.h>
#include <ESPmDNS.h>
#include <esp_partition.h>
#include <esp_ota_ops.h>
#include <esp_flash_encrypt.h>
#include <soc/uart_pins.h>

#include <DNSServer.h>
#include "server.h"

volatile WiFiMode_t wifiMode = WIFI_OFF;
const char* wifi_hostname    = "esp32_demo";

static void startMDNS() {
	MDNS.end();
	if (!MDNS.begin(wifi_hostname)) {
		Serial.println("Error starting mDNS");
		return;
	}
	String instance = String(wifi_hostname) + "_" + WiFi.macAddress();
	instance.replace(":", "");

	MDNS.setInstanceName(instance);
	MDNS.addService("http", "tcp", 80);
	MDNS.addServiceTxt("http", "tcp", "vendor", "");
	Serial.printf("mDNS started for %s.local\n\n", wifi_hostname);
}


void startWiFiFront() {
	Serial.print("Setting AP (Access Point) mode...\n");

	WiFi.disconnect();
	wifiMode = WIFI_AP;
	WiFi.setHostname(wifi_hostname); // hostname must be set before the mode is set to STA
	WiFi.mode(wifiMode);
	WiFi.setTxPower(WIFI_POWER_19_5dBm);
	WiFi.softAP(AP_SSID, AP_PASSWORD);
	IPAddress IP = WiFi.softAPIP();
	Serial.print("AP IP address: "); Serial.println(IP);

	startMDNS();
	startServer();
}

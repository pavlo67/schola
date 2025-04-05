#include <Arduino.h>

#include "front/front.h"
#include "front/server.h"

#include "targets.h"

uint32_t LED_DELAY_MS = 1000;
uint32_t ledChangeNextAt;
uint8_t  ledOn        = 0;
bool     frontStarted = false;

void setup() {  
  Serial.begin(MONITOR_SPEED);
  while (!Serial) {}

  delayMicroseconds(2e6); // wait for the connection is established
  Serial.printf("Serial is ok at %d!\n\n", MONITOR_SPEED);

  pinMode(GPIO_PIN_LED, OUTPUT);
  ledChangeNextAt = millis();
}

void loop() {
  if (!frontStarted) {
    startWiFiFront();
    frontStarted = true;
  }
  
  unsigned long at = millis();
  if (at >= ledChangeNextAt) {
    ledChangeNextAt = at + LED_DELAY_MS;
    ledOn     = 1 - ledOn;
    digitalWrite(GPIO_PIN_LED, ledOn);
  }
}

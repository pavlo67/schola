#include "cli.h"

bool confirm(const char* prompt) {
    Serial.printf("%s [y/n]?", prompt);
    int c = -1;
    while (c < 0 && c != 'y' && c != 'Y' && c != 'n' && c != 'N') {
        c = Serial.read(); 
    }
    Serial.println();

    return c == 'y' || c == 'Y';
}    


void getKeyPress(const char* prompt) {
    Serial.printf("%s: waiting for any key...", prompt);
    int c = -1;
    while (c < 0) {
        c = Serial.read(); 
    }
    Serial.read();     // for 2nd symbol in "\r\n" 
    Serial.println();
}    

int64_t getInteger(const char* prompt, uint8_t base) {
    Serial.printf("%s: ", prompt);

    bool negative = false;
    int  digits   = 0;
    int64_t value = 0;
    
    while (true) {
        int c = Serial.read(); 
        if (c < 0) {
            // do nothing
        } else if (c >= '0' && c <= '9') {
            value  = value * base + (c - '0');
            digits++;
            Serial.print((char)c);
        } else if (base > 10 && (c >= 'a' && (c - 'a' < base)) ) {
            value  = value * base + (10 + c - 'a');
            digits++;
            Serial.print((char)c);
        } else if (base > 10 && (c >= 'A' && (c - 'A' < base)) ) {
            value  = value * base + (10 + c - 'A');
            digits++;
            Serial.print((char)c);
        } else if (digits) {
            if (c == 0x08) {  // backspace
                Serial.print('\\');
                value /= base;
                digits--;
            } else {
                Serial.read();   // for 2nd symbol in "\r\n" 
                Serial.println();
                return negative ? -value : value;
            }
        } else if (negative) {
            if (c == 0x08) {  // backspace
                Serial.print('\\');
                negative = false;
            } else {
                Serial.printf("\nwrong char after minus sign: %c\n", c);
            }
        } else if (c == '-') {
            negative = true;
            Serial.print('-');
        }
    }
}

void showBuffer(uint8_t* buffer, int size) {
    for (int i = 0; i < size; i++) {
        Serial.printf(" %02x", buffer[i]);
    }
}    

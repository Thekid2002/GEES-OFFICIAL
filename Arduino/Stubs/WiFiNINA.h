#pragma once

// Minimal stub for WiFiClient used in WiFiSender.h
#define WL_CONNECTED 1

struct {
    int begin(const char* ssid, const char* pass) {
        return WL_CONNECTED;
    }
} WiFi;

class WiFiClient {
public:
    WiFiClient() = default;

    void flush() {}

    int connect(const char* host, uint16_t port) {
        return 1; // pretend connection succeeded
    }

    size_t write(const uint8_t* buffer, size_t size) {
        return size;
    }

    size_t write(uint8_t val) {
        return 1;
    }

    int available() {
        return 0;
    }

    int read() {
        return -1;
    }

    void stop() {}

    bool connected() {
        return true;
    }

    operator bool() const {
        return true;
    }

    void print(const char* str) {}
    void println(const char* str) {}
    void println() {}

    void print(size_t val) {}
    void println(size_t val) {}

    void print(int val) {}
    void println(int val) {}

    void print(long val) {}
    void println(long val) {}
};

#pragma once

#include <cstdint>
#include <cstring>
#include <cmath>

#include <string>
#include <sstream>

// Replace basic Arduino types/macros if needed
typedef uint8_t byte;
typedef unsigned long millis_t;

void delay(unsigned long);
unsigned long millis();
long map(long x, long in_min, long in_max, long out_min, long out_max);

struct {
    void print(const char*) {}
    void println(const char*) {}
    void println() {}
} Serial;

class String {
    public:
        std::string str;
    
        String() = default;
    
        String(const char* s) : str(s) {}
    
        String(float val, int precision = 2) {
            std::ostringstream out;
            out.precision(precision);
            out << std::fixed << val;
            str = out.str();
        }
    
        String operator+(const char* other) const {
            return String((str + other).c_str());
        }
    
        String operator+(const String& other) const {
            return String((str + other.str).c_str());
        }
    
        friend String operator+(const char* lhs, const String& rhs) {
            return String((std::string(lhs) + rhs.str).c_str());
        }
    
        String& operator+=(const String& other) {
            str += other.str;
            return *this;
        }
    
        String& operator+=(const char* other) {
            str += other;
            return *this;
        }
    
        operator const char*() const {
            return str.c_str();
        }
    
        void reserve(size_t size) {}
 
        size_t length() const {
            return str.length();
        }
    
        int indexOf(const char* substr) const {
            size_t pos = str.find(substr);
            return (pos == std::string::npos) ? -1 : static_cast<int>(pos);
        }
    };
    

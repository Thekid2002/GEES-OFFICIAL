#ifndef WiFiSender_h
#define WiFiSender_h

#include <Arduino.h>
#include <WiFiNINA.h>

class WiFiSender {
  public:
    WiFiSender(const char* ssid, const char* pass, const char* server, int port);
    
    void connect();
    void sendData(String path, String &jsonData);  
    void post(String path, String &jsonData);
    String waitForResponse();

  private:
    const char* _ssid;
    const char* _pass;
    const char* _server;
    int _port;
    WiFiClient _client;
    unsigned long _lastPostTime;
    const int _postInterval = 100;

    void connectToWiFi();
    void connectToServer();
};

#endif

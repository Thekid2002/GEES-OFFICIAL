#ifndef RGB_H
#define RGB_H

extern const int RED_PIN;
extern const int GREEN_PIN;
extern const int BLUE_PIN;

void setupLED();
void enableLED(int red, int green, int blue);
void disableLED();
void blinkLED(int red, int green, int blue, int activeTime);

#endif  // RGB_H

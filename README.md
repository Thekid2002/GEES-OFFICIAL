# GEES-OFFICIAL
This is the GEES (Gesture Estimation and Evaluation System) - Official repository.

## Equipment

The equipment required for this project is the following:

| *Device*                | *Specification*                                                |
|-------------------------|----------------------------------------------------------------|
| Arduino                 | Arduino UNO WiFi Rev2 [^1]                                     |
| IMU Module              | LSM6DS3TR module [^1]                                          |
| WiFi & Bluetooth Module | NINA-W102 module from u-Blox [^1][^2]                          |
| Computer                | A modern personal computer                                     |

[^1]: Arduino UNO WiFi Rev2 – [source](https://store.arduino.cc/products/arduino-uno-wifi-rev2)  
[^2]: NINA-W102 – [source](https://www.u-blox.com/en/product/nina-w10-series)

## How To Run
The project is run by first starting docker desktop

Then CD into the Gees_Backend directory

```
cd /Gees_Backend
```

Then run the following command

```
docker compose up
```

A seperate arduino MUST be used and it should be run with the .ino file found in the directory /Arduino/IMU_Reader

Remember to change the .ino to have the corrrect WiFi SSID, Password, and the Correct IP of your running backend. 
## About us
This project was developed by group CS-25-SW6-07 from the Software Bachelor at Aalborg University
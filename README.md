# meteorpi

Raspberry Pi weather station.

Listens and periodically reads data, saves it to a CSV, makes it pretty with Gnuplot.

## Sensors

* BMP180 (locally over I2C)
* Inside DHT22 (locally via a custom AVR serial encoder thingy which isn't really
  needed)
* Outside DHT22 (Solar powered ESP8266, see
  [esp8266-dht22](https://github.com/lucaspiller/esp8266-dht22) for firmware /
  details)

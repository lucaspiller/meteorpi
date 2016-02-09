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

## Pretty things

![screen shot 2016-02-09 at 19 56
38](https://cloud.githubusercontent.com/assets/18404/12927217/49d50400-cf67-11e5-877e-d5e592e5246c.png)
![screen shot 2016-02-09 at 19 56
18](https://cloud.githubusercontent.com/assets/18404/12927218/4aaefbba-cf67-11e5-9dbc-8a349aa85d42.png)

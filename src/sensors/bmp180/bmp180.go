package bmp180

import (
	"fmt"
	"time"

	"github.com/kidoman/embd/host/generic"
	sensor "github.com/kidoman/embd/sensor/bmp180"
	t "github.com/lucaspiller/meteorpi/src/sensors/types"
)

var MeasurementInterval = 30 * time.Second

// start the sensor as a new Gorouting, and return a Channel for it to spit out
// measurements
func Start(data chan *t.Measurement) {
	go start(data)
}

func start(data chan *t.Measurement) {
	// get i2c bus (/dev/i2c-1)
	bus := generic.NewI2CBus(1)
	baro := sensor.New(bus)
	defer baro.Close()

	// set highest sensitivity
	baro.SetOss(3)

	for {
		measure(baro, data)
		time.Sleep(MeasurementInterval)
	}
}

func measure(baro *sensor.BMP180, data chan *t.Measurement) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("bmp180 caught panic while measuring:", r)
		}
	}()

	measurement := t.NewMeasurement("BMP180")
	measurement.AddReading(measureTemperature(baro))
	measurement.AddReading(measurePressure(baro))
	data <- measurement
}

func measureTemperature(baro *sensor.BMP180) (reading t.Reading) {
	temp, err := baro.Temperature()
	if err != nil {
		panic(err)
	}
	return t.Reading{"temperature", temp}
}

func measurePressure(baro *sensor.BMP180) (reading t.Reading) {
	pressure, err := baro.Pressure()
	if err != nil {
		panic(err)
	}
	return t.Reading{"pressure", pressure}
}

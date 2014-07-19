// +build !linux

package random

import "time"
import t "github.com/lucaspiller/meteorpi/src/sensors/types"

var MeasurementInterval = 5 * time.Second

// start the sensor as a new Gorouting, and return a Channel for it to spit out
// measurements
func Start(data chan *t.Measurement) {
	go start(data)
}

// take a measurement every MeasurementInterval forever
func start(data chan *t.Measurement) {
	for {
		measurement := t.NewMeasurement("Random")
		measurement.AddReading(measure())
		data <- measurement

		time.Sleep(MeasurementInterval)
	}
}

// perform the actual measurement, here we just return a dummy value but this
// is where you should query the sensor which should be done in a blocking way
func measure() (reading t.Reading) {
	return t.Reading{"random_reading", "foo"}
}

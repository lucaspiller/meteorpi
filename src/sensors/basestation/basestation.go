// +build linux

package basestation

import (
	"fmt"
	"strconv"
	"strings"

	t "github.com/lucaspiller/meteorpi/src/sensors/types"
)

// start the sensor as a new Gorouting, and return a Channel for it to spit out
// measurements
func Start(data chan *t.Measurement) {
	go start(data)
}

//-----------------------------------------------------------------------------

func start(data chan *t.Measurement) {
	// open serial port
	serial := SerialOpen()
	defer serial.Close()

	// read each line, blocking until we receive more data
	for serial.Scan() {
		line := serial.Text()
		switch {
		case strings.HasPrefix(line, "dht22"):
			createDHT22Measurement(line, data)
		default:
			fmt.Println("Unknown data format:", line)
		}
	}
}

func createDHT22Measurement(line string, data chan *t.Measurement) {
	fields := strings.Split(line, ",")

	if fields[1] == "0" { // status OK
		parsed, _ := strconv.ParseInt(fields[2], 10, 16)
		temperature := float32(parsed) / 10

		parsed, _ = strconv.ParseInt(fields[3], 10, 16)
		humidity := float32(parsed) / 10

		measurement := t.NewMeasurement("DHT22")
		measurement.AddReading(t.Reading{"temperature", temperature})
		measurement.AddReading(t.Reading{"humidity", humidity})
		data <- measurement
	} else {
		fmt.Println("DHT22 error")
	}
}

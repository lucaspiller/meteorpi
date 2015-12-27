package remote

import (
	"bufio"
	"fmt"
	"net"
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
	// listen on all interfaces
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		// accept connection
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		// will listen for message to process ending in newline (\n)
		message, _, _ := bufio.NewReader(conn).ReadLine()
		handleMessage(string(message), data)
	}
}

func handleMessage(message string, data chan *t.Measurement) {
	fields := strings.Split(message, ",")

	if fields[0] == "1" { // DHT22 sensor
		parsed, _ := strconv.ParseInt(fields[1], 10, 16)
		temperature := float32(parsed) / 10

		parsed, _ = strconv.ParseInt(fields[2], 10, 16)
		humidity := float32(parsed) / 10

		parsed, _ = strconv.ParseInt(fields[3], 10, 16)
		vcc := float32(parsed) / 1000

		measurement := t.NewMeasurement("Remote1")
		measurement.AddReading(t.Reading{"temperature", temperature})
		measurement.AddReading(t.Reading{"humidity", humidity})
		measurement.AddReading(t.Reading{"vcc", vcc})
		data <- measurement
	} else {
		fmt.Println("Unhandled data format:", message)
	}
}

package main

import (
	"flag"
	"fmt"

	"github.com/lucaspiller/meteorpi/src/datastore"
	"github.com/lucaspiller/meteorpi/src/sensors/basestation"
	"github.com/lucaspiller/meteorpi/src/sensors/bmp180"
	//"github.com/lucaspiller/meteorpi/src/sensors/random"
	"github.com/lucaspiller/meteorpi/src/sensors/remote"
	t "github.com/lucaspiller/meteorpi/src/sensors/types"
)

func main() {
	run()
}

func run() {
	flag.Parse()

	store := datastore.OpenWriter()
	defer store.CloseWriter()

	data := make(chan *t.Measurement)
	//random.Start(data)
	bmp180.Start(data)
	basestation.Start(data)
	remote.Start(data)
	for {
		select {
		case measurement := <-data:
			fmt.Println("Got data:", measurement)
			store.WriteMeasurement(measurement)
		}
	}
}

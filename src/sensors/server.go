package main

import (
	"flag"
	"fmt"

	bmp180 "github.com/lucaspiller/meteorpi/src/sensors/bmp180"
	"github.com/lucaspiller/meteorpi/src/sensors/random"
	t "github.com/lucaspiller/meteorpi/src/sensors/types"
)

func main() {
	run()
}

func run() {
	flag.Parse()

	data := make(chan *t.Measurement)
	random.Start(data)
	bmp180.Start(data)
	for {
		select {
		case measurement := <-data:
			fmt.Println("Got data:", measurement)
		}
	}
}

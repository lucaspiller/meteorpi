package main

import "fmt"
import "github.com/lucaspiller/meteorpi/sensors/random"
import t "github.com/lucaspiller/meteorpi/sensors/types"

func main() {
	run()
}

func run() {
	data := make(chan *t.Measurement)
	random.Start(data)
	for {
		select {
		case measurement := <-data:
			fmt.Println("Got data: ", measurement)
		}
	}
}

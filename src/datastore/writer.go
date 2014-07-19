package datastore

import (
	"encoding/csv"
	"fmt"
	"os"

	t "github.com/lucaspiller/meteorpi/src/sensors/types"
)

type Writer struct {
	File *os.File
	Csv  *csv.Writer
}

func OpenWriter() *Writer {
	file, err := os.OpenFile("data/data.csv", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	csv := csv.NewWriter(file)

	return &Writer{File: file, Csv: csv}
}

func (w *Writer) WriteMeasurement(measurement *t.Measurement) {
	length := 2 + len(measurement.Readings)
	data := make([]string, length)

	data[0] = fmt.Sprintf("%v", measurement.Time)
	data[1] = measurement.SensorId

	for i := 0; i < len(measurement.Readings); i++ {
		data[i+2] = fmt.Sprintf("%v", measurement.Readings[i].Value)
	}

	err := w.Csv.Write(data)
	if err != nil {
		panic(err)
	}

	w.Csv.Flush()
}

func (w *Writer) CloseWriter() {
	w.File.Close()
}

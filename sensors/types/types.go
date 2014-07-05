package types

import "time"

type Reading struct {
	Type  string
	Value interface{}
}

type Measurement struct {
	SensorId string
	Time     time.Time
	Readings []Reading
}

func NewMeasurement(sensorId string) *Measurement {
	return &Measurement{SensorId: sensorId, Time: time.Now()}
}

func (m *Measurement) AddReading(Reading Reading) {
	m.Readings = append(m.Readings, Reading)
}

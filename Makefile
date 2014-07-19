.PHONY: build build-rpi

build:
	go build -o bin/meteorpi-server src/sensors/server.go

build-rpi:
	GOARCH=arm GOOS=linux go build -o bin/meteorpi-server src/sensors/server.go

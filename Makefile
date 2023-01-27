.EXPORT_ALL_VARIABLES:
GOOS ?= $(uname -s)
GOARCH ?= $(uname -m)
LD_FLAGS := -ldflags="-s -w -X 'main.BuildDate=$(shell date)'"

build:
	mkdir -p bin/
	go build -v -o bin/peep peep.go

run: build
	chmod +x bin/peep
	./bin/peep -d hum

test:
	go test -v ./...
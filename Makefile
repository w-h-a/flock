.PHONY: build test

build:
	go build -o bin/flock ./cmd/flock

test:
	go test ./...

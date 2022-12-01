.PHONY: test lint

lint:
	golangci-lint run

test:
	go test -v -race -timout 30s ./...

.DEFAULT_GOAL := test

.DEFAULT_GOAL := test
test:
	go test -v -race -timout 30s ./...
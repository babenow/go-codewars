.PHONY: test lint

lint:
	golangci-lint run

test:
	go test -v -race -timout 30s ./...

.DEFAULT_GOAL := test

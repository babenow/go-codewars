.PHONY: test

test:
	go test -v -race -timout 30s ./...
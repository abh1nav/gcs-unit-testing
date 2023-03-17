.PHONY: all build

build:
	@go build -o build/giraffe.bin cmd/giraffe/main.go

run:
	@go run cmd/giraffe/main.go

test:
	@go test -race -covermode=atomic ./...
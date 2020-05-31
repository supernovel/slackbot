build:
	@go build -o bin/server ./cmd/main.go

test:
	go test -v

run:
	go run ./cmd/main.go

all: build test
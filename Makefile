build:
	@go build -o bin/server ./cmd/main.go

test:
	go test -v

all: build test
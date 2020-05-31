GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build: clean
	$(GO_BUILD_ENV) go build -o bin/server ./cmd/server.go

test:
	go test -v

run:
	go run ./cmd/server.go

clean:
	rm -rf bin/

all: build test

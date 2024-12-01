all: deps build

deps:
	go mod tidy

build:
	go build -o bin/aoc cmd/main.go

test:
	go test ./...

run:
	bin/aoc

clean:
	go clean
	rm bin/aoc

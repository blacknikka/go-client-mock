all: build

build:
	go build -o app main.go

cover:
	go test -coverprofile=cover.out ./...
	go tool cover -func=cover.out

test:
	go test -v ./...

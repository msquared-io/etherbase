.PHONY: dev reader writer build-all

dev:
	go run cmd/dev/main.go

reader:
	go run cmd/reader/main.go

writer:
	go run cmd/writer/main.go

build-all:
	go build -o bin/reader cmd/reader/main.go
	go build -o bin/writer cmd/writer/main.go 
.PHONY: test test-coverage

test:
	go test -v ./...
	goreportcard-cli -v

test-coverage:
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html
	open cover.html

run-server:
	go run ./cmd/rover_server

run-client:
	open http://127.0.0.1:8080
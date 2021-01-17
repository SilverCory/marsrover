.PHONY: test test-coverage

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html
	open cover.html
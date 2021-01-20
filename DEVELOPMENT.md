Development
===

## Requirements
- Go version 1.15.2+
- GNU Make
- Git
- goreportcard-cli
    - can be installed using `$ go install github.com/gojp/goreportcard/cmd/goreportcard-cli`

## Tests
### Writing tests
Code coverage can be monitored by using the `test-coverage` make command.

Simply run `$ make test-coverage`, a html document will be opened displaying the test coverage.

### Running tests
Tests can be executed by running the `test` make command.

`$ make test`

This will also vet the project and run it against `goreportcard-cli` to bring up any additional issues.

## Code format
The code formatting is standard go formatting and should be done using the `$ go fmt` command.
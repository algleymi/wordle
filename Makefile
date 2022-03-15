.PHONY: get lint test deps
all: get lint test

lint:
	go fmt ./...
get:
	go get -v -t -d ./...
test:
	go test ./... -cover
	go test --bench=./... -cover
deps:
	go list -m -u all

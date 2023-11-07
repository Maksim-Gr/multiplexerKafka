BINARY_NAME=consumer

build:
	GOARCH=amd64 GOS=darwin go build -o build/${BINARY_NAME}-darwin cmd/consumer/main.go
	GOARCH=amd64 GOS=linux go build -o build/${BINARY_NAME}-linux cmd/consumer/main.go
	GOARCH=amd64 GOS=windows go build -o build/${BINARY_NAME}-windows cmd/consumer/main.go

run:
	build/${BINARY_NAME}-darwin

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

dep:
	go mod download
vet:
	go vet

lint:
	golang-lint run --enable-all

test:
	go test -race ./..

test_coverage:
	go test ./... -coverprofile=coverage.out
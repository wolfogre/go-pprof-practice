.PHONY: dep test check build run clean

BINARY_NAME="pprof"

MAIN_GO=main.go

dep:
	go mod download

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

check:
	go fmt ./...
	go vet ./...

build: dep check test
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME}-arm64-darwin ${MAIN_GO}
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-amd64-darwin ${MAIN_GO}
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-amd64-linux ${MAIN_GO}
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-amd64-windows ${MAIN_GO}

run:
	./${BINARY_NAME}-linux

clean:
	go clean
	rm -f ${BINARY_NAME}-arm64-darwin
	rm -f ${BINARY_NAME}-amd64-darwin
	rm -f ${BINARY_NAME}-amd64-linux
	rm -f ${BINARY_NAME}-amd64-windows

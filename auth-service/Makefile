BINARY_NAME=authApp

## build: Build binary
build:
	@echo "Building auth app..."
	env CGO_ENABLED=0  go build -ldflags="-s -w" -o ${BINARY_NAME} ./cmd/api
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	./${BINARY_NAME} &
	@echo "Started"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

## test: runs all tests
.PHONY: test
test:
	go test -v ./cmd/api

BINARY_NAME=logApp
DSN="host=localhost user=root password=root dbname=go_db sslmode=disable timezone=UTC connect_timeout=5"

## build: Build binary
build:
	@echo "Building log app..."
	env CGO_ENABLED=0  go build -ldflags="-s -w" -o ${BINARY_NAME} ./cmd/api
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	@env DSN=${DSN} ./${BINARY_NAME} &
	@echo "Started"

start: run

stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

restart: stop start

## test: runs all tests
.PHONY: test
test:
	go test -v ./cmd/api

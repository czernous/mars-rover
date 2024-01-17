MAIN_PACKAGE_PATH := ./cmd
BINARY_NAME := rover
INPUT_FILE_PATH := ./data/input.txt

build:
	@go build -o=bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

run: build
	@./bin/rover ${INPUT_FILE_PATH}

test:
	@go test -v -race -buildvcs ./...


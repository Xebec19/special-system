BINARY=wc-clone
VERSION=$(shell git describe --tags --abbrev=0 2>/dev/null || echo "dev")
COMMIT=$(shell git rev-parse --short HEAD)
DATE=$(shell date -u +%Y-%m-%d)

help:
	@echo "Usage: go run <file name>"
	@echo ""
	@echo "Output:"
	@echo "It would print below parameters"
	@echo "\t number of lines"
	@echo "\t number of bytes"
	@echo "\t number of characters"
	@echo "\t filename"

test:
	go run cmd/wc-clone/main.go sample.txt

build:
	go build -ldflags "-X github.com/Xebec19/special-system/internal/version.Version=$(VERSION) \
	-X github.com/Xebec19/special-system/internal/version.Commit=$(COMMIT) \
	-X github.com/Xebec19/special-system/internal/version.BuildDate=$(DATE)" \
	-o bin/$(BINARY) ./cmd/$(BINARY)

install:
	go install ./cmd/$(BINARY)

.PHONY: help test
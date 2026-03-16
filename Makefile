BINARY=wc-clone
VERSION=1.0.0
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
	go build -ldflags "-X /internal/version.Version=$(VERSION) \
	-X /internal/version.Commit=$(COMMIT) \
	-X /internal/version.BuildDate=$(DATE)" \
	-o bin/$(BINARY) ./cmd/$(BINARY)

install:
	build && go install ./bin/wc-clone

.PHONY: help test
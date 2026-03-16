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

.PHONY: test
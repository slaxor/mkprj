.DEFAULT_GOAL := test

.PHONY: debug
debug:
	dlv test -v ./...

.PHONY: test
test:
	go test -v ./...


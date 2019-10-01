
.PHONY: setup
setup:
	mkdir bin

.PHONY: build
build:
	go build \
		-o bin/synapse \
		cmd/synapse/main.go

.PHONY: run
run: build
	./bin/synapse
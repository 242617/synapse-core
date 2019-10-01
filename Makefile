
.PHONY: setup
setup:
	mkdir bin

.PHONY: build
build:
	go build \
		-o bin/mirror \
		cmd/mirror/main.go

.PHONY: run
run: build
	./bin/mirror
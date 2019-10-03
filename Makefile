
APPLICATION := synapse
ENVIRONMENT := production
PROJECT := github.com/242617/${APPLICATION}

.PHONY: setup
setup:
	mkdir bin

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build \
		-o bin/synapse \
		-ldflags "\
			-X '${PROJECT}/version.Application=${APPLICATION}'\
			-X '${PROJECT}/version.Environment=${ENVIRONMENT}'\
		"\
		cmd/synapse/main.go

.PHONY: run
run: build
	./bin/synapse
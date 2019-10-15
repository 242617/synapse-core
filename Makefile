
APPLICATION := synapse
ENVIRONMENT := production
PROJECT := github.com/242617/${APPLICATION}

.PHONY: setup
setup:
	mkdir -p build

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build \
		-o build/synapse \
		-ldflags "\
			-X '${PROJECT}/version.Application=${APPLICATION}'\
			-X '${PROJECT}/version.Environment=${ENVIRONMENT}'\
		"\
		cmd/synapse/main.go

.PHONY: run
run: build
	./build/synapse

.PHONY: stat\:build
stat\:build:
	GOOS=linux GOARCH=amd64 go build \
		-o build/stat \
		cmd/stat/main.go
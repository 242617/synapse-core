
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
	cp config.template.yaml build/config.yaml

.PHONY: run
run: build
	./build/synapse


DOCKER_CONTAINER_NAME := synapse-core
DOCKER_IMAGE_NAME := 242617/synapse-core:1.0.0

.PHONY: docker\:build
docker\:build:
	docker build \
		-t ${DOCKER_IMAGE_NAME} \
		.

.PHONY: docker\:test
docker\:test: docker\:build
	docker run \
		--rm \
		-p 8080:8080 \
		--name ${DOCKER_CONTAINER_NAME}\
		${DOCKER_IMAGE_NAME}

.PHONY: docker\:save
docker\:save:
	docker save ${DOCKER_IMAGE_NAME} > ${DOCKER_CONTAINER_NAME}.tar
	du -h ${DOCKER_CONTAINER_NAME}.tar
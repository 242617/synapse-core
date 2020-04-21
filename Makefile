
APPLICATION ?= synapse-core
ENVIRONMENT ?= production
PROJECT ?= github.com/242617/${APPLICATION}
VERSION ?= 1.0.0

.PHONY: setup
setup:
	mkdir -p build

.PHONY: test
test:
	go test ./...

.PHONY: config
config:
	. ./env.sh; envsubst < config.template.yaml > build/config.yaml

.PHONY: proto
proto:
	protoc --proto_path api/proto api/proto/system.proto --go_out=plugins=grpc:api
	protoc --proto_path api/proto api/proto/tasks.proto --go_out=plugins=grpc:api

.PHONY: build
build: config proto
	go generate ./...
	# GOARCH=amd64 GOOS=linux
	go build \
		-o build/core \
		-ldflags "\
			-X '${PROJECT}/version.Application=${APPLICATION}'\
			-X '${PROJECT}/version.Environment=${ENVIRONMENT}'\
			-X '${PROJECT}/version.Version=${VERSION}'\
		"\
		cmd/core/main.go

.PHONY: run
run: build
	. ./env.sh; cd build && \
		./core \
		--config config.yaml


DOCKER_CONTAINER_NAME := synapse-core
DOCKER_IMAGE_NAME := 242617/synapse-core

docker\:build:
	make config
	docker build \
		--build-arg APPLICATION=${APPLICATION} \
		--build-arg ENVIRONMENT=${ENVIRONMENT} \
		--build-arg PROJECT=${PROJECT} \
		--build-arg VERSION=${VERSION} \
		-t ${DOCKER_IMAGE_NAME} \
		.

docker\:debug:
	make docker\:build
	docker run \
		--rm \
		-p 8080:8080 \
		-e TOKEN=s.ThaBjyEw5VTI8rha4TQNmxy4 \
		--name ${DOCKER_CONTAINER_NAME}\
		${DOCKER_IMAGE_NAME}

docker\:save:
	docker save -o ${DOCKER_CONTAINER_NAME}.tar ${DOCKER_IMAGE_NAME}
	du -h ${DOCKER_CONTAINER_NAME}.tar


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
	. ./env.sh
	envsubst < config.template.yaml > build/config.yaml

.PHONY: build
build: config
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
	cd build && ./core \
		--config config.yaml


DOCKER_CONTAINER_NAME := synapse-core
DOCKER_IMAGE_NAME := 242617/synapse-core

.PHONY: docker-build
docker-build: config
	docker build \
		--build-arg APPLICATION=${APPLICATION} \
		--build-arg ENVIRONMENT=${ENVIRONMENT} \
		--build-arg PROJECT=${PROJECT} \
		--build-arg VERSION=${VERSION} \
		-t ${DOCKER_IMAGE_NAME} \
		.

.PHONY: docker-debug
docker-debug: docker-build
	docker run \
		--rm \
		-p 8080:8080 \
		--name ${DOCKER_CONTAINER_NAME}\
		${DOCKER_IMAGE_NAME}

.PHONY: docker-save
docker-save:
	docker save -o ${DOCKER_CONTAINER_NAME}.tar ${DOCKER_IMAGE_NAME}
	du -h ${DOCKER_CONTAINER_NAME}.tar


.PHONY: deploy
deploy: docker-build docker-save
	. ./env.sh; \
		rsync -Pav -e ssh synapse-core.tar $${SYNAPSE_USER}@$${SYNAPSE_HOST}:/home/synapse-core; \
		ssh -t $${SYNAPSE_USER}@$${SYNAPSE_HOST} '\
			docker load -i /home/synapse-core/synapse-core.tar && \
			systemctl restart synapse-core && \
			rm /home/synapse-core/synapse-core.tar \
		'
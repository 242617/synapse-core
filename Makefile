
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

.PHONY: build
build:
	go build \
		-o build/core \
		-ldflags "\
			-X '${PROJECT}/version.Application=${APPLICATION}'\
			-X '${PROJECT}/version.Environment=${ENVIRONMENT}'\
			-X '${PROJECT}/version.Version=${VERSION}'\
		"\
		cmd/core/main.go
	cp config.template.yaml build/config.yaml

.PHONY: run
run: build
	cd build && ./core \
		--config config.yaml


DOCKER_CONTAINER_NAME := synapse-core
DOCKER_IMAGE_NAME := 242617/synapse-core


.PHONY: docker-build
docker-build:
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


SYNAPSE ?= ${SYNAPSE_USER}@${SYNAPSE_HOST}

.PHONY: deploy
deploy: docker-build docker-save
	rsync -Pav -e ssh synapse-core.tar ${SYNAPSE}:/home/synapse-core
	ssh -t ${SYNAPSE} 'docker load -i /home/synapse-core/synapse-core.tar'
	ssh -t ${SYNAPSE} 'systemctl restart synapse-core'
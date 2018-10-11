IMAGE_NAME := image-name
DOCKER_REGISTRY := registry.com

.PHONY: all build push

all: build push

build:
	docker build -t ${DOCKER_REGISTRY}/${IMAGE_NAME} .

push:
	docker push ${DOCKER_REGISTRY}/${IMAGE_NAME}

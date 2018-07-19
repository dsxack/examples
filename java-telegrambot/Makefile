DOCKER_IMAGE_NAME := bots-java-telegram-smotrovibot
DOCKER_REGISTRY_HOST := registry.dsxack.com
ANSIBLE_DEPLOY_HOST := prod1.dsxack.com

.PHONY: docker-build docker-push deploy all

all: docker-build docker-push deploy

docker-build:
	docker build -t ${DOCKER_IMAGE_NAME} .

docker-push:
	docker tag ${DOCKER_IMAGE_NAME}:latest ${DOCKER_REGISTRY_HOST}/${DOCKER_IMAGE_NAME}
	docker push ${DOCKER_REGISTRY_HOST}/${DOCKER_IMAGE_NAME}

deploy:
	ansible-playbook ./play-deploy.yml \
	--extra-vars="container_name=${DOCKER_IMAGE_NAME} image_name=${DOCKER_REGISTRY_HOST}/${DOCKER_IMAGE_NAME}:latest" \
	-i ${ANSIBLE_DEPLOY_HOST},

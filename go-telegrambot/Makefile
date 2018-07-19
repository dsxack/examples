PROJECT_NAME := bots-telegram-go-smotrovbot
DOCKER_REGISTRY := dsxack.com:5000
DEPLOY_HOST := prod1.dsxack.com

.PHONY: docker-build docker-push

all: docker-build docker-push deploy

docker-build:
	docker build -t ${PROJECT_NAME} .

docker-push:
	docker tag ${PROJECT_NAME}:latest ${DOCKER_REGISTRY}/${PROJECT_NAME}
	docker push ${DOCKER_REGISTRY}/${PROJECT_NAME}

deploy:
	ansible-playbook ./play-deploy.yml \
	--extra-vars="container_name=${PROJECT_NAME} image_name=${DOCKER_REGISTRY}/${PROJECT_NAME}:latest" \
	-i ${DEPLOY_HOST},

COMPOSE_FILE := docker-compose.yml
COMPOSE_PROJECT_NAME := 001_haproxy_scaling_containers

up:
	@COMPOSE_FILE=$(COMPOSE_FILE) COMPOSE_PROJECT_NAME=$(COMPOSE_PROJECT_NAME) \
		COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker-compose up --build --scale server=5

down:
	@COMPOSE_FILE=$(COMPOSE_FILE) COMPOSE_PROJECT_NAME=$(COMPOSE_PROJECT_NAME) \
		COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker-compose down -v --rmi local

reload: down up

test:
	@while true; do curl -H 'X-REQUEST-ID: e8616fab-dff0-493c-9983-ebf7250c928c' http://localhost;echo; sleep 1; done

.PHONY: up down test

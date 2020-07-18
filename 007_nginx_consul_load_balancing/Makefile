COMPOSE_FILE := docker-compose.yml
COMPOSE_PROJECT_NAME := 007_nginx_consul_load_balancing

build:
	DOCKER_BUILDKIT=1 docker build \
		-f server.Dockerfile \
		-t server:v1.0.0 .

up:
	@COMPOSE_FILE=$(COMPOSE_FILE) COMPOSE_PROJECT_NAME=$(COMPOSE_PROJECT_NAME) \
		COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker-compose up --build

down:
	@COMPOSE_FILE=$(COMPOSE_FILE) COMPOSE_PROJECT_NAME=$(COMPOSE_PROJECT_NAME) \
		COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker-compose down -v --rmi local

reload: down up

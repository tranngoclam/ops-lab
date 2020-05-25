build:
	DOCKER_BUILDKIT=1 docker build \
		--ssh default \
		--build-arg FLAVOR=$(FLAVOR) \
		--secret id=secret,src=.secret \
		--progress plain \
		--no-cache \
		-f ./docker/Dockerfile \
		-t coffee-$(FLAVOR) .

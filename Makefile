build:
	DOCKER_BUILDKIT=1 docker build \
		--ssh default \
		--build-arg FLAVOR=$(FLAVOR) \
		--secret id=keyone,src=./etc/keyone \
		--progress plain \
		--no-cache \
		-f ./Dockerfile \
		-t coffee-$(FLAVOR) .

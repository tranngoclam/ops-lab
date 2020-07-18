build:
	DOCKER_BUILDKIT=1 docker build \
		--no-cache \
		-f Dockerfile \
		-t foobar:0.0.3 .

validate:
	nomad job validate foobar.nomad
run:
	nomad job run foobar.nomad
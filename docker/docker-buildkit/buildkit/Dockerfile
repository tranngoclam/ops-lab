# syntax=docker/dockerfile:1.2.1

ARG BUILD_DIR=/go/src/github.com/tranngoclam/go-coffee
ARG RUN_DIR=/run
ARG FLAVOR=default

FROM golang:1.17-alpine AS builder

LABEL maintainer="tranngoclam288@gmail.com"

ARG BUILD_DIR
ARG RUN_DIR
ARG SECRET

WORKDIR $BUILD_DIR

COPY docker $BUILD_DIR/

RUN --mount=type=cache,target=/root/.cache/go-build \
	--mount=type=secret,id=keyone \
	CGO_ENABLED=0 GOOS=linux go build -a \
	-ldflags="-s -w -X 'main.Secret=$SECRET' \
					-X 'github.com/tranngoclam/buildkit/secret.Secret=$(cat etc/keyone)' \
	-installsuffix cgo -o $RUN_DIR/app main.go

FROM alpine:3.15

ARG RUN_DIR

WORKDIR $RUN_DIR

COPY --from=builder $RUN_DIR/ $RUN_DIR/

CMD ["./app"]

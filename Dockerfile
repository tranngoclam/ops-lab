# syntax=docker/dockerfile:experimental

ARG BUILD_DIR=/build
ARG RUN_DIR=/run
ARG FLAVOR=default

FROM golang:1.12.7-alpine AS builder

LABEL maintainer="tranngoclam288@gmail.com"

ARG BUILD_DIR
ARG RUN_DIR
ARG FLAVOR

WORKDIR $BUILD_DIR

RUN --mount=target=. \
		--mount=type=cache,target=/root/.cache/go-build \
		--mount=type=secret,id=keyone,dst=etc/keyone \
		CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w -X 'main.Flavor=$FLAVOR' -X 'main.KeyOne=$(cat etc/keyone)' -X 'github.com/tranngoclam/go-coffee/app/build.KeyOne=$(cat etc/keyone)' " -installsuffix cgo -o $RUN_DIR/coffee app/main.go

FROM alpine:3.11.3

ARG RUN_DIR

WORKDIR $RUN_DIR

COPY --from=builder $RUN_DIR/ $RUN_DIR/

CMD ["tail", "-f", "/dev/null"]

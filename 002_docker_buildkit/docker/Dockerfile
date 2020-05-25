# syntax=docker/dockerfile:experimental

ARG BUILD_DIR=/go/src/github.com/tranngoclam/go-coffee
ARG RUN_DIR=/run
ARG FLAVOR=default

FROM golang:1.12.7-alpine AS builder

LABEL maintainer="tranngoclam288@gmail.com"

ARG BUILD_DIR
ARG RUN_DIR
ARG FLAVOR

RUN apk add --no-cache git openssh-client
RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir -p ~/.ssh && \
		chmod 0600 ~/.ssh && \
		ssh-keyscan github.com >> ~/.ssh/known_hosts

WORKDIR $BUILD_DIR

COPY Gopkg.toml Gopkg.lock $BUILD_DIR/
# hadolint ignore=SC2215
RUN --mount=type=ssh dep ensure --vendor-only

COPY . $BUILD_DIR/

RUN --mount=type=cache,target=/root/.cache/go-build \
		--mount=type=secret,id=keyone,dst=etc/keyone \
		CGO_ENABLED=0 GOOS=linux go build -a \
		-ldflags="-s -w -X 'main.Flavor=$FLAVOR' \
										-X 'github.com/tranngoclam/go-coffee.KeyOne=$(cat etc/keyone)' \
										-X 'github.com/tranngoclam/go-coffee/app/build.KeyOne=$(cat etc/keyone)'" \
		-installsuffix cgo -o $RUN_DIR/coffee app/main.go

FROM alpine:3.11.3

ARG RUN_DIR

WORKDIR $RUN_DIR

COPY --from=builder $RUN_DIR/ $RUN_DIR/

CMD ["./coffee"]

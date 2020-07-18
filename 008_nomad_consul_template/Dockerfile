ARG DIR=/go/src/github.com/tranngoclam/ops-lab/008_nomad_consul_template

FROM golang:1.14.4-alpine AS builder

LABEL maintainer="tranngoclam288@gmail.com"

ARG DIR
WORKDIR $DIR

COPY . $DIR/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $DIR/foobar cmd/main.go

FROM alpine:3.12.0

ARG DIR

WORKDIR $DIR/

COPY --from=builder $DIR/foobar $DIR/foobar

CMD ["./foobar"]

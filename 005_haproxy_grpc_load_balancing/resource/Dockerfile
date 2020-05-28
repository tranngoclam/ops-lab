ARG BUILD_DIR=/build

FROM golang:1.14.3 AS build

ARG BUILD_DIR

WORKDIR $BUILD_DIR

COPY go.mod go.sum $BUILD_DIR/

RUN GO111MODULE=on go mod download

COPY . $BUILD_DIR/

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./server cmd/main.go

FROM alpine:3.11.6

ARG BUILD_DIR

COPY --from=build $BUILD_DIR/server ./

CMD ["./server"]

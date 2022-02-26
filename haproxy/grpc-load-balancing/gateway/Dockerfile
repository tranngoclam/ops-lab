FROM golang:1.17 AS build

COPY go.mod go.sum ./

RUN GO111MODULE=on go mod download

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /server cmd/main.go

FROM alpine:3.15

ARG BUILD_DIR

COPY --from=build /server /

CMD ["./server"]

FROM golang:1.17 AS build

COPY server.go /go/src/server.go

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /server /go/src/server.go

FROM alpine:3.15

COPY --from=build /server ./

CMD ["./server"]

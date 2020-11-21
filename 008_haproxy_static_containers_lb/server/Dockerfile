FROM golang:1.15.4 AS build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /server main.go

FROM alpine:3.12.0

COPY --from=build /server ./

CMD ["./server"]

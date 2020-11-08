FROM golang:1.15.4 AS build

WORKDIR /src

COPY go.mod go.sum /src/

RUN GO111MODULE=on go mod download

COPY . /src/

RUN ls

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/server /src/server/main.go

FROM alpine:3.12.0

WORKDIR /src

COPY --from=build /app/server ./

CMD ["./server"]

FROM golang:1.15.4 AS build

WORKDIR /src

RUN go get github.com/grpc-ecosystem/grpc-health-probe && go install github.com/grpc-ecosystem/grpc-health-probe

COPY go.mod go.sum /src/

RUN GO111MODULE=on go mod download

COPY . /src/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/server /src/server/main.go

CMD ["/app/server"]

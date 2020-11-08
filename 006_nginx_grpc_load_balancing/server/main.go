package main

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	hw "google.golang.org/grpc/examples/helloworld/helloworld"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os"
)

const (
	port = ":50051"
)

var (
	logger *zap.Logger
)

type server struct {
	hw.UnimplementedGreeterServer
	health.UnimplementedHealthServer

	hostname string
}

func (s *server) Watch(_ *health.HealthCheckRequest, _ health.Health_WatchServer) error {
	panic("implement me")
}

func (s *server) Check(_ context.Context, _ *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func (s *server) SayHello(_ context.Context, req *hw.HelloRequest) (*hw.HelloReply, error) {
	message := fmt.Sprintf("Hello %s from %s", req.Name, s.hostname)
	logger.Debug(message)
	return &hw.HelloReply{
		Message: message,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(logger),
			),
		),
	)
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	server := &server{hostname: hostname}
	health.RegisterHealthServer(s, server)
	hw.RegisterGreeterServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	hw "google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "", "address of resource server")
	name = flag.String("name", "grpc", "name")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	cc, err := grpc.DialContext(ctx, *addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func() {
		err := cc.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	client := hw.NewGreeterClient(cc)
	req := &hw.HelloRequest{Name: *name}
	for {
		res, err := client.SayHello(ctx, req)
		if err != nil {
			panic(err)
		}
		fmt.Println(res.GetMessage())
		time.Sleep(time.Second)
	}
}

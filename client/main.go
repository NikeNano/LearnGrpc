package main

import (
	"context"
	"fmt"
	"log"
	"os"

	api "github.com/NikeNano/LearnGrpc/api"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Starting client...")

	hostname := os.Getenv("SVC_HOST_NAME")

	if len(hostname) <= 0 {
		hostname = "0.0.0.0"
	}

	port := os.Getenv("SVC_PORT")

	if len(port) <= 0 {
		port = "50051"
	}

	fmt.Printf("The port is %v", port)
	cc, err := grpc.Dial(hostname+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := api.NewNanoServiceClient(cc)

	callService(c)

}

func callService(c api.NanoServiceClient) {
	fmt.Println("callService...")
	req := &api.NanoRequest{
		Input: "Niklas testy",
	}
	res, err := c.NanoService(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling gRPC: %v", err)
	}
	log.Printf("Response from Service: %v", res.Response)

	res, err = c.NanoSuperService(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling gRPC: %v", err)
	}
	log.Printf("Response from Service: %v", res.Response)

}

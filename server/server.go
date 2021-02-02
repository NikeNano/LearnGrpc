package main

import (
	"context"
	"fmt"
	"os/signal"

	"log"
	"net"
	"os"

	api "github.com/NikeNano/LearnGrpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type nanoServer struct{}

func (*nanoServer) NanoService(ctx context.Context, req *api.NanoRequest) (*api.NanoResponse, error) {
	fmt.Printf("NanoServer %v\n", req)
	name, _ := os.Hostname()

	input := req.GetInput()
	result := "Got input " + input + " server host: " + name
	res := &api.NanoResponse{
		Response: result,
	}
	return res, nil
}

func (*nanoServer) NanoSuperService(ctx context.Context, req *api.NanoRequest) (*api.NanoResponse, error) {
	res := &api.NanoResponse{
		Response: "Success Success oops dont care what you send",
	}
	return res, nil
}

func main() {
	fmt.Println("Starting Server...")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	hostname := os.Getenv("SVC_HOST_NAME")

	if len(hostname) <= 0 {
		hostname = "0.0.0.0"
	}

	port := os.Getenv("SVC_PORT")

	if len(port) <= 0 {
		port = "50051"
	}

	lis, err := net.Listen("tcp", hostname+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	api.RegisterNanoServiceServer(s, &nanoServer{})

	// reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Server running on ", (hostname + ":" + port))
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Server Shutdown")

}

package client

import (
	"context"
	"log"

	pb "grpc_client/grpc_example.v1"

	"google.golang.org/grpc"
)

func RunClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewExampleServiceClient(conn)
	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	log.Printf("Response from server: %s", response.GetGreeting())
}

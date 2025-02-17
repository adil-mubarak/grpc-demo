package main

import (
	"context"
	"log"
	"time"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Akhil", "Alice", "Bob"},
	}
	// callSayHello(client)
	// callSayHelloServerStream(client,names)
	// callSayHelloClientStream(client,names)
	callSayHelloBidirectionStream(client, names)
}

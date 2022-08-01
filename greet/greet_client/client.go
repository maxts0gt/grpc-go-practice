package main

import (
	"fmt"
	"log"

	pb "github.com/maxts0gt/grpc-go-practice/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := pb.NewGreetServiceClient(cc)

	fmt.Printf("Created client: %f", c)

}

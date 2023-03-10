package main

import (
	"log"
	"time"

	pb "github.com/daniloscipioni/grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	//doGreet(c)
	//doGreetManyTimes(c)
	//doGreetWithDeadline(c, 1*time.Second)
	//doGreetWithDeadline(c, 2*time.Second)
	//doGreetWithDeadline(c, 3*time.Second)
	doGreetWithDeadline(c, 4*time.Second)
	//doGreetWithDeadline(c, 5*time.Second)
}

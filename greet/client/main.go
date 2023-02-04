package main

import (
	"log"
	"path/filepath"

	pb "github.com/daniloscipioni/grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {

	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile, _ := filepath.Abs("greet/ssl/ca.crt")

		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))

	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	//doGreetManyTimes(c)
	//doGreetWithDeadline(c, 1*time.Second)
	//doGreetWithDeadline(c, 2*time.Second)
	//doGreetWithDeadline(c, 3*time.Second)
	//doGreetWithDeadline(c, 4*time.Second)
	//doGreetWithDeadline(c, 5*time.Second)
}

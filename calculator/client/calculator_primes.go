package main

import (
	"context"
	"io"
	"log"

	pb "github.com/daniloscipioni/grpc/calculator/proto"
)

func doCalculatorPrimes(c pb.CalculatorServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.CalculatorPrimeRequest{
		Value: 120,
	}

	stream, err := c.CalculatorPrimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Calculator Primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("PRIME: %v\n", msg.Result)

	}

}

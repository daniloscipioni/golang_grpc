package main

import (
	"context"
	"log"

	pb "github.com/daniloscipioni/grpc/calculator/proto"
)

func doCalculator(c pb.CalculatorServiceClient) {
	log.Println("doCalculator was invoked")

	res, err := c.Calculator(context.Background(), &pb.CalculatorRequest{
		Value1: 5,
		Value2: 10,
	})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("Calculation: %v\n", res.Result)
}

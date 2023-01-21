package main

import (
	"context"
	"log"

	pb "github.com/daniloscipioni/grpc/calculator/proto"
)

func (s *Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)

	return &pb.CalculatorResponse{Result: in.Value1 + in.Value2}, nil
}

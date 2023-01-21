package main

import (
	"log"

	pb "github.com/daniloscipioni/grpc/calculator/proto"
)

func (s *Server) CalculatorPrimes(in *pb.CalculatorPrimeRequest, stream pb.CalculatorService_CalculatorPrimesServer) error {
	log.Printf("Calculator Primes was invoked: %v\n", in)

	k := int32(2)
	n := in.Value
	for n > 1 {
		result := n % k
		if result == 0 {
			stream.Send(&pb.CalculatorResponse{
				Result: k,
			})
			n = n / k
		} else {
			k = k + 1
		}
	}

	return nil
}

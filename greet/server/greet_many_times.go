package main

import (
	"fmt"
	"log"

	pb "github.com/daniloscipioni/grpc/greet/proto"
)

func (s *Server) GreetManyTymes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTymesServer) error {
	log.Printf("Greet many times was invoked: %v\n", in)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}

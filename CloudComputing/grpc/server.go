package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "./common"
)

const (
	port = ":50051"
)

type server struct {}

func (s *server) ConvertLength(ctx context.Context, in *pb.LengthConvertorRequest) (*pb.LengthConvertorResponse, error) {
	log.Printf("Received: %v", in.Value)
	return &pb.LengthConvertorResponse{Value: 10}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterConvertorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
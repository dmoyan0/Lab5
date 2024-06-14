package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/dmoyan0/Lab5/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFulcrumServer
}

func (s *server) ProcessCommand(ctx context.Context, req *pb.CommandRequest) (*pb.VectorClockResponse, error) {
	fmt.Printf("Fulcrum received command: %d\n", req.Command)
	vectorClock := []int32{1, 2, 3} // Aquí debería estar la lógica para generar el reloj vectorial
	return &pb.VectorClockResponse{VectorClock: vectorClock}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":60053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFulcrumServer(s, &server{})

	fmt.Println("Fulcrum is running on port 60053...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/dmoyan0/Lab5/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFulcrumServer
	vectorClock []int32
	mu          sync.Mutex
}

func newServer() *server {
	return &server{
		vectorClock: []int32{0, 0, 0},
	}
}

func (s *server) ProcessCommand(ctx context.Context, req *pb.CommandRequest) (*pb.VectorClockResponse, error) {
	fmt.Printf("Fulcrum received command: %d\n", req.Command)

	s.mu.Lock()
	s.vectorClock[2]++
	vectorClockCopy := append([]int32(nil), s.vectorClock...)
	s.mu.Unlock()

	return &pb.VectorClockResponse{VectorClock: vectorClockCopy}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":60053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFulcrumServer(s, newServer())

	fmt.Println("Fulcrum is running on port 60053...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

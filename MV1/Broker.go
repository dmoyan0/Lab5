package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/dmoyan0/Lab5/grpc"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBrokerServer
}

func (s *server) SendDecision(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	fmt.Printf("Broker received decision: %d\n", req.Command)
	possibleAddresses := []int32{1, 2, 3} //Cambiar por strings

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomAddress := possibleAddresses[rng.Intn(len(possibleAddresses))]
	return &pb.CommandResponse{Address: randomAddress}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBrokerServer(s, &server{})

	fmt.Println("Broker is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

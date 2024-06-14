package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/dmoyan0/Lab5"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBrokerServer
}

func (s *server) SendDecision(ctx context.Context, req *pb.DecisionRequest) (*pb.DecisionResponse, error) {
	fmt.Printf("Broker received decision: %d\n", req.Decision)
	possibleAddresses := []string{"Address1", "Address2", "Address3"}
	rand.Seed(time.Now().UnixNano())
	randomAddress := possibleAddresses[rand.Intn(len(possibleAddresses))]
	return &pb.DecisionResponse{Address: randomAddress}, nil
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

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

func (s *server) SendAddress(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	fmt.Printf("Broker received decision: %d\n", req.Command)
	possibleAddresses := []string{":60051", ":60052", ":60053"} //Cambiar por strings

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomAddress := possibleAddresses[rng.Intn(len(possibleAddresses))]
	return &pb.CommandResponse{Address: randomAddress}, nil
}

func (s *server) NotifyInconsistency(ctx context.Context, req *pb.InconsistencyRequest) (*pb.InconsistencyResponse, error) {
	log.Printf("Inconsistencia reportada: Sector: %s, Base: %s, Cliente: %s, Error: %s",
		req.Sector, req.Base, req.ClientAddress, req.ErrorMessage)
	// Aquí puedes implementar la lógica para manejar la inconsistencia
	// Por ejemplo, notificar a otros servidores Fulcrum, registrar el problema, etc.
	messageResponse, err := s.fulcrumClient.merge(ctx, req)
	if err != nil {
		return &pb.InconsistencyResponse{Success: false}, err
	}
	return &pb.InconsistencyResponse{Success: true}, nil
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

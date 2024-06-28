package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/dmoyan0/Lab5/PROTO/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	// Log de la inconsistencia reportada
	log.Printf("Inconsistencia reportada: Sector: %s, Base: %s, Fulcrum: %s, Error: %s",
		req.Sector, req.Base, req.ClientAddress, req.ErrorMessage)

	// Crear un cliente Fulcrum para el servidor específico que maneja el merge
	mergeFulcrumAddress := ":60051" // Dirección del Fulcrum que maneja el merge
	conn, err := grpc.Dial(mergeFulcrumAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Fulcrum at %s: %v", mergeFulcrumAddress, err)
	}
	defer conn.Close()

	fulcrumClient := pb.NewFulcrumClient(conn)

	// Llamar al método Merge en el cliente Fulcrum
	_, err = fulcrumClient.Merge(ctx, &pb.MergeRequest{})
	if err != nil {
		return &pb.InconsistencyResponse{Success: false}, err
	}

	// Registro adicional para notificar que la inconsistencia fue manejada
	log.Println("La inconsistencia fue manejada correctamente.")

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

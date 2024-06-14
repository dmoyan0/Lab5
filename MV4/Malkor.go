package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/dmoyan0/Lab5/grpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBrokerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	command := &pb.CommandRequest{Command: 1}
	BrokerResponse, err := c.SendAddress(ctx, command)
	if err != nil {
		log.Fatalf("could not send decision: %v", err)
	}
	fmt.Printf("Jeth received address: %s\n", BrokerResponse.Address)

	//Conexión al Fulcrum
	fulcrumConn, err := grpc.Dial(BrokerResponse.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error de conexión con Fulcrum: %v", err)
	}
	defer fulcrumConn.Close()
	fulcrumClient := pb.NewFulcrumClient(fulcrumConn)

	fulcrumResp, err := fulcrumClient.ProcessCommand(ctx, command)
	if err != nil {
		log.Fatalf("No se pudo procesar el comando %v", err)
	}
	fmt.Printf("Malkor recibio el reloj vectorial: %v\n", fulcrumResp.VectorClock)
}

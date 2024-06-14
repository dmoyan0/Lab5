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

	decision := &pb.CommandRequest{Command: 1} //Ajustar comando a enviar
	r, err := c.SendAddress(ctx, decision)
	if err != nil {
		log.Fatalf("could not send decision: %v", err)
	}
	fmt.Printf("Jeth received address: %s\n", r.Address)
}

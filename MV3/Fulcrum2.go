package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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
	//Hacer algo con el comando

	//Actualizacion del reloj vectorial
	s.mu.Lock()
	defer s.mu.Unlock()
	switch req.Command {
	case 1: // AgregarBase
		s.agregarBase(req.Sector, req.Base, req.Value)
	case 2: // RenombrarBase
		s.renombrarBase(req.Sector, req.Base, req.NewName)
	case 3: // ActualizarValor
		s.actualizarValor(req.Sector, req.Base, req.Value)
	case 4: // BorrarBase
		s.borrarBase(req.Sector, req.Base)
	}

	s.vectorClock[1]++
	vectorClockCopy := append([]int32(nil), s.vectorClock...)
	//s.mu.Unlock()

	return &pb.VectorClockResponse{VectorClock: vectorClockCopy}, nil
}

func (s *server) agregarBase(sector, base string, value int32) {
	filename := sector + ".txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("%s %d\n", base, value)); err != nil {
		log.Fatalf("failed to write to file: %v", err)
	}
}

func (s *server) renombrarBase(sector, base, newName string) {
	filename := sector + ".txt"
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	output := ""
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, base) {
			output += fmt.Sprintf("%s %s\n", newName, strings.TrimSpace(line[len(base):]))
		} else {
			output += line + "\n"
		}
	}

	if err = os.WriteFile(filename, []byte(output), 0644); err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
}

func (s *server) actualizarValor(sector, base string, value int32) {
	filename := sector + ".txt"
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	output := ""
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, base) {
			output += fmt.Sprintf("%s %d\n", base, value)
		} else {
			output += line + "\n"
		}
	}

	if err = os.WriteFile(filename, []byte(output), 0644); err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
}

func (s *server) borrarBase(sector, base string) {
	filename := sector + ".txt"
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	output := ""
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if !strings.HasPrefix(line, base) {
			output += line + "\n"
		}
	}

	if err = os.WriteFile(filename, []byte(output), 0644); err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":60052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFulcrumServer(s, newServer())

	fmt.Println("Fulcrum is running on port 60052...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

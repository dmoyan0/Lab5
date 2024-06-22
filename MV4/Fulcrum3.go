package main

import (
	"context"
	"fmt"
	"io/ioutil"
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

func (s *server) GetVectorClock(ctx context.Context, req *pb.CommandRequest) (*pb.VectorClockResponse, error) {
	fmt.Printf("Fulcrum vector. . . .\n")
	vectorClockCopy := append([]int32(nil), s.vectorClock...)
	return &pb.VectorClockResponse{VectorClock: vectorClockCopy}, nil
}

func (s *server) ProcessCommand(ctx context.Context, req *pb.CommandRequest) (*pb.VectorClockResponse, error) {
	fmt.Printf("Fulcrum received command: %d\n", req.Command)

	s.mu.Lock()
	defer s.mu.Unlock()
	for i, v := range req.VectorClock {
		if s.vectorClock[i] < v {
			s.vectorClock[i] = v
		}
	}
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
	s.vectorClock[2]++
	vectorClockCopy := append([]int32(nil), s.vectorClock...)
	//s.mu.Unlock()

	return &pb.VectorClockResponse{VectorClock: vectorClockCopy}, nil
}

func (s *server) agregarBase(sector, base string, value int32) {
	Log := "LogF3.txt"
	filelog, err := os.OpenFile(Log, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log: %v", err)
	}
	defer filelog.Close()

	if _, err := filelog.WriteString(fmt.Sprintf("AgregarBase %s %s %d\n", sector, base, value)); err != nil {
		log.Fatalf("failed to write to log: %v", err)
	}

	filename := sector + ".txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("%s %s %d\n", sector, base, value)); err != nil {
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
		parts := strings.Fields(line)
		if len(parts) == 3 && parts[1] == base {
			output += fmt.Sprintf("%s %s %s\n", sector, newName, parts[2])
		} else {
			output += line + "\n"
		}
	}

	if err = os.WriteFile(filename, []byte(output), 0644); err != nil {
		log.Fatalf("failed to write file: %v", err)
	}

	Log := "LogF3.txt"
	file, err := os.OpenFile(Log, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("RenombrarBase %s %s %d\n", sector, base, newName)); err != nil {
		log.Fatalf("failed to write to log: %v", err)
	}
}

func (s *server) actualizarValor(sector, base string, value int32) {
	Log := "LogF3.txt"
	file, err := os.OpenFile(Log, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("ActualizarValor %s %s %d\n", sector, base, value)); err != nil {
		log.Fatalf("failed to write to log: %v", err)
	}

	filename := sector + ".txt"
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	output := ""
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, sector+" "+base) {
			output += fmt.Sprintf("%s %d\n", sector+" "+base, value)
		} else {
			output += line + "\n"
		}
	}

	if err = os.WriteFile(filename, []byte(output), 0644); err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
}

func (s *server) borrarBase(sector, base string) {
	Log := "LogF3.txt"
	file, err := os.OpenFile(Log, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("BorrarBase %s %s\n", sector, base)); err != nil {
		log.Fatalf("failed to write to log: %v", err)
	}

	filename := sector + ".txt"
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	output := ""
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if !strings.HasPrefix(line, sector+" "+base) {
			output += line + "\n"
		}
	}

	if err = os.WriteFile(filename, []byte(output), 0644); err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
}

func (s *server) GetEnemies(ctx context.Context, req *pb.EnemyRequest) (*pb.EnemyResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filename := req.Sector + ".txt"
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error al leer el archivo: %v", err)
	}

	var enemies int32
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 3 && parts[1] == req.Base {
			fmt.Sscanf(parts[2], "%d", &enemies)
			break
		}
	}

	vectorClockCopy := append([]int32(nil), s.vectorClock...)
	return &pb.EnemyResponse{Enemies: enemies, VectorClock: vectorClockCopy}, nil
}

func (s *server) GetFile(ctx context.Context, req *pb.FileRequest) (*pb.FileResponse, error) {
	content, err := ioutil.ReadFile(req.GetFilename())
	if err != nil {
		return nil, err
	}
	return &pb.FileResponse{Content: content}, nil
}

// Recibe archivo merge y lo lee para actualizar info de los sectores
func (s *server) ReceiveMergedFile(ctx context.Context, req *pb.FileRequest) (*pb.FileResponse, error) {
	err := ioutil.WriteFile(req.GetFilename(), req.GetContent(), 0644)
	if err != nil {
		return nil, err
	}

	// Reescribir los archivos de los sectores con el nuevo Log
	content, err := ioutil.ReadFile(req.GetFilename())
	if err != nil {
		return nil, fmt.Errorf("Error al leer el archivo merged: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue
		}
		command := parts[0]
		sector := parts[1]
		base := parts[2]
		var value int32
		if len(parts) > 3 {
			fmt.Sscanf(parts[3], "%d", &value)
		}

		switch command {
		case "AgregarBase":
			s.agregarBase(sector, base, value)
		case "RenombrarBase":
			s.renombrarBase(sector, base, base)
		case "ActualizarValor":
			s.actualizarValor(sector, base, value)
		case "BorrarBase":
			s.borrarBase(sector, base)
		}
	}

	// Actualizar el reloj vectorial después de leer el archivo merge
	s.mu.Lock()
	for i := range s.vectorClock {
		s.vectorClock[i]++
	}
	s.mu.Unlock()

	return &pb.FileResponse{Content: []byte("Se actualizo la info correctamente")}, nil
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

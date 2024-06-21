package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

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
	s.vectorClock[0]++
	vectorClockCopy := append([]int32(nil), s.vectorClock...)
	//s.mu.Unlock()

	return &pb.VectorClockResponse{VectorClock: vectorClockCopy}, nil
}

func (s *server) agregarBase(sector, base string, value int32) {
	Log := "LogF1.txt"
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

	Log := "LogF1.txt"
	file, err := os.OpenFile(Log, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("RenombrarBase %s %s %s\n", sector, base, newName)); err != nil {
		log.Fatalf("failed to write to log: %v", err)
	}
}

func (s *server) actualizarValor(sector, base string, value int32) {
	Log := "LogF1.txt"
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
	Log := "LogF1.txt"
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

func getFileFromServer(address, filename string) ([]byte, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFulcrumClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.FileRequest{Filename: filename}
	res, err := client.GetFile(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not get file: %v", err)
	}

	return res.GetContent(), nil
}

func splitContentIntoLines(content []byte) []string {
	return strings.Split(string(content), "\n")
}

func merge() error {
	servers := []string{":60052", ":60053"}
	filename := "Log.txt"

	var allLines [][]string

	for _, server := range servers {
		content, err := getFileFromServer(server, filename)
		if err != nil {
			log.Printf("Error getting log file from %s: %v\n", server, err)
		}

		lines := splitContentIntoLines(content)
		allLines = append(allLines, lines)
	}

	content, err := ioutil.ReadFile("LogF1.txt")
	if err != nil {
		log.Printf("Error getting log file F1: %v\n", err)
	}
	lines := splitContentIntoLines(content)
	allLines = append(allLines, lines)

	var merged []string
	for _, array := range allLines {
		merged = append(merged, array...)
	}

	file, err := os.Create("Log.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range merged {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":60051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFulcrumServer(s, newServer())

	fmt.Println("Fulcrum is running on port 60051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

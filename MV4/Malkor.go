package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/dmoyan0/Lab5/grpc"
	"google.golang.org/grpc"
)

type SectorRecord struct {
	Sector        string
	VectorClock   []int32
	FulcrumServer string
}

type Malkor struct {
	Sectors map[string]SectorRecord
}

func main() {
	malkor := Malkor{Sectors: make(map[string]SectorRecord)}
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBrokerClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Ingrese alguno de los comandos siguiendo el formato: \n (1) AgregarBase <nombre sector> <nombre base> [valor] \n (2) RenombrarBase <nombre sector> <nombre base> <nuevo nombre> \n (3) ActualizarValor <nombre sector> <nombre base> <nuevo valor> \n (4) BorrarBase <nombre sector> <nombre base>")
		commandStr, _ := reader.ReadString('\n')
		commandStr = strings.TrimSpace(commandStr)     //Quitamos espacios en blanco innecesarios
		commandParts := strings.Split(commandStr, " ") //Separamos por ' '

		if len(commandParts) < 2 {
			fmt.Println("Comando invalido, favor intentelo de nuevo")
			continue
		}

		commandType := commandParts[0] //AgregarBase, RenombrarBase, ActualizarValor y BorrarBase
		sector := commandParts[1]
		base := ""
		var value int32
		var newName string

		switch commandType {
		case "AgregarBase":
			if len(commandParts) < 3 {
				fmt.Println("Comando invalido, favor intentelo de nuevo")
				continue
			}
			base = commandParts[2]
			if len(commandParts) > 3 {
				fmt.Sscanf(commandParts[3], "%d", &value)
			}
		case "RenombrarBase":
			if len(commandParts) < 4 {
				fmt.Println("Comando invalido, favor intentelo de nuevo")
				continue
			}
			base = commandParts[2]
			newName = commandParts[3]
		case "ActualizarValor":
			if len(commandParts) < 4 {
				fmt.Println("Comando invalido, favor intentelo de nuevo")
				continue
			}
			base = commandParts[2]
			fmt.Sscanf(commandParts[3], "%d", &value)
		case "BorrarBase":
			if len(commandParts) < 3 {
				fmt.Println("Comando invalido, favor intentelo de nuevo")
				continue
			}
			base = commandParts[2]
		default:
			fmt.Println("Error comando")
			continue
		}

		var command *pb.CommandRequest

		switch commandType {
		case "AgregarBase":
			command = &pb.CommandRequest{Command: 1, Sector: sector, Base: base, Value: value}
		case "RenombrarBase":
			command = &pb.CommandRequest{Command: 2, Sector: sector, Base: base, NewName: newName}
		case "ActualizarValor":
			command = &pb.CommandRequest{Command: 3, Sector: sector, Base: base, Value: value}
		case "BorrarBase":
			command = &pb.CommandRequest{Command: 4, Sector: sector, Base: base}
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		//command := &pb.CommandRequest{Command: 1}
		/*command := &pb.CommandRequest{
			Command: 2, // RenombrarBase
			Sector:  "SectorAlpha",
			Base:    "Campamento1",
			NewName: "CampamentoRenombrado",
		}*/

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

		//Validacion del reloj vectorial
		record, exists := malkor.Sectors[sector]
		if exists {
			clockResp, err := fulcrumClient.GetVectorClock(ctx, command)
			if err != nil {
				log.Fatalf("No se pudo obtener el reloj vectorial: %v", err)
			}
			if !compareVectorClock(clockResp.VectorClock, record.VectorClock) {
				fmt.Printf("El reloj vectorial no coincide, elegir otro comando o esperar a que el Broker envie una direccion correcta.")
				continue //Otra opcion es pedir inmediatamente otra direccion al Broker
			}
		}

		fulcrumResp, err := fulcrumClient.ProcessCommand(ctx, command)
		if err != nil {
			log.Fatalf("No se pudo procesar el comando %v", err)
		}
		fmt.Printf("Malkor recibio el reloj vectorial: %v\n", fulcrumResp.VectorClock)

		malkor.Sectors[sector] = SectorRecord{
			Sector:        sector,
			VectorClock:   fulcrumResp.VectorClock,
			FulcrumServer: BrokerResponse.Address,
		}
	}
}

func compareVectorClock(vc1, vc2 []int32) bool {
	for i := range vc1 {
		if vc1[i] != vc2[i] {
			return false
		}
	}
	return true
}

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/dmoyan0/Lab5/PROTO/grpc"
	"google.golang.org/grpc"
)

type SectorRecord struct {
	Sector        string
	VectorClock   []int32
	FulcrumServer string
}

type Jeth struct {
	Sectors map[string]SectorRecord
}

func main() {
	jeth := Jeth{Sectors: make(map[string]SectorRecord)}
	fmt.Println("Iniciado")
	// Conexión al Broker
	conn, err := grpc.Dial("dist029:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBrokerClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Ingrese alguno de los comandos siguiendo el formato: \n (1) AgregarBase <nombre sector> <nombre base> [valor] \n (2) RenombrarBase <nombre sector> <nombre base> <nuevo nombre> \n (3) ActualizarValor <nombre sector> <nombre base> <nuevo valor> \n (4) BorrarBase <nombre sector> <nombre base>")
		commandStr, _ := reader.ReadString('\n')
		commandStr = strings.TrimSpace(commandStr)     // Quitamos espacios en blanco innecesarios
		commandParts := strings.Split(commandStr, " ") // Separamos por ' '

		if len(commandParts) < 2 {
			fmt.Println("Comando invalido, favor intentelo de nuevo")
			continue
		}

		commandType := commandParts[0] // AgregarBase, RenombrarBase, ActualizarValor y BorrarBase
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

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// Loop para nueva dirección hasta que el reloj vectorial coincida
		for {
			BrokerResponse, err := c.SendAddress(ctx, command)
			if err != nil {
				log.Fatalf("could not send command: %v", err)
			}
			fmt.Printf("Jeth received address: %s\n", BrokerResponse.Address)

			// Conexión al Fulcrum
			fulcrumConn, err := grpc.Dial(BrokerResponse.Address, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("Error de conexión con Fulcrum: %v", err)
			}
			defer fulcrumConn.Close()
			fulcrumClient := pb.NewFulcrumClient(fulcrumConn)

			// Validacion del reloj vectorial
			record, exists := jeth.Sectors[sector]
			if exists {
				clockResp, err := fulcrumClient.GetVectorClock(ctx, &pb.CommandRequest{Sector: sector, Base: base})
				if err != nil {
					log.Fatalf("No se pudo obtener el reloj vectorial: %v", err)
				}
				if !compareVectorClock(clockResp.VectorClock, record.VectorClock) {
					notifyInconsistency(c, sector, base, BrokerResponse.Address, "Reloj vectorial inconsistente")
					fmt.Printf("El reloj vectorial no coincide, elegir otro comando o esperar a que el Broker envie una direccion correcta.")
					continue // Otra opción es pedir inmediatamente otra dirección al Broker
				}
			}

			command.VectorClock = jeth.Sectors[sector].VectorClock

			fulcrumResp, err := fulcrumClient.ProcessCommand(ctx, command)
			if err != nil {
				log.Fatalf("No se pudo procesar el comando %v", err)
			}
			fmt.Printf("Jeth recibio el reloj vectorial: %v\n", fulcrumResp.VectorClock)

			jeth.Sectors[sector] = SectorRecord{
				Sector:        sector,
				VectorClock:   fulcrumResp.VectorClock,
				FulcrumServer: BrokerResponse.Address,
			}
			break
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

func notifyInconsistency(brokerClient pb.BrokerClient, sector, base, clientAddress, errorMessage string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.InconsistencyRequest{
		Sector:        sector,
		Base:          base,
		ClientAddress: clientAddress,
		ErrorMessage:  errorMessage,
	}

	_, err := brokerClient.NotifyInconsistency(ctx, req)
	if err != nil {
		log.Printf("Error al notificar al Broker: %v", err)
	} else {
		fmt.Println("Inconsistencia notificada al Broker.")
	}
}

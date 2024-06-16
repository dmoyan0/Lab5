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

type InfoSector struct {
	Sector        string
	Base          string
	VectorClock   []int32
	FulcrumServer string
}

// Se almacena sectores consultados
type Commander struct {
	Sectors map[string]InfoSector
}

func main() {
	commander := Commander{Sectors: make(map[string]InfoSector)}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBrokerClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Ingrese el comando: GetEnemigos <nombre sector> <nombre base>, para consultar info del sector.")
		commandStr, _ := reader.ReadString('\n')
		commandStr = strings.TrimSpace(commandStr)
		commandParts := strings.Split(commandStr, " ")

		if len(commandParts) != 3 || commandParts[0] != "GetEnemigos" {
			fmt.Println("Comando invalido, favor intentelo de nuevo")
			continue
		}

		sector := commandParts[1]
		base := commandParts[2]

		request := &pb.EnemyRequest{Sector: sector, Base: base}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		for {
			BrokerResponse, err := c.SendAddress(ctx, &pb.CommandRequest{})
			if err != nil {
				log.Fatalf("Error enviar el comando: %v", err)
			}

			fulcrumConn, err := grpc.Dial(BrokerResponse.Address, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("Error de conexión con Fulcrum: %v", err)
			}
			defer fulcrumConn.Close()
			fulcrumClient := pb.NewFulcrumClient(fulcrumConn)

			record, exists := commander.Sectors[sector]
			if exists {
				clockResp, err := fulcrumClient.GetVectorClock(ctx, &pb.CommandRequest{Sector: sector, Base: base})
				if err != nil {
					log.Fatalf("No se pudo obtener el reloj vectorial: %v", err)
				}
				if !compareVectorClock(clockResp.VectorClock, record.VectorClock) {
					fmt.Println("El reloj vectorial no coincide, solicitando otra dirección...")
					continue
				}
			}

			enemyResp, err := fulcrumClient.GetEnemies(ctx, request)
			if err != nil {
				log.Fatalf("No se pudo obtener la información: %v", err)
			}

			fmt.Printf("Cantidad de enemigos en %s %s: %d\n", sector, base, enemyResp.Enemies)
			commander.Sectors[sector] = InfoSector{
				Sector:        sector,
				Base:          base,
				VectorClock:   enemyResp.VectorClock,
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

package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	pb "go-client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	target := "go-server-svc.mumnk8s.svc.cluster.local:50051"

	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("No se pudo conectar a gRPC: %v", err)
	}
	defer conn.Close()

	c := pb.NewWarReportServiceClient(conn)
	log.Println("Cliente Go conectado a gRPC. Iniciando envío automático...")

	// Todos los países disponibles en el enum
	countries := []pb.Countries{
		pb.Countries_usa,
		pb.Countries_rus,
		pb.Countries_chn,
		pb.Countries_esp,
		pb.Countries_gmt,
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		req := &pb.WarReportRequest{
			Country:         countries[rng.Intn(len(countries))],
			WarplanesInAir:  int32(rng.Intn(50) + 1),
			WarshipsInWater: int32(rng.Intn(20) + 1),
			Timestamp:       time.Now().UTC().Format(time.RFC3339),
		}

		res, err := c.SendReport(context.Background(), req)
		if err != nil {
			log.Printf("Error al llamar SendReport: %v", err)
		} else {
			log.Printf("Reporte enviado → País: %s | Aviones: %d | Barcos: %d | Status: %s",
				req.Country, req.WarplanesInAir, req.WarshipsInWater, res.Status)
		}

		time.Sleep(3 * time.Second)
	}
}

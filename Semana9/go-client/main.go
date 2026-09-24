package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	// Paquete generado de Protobuf para el cliente
	pb "go-client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Direccion DNS interna del servicio gRPC en Kubernetes
	target := "go-server-svc.mumnk8s.svc.cluster.local:50051"

	// Establecer conexion gRPC sin TLS dentro del cluster
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("No se pudo conectar a gRPC: %v", err)
	}
	defer conn.Close()

	// Crear cliente gRPC a partir de la conexion establecida
	c := pb.NewWarReportServiceClient(conn)
	log.Println("Cliente Go conectado a gRPC. Iniciando envío automático...")

	// Lista de paises definidos en el enum de Protobuf
	countries := []pb.Countries{
		pb.Countries_usa,
		pb.Countries_rus,
		pb.Countries_chn,
		pb.Countries_esp,
		pb.Countries_gmt,
	}

	// Generador de numeros aleatorios con semilla temporal
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Bucle infinito para emision periodica de reportes
	for {
		// Construir mensaje con valores aleatorios simulados
		req := &pb.WarReportRequest{
			Country:         countries[rng.Intn(len(countries))],
			WarplanesInAir:  int32(rng.Intn(50) + 1),
			WarshipsInWater: int32(rng.Intn(20) + 1),
			Timestamp:       time.Now().UTC().Format(time.RFC3339),
		}

		// Enviar reporte al servidor mediante llamada RPC
		res, err := c.SendReport(context.Background(), req)
		if err != nil {
			log.Printf("Error al llamar SendReport: %v", err)
		} else {
			log.Printf("Reporte enviado -> País: %s | Aviones: %d | Barcos: %d | Status: %s",
				req.Country, req.WarplanesInAir, req.WarshipsInWater, res.Status)
		}

		// Intervalo de espera entre envios
		time.Sleep(3 * time.Second)
	}
}

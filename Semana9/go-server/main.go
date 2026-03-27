package main

import (
	"context"
	"log"
	"net"

	pb "go-server/proto"

	"google.golang.org/grpc"
)

// Estructura del servidor que implementa la interfaz gRPC
type server struct {
	pb.UnimplementedWarReportServiceServer
}

// Función que se ejecuta cuando el cliente hace la llamada gRPC
func (s *server) SendReport(ctx context.Context, req *pb.WarReportRequest) (*pb.WarReportResponse, error) {
	log.Printf("Servidor recibió reporte: País=%s, Aviones=%d, Barcos=%d", req.Country, req.WarplanesInAir, req.WarshipsInWater)

	return &pb.WarReportResponse{Status: "1"}, nil
}

func main() {
	// abrir el puerto 50051 (estándar gRPC)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Fallo al escuchar puerto: %v", err)
	}

	// crear servidor gRPC y registramos nuestro servicio
	s := grpc.NewServer()
	pb.RegisterWarReportServiceServer(s, &server{})

	log.Printf("Servidor gRPC en Go escuchando en puerto 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fallo al iniciar servidor: %v", err)
	}
}

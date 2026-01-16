package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"semana9/proto"

	"google.golang.org/grpc"
)

// definimos nuestro servidor
type server struct {
	proto.UnimplementedSystemMonitorServer
}

// Implementar función del proto
func (s *server) GetRamStatus(ctx context.Context, req *proto.RamRequest) (*proto.RamResponse, error) {
	fmt.Printf("--> [Go Server] Petición recibida desde: %s\n", req.GetHostName())

	// simular
	return &proto.RamResponse{
		UsedPercent: 65,
		Message:     "Servidor Go, funcionando",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterSystemMonitorServer(s, &server{})

	fmt.Println("Servidor gRPC (Go) corriendo en puerto 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}

package cars

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/apavanello/rocketcarsback/pkg/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeCarsServiceServer
}

func (s *server) GetCar(ctx context.Context, in *pb.CarRequest) (*pb.CarResponse, error) {

	log.Printf("Received: %v", in.GetName())
	return &pb.CarResponse{

		Name:      "Octane",
		Length:    108.10,
		Width:     88.20,
		Height:    45.80,
		SA:        22.00,
		Elevation: 1,
	}, nil
}

// Start gRPC server in port(int)
func StartServer(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCarsServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

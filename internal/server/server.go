package server

import (
	"context"
	"fmt"
	"log"
	"net"

	ds "github.com/apavanello/rocketcarsback/internal/dataset"
	pb "github.com/apavanello/rocketcarsback/pkg/pbcars"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) GetCar(ctx context.Context, in *pb.CarRequest) (car *pb.CarResponse, err error) {

	ds.GetCarByName(in.Name, car)
	err = nil

	return
}

func (s *Server) GetCarsList(ctx context.Context, in *pb.CarListRequest) (carsList *pb.CarsListResponse, err error) {

	carsList = ds.GetList()
	err = nil

	return carsList, err
}

func (s *Server) GetServerVersion(ctx context.Context, in *pb.ServerVersionRequest) (vsr *pb.ServerVersionResponse, err error) {

	return nil, nil
}

// Start gRPC server in port(int)
func StartServer(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCarsServiceServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

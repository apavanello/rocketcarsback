package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "github.com/apavanello/rocketcarsback/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "localhost:50051"
)

func main() {
	done := make(chan bool)
	for i := 0; i < 100; i++ {
		go goGetCar(done)
	}
	println(<-done)
}

func goGetCar(done chan bool) {
	for i := 0; i < 100; i++ {
		// Set up a connection to the server.
		opts := grpc.WithTransportCredentials(insecure.NewCredentials())
		conn, err := grpc.Dial(addr, opts)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewCarsServiceClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		s := strconv.Itoa(i)
		car := pb.CarRequest{
			Name: s,
		}

		carResponse, err := client.GetCar(ctx, &car)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		j, _ := json.Marshal(carResponse)
		fmt.Println(string(j))

	}
	done <- true
}

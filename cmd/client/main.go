package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/apavanello/rocketcarsback/pkg/pbcars"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "localhost:50051"
)

func main() {
	rt := flag.String("rt", "", "Run type")
	carName := flag.String("car-name", "vulcan", "")
	flag.Parse()

	fmt.Println(*rt)

	switch *rt {
	case "getcar":
		done := make(chan bool)
		for i := 0; i < 1; i++ {
			go goGetCar(*carName, done)
		}
		println(<-done)

	case "getlist":
		goGetList()
	}
}

func connector() (client pb.CarsServiceClient, conn *grpc.ClientConn, err error) {

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err = grpc.Dial(addr, opts)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = pb.NewCarsServiceClient(conn)

	return
}

func goGetCar(carName string, done chan bool) {
	for i := 0; i < 100; i++ {

		client, conn, err := connector()
		if err != nil {
			log.Fatalf("could not connect: %v", err)
		}
		defer conn.Close()

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		car := pb.CarRequest{
			Name: carName,
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

func goGetList() {

	client, cc, err := connector()

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	req := &pb.CarListRequest{}

	carListResponse, err := client.GetCarsList(ctx, req)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	j, _ := json.Marshal(carListResponse)
	fmt.Println(string(j))

}

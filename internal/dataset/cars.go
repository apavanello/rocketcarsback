package dataset

import (
	"log"

	"github.com/apavanello/rocketcarsback/internal/utils"
	pb "github.com/apavanello/rocketcarsback/pkg/pbcars"
)

func GetCarByName(carName string, cr *pb.CarResponse) {

	records, err := readCSV(carsDataCSV)
	utils.PrintError(err)

	for _, record := range records {
		if utils.UperCase(record[0]) == utils.UperCase(carName) {
			cr.Car = carAssing(record).Car
		}
	}
}

func GetList(i int) (*pb.CarsListResponse, bool) {

	EOF := false
	clr := &pb.CarsListResponse{
		Cars: &pb.Car{},
	}

	records, err := readCSV(carsDataCSV)

	if err != nil {
		log.Fatalf("%v", err)
	}

	if len(records)-1 == i {
		EOF = true
		return nil, EOF
	}
	record := records[i]
	cr := carAssing(record)
	clr.Cars = cr.Car

	return clr, EOF
}

func carAssing(record []string) *pb.CarResponse {

	cr := &pb.CarResponse{
		Car: &pb.Car{
			Name:      record[0],
			Length:    utils.ConvertFloat(record[1]),
			Width:     utils.ConvertFloat(record[2]),
			Height:    utils.ConvertFloat(record[3]),
			SA:        utils.ConvertFloat(record[4]),
			Elevation: utils.ConvertFloat(record[5]),
		},
	}

	return cr
}

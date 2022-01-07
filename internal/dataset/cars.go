package dataset

import (
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

func GetList() *pb.CarsListResponse {

	clr := &pb.CarsListResponse{
		Cars: make([]*pb.Car, 0),
	}

	records, err := readCSV(carsDataCSV)
	utils.PrintError(err)

	for _, record := range records {
		cr := carAssing(record)
		clr.Cars = append(clr.Cars, cr.Car)
	}
	return clr
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

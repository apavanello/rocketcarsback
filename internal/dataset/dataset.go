package dataset

import (
	"encoding/csv"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/apavanello/rocketcarsback/internal/utils"
	pb "github.com/apavanello/rocketcarsback/pkg/pb"
)

var carsDataCSV = path.Join("C:\\Users\\apava\\go\\src\\github.com\\apavanello\\rocketcarsback\\data\\cars.csv")

func GetList(carList *pb.CarsListResponse) (err error) {

	records, err := readCSV(carsDataCSV)
	utils.PrintError(err)

	for _, record := range records {
		car := &pb.CarResponse{}
		carAssing(car, record)
		carList.Car = append(carList.Car, car)
	}
	return
}

func GetCarByName(car *pb.CarResponse, name string) (err error) {

	records, err := readCSV(carsDataCSV)
	utils.PrintError(err)
	for _, record := range records {
		if uperCase(record[0]) == uperCase(name) {
			carAssing(car, record)
			return
		}
	}
	return
}

func uperCase(input string) string {
	return strings.ToUpper(input)
}

func carAssing(c *pb.CarResponse, record []string) {

	c.Name = record[0]
	c.Length = convertFloat(record[1])
	c.Width = convertFloat(record[2])
	c.Height = convertFloat(record[3])
	c.SA = convertFloat(record[4])
	c.Elevation = convertFloat(record[5])
}

func convertFloat(input string) float64 {
	output, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func readCSV(filename string) (records [][]string, err error) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	r := csv.NewReader(f)
	r.Comma = ';'

	records, err = r.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return
}

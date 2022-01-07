package dataset

import (
	"encoding/csv"
	"log"
	"os"
	"path"
)

var carsDataCSV = path.Join("C:", "Users", "apava", "go", "src", "github.com", "apavanello", "rocketcarsback", "data", "cars.csv")

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

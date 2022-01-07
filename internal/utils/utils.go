package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func PrintError(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func UperCase(input string) string {
	return strings.ToUpper(input)
}

func ConvertFloat(input string) float64 {
	output, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

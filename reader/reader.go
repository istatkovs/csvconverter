package reader

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadFile() [][]string {
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

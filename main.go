package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/istatkovs/csvconverter/builder"
	"github.com/istatkovs/csvconverter/reader"
	"log"
	"os"
)

func FormatData(data [][]string) []builder.Employee {

	var format []builder.Employee

	for i, row := range data {

		if i > 0 {

			rec := builder.NewEmployee().WithTitle(row[0]).WithFirstName(row[1]).WithSn(
				row[2]).WithFullName(row[3]).WithUid(row[4]).WithPersonalNumber(row[5]).WithMail(
				row[6]).WithCompany(row[7]).WithManager(row[8]).WithFunction(row[9])

			format = append(format, rec)
		}
	}
	return format
}

func main() {

	file := reader.ReadFile()

	data := FormatData(file)

	for _, empl := range data {

		e, err := json.MarshalIndent(empl, "", "	")
		if err != nil {
			log.Fatal(err)
		}
		filename := "./output/" + empl.Uid + ".json"
		er := os.WriteFile(filename, e, 0644)

		if er != nil {
			fmt.Println(errors.New("could not write file"))
		}
	}

}

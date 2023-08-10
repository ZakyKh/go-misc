package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readCsv(filename string, delimiter rune, hasHeader bool) ([][]string, error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(csvFile)
	reader.Comma = delimiter

	rows, err := reader.ReadAll() 
	if err != nil {
		return rows, err
	}

	if hasHeader {
		rows = rows[1:]
	}

	return rows, nil
}

func main() {
	records, err := readCsv("data.csv", ';', true)
	if err != nil {
		panic(err)
	}

	fmt.Println("data.csv contents (excluding header, if exist):")
	for idx, rec := range records {
		fmt.Printf("Row %05d: %s\n", idx, rec)
	}
}
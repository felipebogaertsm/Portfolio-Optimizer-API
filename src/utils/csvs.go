// Author: Felipe Bogaerts de Mattos
// email: felipebogaerts@gmail.com
// License: MIT

package utils

import (
	"encoding/csv"
	"io"
	"os"
)

func ReadCsvFile(filePath string) [][]string {
	// Creating data slice:
	data := make([][]string, 0)

	// Loading the CSV file:
	f, _ := os.Open(filePath)

	// Reading from CSV:
	reader := csv.NewReader(f)

	for {
		record, err := reader.Read()

		// Stop at EOF:
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		// Appending to data slice:
		row := make([]string, 0)
		for value := range record {
			row = append(row, record[value])
		}
		data = append(data, row)
	}

	return data
}

// Author: Felipe Bogaerts de Mattos
// email: felipebogaerts@gmail.com
// License: MIT

package utils

import (
	"bytes"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func ReadCSVFile(filePath string) string {
	// Opening csv file:
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	// Using gota to read from the CSV:
	df := dataframe.ReadCSV(file)

	buffer := new(bytes.Buffer)
	_ = df.WriteJSON(buffer)

	return buffer.String()
}

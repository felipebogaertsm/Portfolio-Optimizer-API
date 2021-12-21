// Author: Felipe Bogaerts de Mattos
// email: felipebogaerts@gmail.com
// License: MIT

package utils

import (
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func ReadCSVFile(filePath string) dataframe.DataFrame {
	// Opening csv file:
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	// Using gota to read from the CSV:
	df := dataframe.ReadCSV(file)
	return df
}

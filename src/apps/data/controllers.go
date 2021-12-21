// Author: Felipe Bogaerts de Mattos
// email: felipebogaerts@gmail.com
// License: MIT

package data

import (
	"bytes"
	. "popt/utils"

	"github.com/gin-gonic/gin"
)

const filePath string = "public/stocks.csv"

func DataReadCSVController(c *gin.Context) {
	csvFileDataFrame := ReadCSVFile(filePath)

	// Exporting dataframe to JSON:
	buffer := new(bytes.Buffer)
	_ = csvFileDataFrame.WriteJSON(buffer)

	c.JSON(200, gin.H{"data": buffer.String()})
}

func DataMeanController(c *gin.Context) {
	meanData := GetAnualisedReturns(filePath)

	// Exporting dataframe to JSON:
	buffer := new(bytes.Buffer)
	_ = meanData.WriteJSON(buffer)

	c.JSON(200, gin.H{"data": buffer.String()})
}

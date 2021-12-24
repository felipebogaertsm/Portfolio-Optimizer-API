// Author: Felipe Bogaerts de Mattos
// email: felipebogaerts@gmail.com
// License: MIT

package data

import (
	"bytes"
	"os"
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
	csvFileDataFrame := ReadCSVFile(filePath)
	meanData := GetAnualisedReturns(csvFileDataFrame)

	// Exporting dataframe to JSON:
	buffer := new(bytes.Buffer)
	_ = meanData.WriteJSON(buffer)

	c.JSON(200, gin.H{"data": buffer.String()})
}

func DataCovarianceController(c *gin.Context) {
	csvFileDataFrame := ReadCSVFile(filePath)
	covarianceData := GetCovariance(csvFileDataFrame)

	// Exporting dataframe to JSON:
	buffer := new(bytes.Buffer)
	_ = covarianceData.WriteJSON(buffer)

	c.JSON(200, gin.H{"data": buffer.String()})
}

func MarkowitzOutputController(c *gin.Context) {
	csvFileDataFrame := ReadCSVFile(filePath)
	df := MarkowitzOptimizer(csvFileDataFrame)

	// Exporting dataframe to JSON:
	buffer := new(bytes.Buffer)
	_ = df.WriteJSON(buffer)

	// Exporting dataframe to JSON file:
	file, _ := os.Create("plotter/return_variance_data.json")
	df.WriteJSON(file)

	c.JSON(200, gin.H{"data": buffer.String()})
}

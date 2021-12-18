// Author: Felipe Bogaerts de Mattos
// email: felipebogaerts@gmail.com
// License: MIT

package data

import (
	. "popt/utils"

	"github.com/gin-gonic/gin"
)

func DataReadCsvController(c *gin.Context) {
	filePath := "public/stocks.csv"

	csvFileData := ReadCsvFile(filePath)

	c.JSON(200, gin.H{"data": csvFileData})
}

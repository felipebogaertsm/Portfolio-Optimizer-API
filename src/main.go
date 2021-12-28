// Author: Felipe Bogaerts de Mattos
// email: felipebogaerts@gmail.com
// License: MIT

package main

import (
	. "popt/apps/data"

	"github.com/gin-gonic/gin"
)

const portNumber string = "8080"
const inProduction = false

// main is the main function
func main() {
	router := gin.Default()

	router.GET("/data/read", DataReadCSVController)
	router.GET("/calculations/mean", DataMeanController)
	router.GET("/calculations/covariance", DataCovarianceController)
	router.GET("/optimizer/markowitz", MarkowitzOutputController)

	router.Run(":" + portNumber)
}

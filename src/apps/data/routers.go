// Author: Felipe Bogaerts de Mattos
// email: felipebogaerts@gmail.com
// License: MIT

package data

import "github.com/gin-gonic/gin"

func DataReadCsv(router *gin.RouterGroup) {
	router.GET("/get", DataReadCSVController)
}

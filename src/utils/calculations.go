// Author: Felipe Bogaerts de Mattos
// email: felipebogaerts@gmail.com
// License: MIT

package utils

import (
	"math"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"gonum.org/v1/gonum/stat"
)

func GetAnualisedReturns(df dataframe.DataFrame) dataframe.DataFrame {
	df = df.Drop(0) // deleting date column

	df = df.Capply(func(s series.Series) series.Series {
		seriesFloat := s.Float()
		mean := math.Pow(stat.Mean(seriesFloat, nil)+1, 252) - 1
		return series.Floats(mean)
	})

	return df
}

func GetCovariance(df dataframe.DataFrame) [][]float64 {
	df = df.Drop(0)

	var covarianceMatrix [][]float64

	colNames := df.Names()
	for i := 0; i < len(colNames); i++ {
		var covarianceMatrixRow []float64
		for j := 0; j < len(colNames); j++ {
			covarianceMatrixRow = append(covarianceMatrixRow, 0.3)
		}
	}

	return covarianceMatrix
}

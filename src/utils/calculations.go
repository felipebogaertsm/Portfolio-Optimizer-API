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

func GetMean(df dataframe.DataFrame) dataframe.DataFrame {
	df = df.Drop(0) // deleting date column

	df = df.Capply(func(s series.Series) series.Series {
		seriesFloat := s.Float()
		mean := stat.Mean(seriesFloat, nil)
		return series.Floats(mean)
	})

	return df
}

// Source:
// https://www.investopedia.com/articles/financial-theory/11/calculating-covariance.asp
func GetCovariance(df dataframe.DataFrame) dataframe.DataFrame {
	df = df.Drop(0)
	colNames := df.Names()

	covarianceDf := df.Capply(func(s series.Series) series.Series {
		var covarianceDfCol []float64
		for i := 0; i < len(colNames); i++ {
			colYName := colNames[i]
			sum := 0.0 // sum of all elements
			// Iterate through all elements of the series (col)
			for j := 0; j < s.Len(); j++ {
				sum += (s.Elem(j).Float() - s.Mean()) * (df.Col(colYName).Elem(j).Float() - df.Col(colYName).Mean())
			}
			covarianceDfCol = append(covarianceDfCol, sum/(float64(len(colNames)-1)))
		}
		return series.Floats(covarianceDfCol)
	})

	return covarianceDf
}

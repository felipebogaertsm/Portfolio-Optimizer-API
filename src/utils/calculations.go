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

func GetAnualisedReturns(filePath string) dataframe.DataFrame {
	df := ReadCSVFile(filePath)

	df = df.Capply(func(s series.Series) series.Series {
		seriesFloat := s.Float()
		mean := math.Pow(stat.Mean(seriesFloat, nil)+1, 252) - 1
		return series.Floats(mean)
	})

	return df
}

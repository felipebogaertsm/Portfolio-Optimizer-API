// Author: Felipe Bogaerts de Mattos
// email: felipebogaerts@gmail.com
// License: MIT

package utils

import (
	"math/rand"

	"github.com/go-gota/gota/dataframe"
)

func MarkowitzOptimizer(data dataframe.DataFrame) dataframe.DataFrame {
	// Deleting date column:
	data = data.Drop(data.Select("Date"))

	numberOfPortfolios := 1000        // number of portfolios to sample
	assetNames := data.Names()        // data column names or asset names
	numberOfAssets := len(assetNames) // number of assets
	rand.Seed(75)                     // fixed seed, for reproducibility

	// Generate portfolios:
	for i := 0; i < numberOfPortfolios; i++ {

		// Choose random asset:
		randomAsset := assetNames[rand.Intn(numberOfAssets)]
		dataCol := data.Col(randomAsset)

	}

	return data
}

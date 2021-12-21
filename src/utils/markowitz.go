// Author: Felipe Bogaerts de Mattos
// email: felipebogaerts@gmail.com
// License: MIT

package utils

import (
	"fmt"
	"math/rand"

	"github.com/go-gota/gota/dataframe"
)

func MarkowitzOptimizer(df dataframe.DataFrame) dataframe.DataFrame {
	// Deleting date column:
	df = df.Drop(0)

	// Getting mean values:
	dfAnualisedReturns := GetAnualisedReturns(df)
	covarianceSlice := GetCovariance(df)

	numberOfPortfolios := 1000        // number of portfolios to sample
	assetNames := df.Names()          // data column names or asset names
	numberOfAssets := len(assetNames) // number of assets
	rand.Seed(75)                     // fixed seed, for reproducibility

	// Generate portfolios:
	for i := 0; i < numberOfPortfolios; i++ {

		// Choose random asset:
		randomAssetIndexes := rand.Perm(numberOfAssets)
		randomAsset := assetNames[rand.Intn(numberOfAssets)]
		dataCol := df.Col(randomAsset)

		// Generate random weights
		var weights []float64
		var weightSum float64 = 0
		for j := 0; j < numberOfAssets; j++ {
			randomNumber := rand.Float64()
			weights = append(weights, randomNumber)
			weightSum += randomNumber
		}
		for j := 0; j < len(weights); j++ {
			weights[j] = weights[j] / weightSum
		}

		// Calculating portfolio return and variance:
		portfolioReturn := 0.0
		portfolioVariance := 0.0

		for j := 0; j < numberOfAssets; j++ {
			anualisedReturnsSlice := dfAnualisedReturns.Col(assetNames[randomAssetIndexes[j]]).Float()
			weightedAnualisedReturns := 0.0
			for k := 0; k < len(anualisedReturnsSlice); k++ {
				weightedAnualisedReturns += weights[j] * anualisedReturnsSlice[k]
			}
			portfolioReturn += weightedAnualisedReturns

			for k := 0; k < numberOfAssets; j++ {
				weightedPortfolioVariance := 0.0
				for l := 0; l < len(covarianceSlice); l++ {
					weightedPortfolioVariance += weights[j] * weights[k] * covarianceSlice[j][k]
				}
				portfolioVariance += weightedPortfolioVariance
			}
		}

		fmt.Println(i, dataCol, weights, portfolioReturn, portfolioVariance)

	}

	return df
}

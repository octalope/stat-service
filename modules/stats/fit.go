package stats

import (
	"math"
)

type Statistics struct {
	n     float64 // 	Number of records
	Sx    float64 // 	Sum of x values
	Sy    float64 // 	Sum of y values
	meanX float64 // 	Mean of x values
	meanY float64 // 	Mean of y values
	SSxx  float64 // 	SSxx: Sum of squares of x deviations
	SSyy  float64 // 	SSyy: Sum of squares of y deviations
	SSxy  float64 // 	SSxy: Sum of products of x and y deviations
}

func calculateStats(records [][]float64, xIndex int, yIndex int) Statistics {
	n := float64(len(records))

	var Sx, Sy float64
	for _, record := range records {
		Sx += record[xIndex]
		Sy += record[yIndex]
	}

	meanX := Sx / n
	meanY := Sy / n

	SSxx := 0.0
	SSyy := 0.0
	SSxy := 0.0
	for _, record := range records {
		SSxx += (record[xIndex] - meanX) * (record[xIndex] - meanX)
		SSyy += (record[yIndex] - meanY) * (record[yIndex] - meanY)
		SSxy += (record[xIndex] - meanX) * (record[yIndex] - meanY)
	}

	return Statistics{n, Sx, Sy, meanX, meanY, SSxx, SSyy, SSxy}
}

func LeastSquaresFit(records [][]float64, xIndex int, yIndex int) (float64, float64, float64, float64, float64) {
	stats := calculateStats(records, xIndex, yIndex)

	m := stats.SSxy / stats.SSxx
	b := stats.meanY - m*stats.meanX

	// Calculate the standard error of the estimate (SSE)
	SSE := 0.0
	for _, record := range records {
		y_hat := m*record[xIndex] + b
		SSE += ((record[yIndex] - y_hat) * (record[yIndex] - y_hat))
	}

	// Calculate the variance of the residuals
	var_yx := SSE / (stats.n - 2)

	// Calculate the variance and standard deviation of the slope
	var_m := var_yx / stats.SSxx
	s_m := math.Sqrt(var_m)

	// 95% confidence interval for the slope
	dm := 2 * s_m

	// Calculate the variance and standard deviation of the intercept
	var_b := var_yx * (1.0/stats.n + (stats.meanX * stats.meanX / stats.SSxx))
	s_b := math.Sqrt(var_b)

	// 95% confidence interval for the intercept
	db := 2 * s_b

	// Calculate R-squared
	SSR := stats.SSyy - SSE
	rSquared := SSR / stats.SSyy

	return m, dm, b, db, rSquared
}

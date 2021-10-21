package main

import (
	"math"
	"math/rand"
)

func uniformRand(from int64, to int64) int64 {
	return rand.Int63n(to - from) + from
}

func lognormalRand(mean float64, stdev float64) int64 {
	// The following alterations are done so that the mean and stdev parameters behave as expected if it were a normal distribution
	// see: https://en.wikipedia.org/wiki/Log-normal_distribution#:~:text=In%20probability%20theory%2C%20a%20log,X
	alteredMean := math.Log(math.Pow(mean, 2) / math.Sqrt(math.Pow(mean, 2) + math.Pow(stdev, 2)))
	alteredVariance := math.Log(1 + (math.Pow(stdev, 2) / math.Pow(mean, 2)))
	return int64(math.Exp(alteredMean + alteredVariance * rand.NormFloat64()))
}
package statistics

import "math"

func Sigmoid(t float64) float64 {
	return 1 / (1 + math.Pow(math.E, -t))
}

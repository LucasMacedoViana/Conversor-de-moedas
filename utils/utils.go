package utils

import "math"

// Round - Arredonda um número float64 para um número de casas decimais
func Round(num float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(num*shift) / shift
}

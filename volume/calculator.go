package volume

import (
	"github.com/alancesar/homebrew/temperature"
	"math"
)

func CO2(kgfcm2 float64, temperature temperature.Temperature) float64 {
	celsius := temperature.Celsius
	bar := kgfcm2 * 0.9806652048
	x := (bar * 14.50377) + 14.695
	y := 0.01821 + (0.090115 * math.Exp(-((celsius * 1.8) / 43.11)))
	return (x * y) - 0.003342
}

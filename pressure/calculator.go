package pressure

import (
	"github.com/alancesar/homebrew/temperature"
	"math"
)

const (
	atmosphericPressureAtSeaLevel = 14.695
	fudgeFactor                   = 0.003342
)

func CalculateVolumesOfCarbonation(pressure Pressure, temperature temperature.Temperature) float64 {
	x := pressure.PSI + atmosphericPressureAtSeaLevel
	y := 0.01821 + (0.090115 * math.Exp(-((temperature.Celsius * 1.8) / 43.11)))
	return (x * y) - fudgeFactor
}

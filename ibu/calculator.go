package ibu

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/volume"
	"math"
	"sort"
)

const (
	cutoffGravity   = 1.050
	maxDanielFactor = .27
)

var (
	danielUtilizationFactor = map[int]float64{
		10: .05,
		20: .12,
		30: .15,
		45: .19,
		60: .22,
		75: .24,
		90: .27,
	}
	danielUtilizationFactorKeys []int
)

type Calculator func(hops []hop.Hop, wortGravity density.Density, batchSize volume.Volume) (ibu float64)

func init() {
	for key := range danielUtilizationFactor {
		danielUtilizationFactorKeys = append(danielUtilizationFactorKeys, key)
	}

	sort.Ints(danielUtilizationFactorKeys)
}

func CalculateTinseth(hops []hop.Hop, wortGravity density.Density, batchSize volume.Volume) (ibu float64) {
	bignessFactor := calculateBignessFactor(wortGravity)

	for _, input := range hops {
		mglAlphaAcid := input.AlphaAcids * input.Quantity.Grams * 1000
		boilTimeFactor := calculateBoilTimeFactor(input.BoilTime)
		alphaAcidUtilization := bignessFactor * boilTimeFactor
		if input.Pellet {
			alphaAcidUtilization = alphaAcidUtilization * 1.1
		}

		ibu += alphaAcidUtilization * (mglAlphaAcid / batchSize.Liters)
	}

	return ibu
}

func calculateBignessFactor(wortGravity density.Density) float64 {
	return 1.65 * math.Pow(0.000125, wortGravity.SG-1)
}

func calculateBoilTimeFactor(boilTime int) float64 {
	return (1 - math.Pow(math.E, -(0.04*float64(boilTime)))) / 4.15
}

func CalculateRager(hops []hop.Hop, wortGravity density.Density, batchSize volume.Volume) (ibu float64) {
	adjustment := calculateGravityAdjustment(wortGravity)
	gravityAdjustmentFactor := batchSize.Liters * (1 + adjustment)

	for _, input := range hops {
		utilizationFactor := calculateUtilizationFactor(input.BoilTime)
		ibu += (input.Quantity.Grams * utilizationFactor * input.AlphaAcids * 1000) / gravityAdjustmentFactor
	}

	return ibu
}

func calculateGravityAdjustment(wortGravity density.Density) float64 {
	if wortGravity.SG > cutoffGravity {
		return (wortGravity.SG - cutoffGravity) / 0.2
	}

	return 0
}

func calculateUtilizationFactor(boilTime int) float64 {
	return (18.11 + (13.86 * math.Tanh((float64(boilTime)-31.32)/18.27))) / 100
}

func CalculateGaretz(hops []hop.Hop, wortGravity density.Density, wortCollected, batchSize volume.Volume) (ibu float64) {
	caFactor := batchSize.Gallons * calculateCombinedAdjustment(wortGravity, wortCollected, batchSize)
	for _, input := range hops {
		utilization := 7.2994 + (15.0746 * math.Tanh((float64(input.BoilTime)-21.86)/24.71))
		ibu += (utilization * (input.AlphaAcids * 100) * input.Quantity.Ounces * 0.749) / caFactor
	}

	return ibu
}

func calculateCombinedAdjustment(wortGravity density.Density, wortCollected, batchSize volume.Volume) float64 {
	cf := batchSize.Gallons / wortCollected.Gallons
	bg := (cf * (wortGravity.SG - 1)) + 1
	gf := (bg-cutoffGravity)/0.2 + 1
	// 25 = expected IBU
	hf := ((cf * 25) / 260) + 1
	// 2500 = elevation in feet
	tf := ((2500 / 550) * 0.02) + 1

	return gf * hf * tf
}

func CalculateDaniel(hops []hop.Hop, _ density.Density, batchSize volume.Volume) (ibu float64) {
	for _, input := range hops {
		uf := getUtilizationFactor(input.BoilTime)
		ibu += uf * (input.AlphaAcids * input.Quantity.Grams * 1000) / batchSize.Liters
	}

	return ibu
}

func getUtilizationFactor(boilTime int) float64 {
	for _, time := range danielUtilizationFactorKeys {
		if boilTime < time {
			return danielUtilizationFactor[time]
		}
	}

	return maxDanielFactor
}

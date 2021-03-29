package ibu

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/volume"
	"math"
)

const (
	alphaAcidFactor = 7490
	cutoffGravity   = 1.050
)

func CalculateTinseth(hops []hop.Hop, wortGravity density.Density, batchSize volume.Volume) (ibu float64) {
	bignessFactor := calculateBignessFactor(wortGravity)

	for _, input := range hops {
		mglAlphaAcid := input.AlphaAcids * input.Quantity.Ounces * alphaAcidFactor
		boilTimeFactor := calculateBoilTimeFactor(input.BoilTime)
		alphaAcidUtilization := bignessFactor * boilTimeFactor
		if input.Pellet {
			alphaAcidUtilization = alphaAcidUtilization * 1.1
		}

		ibu += alphaAcidUtilization * (mglAlphaAcid / batchSize.Gallons)
	}

	return ibu
}

func calculateBignessFactor(wortGravity density.Density) float64 {
	return 1.65 * math.Pow(0.000125, wortGravity.SG-1)
}

func calculateBoilTimeFactor(boilTime int) float64 {
	return (1 - math.Pow(math.E, -(0.04*float64(boilTime)))) / 4.15
}

func CalculateRoger(hops []hop.Hop, wortGravity density.Density, batchSize volume.Volume) (ibu float64) {
	adjustment := calculateGravityAdjustment(wortGravity)
	gravityAdjustmentFactor := batchSize.Gallons * (1 + adjustment)

	for _, input := range hops {
		utilizationFactor := calculateUtilizationFactor(input.BoilTime)
		ibu += (input.Quantity.Ounces * utilizationFactor * input.AlphaAcids * 7462) / gravityAdjustmentFactor
	}

	return ibu
}

func calculateGravityAdjustment(wortGravity density.Density) float64 {
	if wortGravity.SG > cutoffGravity {
		return wortGravity.SG - cutoffGravity/0.2
	}

	return 0
}

func calculateUtilizationFactor(boilTime int) float64 {
	return (18.11 + (13.86 * math.Tanh((float64(boilTime)-31.32)/18.27))) / 100
}

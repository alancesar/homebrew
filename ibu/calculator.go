package ibu

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/volume"
	"math"
)

const (
	alphaAcidFactor = 7490
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

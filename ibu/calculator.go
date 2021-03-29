package ibu

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/volume"
	"math"
)

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

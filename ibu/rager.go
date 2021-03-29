package ibu

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/volume"
	"math"
)

const cutoffGravity = 1.050

type Rager struct {
}

func NewRagerCalculator() *Rager {
	return &Rager{}
}

func (r *Rager) Calculate(hops []hop.Hop, wortGravity density.Density, batchSize volume.Volume) (ibu float64) {
	adjustment := r.calculateGravityAdjustment(wortGravity)
	gravityAdjustmentFactor := batchSize.Liters * (1 + adjustment)

	for _, input := range hops {
		utilization := r.calculateUtilization(input.BoilTime)
		ibu += (input.Quantity.Grams * utilization * input.AlphaAcids * 1000) / gravityAdjustmentFactor
	}

	return ibu
}

func (*Rager) calculateGravityAdjustment(wortGravity density.Density) float64 {
	if wortGravity.SG > cutoffGravity {
		return (wortGravity.SG - cutoffGravity) / 0.2
	}

	return 0
}

func (*Rager) calculateUtilization(boilTime int) float64 {
	return (18.11 + (13.86 * math.Tanh((float64(boilTime)-31.32)/18.27))) / 100
}

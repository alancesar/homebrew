package ibu

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/volume"
	"math"
)

type Tinseth struct {
}

func NewTinsethCalculator() *Tinseth {
	return &Tinseth{}
}

func (t *Tinseth) Calculate(hops []hop.Hop, wortGravity density.Density, batchSize volume.Volume) (ibu float64) {
	bignessFactor := t.calculateBignessFactor(wortGravity)

	for _, input := range hops {
		mglAlphaAcid := input.AlphaAcids * input.Quantity.Grams * 1000
		utilization := t.calculateUtilization(bignessFactor, input.BoilTime)
		if input.Pellet {
			utilization = utilization * 1.1
		}

		ibu += utilization * (mglAlphaAcid / batchSize.Liters)
	}

	return ibu
}

func (t *Tinseth) calculateUtilization(bignessFactor float64, boilTime int) float64 {
	boilTimeFactor := t.calculateBoilTimeFactor(boilTime)
	return bignessFactor * boilTimeFactor
}

func (*Tinseth) calculateBignessFactor(wortGravity density.Density) float64 {
	return 1.65 * math.Pow(0.000125, wortGravity.SG-1)
}

func (*Tinseth) calculateBoilTimeFactor(boilTime int) float64 {
	return (1 - math.Pow(math.E, -(0.04*float64(boilTime)))) / 4.15
}

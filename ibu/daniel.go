package ibu

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/volume"
)

const maxDanielFactor = .27

type danielUtilizationFactor struct {
	time  int
	value float64
}

var danielUtilizationFactorTable = []danielUtilizationFactor{
	{10, .05},
	{20, .12},
	{30, .15},
	{45, .19},
	{60, .22},
	{75, .24},
	{90, .27},
}

type Daniel struct {
}

func NewDanielCalculator() *Daniel {
	return &Daniel{}
}

func (d *Daniel) Calculate(hops []hop.Hop, _ density.Density, batchSize volume.Volume) (ibu float64) {
	for _, input := range removeDryHopping(hops) {
		uf := d.getUtilizationFactor(input.BoilTime)
		ibu += uf * (input.AlphaAcids * input.Quantity.Grams * 1000) / batchSize.Liters
	}

	return ibu
}

func (d *Daniel) getUtilizationFactor(boilTime int) float64 {
	for _, item := range danielUtilizationFactorTable {
		if boilTime < item.time {
			return item.value
		}
	}

	return maxDanielFactor
}

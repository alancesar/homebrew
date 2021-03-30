package api

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/fermentable"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/recipe"
	"github.com/alancesar/homebrew/volume"
)

const (
	minimalExpectedBrixValue = 1.130
	minimalExpectedSGValue   = 1
)

func bindPercentValue(value int, binder func(float64) *recipe.Recipe) {
	if value != 0 {
		binder(float64(value) / 100)
	}
}

func bindVolume(value string, binder func(volume.Volume) *recipe.Recipe) {
	if value != "" {
		binder(volume.NewFrom(value))
	}
}

func bindDensity(value float64, binder func(density density.Density) *recipe.Recipe) {
	if value > minimalExpectedBrixValue {
		binder(density.NewFromBrix(value))
	} else if value > minimalExpectedSGValue {
		binder(density.NewFromSG(value))
	}
}

func bindFermentable(input []Fermentable, binder func(...fermentable.Fermentable) *recipe.Recipe) {
	var fermentableList []fermentable.Fermentable
	for _, f := range input {
		fermentableList = append(fermentableList, parseFermentable(f))
	}
	binder(fermentableList...)
}

func bindHops(input []Hop, binder func(...hop.Hop) *recipe.Recipe) {
	var hops []hop.Hop
	for _, h := range input {
		hops = append(hops, parseHop(h))
	}
	binder(hops...)
}

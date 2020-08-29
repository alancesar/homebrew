package recipe

import (
	"github.com/alancesar/homebrew/color"
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/mass"
	"github.com/alancesar/homebrew/volume"
)

type Recipe struct {
	Og     density.Density
	Fg     density.Density
	Inputs []Input
}

type Input struct {
	Quantity  mass.Mass
	Color     color.Color
	BatchSize volume.Volume
}

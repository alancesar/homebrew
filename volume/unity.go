package volume

import (
	"github.com/alancesar/homebrew/units"
)

const (
	milliliterSymbol = "ml"
	literSymbol      = "l"
	gallonSymbol     = "gal"
)

type Volume struct {
	Milliliters units.Unity
	Liters      units.Unity
	Gallons     units.Unity
}

func create(from *from) Volume {
	return Volume{
		Milliliters: units.New(from.ToMilliliter(), milliliterSymbol),
		Liters:      units.New(from.ToLiter(), literSymbol),
		Gallons:     units.New(from.ToGallon(), gallonSymbol),
	}
}

func Milliliter(value float64) Volume {
	return create(FromMilliliter(value))
}

func Liter(value float64) Volume {
	return create(FromLiter(value))
}

func Gallon(value float64) Volume {
	return create(FromGallon(value))
}

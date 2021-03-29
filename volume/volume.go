package volume

import "github.com/alancesar/homebrew/measure"

const (
	millilitersInLiters = 1000
	gallonsInLiters     = 0.26417205235815
)

type (
	Volume struct {
		Milliliters float64
		Liters      float64
		Gallons     float64
	}

	volumeConstructor func(value float64) Volume
)

var constructorsMap = map[string]volumeConstructor{
	"ml":  NewFromMilliliter,
	"l":   NewFromLiter,
	"gal": NewFromGallon,
}

func NewFrom(input string) Volume {
	symbol, value, err := measure.ExtractSymbolAndValue(input)
	if err != nil {
		return Volume{}
	}

	constructor, exists := constructorsMap[symbol]
	if !exists {
		return Volume{}
	}

	return constructor(value)
}

func NewFromMilliliter(value float64) Volume {
	return createFromLiters(value / millilitersInLiters)
}

func NewFromLiter(value float64) Volume {
	return createFromLiters(value)
}

func NewFromGallon(value float64) Volume {
	return createFromGallons(value)
}

func createFromLiters(liters float64) Volume {
	return Volume{
		Milliliters: liters * millilitersInLiters,
		Liters:      liters,
		Gallons:     liters * gallonsInLiters,
	}
}

func createFromGallons(gallons float64) Volume {
	return Volume{
		Milliliters: gallons / gallonsInLiters * millilitersInLiters,
		Liters:      gallons / gallonsInLiters,
		Gallons:     gallons,
	}
}

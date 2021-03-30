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

func (v Volume) IsZero() bool {
	return v.Liters == 0 && v.Gallons == 0
}

var constructorsMap = map[string]volumeConstructor{
	"ml":  NewFromMilliliter,
	"l":   NewFromLiter,
	"gal": NewFromGallon,
}

func NewFrom(input string) Volume {
	volume := Volume{}
	measure.NewFrom(input, func(symbol string, value float64) {
		if constructor, exists := constructorsMap[symbol]; exists {
			volume = constructor(value)
		}
	})

	return volume
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

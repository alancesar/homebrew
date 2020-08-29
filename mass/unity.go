package mass

import (
	"github.com/alancesar/homebrew/units"
)

const (
	milligram = "mg"
	gram      = "g"
	kilogram  = "kg"
	pound     = "lb"
	ounce     = "oz"
)

type Mass struct {
	Milligrams units.Unity
	Grams      units.Unity
	Kilograms  units.Unity
	Pounds     units.Unity
	Ounces     units.Unity
}

func create(from *from) Mass {
	return Mass{
		Milligrams: units.New(from.ToMilligram(), milligram),
		Grams:      units.New(from.ToGram(), gram),
		Kilograms:  units.New(from.ToKilogram(), kilogram),
		Pounds:     units.New(from.ToPounds(), pound),
		Ounces:     units.New(from.ToOunces(), ounce),
	}
}

func Milligram(value float64) Mass {
	return create(FromMilligram(value))
}

func Gram(value float64) Mass {
	return create(FromGram(value))
}

func Kilogram(value float64) Mass {
	return create(FromKilogram(value))
}

func Pound(value float64) Mass {
	return create(FromPound(value))
}

func Ounce(value float64) Mass {
	return create(FromOunce(value))
}

package mass

import "github.com/alancesar/homebrew/measure"

const (
	milligramsInGrams = 1000
	kilogramsInGrams  = 0.001
	poundsInGrams     = 453.592
	ouncesInGrams     = 28.3495
	poundsInOunces    = 16
)

type (
	Mass struct {
		Milligrams float64
		Grams      float64
		Kilograms  float64
		Pounds     float64
		Ounces     float64
	}

	massConstructor func(value float64) Mass
)

func (m Mass) IsZero() bool {
	return m.Kilograms == 0 && m.Pounds == 0
}

var constructorsMap = map[string]massConstructor{
	"mg": NewFromMilligram,
	"g":  NewFromGram,
	"kg": NewFromKilogram,
	"lb": NewFromPound,
	"oz": NewFromOunce,
}

func NewFrom(input string) Mass {
	mass := Mass{}
	measure.NewFrom(input, func(symbol string, value float64) {
		if constructor, exists := constructorsMap[symbol]; exists {
			mass = constructor(value)
		}
	})

	return mass
}

func NewFromMilligram(value float64) Mass {
	return createFromGrams(value / milligramsInGrams)
}

func NewFromGram(value float64) Mass {
	return createFromGrams(value)
}

func NewFromKilogram(value float64) Mass {
	return createFromGrams(value / kilogramsInGrams)
}

func NewFromPound(value float64) Mass {
	return createFromPounds(value)
}

func NewFromOunce(value float64) Mass {
	return createFromPounds(value / poundsInOunces)
}

func createFromGrams(grams float64) Mass {
	return Mass{
		Milligrams: grams * milligramsInGrams,
		Grams:      grams,
		Kilograms:  grams * kilogramsInGrams,
		Pounds:     grams / poundsInGrams,
		Ounces:     grams / ouncesInGrams,
	}
}

func createFromPounds(pounds float64) Mass {
	return Mass{
		Milligrams: pounds * poundsInGrams * milligramsInGrams,
		Grams:      pounds * poundsInGrams,
		Kilograms:  pounds * poundsInGrams * kilogramsInGrams,
		Pounds:     pounds,
		Ounces:     pounds * poundsInOunces,
	}
}

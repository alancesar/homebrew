package temperature

import "github.com/alancesar/homebrew/measure"

type (
	Temperature struct {
		Celsius    float64
		Fahrenheit float64
	}

	temperatureConstructor func(value float64) Temperature
)

func (t Temperature) IsZero() bool {
	return t.Celsius == 0 && t.Fahrenheit == 0
}

var constructorsMap = map[string]temperatureConstructor{
	"c":  NewFromCelsius,
	"ºc": NewFromCelsius,
	"f":  NewFromFahrenheit,
	"ºf": NewFromFahrenheit,
}

func NewFrom(input string) Temperature {
	symbol, value, err := measure.ExtractSymbolAndValue(input)
	if err != nil {
		return Temperature{}
	}

	constructor, exists := constructorsMap[symbol]
	if !exists {
		return Temperature{}
	}

	return constructor(value)
}

func NewFromCelsius(value float64) Temperature {
	return Temperature{
		Celsius:    value,
		Fahrenheit: (value * 1.8) + 32,
	}
}

func NewFromFahrenheit(value float64) Temperature {
	return Temperature{
		Celsius:    (value - 32) / 1.8,
		Fahrenheit: value,
	}
}

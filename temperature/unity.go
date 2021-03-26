package temperature

type Temperature struct {
	Celsius    float64
	Fahrenheit float64
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

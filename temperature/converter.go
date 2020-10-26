package temperature

type from struct {
	celsius    float64
	fahrenheit float64
}

func (f *from) ToCelsius() float64 {
	return f.celsius
}

func (f *from) ToFahrenheit() float64 {
	return f.fahrenheit
}

func FromCelsius(value float64) *from {
	return &from{
		celsius:    value,
		fahrenheit: (value * 1.8) + 32,
	}
}

func FromFahrenheit(value float64) *from {
	return &from{
		celsius:    (value - 32) / 1.8,
		fahrenheit: value,
	}
}

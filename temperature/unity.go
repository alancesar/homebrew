package temperature

type Temperature struct {
	Celsius    float64
	Fahrenheit float64
}

func create(from *from) Temperature {
	return Temperature{
		Celsius:    from.ToCelsius(),
		Fahrenheit: from.ToFahrenheit(),
	}
}

func Celsius(value float64) Temperature {
	return create(FromCelsius(value))
}

func Fahrenheit(value float64) Temperature {
	return create(FromFahrenheit(value))
}

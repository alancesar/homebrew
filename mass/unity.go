package mass

type Mass struct {
	Milligrams float64
	Grams      float64
	Kilograms  float64
	Pounds     float64
	Ounces     float64
}

func create(from *from) Mass {
	return Mass{
		Milligrams: from.ToMilligram(),
		Grams:      from.ToGram(),
		Kilograms:  from.ToKilogram(),
		Pounds:     from.ToPounds(),
		Ounces:     from.ToOunces(),
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

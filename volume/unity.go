package volume

type Volume struct {
	Milliliters float64
	Liters      float64
	Gallons     float64
}

func create(from *from) Volume {
	return Volume{
		Milliliters: from.ToMilliliter(),
		Liters:      from.ToLiter(),
		Gallons:     from.ToGallon(),
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

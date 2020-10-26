package density

type Density struct {
	Sg   float64
	Brix float64
}

func create(from *from) Density {
	return Density{
		Sg:   from.ToSg(),
		Brix: from.ToBrix(),
	}
}

func Sg(value float64) Density {
	return create(FromSg(value))
}

func Brix(value float64) Density {
	return create(FromBrix(value))
}

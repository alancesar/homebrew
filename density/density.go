package density

type Density struct {
	Sg   float64
	Brix float64
}

func NewFromSg(value float64) Density {
	return createFromSg(value)
}

func NewFromBrix(value float64) Density {
	return createFromBrix(value)
}

func createFromSg(sg float64) Density {
	return Density{
		Sg:   sg,
		Brix: (((((182.4601 * sg) - 775.6821) * sg) + 1262.7794) * sg) - 669.5622,
	}
}

func createFromBrix(brix float64) Density {
	return Density{
		Sg:   (brix / (258.6 - ((brix / 258.2) * 227.1))) + 1,
		Brix: brix,
	}
}

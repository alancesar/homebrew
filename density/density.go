package density

type Density struct {
	SG   float64
	Brix float64
}

func (d Density) IsZero() bool {
	return d.SG == 0 && d.Brix == 0
}

func (d Density) IsNotZero() bool {
	return !d.IsZero()
}

func NewFromSG(value float64) Density {
	return createFromSG(value)
}

func NewFromBrix(value float64) Density {
	return createFromBrix(value)
}

func createFromSG(sg float64) Density {
	return Density{
		SG:   sg,
		Brix: (((((182.4601 * sg) - 775.6821) * sg) + 1262.7794) * sg) - 669.5622,
	}
}

func createFromBrix(brix float64) Density {
	return Density{
		SG:   (brix / (258.6 - ((brix / 258.2) * 227.1))) + 1,
		Brix: brix,
	}
}

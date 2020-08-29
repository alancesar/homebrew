package density

func FromBrix(value float64) *from {
	sg := (value / (258.6 - ((value / 258.2) * 227.1))) + 1
	return &from{
		sg: sg,
	}
}

func FromSg(value float64) *from {
	return &from{
		sg: value,
	}
}

type from struct {
	sg float64
}

func (f from) ToBrix() float64 {
	return (((((182.4601 * f.sg) - 775.6821) * f.sg) + 1262.7794) * f.sg) - 669.5622
}

func (f from) ToSg() float64 {
	return f.sg
}

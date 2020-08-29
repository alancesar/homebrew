package converter

type FromDensity struct {
	input float64
}

func (from FromDensity) FromBrix() *ToDensity {
	sg := (from.input / (258.6 - ((from.input / 258.2) * 227.1))) + 1

	return &ToDensity{
		sg: sg,
	}
}

func (from FromDensity) FromSg() *ToDensity {
	return &ToDensity{
		sg: from.input,
	}
}

type ToDensity struct {
	sg float64
}

func (to ToDensity) ToBrix() float64 {
	return (((((182.4601 * to.sg) - 775.6821) * to.sg) + 1262.7794) * to.sg) - 669.5622
}

func (to ToDensity) ToSg() float64 {
	return to.sg
}

func ConvertDensity(value float64) *FromDensity {
	return &FromDensity{
		input: value,
	}
}

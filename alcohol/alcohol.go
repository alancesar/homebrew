package alcohol

type Alcohol struct {
	ABV         float64
	ABW         float64
	Attenuation float64
}

func (a Alcohol) IsZero() bool {
	return a.ABV == 0 && a.ABW == 0
}

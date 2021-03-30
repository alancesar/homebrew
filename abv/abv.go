package abv

type Abv struct {
	Abv         float64
	Abw         float64
	Attenuation float64
}

func (a Abv) IsZero() bool {
	return a.Abv == 0 && a.Abw == 0
}

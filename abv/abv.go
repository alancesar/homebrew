package abv

import "github.com/alancesar/homebrew/density"

type Abv struct {
	Abv         float64
	Abw         float64
	Attenuation float64
}

func Calculate(og, fg density.Density) Abv {
	abv := (((76.08 * (og.SG - fg.SG)) / (1.775 - og.SG)) * (fg.SG / 0.794)) / 100
	abw := abv * 0.8
	att := (og.SG - fg.SG) / (og.SG - 1)

	return Abv{
		Abv:         abv,
		Abw:         abw,
		Attenuation: att,
	}
}

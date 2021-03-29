package abv

import "github.com/alancesar/homebrew/density"

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

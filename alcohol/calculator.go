package alcohol

import "github.com/alancesar/homebrew/density"

func Calculate(og, fg density.Density) Alcohol {
	abv := (((76.08 * (og.SG - fg.SG)) / (1.775 - og.SG)) * (fg.SG / 0.794)) / 100
	abw := abv * 0.8
	att := (og.SG - fg.SG) / (og.SG - 1)

	return Alcohol{
		ABV:         abv,
		ABW:         abw,
		Attenuation: att,
	}
}

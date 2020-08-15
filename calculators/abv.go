package calculators

import "github.com/alancesar/homebrew/entities"

type AbvCalculator struct{}

func (calculator AbvCalculator) Calculate(recipe entities.Recipe) entities.Abv {
	ogInSg := recipe.Og.Sg.Value
	fgInSg := recipe.Fg.Sg.Value
	abv := (((76.08 * (ogInSg - fgInSg)) / (1.775 - ogInSg)) * (fgInSg / 0.794)) / 100
	abw := abv * 0.8
	attenuation := (ogInSg - fgInSg) / (ogInSg - 1)

	return entities.Abv{
		Abv:         abv,
		Abw:         abw,
		Attenuation: attenuation,
	}
}

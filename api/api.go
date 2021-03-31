package api

import (
	"github.com/alancesar/homebrew/recipe"
)

type Response struct {
	Bitterness []Bitterness `json:"bitterness,omitempty"`
	Alcohol    Alcohol      `json:"alcohol,omitempty"`
	Color      Color        `json:"color,omitempty"`
	Expected   Expected     `json:"expected"`
}

type Expected struct {
	PBG     Density `json:"pbg,omitempty"`
	OG      Density `json:"og,omitempty"`
	FG      Density `json:"fg,omitempty"`
	Alcohol Alcohol `json:"alcohol,omitempty"`
}

func BuildRecipe(r Recipe) *recipe.Recipe {
	newRecipe := recipe.NewRecipe(r.Name)

	bindPercentValue(r.Efficiency, newRecipe.WithEfficiency)
	bindPercentValue(r.Attenuation, newRecipe.WithAttenuation)
	bindDensity(r.OG, newRecipe.WithOG)
	bindDensity(r.FG, newRecipe.WithFG)
	bindVolume(r.WortCollected, newRecipe.WithWortCollected)
	bindVolume(r.BatchSize, newRecipe.WithBatchSize)
	bindFermentable(r.Fermentable, newRecipe.WithFermentable)
	bindHops(r.Hops, newRecipe.WithHops)
	return newRecipe
}

func BuildResponse(r recipe.Recipe) Response {
	return Response{
		Bitterness: parseFromBitterness(r.Bitterness()),
		Color:      parseFromColor(r.Color()),
		Alcohol:    parseFromAlcohol(r.Alcohol()),
		Expected: Expected{
			PBG:     parseFromDensity(r.ExpectedPBG()),
			OG:      parseFromDensity(r.ExpectedOG()),
			FG:      parseFromDensity(r.ExpectedFG()),
			Alcohol: parseFromAlcohol(r.ExpectedAlcohol()),
		},
	}
}

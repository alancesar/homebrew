package api

import (
	"github.com/alancesar/homebrew/recipe"
)

type Response struct {
	Bitterness             []Bitterness `json:"bitterness,omitempty"`
	Alcohol                Alcohol      `json:"alcohol,omitempty"`
	Color                  Color        `json:"color,omitempty"`
	ExpectedPreBoilDensity Density      `json:"expected_pre_boil_density,omitempty"`
	ExpectedOG             Density      `json:"expected_og,omitempty"`
	ExpectedFG             Density      `json:"expected_fg,omitempty"`
	ExpectedABV            Alcohol      `json:"expected_abv,omitempty"`
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
		Bitterness:             parseFromBitterness(r.Bitterness()),
		Color:                  parseFromColor(r.Color()),
		Alcohol:                parseFromAlcohol(r.Alcohol()),
		ExpectedPreBoilDensity: parseFromDensity(r.ExpectedPreBoilDensity()),
		ExpectedOG:             parseFromDensity(r.ExpectedOG()),
		ExpectedFG:             parseFromDensity(r.ExpectedFG()),
		ExpectedABV:            parseFromAlcohol(r.ExpectedAlcohol()),
	}
}

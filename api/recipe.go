package api

import (
	alcohol "github.com/alancesar/homebrew/abv"
	"github.com/alancesar/homebrew/bitterness"
	"github.com/alancesar/homebrew/color"
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/fermentable"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/mass"
	"github.com/alancesar/homebrew/recipe"
	"github.com/alancesar/homebrew/volume"
)

const (
	minimalExpectedBrixValue = 2
	minimalExpectedSGValue   = 1
)

type Recipe struct {
	UUID          string        `json:"uuid"`
	Name          string        `json:"name"`
	Efficiency    int           `json:"efficiency"`
	Attenuation   int           `json:"attenuation"`
	OG            float64       `json:"og"`
	FG            float64       `json:"fg"`
	WortCollected string        `json:"wort_collected"`
	BatchSize     string        `json:"batch_size"`
	Hops          []Hop         `json:"hops"`
	Fermentable   []Fermentable `json:"fermentable"`
}

type Response struct {
	ExpectedBitterness     bitterness.Table `json:"expected_bitterness"`
	ExpectedPreBoilDensity density.Density  `json:"expected_pre_boil_density"`
	ExpectedOG             density.Density  `json:"expected_og"`
	ExpectedFG             density.Density  `json:"expected_fg"`
	ExpectedABV            alcohol.Abv      `json:"expected_abv"`
}

type Hop struct {
	Quantity   string `json:"quantity"`
	BoilTime   int    `json:"boil_time"`
	AlphaAcids int    `json:"alpha_acids"`
	Pellet     bool   `json:"pellet"`
	DryHopping bool   `json:"dry_hopping"`
}

type Fermentable struct {
	Quantity string  `json:"quantity"`
	Lovibond float64 `json:"lovibond"`
	PPG      float64 `json:"ppg"`
	Mashing  bool    `json:"mashing"`
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
		ExpectedBitterness:     r.ExpectedIBU(),
		ExpectedPreBoilDensity: r.ExpectedPreBoilDensity(),
		ExpectedOG:             r.ExpectedOG(),
		ExpectedFG:             r.ExpectedFG(),
		ExpectedABV:            r.ExpectedABV(),
	}
}

func bindPercentValue(value int, binder func(float64) *recipe.Recipe) {
	if value != 0 {
		binder(float64(value) / 100)
	}
}

func bindVolume(value string, binder func(volume.Volume) *recipe.Recipe) {
	if value != "" {
		binder(volume.NewFrom(value))
	}
}

func bindDensity(value float64, binder func(density density.Density) *recipe.Recipe) {
	if value > minimalExpectedBrixValue {
		binder(density.NewFromBrix(value))
	} else if value > minimalExpectedSGValue {
		binder(density.NewFromSG(value))
	}
}

func bindFermentable(input []Fermentable, binder func(...fermentable.Fermentable) *recipe.Recipe) {
	var fermentableList []fermentable.Fermentable
	for _, f := range input {
		fermentableList = append(fermentableList, parseFermentable(f))
	}
	binder(fermentableList...)
}

func bindHops(input []Hop, binder func(...hop.Hop) *recipe.Recipe) {
	var hops []hop.Hop
	for _, h := range input {
		hops = append(hops, parseHop(h))
	}
	binder(hops...)
}

func parseFermentable(f Fermentable) fermentable.Fermentable {
	return fermentable.Fermentable{
		Quantity: mass.NewFrom(f.Quantity),
		Color:    color.NewFromLovibond(f.Lovibond),
		PPG:      f.PPG,
		Mashing:  f.Mashing,
	}
}

func parseHop(h Hop) hop.Hop {
	return hop.Hop{
		Quantity:   mass.NewFrom(h.Quantity),
		BoilTime:   h.BoilTime,
		AlphaAcids: float64(h.AlphaAcids) / 100,
		Pellet:     h.Pellet,
		DryHopping: h.DryHopping,
	}
}

package api

import (
	"github.com/alancesar/homebrew/alcohol"
	"github.com/alancesar/homebrew/bitterness"
	"github.com/alancesar/homebrew/color"
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/fermentable"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/mass"
	"math"
)

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
		AlphaAcids: h.AlphaAcids / 100,
		Pellet:     h.Pellet,
		DryHopping: h.DryHopping,
	}
}

func parseFromColor(c color.Color) Color {
	return Color{
		SRM:      math.Round(c.SRM),
		EBC:      math.Round(c.EBC),
		Lovibond: math.Round(c.Lovibond),
	}
}

func parseFromBitterness(b bitterness.Table) (output []Bitterness) {
	for _, item := range b {
		output = append(output, Bitterness{
			Method: string(item.Method),
			Value: BitternessValue{
				IBU: int(math.Round(item.Value.IBU)),
			},
		})
	}

	return output
}

func parseFromDensity(d density.Density) Density {
	return Density{
		SG:   math.Round(d.SG*1000) / 1000,
		Brix: math.Round(d.Brix),
	}
}

func parseFromAlcohol(a alcohol.Alcohol) Alcohol {
	return Alcohol{
		ABV:         math.Round(a.ABV*10000) / 100,
		ABW:         math.Round(a.ABW*10000) / 100,
		Attenuation: math.Round(a.Attenuation*10000) / 100,
	}
}

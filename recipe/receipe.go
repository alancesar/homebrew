package recipe

import (
	alcohol "github.com/alancesar/homebrew/abv"
	"github.com/alancesar/homebrew/color"
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/mass"
	"github.com/alancesar/homebrew/volume"
	"math"
)

type Recipe struct {
	Og            density.Density
	Fg            density.Density
	BatchSize     volume.Volume
	WortCollected volume.Volume
	Grains        []Grain
	Hops          []Hop
	Efficiency    float64
	Attenuation   float64
}

type Grain struct {
	Quantity mass.Mass
	Color    color.Color
	Ppg      float64
	Mashing  bool
}

type Hop struct {
	Quantity   mass.Mass
	BoilTime   int
	AlphaAcids float64
}

func (r *Recipe) Color() color.Color {
	var mcu float64
	for _, input := range r.Grains {
		mcu += input.Quantity.Pounds * input.Color.Lovibond / r.BatchSize.Gallons
	}
	srm := 1.4922 * math.Pow(mcu, 0.6859)
	return color.NewFromSrm(srm)
}

func (r *Recipe) Abv() alcohol.Abv {
	return alcohol.Calculate(r.Og, r.Fg)
}

func (r *Recipe) Ibu() float64 {
	boilGravity := (r.BatchSize.Gallons / r.WortCollected.Gallons) * (r.Og.SG - 1)
	var ibu float64

	for _, input := range r.Hops {
		boilFactor := 1.65 * math.Pow(0.000125, boilGravity)
		timeFactor := (1 - math.Pow(math.E, -(0.04*float64(input.BoilTime)))) / 4.15
		ibu += ((boilFactor * timeFactor) * 1.1) * (((input.AlphaAcids) *
			input.Quantity.Ounces * 7490) / r.BatchSize.Gallons)
	}

	return ibu
}

func (r *Recipe) ExpectedGravity() (preBoilOg, og, fg density.Density, abv alcohol.Abv) {
	var mashingPoints, notMashingPoints float64

	for _, input := range r.Grains {
		if input.Mashing {
			mashingPoints += input.Ppg * input.Quantity.Pounds
		} else {
			notMashingPoints += input.Ppg * input.Quantity.Pounds
		}
	}

	points := (mashingPoints * r.Efficiency) + notMashingPoints
	preBoilOg = density.NewFromSG(((points / r.WortCollected.Gallons) * 0.001) + 1)
	og = density.NewFromSG(((points / r.BatchSize.Gallons) * 0.001) + 1)
	fg = density.NewFromSG(((og.SG - 1) * (1 - r.Attenuation)) + 1)
	abv = alcohol.Calculate(og, fg)
	return
}

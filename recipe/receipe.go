package recipe

import (
	alcohol "github.com/alancesar/homebrew/abv"
	"github.com/alancesar/homebrew/color"
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/volume"
	"math"
)

type Recipe struct {
	OG density.Density
	FG density.Density

	Efficiency  float64
	Attenuation float64

	name          string
	wortCollected volume.Volume
	batchSize     volume.Volume
	hops          []Hop
	grains        []Grain

	expectedPreBoilDensity density.Density
	expectedOG             density.Density
	expectedFG             density.Density
	expectedABV            alcohol.Abv
}

func NewRecipe(name string) *Recipe {
	return &Recipe{
		name: name,
	}
}

func (r *Recipe) WithWortCollected(wortCollected volume.Volume) *Recipe {
	r.wortCollected = wortCollected
	r.calculateExpectedGravity()
	return r
}

func (r *Recipe) WithBatchSize(batchSize volume.Volume) *Recipe {
	r.batchSize = batchSize
	r.calculateExpectedGravity()
	return r
}

func (r *Recipe) WithHops(hops ...Hop) *Recipe {
	r.hops = hops
	r.calculateExpectedGravity()
	return r
}

func (r *Recipe) WithGrains(grains ...Grain) *Recipe {
	r.grains = grains
	r.calculateExpectedGravity()
	return r
}

func (r *Recipe) Color() color.Color {
	var mcu float64
	for _, input := range r.grains {
		mcu += input.Quantity.Pounds * input.Color.Lovibond / r.batchSize.Gallons
	}
	srm := 1.4922 * math.Pow(mcu, 0.6859)
	return color.NewFromSRM(srm)
}

func (r *Recipe) ABV() alcohol.Abv {
	return alcohol.Calculate(r.OG, r.FG)
}

func (r *Recipe) IBU() float64 {
	boilGravity := (r.batchSize.Gallons / r.wortCollected.Gallons) * (r.OG.SG - 1)
	var ibu float64

	for _, input := range r.hops {
		boilFactor := 1.65 * math.Pow(0.000125, boilGravity)
		timeFactor := (1 - math.Pow(math.E, -(0.04*float64(input.BoilTime)))) / 4.15
		ibu += ((boilFactor * timeFactor) * 1.1) * (((input.AlphaAcids) *
			input.Quantity.Ounces * 7490) / r.batchSize.Gallons)
	}

	return ibu
}

func (r *Recipe) ExpectedPreBoilDensity() density.Density {
	return r.expectedPreBoilDensity
}

func (r *Recipe) ExpectedOG() density.Density {
	return r.expectedOG
}

func (r *Recipe) ExpectedFG() density.Density {
	return r.expectedFG
}

func (r *Recipe) ExpectedABV() alcohol.Abv {
	return r.expectedABV
}

func (r *Recipe) calculateExpectedGravity() {
	var mashingPoints, notMashingPoints float64

	for _, input := range r.grains {
		if input.Mashing {
			mashingPoints += input.PPG * input.Quantity.Pounds
		} else {
			notMashingPoints += input.PPG * input.Quantity.Pounds
		}
	}

	points := (mashingPoints * r.Efficiency) + notMashingPoints
	r.expectedPreBoilDensity = density.NewFromSG(((points / r.wortCollected.Gallons) * 0.001) + 1)
	r.expectedOG = density.NewFromSG(((points / r.batchSize.Gallons) * 0.001) + 1)
	r.expectedFG = density.NewFromSG(((r.expectedOG.SG - 1) * (1 - r.Attenuation)) + 1)
	r.expectedABV = alcohol.Calculate(r.expectedOG, r.expectedFG)
	return
}

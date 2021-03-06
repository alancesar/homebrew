package recipe

import (
	"github.com/alancesar/homebrew/alcohol"
	"github.com/alancesar/homebrew/bitterness"
	"github.com/alancesar/homebrew/color"
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/fermentable"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/measure"
	"github.com/alancesar/homebrew/volume"
	"math"
)

const (
	defaultEfficiency  = 0.75
	defaultAttenuation = 0.72
)

type bitternessCalculator interface {
	Calculate(hops []hop.Hop, wortGravity density.Density, batchSize volume.Volume) bitterness.Bitterness
	Method() bitterness.Method
}

var (
	calculators = []bitternessCalculator{
		bitterness.NewTinsethCalculator(),
		bitterness.NewRagerCalculator(),
		bitterness.NewDanielCalculator(),
	}
)

type Recipe struct {
	name          string
	efficiency    float64
	attenuation   float64
	og            density.Density
	fg            density.Density
	wortCollected volume.Volume
	batchSize     volume.Volume
	hops          []hop.Hop
	fermentable   []fermentable.Fermentable

	expectedPBG density.Density
	expectedOG  density.Density
	expectedFG  density.Density
	expectedABV alcohol.Alcohol
}

func NewRecipe(name string) *Recipe {
	return &Recipe{
		name:        name,
		efficiency:  defaultEfficiency,
		attenuation: defaultAttenuation,
	}
}

func (r *Recipe) WithOG(og density.Density) *Recipe {
	r.og = og
	return r
}

func (r *Recipe) WithFG(fg density.Density) *Recipe {
	r.fg = fg
	return r
}

func (r *Recipe) WithEfficiency(efficiency float64) *Recipe {
	r.efficiency = efficiency
	return r
}

func (r *Recipe) WithAttenuation(attenuation float64) *Recipe {
	r.attenuation = attenuation
	return r
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

func (r *Recipe) WithHops(hops ...hop.Hop) *Recipe {
	r.hops = hops
	r.calculateExpectedGravity()
	return r
}

func (r *Recipe) WithFermentable(f ...fermentable.Fermentable) *Recipe {
	r.fermentable = f
	r.calculateExpectedGravity()
	return r
}

func (r *Recipe) Color() color.Color {
	var mcu float64
	for _, input := range r.fermentable {
		mcu += input.Quantity.Pounds * input.Color.Lovibond / r.batchSize.Gallons
	}
	srm := 1.4922 * math.Pow(mcu, 0.6859)
	return color.NewFromSRM(srm)
}

func (r *Recipe) Alcohol() alcohol.Alcohol {
	if measure.HasSomeZeroValue(r.og, r.fg) {
		return alcohol.Alcohol{}
	}

	return alcohol.Calculate(r.og, r.fg)
}

func (r *Recipe) Bitterness() (table bitterness.Table) {
	if measure.HasSomeZeroValue(r.og, r.batchSize, r.wortCollected) {
		return table
	}

	wortGravity := density.NewFromSG(((r.batchSize.Gallons / r.wortCollected.Gallons) * (r.og.SG - 1)) + 1)
	for _, calculator := range calculators {
		table = append(table, bitterness.TableItem{
			Method: calculator.Method(),
			Value:  calculator.Calculate(r.hops, wortGravity, r.batchSize),
		})
	}

	return table
}

func (r *Recipe) ExpectedPBG() density.Density {
	return r.expectedPBG
}

func (r *Recipe) ExpectedOG() density.Density {
	return r.expectedOG
}

func (r *Recipe) ExpectedFG() density.Density {
	return r.expectedFG
}

func (r *Recipe) ExpectedAlcohol() alcohol.Alcohol {
	return r.expectedABV
}

func (r *Recipe) calculateExpectedGravity() {
	if measure.HasSomeZeroValue(r.wortCollected, r.batchSize) {
		return
	}

	mashingPoints, notMashingPoints := calculatePPGPoints(r.fermentable)
	points := (mashingPoints * r.efficiency) + notMashingPoints
	r.expectedPBG = density.NewFromSG(((points / r.wortCollected.Gallons) * 0.001) + 1)
	r.expectedOG = density.NewFromSG(((points / r.batchSize.Gallons) * 0.001) + 1)
	r.expectedFG = density.NewFromSG(((r.expectedOG.SG - 1) * (1 - r.attenuation)) + 1)
	r.expectedABV = alcohol.Calculate(r.expectedOG, r.expectedFG)
}

func calculatePPGPoints(f []fermentable.Fermentable) (mashingPoints float64, notMashingPoints float64) {
	for _, input := range f {
		if input.Mashing {
			mashingPoints += input.PPG * input.Quantity.Pounds
		} else {
			notMashingPoints += input.PPG * input.Quantity.Pounds
		}
	}

	return mashingPoints, notMashingPoints
}

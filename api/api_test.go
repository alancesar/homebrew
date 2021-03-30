package api

import (
	"github.com/alancesar/homebrew/color"
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/fermentable"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/mass"
	"github.com/alancesar/homebrew/recipe"
	"github.com/alancesar/homebrew/volume"
	"reflect"
	"testing"
)

func buildTestRecipe() *recipe.Recipe {
	return recipe.NewRecipe("Test Recipe").
		WithEfficiency(0.75).
		WithAttenuation(0.72).
		WithOG(density.NewFromSG(1.040)).
		WithFG(density.NewFromBrix(6)).
		WithWortCollected(volume.NewFromLiter(64)).
		WithBatchSize(volume.NewFromLiter(40)).
		WithHops(
			hop.Hop{
				Quantity:   mass.NewFromGram(40),
				BoilTime:   45,
				AlphaAcids: 0.15,
				Pellet:     true,
				DryHopping: false,
			},
			hop.Hop{
				Quantity:   mass.NewFromGram(20),
				BoilTime:   15,
				AlphaAcids: 0.05,
				Pellet:     false,
				DryHopping: false,
			},
			hop.Hop{
				Quantity:   mass.NewFromGram(50),
				BoilTime:   0,
				AlphaAcids: 0.02,
				Pellet:     true,
				DryHopping: true,
			},
		).
		WithFermentable(
			fermentable.Fermentable{
				Quantity: mass.NewFromKilogram(4),
				Color:    color.NewFromLovibond(5),
				PPG:      40,
				Mashing:  true,
			},
			fermentable.Fermentable{
				Quantity: mass.NewFromGram(400),
				Color:    color.NewFromLovibond(4),
				PPG:      20,
				Mashing:  false,
			},
		)
}

func TestBuildRecipe(t *testing.T) {
	type args struct {
		r Recipe
	}
	tests := []struct {
		name string
		args args
		want *recipe.Recipe
	}{
		{
			name: "Should parse API entity to Recipe entity",
			args: args{
				r: Recipe{
					Name:          "Test Recipe",
					Efficiency:    75,
					Attenuation:   72,
					OG:            1.040,
					FG:            6,
					WortCollected: "64l",
					BatchSize:     "40l",
					Hops: []Hop{
						{
							Quantity:   "40g",
							BoilTime:   45,
							AlphaAcids: 15,
							Pellet:     true,
							DryHopping: false,
						},
						{
							Quantity:   "20g",
							BoilTime:   15,
							AlphaAcids: 5,
							Pellet:     false,
							DryHopping: false,
						},
						{
							Quantity:   "50g",
							BoilTime:   0,
							AlphaAcids: 2,
							Pellet:     true,
							DryHopping: true,
						},
					},
					Fermentable: []Fermentable{
						{
							Quantity: "4kg",
							Lovibond: 5,
							PPG:      40,
							Mashing:  true,
						},
						{
							Quantity: "400g",
							Lovibond: 4,
							PPG:      20,
							Mashing:  false,
						},
					},
				},
			},
			want: buildTestRecipe(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildRecipe(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildRecipe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildResponse(t *testing.T) {
	type args struct {
		r recipe.Recipe
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Should parse Recipe entity to API response",
			args: args{
				r: *buildTestRecipe(),
			},
			want: Response{
				Bitterness: []Bitterness{
					{
						Method: "Tinseth",
						Value: BitternessValue{
							IBU: 47,
						},
					},
					{
						Method: "Rager",
						Value: BitternessValue{
							IBU: 42,
						},
					},
					{
						Method: "Daniel",
						Value: BitternessValue{
							IBU: 36,
						},
					},
				},
				Alcohol: Alcohol{
					ABV:         2.18,
					ABW:         1.74,
					Attenuation: 40.79,
				},
				Color: Color{
					SRM:      4,
					EBC:      8,
					Lovibond: 4,
				},
				ExpectedPreBoilDensity: Density{
					SG:   1.017,
					Brix: 4,
				},
				ExpectedOG: Density{
					SG:   1.027,
					Brix: 7,
				},
				ExpectedFG: Density{
					SG:   1.007,
					Brix: 2,
				},
				ExpectedABV: Alcohol{
					ABV:         2.48,
					ABW:         1.98,
					Attenuation: 72,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildResponse(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

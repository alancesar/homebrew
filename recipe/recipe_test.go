package recipe

import (
	"github.com/alancesar/homebrew/alcohol"
	"github.com/alancesar/homebrew/bitterness"
	"github.com/alancesar/homebrew/color"
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/fermentable"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/mass"
	"github.com/alancesar/homebrew/volume"
	"reflect"
	"testing"
)

func buildRecipe() *Recipe {
	return NewRecipe("Test Recipe").WithBatchSize(volume.NewFromLiter(40)).
		WithWortCollected(volume.NewFromLiter(64)).
		WithFermentable(
			fermentable.Fermentable{
				PPG:      37,
				Quantity: mass.NewFromKilogram(7),
				Mashing:  true,
			},
			fermentable.Fermentable{
				PPG:      33,
				Quantity: mass.NewFromGram(700),
				Mashing:  true,
			},
			fermentable.Fermentable{
				PPG:      38,
				Quantity: mass.NewFromKilogram(6),
				Mashing:  true,
			})
}

func TestRecipe_Color(t *testing.T) {
	tests := []struct {
		name   string
		recipe *Recipe
		want   color.Color
	}{
		{
			name: "Should calculate color of the recipe",
			recipe: NewRecipe("Test Recipe").
				WithBatchSize(volume.NewFromLiter(40)).
				WithFermentable(
					fermentable.Fermentable{
						Quantity: mass.NewFromKilogram(4),
						Color:    color.NewFromSRM(3),
					},
					fermentable.Fermentable{
						Quantity: mass.NewFromKilogram(4),
						Color:    color.NewFromSRM(5),
					},
					fermentable.Fermentable{
						Quantity: mass.NewFromKilogram(2),
						Color:    color.NewFromSRM(5),
					}),
			want: color.Color{
				SRM:      6.018894942588536,
				EBC:      11.857223036899414,
				Lovibond: 5.004351795798417,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.recipe
			if got := r.Color(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecipe_Alcohol(t *testing.T) {
	tests := []struct {
		name   string
		recipe *Recipe
		want   alcohol.Alcohol
	}{
		{
			name: "Should calculate Alcohol value",
			recipe: NewRecipe("Test Recipe").
				WithOG(density.NewFromSG(1.042)).
				WithFG(density.NewFromSG(1.008)),
			want: alcohol.Alcohol{
				ABV:         0.04480076975680501,
				ABW:         0.03584061580544401,
				Attenuation: 0.8095238095238095,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.recipe
			if got := r.Alcohol(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Alcohol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecipe_Bitterness(t *testing.T) {
	tests := []struct {
		name   string
		recipe *Recipe
		want   bitterness.Table
	}{
		{
			name: "Should calculate expected Bitterness",
			recipe: NewRecipe("Test Recipe").
				WithBatchSize(volume.NewFromLiter(36)).
				WithWortCollected(volume.NewFromLiter(42)).
				WithOG(density.NewFromSG(1.063)).
				WithHops(
					hop.Hop{
						Quantity:   mass.NewFromGram(54),
						AlphaAcids: 0.157,
						BoilTime:   60,
						Pellet:     true,
					},
					hop.Hop{
						Quantity:   mass.NewFromGram(44),
						AlphaAcids: 0.078,
						BoilTime:   5,
						Pellet:     true,
					},
					hop.Hop{
						Quantity:   mass.NewFromGram(42),
						AlphaAcids: 0.122,
						BoilTime:   5,
						Pellet:     true,
					},
					hop.Hop{
						Quantity:   mass.NewFromGram(20),
						AlphaAcids: 0.161,
						BoilTime:   5,
						Pellet:     true,
					},
					hop.Hop{
						Quantity:   mass.NewFromGram(43),
						AlphaAcids: 0.105,
						BoilTime:   0,
						Pellet:     true,
					},
				),
			want: bitterness.Table{
				bitterness.TableItem{
					Method: "Tinseth",
					Value:  bitterness.NewFromIBU(73.6053858587605),
				},
				{
					Method: "Rager",
					Value:  bitterness.NewFromIBU(95.80235532474533),
				},
				{
					Method: "Daniel",
					Value:  bitterness.NewFromIBU(79.1463888888889),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.recipe
			if got := r.Bitterness(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bitterness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecipe_ExpectedPBG(t *testing.T) {
	tests := []struct {
		name   string
		recipe *Recipe
		want   density.Density
	}{
		{
			name:   "Should calculate expected pre boil density",
			recipe: buildRecipe(),
			want:   density.NewFromSG(1.049886651759173),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.recipe
			if got := r.ExpectedPBG(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpectedPBG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecipe_ExpectedOG(t *testing.T) {
	tests := []struct {
		name   string
		recipe *Recipe
		want   density.Density
	}{
		{
			name:   "Should calculate expected OG",
			recipe: buildRecipe(),
			want:   density.NewFromSG(1.0798186428146765),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.recipe
			if got := r.ExpectedOG(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpectedOG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecipe_ExpectedFG(t *testing.T) {
	tests := []struct {
		name   string
		recipe *Recipe
		want   density.Density
	}{
		{
			name:   "Should calculate expected FG",
			recipe: buildRecipe(),
			want:   density.NewFromSG(1.0223492199881095),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.recipe
			if got := r.ExpectedFG(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpectedFG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecipe_ExpectedABV(t *testing.T) {
	tests := []struct {
		name   string
		recipe *Recipe
		want   alcohol.Alcohol
	}{
		{
			name:   "Should calculate expected Alcohol",
			recipe: buildRecipe(),
			want: alcohol.Alcohol{
				ABV:         0.08098190520848729,
				ABW:         0.06478552416678983,
				Attenuation: 0.7199999999999996,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.recipe
			if got := r.ExpectedAlcohol(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpectedAlcohol() = %v, want %v", got, tt.want)
			}
		})
	}
}

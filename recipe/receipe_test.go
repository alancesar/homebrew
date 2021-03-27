package recipe

import (
	"github.com/alancesar/homebrew/abv"
	"github.com/alancesar/homebrew/color"
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/mass"
	"github.com/alancesar/homebrew/volume"
	"reflect"
	"testing"
)

func buildRecipe() *Recipe {
	r := &Recipe{
		Efficiency:  0.75,
		Attenuation: 0.72,
	}

	return r.WithBatchSize(volume.NewFromLiter(40)).
		WithWortCollected(volume.NewFromLiter(64)).
		WithGrains(
			Grain{
				PPG:      37,
				Quantity: mass.NewFromKilogram(7),
				Mashing:  true,
			},
			Grain{
				PPG:      33,
				Quantity: mass.NewFromGram(700),
				Mashing:  true,
			},
			Grain{
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
			recipe: NewRecipe("Test recipe").
				WithBatchSize(volume.NewFromLiter(40)).
				WithGrains(
					Grain{
						Quantity: mass.NewFromKilogram(4),
						Color:    color.NewFromSRM(3),
					},
					Grain{
						Quantity: mass.NewFromKilogram(4),
						Color:    color.NewFromSRM(5),
					},
					Grain{
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

func TestRecipe_ABV(t *testing.T) {
	type fields struct {
		Og density.Density
		Fg density.Density
	}
	tests := []struct {
		name   string
		fields fields
		want   abv.Abv
	}{
		{
			name: "Should calculate ABV value",
			fields: fields{
				Og: density.NewFromSG(1.042),
				Fg: density.NewFromSG(1.008),
			},
			want: abv.Abv{
				Abv:         0.04480076975680501,
				Abw:         0.03584061580544401,
				Attenuation: 0.8095238095238095,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Recipe{
				OG: tt.fields.Og,
				FG: tt.fields.Fg,
			}
			if got := r.ABV(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ABV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecipe_IBU(t *testing.T) {
	type fields struct {
		Og            density.Density
		BatchSize     volume.Volume
		WortCollected volume.Volume
		Hops          []Hop
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should calculate IBU",
			fields: fields{
				Og:            density.NewFromSG(1.063),
				BatchSize:     volume.NewFromLiter(36),
				WortCollected: volume.NewFromLiter(42),
				Hops: []Hop{
					{
						Quantity:   mass.NewFromGram(54),
						AlphaAcids: 0.157,
						BoilTime:   60,
					},
					{
						Quantity:   mass.NewFromGram(44),
						AlphaAcids: 0.078,
						BoilTime:   5,
					},
					{
						Quantity:   mass.NewFromGram(42),
						AlphaAcids: 0.122,
						BoilTime:   5,
					},
					{
						Quantity:   mass.NewFromGram(20),
						AlphaAcids: 0.161,
						BoilTime:   5,
					},
					{
						Quantity:   mass.NewFromGram(43),
						AlphaAcids: 0.105,
						BoilTime:   0,
					},
				},
			},
			want: 73.61378315374786,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Recipe{
				OG:            tt.fields.Og,
				batchSize:     tt.fields.BatchSize,
				wortCollected: tt.fields.WortCollected,
				hops:          tt.fields.Hops,
			}
			if got := r.IBU(); got != tt.want {
				t.Errorf("IBU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecipe_ExpectedPreBoilDensity(t *testing.T) {
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
			if got := r.ExpectedPreBoilDensity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpectedPreBoilDensity() = %v, want %v", got, tt.want)
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
		want   abv.Abv
	}{
		{
			name:   "Should calculate expected ABV",
			recipe: buildRecipe(),
			want: abv.Abv{
				Abv:         0.08098190520848729,
				Abw:         0.06478552416678983,
				Attenuation: 0.7199999999999996,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.recipe
			if got := r.ExpectedABV(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpectedABV() = %v, want %v", got, tt.want)
			}
		})
	}
}

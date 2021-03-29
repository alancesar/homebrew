package ibu

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/mass"
	"github.com/alancesar/homebrew/volume"
	"testing"
)

func TestCalculateTinseth(t *testing.T) {
	type args struct {
		hops        []hop.Hop
		wortGravity density.Density
		batchSize   volume.Volume
	}
	tests := []struct {
		name    string
		args    args
		wantIbu float64
	}{
		{
			name: "Should calculate IBU using Tinseth's formula",
			args: args{
				hops: []hop.Hop{
					{
						Quantity:   mass.NewFromGram(54),
						AlphaAcids: 0.157,
						BoilTime:   60,
						Pellet:     true,
					},
					{
						Quantity:   mass.NewFromGram(44),
						AlphaAcids: 0.078,
						BoilTime:   5,
						Pellet:     true,
					},
					{
						Quantity:   mass.NewFromGram(42),
						AlphaAcids: 0.122,
						BoilTime:   5,
						Pellet:     true,
					},
					{
						Quantity:   mass.NewFromGram(20),
						AlphaAcids: 0.161,
						BoilTime:   5,
						Pellet:     true,
					},
					{
						Quantity:   mass.NewFromGram(43),
						AlphaAcids: 0.105,
						BoilTime:   0,
						Pellet:     true,
					},
				},
				wortGravity: density.NewFromSG(1.054),
				batchSize:   volume.NewFromLiter(36),
			},
			wantIbu: 73.6137831537478,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIbu := CalculateTinseth(tt.args.hops, tt.args.wortGravity, tt.args.batchSize); gotIbu != tt.wantIbu {
				t.Errorf("CalculateTinseth() = %v, want %v", gotIbu, tt.wantIbu)
			}
		})
	}
}

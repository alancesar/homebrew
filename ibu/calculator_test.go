package ibu

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/mass"
	"github.com/alancesar/homebrew/volume"
	"testing"
)

func TestCalculateGaretz(t *testing.T) {
	type args struct {
		hops          []hop.Hop
		wortGravity   density.Density
		wortCollected volume.Volume
		batchSize     volume.Volume
	}
	tests := []struct {
		name    string
		args    args
		wantIbu float64
	}{
		{
			name: "Should calculate IBU using Garetz's formula",
			args: args{
				hops: []hop.Hop{
					{
						Quantity:   mass.NewFromOunce(1.5),
						BoilTime:   45,
						AlphaAcids: 0.064,
						Pellet:     false,
					},
					{
						Quantity:   mass.NewFromOunce(1),
						BoilTime:   15,
						AlphaAcids: 0.05,
						Pellet:     false,
					},
				},
				wortGravity:   density.NewFromSG(1.050),
				wortCollected: volume.NewFromGallon(6.5),
				batchSize:     volume.NewFromGallon(5),
			},
			wantIbu: 26.360484634085594,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIbu := CalculateGaretz(tt.args.hops, tt.args.wortGravity, tt.args.wortCollected, tt.args.batchSize)
			if gotIbu != tt.wantIbu {
				t.Errorf("CalculateGaretz() = %v, wantIbu %v", gotIbu, tt.wantIbu)
			}
		})
	}
}
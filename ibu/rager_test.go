package ibu

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/mass"
	"github.com/alancesar/homebrew/volume"
	"testing"
)

func TestRager_Calculate(t *testing.T) {
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
			name: "Should calculate IBU using Rager's formula",
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
				wortGravity: density.NewFromSG(1.050),
				batchSize:   volume.NewFromGallon(5),
			},
			wantIbu: 44.84610801652261,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rager{}
			if gotIbu := r.Calculate(tt.args.hops, tt.args.wortGravity, tt.args.batchSize); gotIbu != tt.wantIbu {
				t.Errorf("Calculate() = %v, want %v", gotIbu, tt.wantIbu)
			}
		})
	}
}

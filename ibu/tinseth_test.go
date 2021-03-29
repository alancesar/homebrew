package ibu

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/mass"
	"github.com/alancesar/homebrew/volume"
	"testing"
)

func TestTinseth_Calculate(t1 *testing.T) {
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
			wantIbu: 39.0188834297751,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tinseth{}
			if gotIbu := t.Calculate(tt.args.hops, tt.args.wortGravity, tt.args.batchSize); gotIbu != tt.wantIbu {
				t1.Errorf("Calculate() = %v, want %v", gotIbu, tt.wantIbu)
			}
		})
	}
}

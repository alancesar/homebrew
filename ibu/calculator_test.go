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
			wantIbu: 39.02333491210092,
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

func TestCalculateRoger(t *testing.T) {
	type args struct {
		hops        []hop.Hop
		wortGravity density.Density
		batchSize   volume.Volume
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Should calculate IBU using Roger's formula",
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
			want: 44.683556171485115,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateRoger(tt.args.hops, tt.args.wortGravity, tt.args.batchSize); got != tt.want {
				t.Errorf("CalculateRoger() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			wantIbu: 39.0188834297751,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIbu := CalculateTinseth(tt.args.hops, tt.args.wortGravity, tt.args.batchSize); gotIbu != tt.wantIbu {
				t.Errorf("CalculateTinseth() = %v, wantIbu %v", gotIbu, tt.wantIbu)
			}
		})
	}
}

func TestCalculateRager(t *testing.T) {
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
			if got := CalculateRager(tt.args.hops, tt.args.wortGravity, tt.args.batchSize); got != tt.wantIbu {
				t.Errorf("CalculateRager() = %v, wantIbu %v", got, tt.wantIbu)
			}
		})
	}
}

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

func TestCalculateDaniel(t *testing.T) {
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
			name: "Should calculate IBU using Daniel's formula",
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
			wantIbu: 40.621125725327666,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIbu := CalculateDaniel(tt.args.hops, tt.args.wortGravity, tt.args.batchSize); gotIbu != tt.wantIbu {
				t.Errorf("CalculateDaniel() = %v, want %v", gotIbu, tt.wantIbu)
			}
		})
	}
}

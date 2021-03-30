package bitterness

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/mass"
	"github.com/alancesar/homebrew/volume"
	"reflect"
	"testing"
)

type bitternessTestArgs struct {
	hops        []hop.Hop
	wortGravity density.Density
	batchSize   volume.Volume
}

func buildBasicTestArgs() bitternessTestArgs {
	return bitternessTestArgs{
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
		name string
		args args
		want Bitterness
	}{
		{
			name: "Should calculate Bitterness using Garetz's formula",
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
			want: NewFromIBU(26.360484634085594),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateGaretz(tt.args.hops, tt.args.wortGravity, tt.args.wortCollected, tt.args.batchSize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateGaretz() = %v, want %v", got, tt.want)
			}
		})
	}
}

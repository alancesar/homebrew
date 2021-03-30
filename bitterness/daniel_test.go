package bitterness

import (
	"github.com/alancesar/homebrew/density"
	"github.com/alancesar/homebrew/hop"
	"github.com/alancesar/homebrew/mass"
	"github.com/alancesar/homebrew/volume"
	"reflect"
	"testing"
)

func TestDaniel_Calculate(t *testing.T) {
	type fields struct {
		keys []int
	}
	type args struct {
		hops        []hop.Hop
		wortGravity density.Density
		batchSize   volume.Volume
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Bitterness
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
			want: NewFromIBU(40.621125725327666),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Daniel{}
			got := d.Calculate(tt.args.hops, tt.args.wortGravity, tt.args.batchSize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

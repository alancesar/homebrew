package abv

import (
	"github.com/alancesar/homebrew/density"
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	type args struct {
		og density.Density
		fg density.Density
	}
	tests := []struct {
		name string
		args args
		want Abv
	}{
		{
			name: "Should parse properly",
			args: args{
				og: density.Sg(1.042),
				fg: density.Sg(1.008),
			},
			want: Abv{
				Abv:         0.04480076975680501,
				Abw:         0.03584061580544401,
				Attenuation: 0.8095238095238095,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.og, tt.args.fg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

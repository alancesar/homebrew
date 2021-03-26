package density

import (
	"reflect"
	"testing"
)

func TestRefractometerCorrection(t *testing.T) {
	type args struct {
		og             Density
		fg             Density
		wortCorrection float64
	}
	tests := []struct {
		name string
		args args
		want Density
	}{
		{
			name: "Should calculate properly",
			args: args{
				og:             NewFromSg(1.053),
				fg:             NewFromBrix(8),
				wortCorrection: 1,
			},
			want: NewFromSg(1.0172375951712538),
		},
		{
			name: "Should calculate properly using brix",
			args: args{
				og:             NewFromBrix(12),
				fg:             NewFromBrix(8),
				wortCorrection: 1.04,
			},
			want: NewFromSg(1.019038422451077),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RefractometerCorrection(tt.args.og, tt.args.fg, tt.args.wortCorrection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefractometerCorrection() = %v, want %v", got, tt.want)
			}
		})
	}
}

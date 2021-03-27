package pressure

import (
	"github.com/alancesar/homebrew/temperature"
	"testing"
)

func TestCalculateVolumesOfCarbonation(t *testing.T) {
	type args struct {
		pressure    Pressure
		temperature temperature.Temperature
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Should calculate volumes of carbonation",
			args: args{
				pressure:    NewFromPSI(14.0),
				temperature: temperature.NewFromCelsius(5),
			},
			want: 2.6178265248435264,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateVolumesOfCarbonation(tt.args.pressure, tt.args.temperature); got != tt.want {
				t.Errorf("CalculateVolumesOfCarbonation() = %v, want %v", got, tt.want)
			}
		})
	}
}

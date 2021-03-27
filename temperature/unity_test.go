package temperature

import (
	"reflect"
	"testing"
)

func TestNewFromCelsius(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Temperature
	}{
		{
			name: "Should parse from Celsius",
			args: args{
				value: 50,
			},
			want: Temperature{
				Celsius:    50,
				Fahrenheit: 122,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromCelsius(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromCelsius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromFahrenheit(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Temperature
	}{
		{
			name: "Should parse from Fahrenheit",
			args: args{
				value: 68,
			},
			want: Temperature{
				Celsius:    20,
				Fahrenheit: 68,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromFahrenheit(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromFahrenheit() = %v, want %v", got, tt.want)
			}
		})
	}
}

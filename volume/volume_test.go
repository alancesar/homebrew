package volume

import (
	"reflect"
	"testing"
)

func TestNewFromMilliliter(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Volume
	}{
		{
			name: "Should create from milliliter properly",
			args: args{
				value: 1000.0,
			},
			want: Volume{
				Milliliters: 1000.0,
				Liters:      1,
				Gallons:     0.26417205235815002,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromMilliliter(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromMilliliter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromLiter(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Volume
	}{
		{
			name: "Should create from liter properly",
			args: args{
				value: 10,
			},
			want: Volume{
				Milliliters: 10000.0,
				Liters:      10,
				Gallons:     2.6417205235815002,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromLiter(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromLiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromGallon(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Volume
	}{
		{
			name: "Should create one liter from gallon properly",
			args: args{
				value: 0.26417205235815,
			},
			want: Volume{
				Milliliters: 1000.0,
				Liters:      1,
				Gallons:     0.26417205235815,
			},
		},
		{
			name: "Should create from gallon properly",
			args: args{
				value: 1,
			},
			want: Volume{
				Milliliters: 3785.411783999977,
				Liters:      3.7854117839999772,
				Gallons:     1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromGallon(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromGallon() = %v, want %v", got, tt.want)
			}
		})
	}
}

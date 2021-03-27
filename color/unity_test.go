package color

import (
	"reflect"
	"testing"
)

func TestNewFromSrm(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Color
	}{
		{
			name: "Should parse from SRM",
			args: args{
				value: 1.0,
			},
			want: Color{
				SRM:      1.0,
				EBC:      1.97,
				Lovibond: 1.2992765391997638,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromSRM(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromSRM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromEbc(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Color
	}{
		{
			name: "Should parse from EBC",
			args: args{
				value: 1.0,
			},
			want: Color{
				SRM:      0.508,
				EBC:      1.0,
				Lovibond: 0.9360696884689207,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromEBC(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromEBC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromLovibond(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Color
	}{
		{
			name: "Should parse from Lovibond",
			args: args{
				value: 1.0,
			},
			want: Color{
				SRM:      0.5946,
				EBC:      1.171362,
				Lovibond: 1.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromLovibond(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromLovibond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_RGB(t *testing.T) {
	tests := []struct {
		name  string
		color Color
		want  string
	}{
		{
			name:  "Should convert SRM value to RGB",
			color: NewFromSRM(24),
			want:  "#821E00",
		},
		{
			name:  "Should convert SRM value to RGB with value < 1",
			color: NewFromSRM(0),
			want:  "#FFE699",
		},
		{
			name:  "Should convert SRM value to RGB with value > 40",
			color: NewFromSRM(41),
			want:  "#36080A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.color
			if got := c.RGB(); got != tt.want {
				t.Errorf("RGB() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
				Srm:      1.0,
				Ebc:      1.97,
				Lovibond: 1.2992765391997638,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromSrm(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromSrm() = %v, want %v", got, tt.want)
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
				Srm:      0.508,
				Ebc:      1.0,
				Lovibond: 0.9360696884689207,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromEbc(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromEbc() = %v, want %v", got, tt.want)
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
				Srm:      0.5946,
				Ebc:      1.171362,
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

package pressure

import (
	"reflect"
	"testing"
)

func TestNewFromPSI(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Pressure
	}{
		{
			name: "Should parse from PSI",
			args: args{
				value: 1,
			},
			want: Pressure{
				PSI:    1,
				Kgfcm2: 0.07030865499542993,
				Bar:    0.06894898108836733,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromPSI(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromPSI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromKgfcm2(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Pressure
	}{
		{
			name: "Should parse from Kgf/cm²",
			args: args{
				value: 1,
			},
			want: Pressure{
				PSI:    14.223,
				Kgfcm2: 1,
				Bar:    0.9806613580198487,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromKgfcm2(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromKgfcm2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromBar(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Pressure
	}{
		{
			name: "Should parse from Bar",
			args: args{
				value: 1,
			},
			want: Pressure{
				PSI:    14.50347756,
				Kgfcm2: 1.01972,
				Bar:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromBar(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromBar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFrom(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want Pressure
	}{
		{
			name: "Should parse from '1psi' string",
			args: args{
				input: "1psi",
			},
			want: NewFromPSI(1),
		},
		{
			name: "Should parse from '1bar' string",
			args: args{
				input: "1bar",
			},
			want: NewFromBar(1),
		},
		{
			name: "Should parse from '1psi' string",
			args: args{
				input: "1psi",
			},
			want: NewFromPSI(1),
		},
		{
			name: "Should parse from '1at' string",
			args: args{
				input: "1at",
			},
			want: NewFromKgfcm2(1),
		},
		{
			name: "Should parse from '1kgfcm2' string",
			args: args{
				input: "1kgfcm2",
			},
			want: NewFromKgfcm2(1),
		},
		{
			name: "Should parse from '1kgf/cm2' string",
			args: args{
				input: "1kgf/cm2",
			},
			want: NewFromKgfcm2(1),
		},
		{
			name: "Should parse from '1kgfcm²' string",
			args: args{
				input: "1kgfcm²",
			},
			want: NewFromKgfcm2(1),
		},
		{
			name: "Should parse from '1kgf/cm²' string",
			args: args{
				input: "1kgf/cm²",
			},
			want: NewFromKgfcm2(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFrom(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

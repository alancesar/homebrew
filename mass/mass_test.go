package mass

import (
	"reflect"
	"testing"
)

func TestNewFromMilligram(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Mass
	}{
		{
			name: "Should parse properly from milligram",
			args: args{
				value: 1000000.0,
			},
			want: Mass{
				Milligrams: 1000000.0,
				Grams:      1000.0,
				Kilograms:  1.0,
				Pounds:     2.2046244201837775,
				Ounces:     35.27399072294044,
			},
		},
		{
			name: "Should parse properly from pound",
			args: args{
				value: 453592,
			},
			want: Mass{
				Milligrams: 453592.0,
				Grams:      453.592,
				Kilograms:  0.453592,
				Pounds:     1,
				Ounces:     16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromMilligram(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromMilligram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromGram(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Mass
	}{
		{
			name: "Should parse properly from gram",
			args: args{
				value: 1000.0,
			},
			want: Mass{
				Milligrams: 1000000.0,
				Grams:      1000.0,
				Kilograms:  1.0,
				Pounds:     2.2046244201837775,
				Ounces:     35.27399072294044,
			},
		},
		{
			name: "Should parse properly from pound",
			args: args{
				value: 453.592,
			},
			want: Mass{
				Milligrams: 453592.0,
				Grams:      453.592,
				Kilograms:  0.453592,
				Pounds:     1,
				Ounces:     16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromGram(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromGram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromKilogram(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Mass
	}{
		{
			name: "Should parse properly from kilogram",
			args: args{
				value: 1.0,
			},
			want: Mass{
				Milligrams: 1000000.0,
				Grams:      1000.0,
				Kilograms:  1.0,
				Pounds:     2.2046244201837775,
				Ounces:     35.27399072294044,
			},
		},
		{
			name: "Should parse properly from pound",
			args: args{
				value: 0.453592,
			},
			want: Mass{
				Milligrams: 453592.0,
				Grams:      453.592,
				Kilograms:  0.453592,
				Pounds:     1,
				Ounces:     16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromKilogram(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromKilogram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromPound(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Mass
	}{
		{
			name: "Should parse properly from pound",
			args: args{
				value: 1.0,
			},
			want: Mass{
				Milligrams: 453592,
				Grams:      453.592,
				Kilograms:  0.453592,
				Pounds:     1,
				Ounces:     16,
			},
		},
		{
			name: "Should parse properly from gram",
			args: args{
				value: 2.20462262185,
			},
			want: Mass{
				Milligrams: 999999.1842901852,
				Grams:      999.9991842901852,
				Kilograms:  0.9999991842901852,
				Pounds:     2.20462262185,
				Ounces:     35.2739619496,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromPound(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromPound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromOunce(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Mass
	}{
		{
			name: "Should parse properly from once",
			args: args{
				value: 1.0,
			},
			want: Mass{
				Milligrams: 28349.5,
				Grams:      28.3495,
				Kilograms:  0.0283495,
				Pounds:     0.0625,
				Ounces:     1,
			},
		},
		{
			name: "Should parse properly from gram",
			args: args{
				value: 35.27396195,
			},
			want: Mass{
				Milligrams: 999999.184301525,
				Grams:      999.999184301525,
				Kilograms:  0.999999184301525,
				Pounds:     2.204622621875,
				Ounces:     35.27396195,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromOunce(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromOunce() = %v, want %v", got, tt.want)
			}
		})
	}
}

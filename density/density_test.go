package density

import (
	"reflect"
	"testing"
)

func TestNewFromSg(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Density
	}{
		{
			name: "Should parse from SG to Brix",
			args: args{
				value: 1.092,
			},
			want: Density{
				SG:   1.092,
				Brix: 22.014119055148626,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromSG(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromSG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromBrix(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Density
	}{
		{
			name: "Should parse from Brix to SG",
			args: args{
				value: 22,
			},
			want: Density{
				SG:   1.0919540676449373,
				Brix: 22.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromBrix(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromBrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

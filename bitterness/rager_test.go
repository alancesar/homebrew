package bitterness

import (
	"reflect"
	"testing"
)

func TestRager_Calculate(t *testing.T) {
	tests := []struct {
		name string
		args bitternessTestArgs
		want Bitterness
	}{
		{
			name: "Should calculate Bitterness using Rager's formula",
			args: buildBasicTestArgs(),
			want: NewFromIBU(44.84610801652261),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rager{}
			got := r.Calculate(tt.args.hops, tt.args.wortGravity, tt.args.batchSize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

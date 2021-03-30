package bitterness

import (
	"reflect"
	"testing"
)

func TestTinseth_Calculate(t1 *testing.T) {
	tests := []struct {
		name string
		args bitternessTestArgs
		want Bitterness
	}{
		{
			name: "Should calculate IBU using Tinseth's formula",
			args: buildBasicTestArgs(),
			want: NewFromIBU(39.0188834297751),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tinseth{}
			got := t.Calculate(tt.args.hops, tt.args.wortGravity, tt.args.batchSize)
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

package bitterness

import (
	"reflect"
	"testing"
)

func TestDaniel_Calculate(t *testing.T) {
	type fields struct {
		keys []int
	}

	tests := []struct {
		name   string
		fields fields
		args   bitternessTestArgs
		want   Bitterness
	}{
		{
			name: "Should calculate Bitterness using Daniel's formula",
			args: buildBasicTestArgs(),
			want: NewFromIBU(40.621125725327666),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Daniel{}
			got := d.Calculate(tt.args.hops, tt.args.wortGravity, tt.args.batchSize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

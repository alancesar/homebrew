package converters

import "testing"

const (
	expectedMaxBrixValue = 22.0
	expectedMaxSgValue   = 1.092
)

func TestConvertDensity(t *testing.T) {
	sg := 1.0919
	brix := ConvertDensity(sg).FromSg().ToBrix()

	if brix > expectedMaxBrixValue {
		t.Error("Unexpected conversion")
	}

	brix = 22.0
	sg = ConvertDensity(brix).FromBrix().ToSg()

	if sg > expectedMaxSgValue {
		t.Error("Unexpected conversion")
	}
}

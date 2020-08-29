package density

import "testing"

const (
	expectedMaxBrixValue = 22.0
	expectedMaxSgValue   = 1.092
)

func TestFrom_ToBrix(t *testing.T) {
	sg := 1.0919
	brix := FromSg(sg).ToBrix()

	if brix > expectedMaxBrixValue {
		t.Error("Unexpected conversion")
	}
}

func TestFrom_ToSg(t *testing.T) {
	brix := 22.0
	sg := FromBrix(brix).ToSg()

	if sg > expectedMaxSgValue {
		t.Error("Unexpected conversion")
	}
}

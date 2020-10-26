package density

import (
	"testing"
)

func TestSg(t *testing.T) {
	density := Sg(1.0919)

	if density.Brix > expectedMaxBrixValue {
		t.Error("Error on conversion")
	}
}

func TestBrix(t *testing.T) {
	density := Brix(22.00)

	if density.Sg > expectedMaxSgValue {
		t.Error("Error on conversion")
	}
}

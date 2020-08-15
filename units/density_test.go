package units

import (
	"testing"
)

const (
	expectedMaxBrixValue = 22.0
	expectedMaxSgValue   = 1.092
)

func TestDensity_FromSg(t *testing.T) {
	density := Density{}.FromSg(1.0919)
	brix := density.Brix

	if brix.Value > expectedMaxBrixValue {
		t.Error("Error on conversion")
	}

	if brix.Symbol != brixSymbol {
		t.Error("Error on conversion")
	}
}

func TestDensity_FromBrix(t *testing.T) {
	density := Density{}.FromBrix(22.00)
	sg := density.Sg

	if sg.Value > expectedMaxSgValue {
		t.Error("Error on conversion")
	}

	if sg.Symbol != sgSymbol {
		t.Error("Error on conversion")
	}
}

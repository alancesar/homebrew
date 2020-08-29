package density

import (
	"testing"
)

func TestSg(t *testing.T) {
	density := Sg(1.0919)
	brix := density.Brix

	if brix.Value > expectedMaxBrixValue {
		t.Error("Error on conversion")
	}

	if brix.Symbol != brixSymbol {
		t.Error("Error on conversion")
	}
}

func TestBrix(t *testing.T) {
	density := Brix(22.00)
	sg := density.Sg

	if sg.Value > expectedMaxSgValue {
		t.Error("Error on conversion")
	}

	if sg.Symbol != sgSymbol {
		t.Error("Error on conversion")
	}
}

package volume

import (
	"testing"
)

func TestConvertVolume(t *testing.T) {
	liters := 1.0
	milliliters := FromLiter(liters).ToMilliliter()

	if milliliters != 1000.0 {
		t.Error("Unexpected conversion")
	}

	americanGallons := 1.0
	liters = FromGallon(americanGallons).ToLiter()

	if liters != 3.7854117839999772 {
		t.Error("Unexpected conversion")
	}
}

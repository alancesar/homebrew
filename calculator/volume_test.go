package calculator

import (
	"testing"
)

func TestConvertVolume(t *testing.T) {
	liters := 1.0
	milliliters := ConvertVolume(liters).FromLiter().ToMilliliter()

	if milliliters != 1000.0 {
		t.Error("Unexpected conversion")
	}

	americanGallons := 1.0
	liters = ConvertVolume(americanGallons).FromGallon().ToLiter()

	if liters != 3.7854117839999772 {
		t.Error("Unexpected conversion")
	}
}

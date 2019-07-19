package converter

import "testing"

func TestConvertMass(t *testing.T) {

	kilograms := 1.0
	grams := ConvertMass(kilograms).FromKilogram().ToGram()

	if grams != 1000 {
		t.Error("Unexpected conversion")
	}

	milligrams := 1000.0
	kilograms = ConvertMass(milligrams).FromMilligram().ToKilogram()

	if kilograms != 0.001 {
		t.Error("Unexpected conversion")
	}
}

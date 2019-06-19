package calculator

import "testing"

func TestConvertMass(t *testing.T) {

	kilograms := float64(1)
	grams := ConvertMass(kilograms).FromKilogram().ToGram()

	if grams != 1000 {
		t.Error("Unexpected conversion")
	}

	milligrams := float64(1000)
	kilograms = ConvertMass(milligrams).FromMilligram().ToKilogram()

	if kilograms != 0.001 {
		t.Error("Unexpected conversion")
	}
}

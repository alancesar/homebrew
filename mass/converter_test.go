package mass

import (
	"testing"
)

func TestFrom_ToGram(t *testing.T) {
	kilograms := 1.0
	grams := FromKilogram(kilograms).ToGram()

	if grams != 1000 {
		t.Error("Unexpected conversion")
	}
}

func TestFrom_ToKilogram(t *testing.T) {
	milligrams := 1000.0
	kilograms := FromMilligram(milligrams).ToKilogram()

	if kilograms != 0.001 {
		t.Error("Unexpected conversion")
	}
}

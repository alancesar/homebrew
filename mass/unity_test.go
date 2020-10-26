package mass

import (
	"testing"
)

func TestMilligram(t *testing.T) {
	mass := Milligram(1000000.0)
	if mass.Pounds != 2.2046244201837775 {
		t.Error("error on conversion")
	}
}

func TestGram(t *testing.T) {
	mass := Gram(1000.0)
	if mass.Ounces != 35.27399072294044 {
		t.Error("error on conversion")
	}
}

func TestKilogram(t *testing.T) {
	mass := Kilogram(1.0)
	if mass.Grams != 1000.0 {
		t.Error("error on conversion")
	}
}

func TestPound(t *testing.T) {
	mass := Pound(1.0)
	if mass.Milligrams != 453592 {
		t.Error("error on conversion")
	}
}

func TestOunce(t *testing.T) {
	mass := Ounce(16.0)
	if mass.Kilograms != 0.453592 {
		t.Error("error on conversion")
	}
}

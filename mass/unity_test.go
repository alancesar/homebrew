package mass

import (
	"fmt"
	"testing"
)

const (
	expectedMassResponse         = "{{1e+06 mg} {1000 g} {1 kg} {2.2046244201837775 lb} {35.27399072294044 oz}}"
	expectedImperialMassResponse = "{{453592 mg} {453.592 g} {0.453592 kg} {1 lb} {16 oz}}"
)

func TestMilligram(t *testing.T) {
	mass := Milligram(1000000.0)
	stringValue := fmt.Sprint(mass)

	if stringValue != expectedMassResponse {
		t.Error("Error on conversion")
	}
}

func TestGram(t *testing.T) {
	mass := Gram(1000.0)
	stringValue := fmt.Sprint(mass)

	if stringValue != expectedMassResponse {
		t.Error("Error on conversion")
	}
}

func TestKilogram(t *testing.T) {
	mass := Kilogram(1.0)
	stringValue := fmt.Sprint(mass)

	if stringValue != expectedMassResponse {
		t.Error("Error on conversion")
	}
}

func TestPound(t *testing.T) {
	mass := Pound(1.0)
	stringValue := fmt.Sprint(mass)

	if stringValue != expectedImperialMassResponse {
		t.Error("Error on conversion")
	}
}

func TestOunce(t *testing.T) {
	mass := Ounce(16.0)
	stringValue := fmt.Sprint(mass)

	if stringValue != expectedImperialMassResponse {
		t.Error("Error on conversion")
	}
}

package units

import (
	"fmt"
	"testing"
)

const (
	expectedMassResponse         = "{{1e+06 mg} {1000 g} {1 kg} {2.2046244201837775 lb} {35.27399072294044 oz}}"
	expectedImperialMassResponse = "{{453592 mg} {453.592 g} {0.453592 kg} {1 lb} {16 oz}}"
)

func TestFromMilligram(t *testing.T) {
	mass := Mass{}.FromMilligram(1000000.0)
	stringValue := fmt.Sprint(mass)

	if stringValue != expectedMassResponse {
		t.Error("Error on conversion")
	}
}

func TestFromGram(t *testing.T) {
	mass := Mass{}.FromGram(1000.0)
	stringValue := fmt.Sprint(mass)

	if stringValue != expectedMassResponse {
		t.Error("Error on conversion")
	}
}

func TestFromKilograms(t *testing.T) {
	mass := Mass{}.FromKilograms(1.0)
	stringValue := fmt.Sprint(mass)

	if stringValue != expectedMassResponse {
		t.Error("Error on conversion")
	}
}

func TestFromPounds(t *testing.T) {
	mass := Mass{}.FromPounds(1.0)
	stringValue := fmt.Sprint(mass)

	if stringValue != expectedImperialMassResponse {
		t.Error("Error on conversion")
	}
}

func TestFromOunces(t *testing.T) {
	mass := Mass{}.FromOunces(16.0)
	stringValue := fmt.Sprint(mass)

	if stringValue != expectedImperialMassResponse {
		t.Error("Error on conversion")
	}
}

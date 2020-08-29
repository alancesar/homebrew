package volume

import (
	"fmt"
	"testing"
)

const (
	expectedVolumeResponse = "{{1000 ml} {1 l} {0.26417205235815 gal}}"
)

func Test_Milliliter(t *testing.T) {
	volume := Milliliter(1000.0)
	stringValue := fmt.Sprint(volume)

	if stringValue != expectedVolumeResponse {
		t.Error("Error on conversion")
	}
}

func Test_Liter(t *testing.T) {
	volume := Liter(1.0)
	stringValue := fmt.Sprint(volume)

	if stringValue != expectedVolumeResponse {
		t.Error("Error on conversion")
	}
}

func Test_Gallon(t *testing.T) {
	volume := Gallon(0.26417205235815)
	stringValue := fmt.Sprint(volume)

	if stringValue != expectedVolumeResponse {
		t.Error("Error on conversion")
	}
}

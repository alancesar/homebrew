package volume

import (
	"testing"
)

func Test_Milliliter(t *testing.T) {
	volume := Milliliter(1000.0)

	if volume.Gallons != 0.26417205235815 {
		t.Error("error on conversion")
	}
}

func Test_Liter(t *testing.T) {
	volume := Liter(1.0)

	if volume.Gallons != 0.26417205235815 {
		t.Error("error on conversion")
	}
}

func Test_Gallon(t *testing.T) {
	volume := Gallon(0.26417205235815)

	if volume.Liters != 1 {
		t.Error("error on conversion")
	}
}

package units

import (
	"fmt"
	"strings"
	"testing"
)

const (
	expectedVolumeResponse = "{{1000 ml} {1 l} {0.26417205235815 gal}}"
)

func TestVolume_FromMilliliter(t *testing.T) {
	volume := Volume{}.FromMilliliter(1000.0)
	stringValue := fmt.Sprint(volume)

	if !strings.EqualFold(stringValue, expectedVolumeResponse) {
		t.Error("Error on conversion")
	}
}

func TestVolume_FromLiter(t *testing.T) {
	volume := Volume{}.FromLiter(1.0)
	stringValue := fmt.Sprint(volume)

	if !strings.EqualFold(stringValue, expectedVolumeResponse) {
		t.Error("Error on conversion")
	}
}

func TestVolume_FromGallon(t *testing.T) {
	volume := Volume{}.FromGallon(0.26417205235815)
	stringValue := fmt.Sprint(volume)

	if !strings.EqualFold(stringValue, expectedVolumeResponse) {
		t.Error("Error on conversion")
	}
}

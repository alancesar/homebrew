package abv

import (
	"github.com/alancesar/homebrew/density"
	"testing"
)

func TestCalculate(t *testing.T) {
	result := Calculate(density.Sg(1.042), density.Sg(1.008))

	if result.Abv != 0.04480076975680501 {
		t.Error("unexpected abv value")
	}

	if result.Abw != 0.03584061580544401 {
		t.Error("unexpected aws value")
	}

	if result.Attenuation != 0.8095238095238095 {
		t.Error("unexpected attenuation value")
	}
}

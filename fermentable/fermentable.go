package fermentable

import (
	"github.com/alancesar/homebrew/color"
	"github.com/alancesar/homebrew/mass"
)

type Fermentable struct {
	Quantity mass.Mass
	Color    color.Color
	PPG      float64
	Mashing  bool
}

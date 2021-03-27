package recipe

import (
	"github.com/alancesar/homebrew/color"
	"github.com/alancesar/homebrew/mass"
)

type Grain struct {
	Quantity mass.Mass
	Color    color.Color
	PPG      float64
	Mashing  bool
}

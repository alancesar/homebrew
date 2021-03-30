package hop

import "github.com/alancesar/homebrew/mass"

type Hop struct {
	Quantity   mass.Mass
	BoilTime   int
	AlphaAcids float64
	Pellet     bool
	DryHopping bool
}

package recipe

import "github.com/alancesar/homebrew/mass"

type Hop struct {
	Quantity   mass.Mass
	BoilTime   int
	AlphaAcids float64
}

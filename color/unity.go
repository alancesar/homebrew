package color

import "github.com/alancesar/homebrew/units"

const (
	srmSymbol      = "srm"
	ebcSymbol      = "ebc"
	lovibondSymbol = "ºl"
)

type Color struct {
	Srm      units.Unity
	Ebc      units.Unity
	Lovibond units.Unity
}

func create(from *from) Color {
	return Color{
		Srm:      units.New(from.ToSrm(), srmSymbol),
		Ebc:      units.New(from.ToEbc(), ebcSymbol),
		Lovibond: units.New(from.ToLovibond(), lovibondSymbol),
	}
}

func Srm(value float64) Color {
	return create(FromSrm(value))
}

func Ebc(value float64) Color {
	return create(FromEbc(value))
}

func (c *Color) FromLovibond(value float64) Color {
	return create(FromLovibond(value))
}

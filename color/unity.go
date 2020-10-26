package color

type Color struct {
	Srm      float64
	Ebc      float64
	Lovibond float64
}

func create(from *from) Color {
	return Color{
		Srm:      from.ToSrm(),
		Ebc:      from.ToEbc(),
		Lovibond: from.ToLovibond(),
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

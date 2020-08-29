package color

type from struct {
	l   float64
	srm float64
	ebc float64
}

func (f *from) ToLovibond() float64 {
	return f.l
}

func (f *from) ToSrm() float64 {
	return f.srm
}

func (f *from) ToEbc() float64 {
	return f.ebc
}

func FromLovibond(value float64) *from {
	return &from{
		l:   value,
		srm: (1.3546 * value) - 0.76,
		ebc: ((1.3546 * value) - 0.76) * 1.97,
	}
}

func FromEbc(value float64) *from {
	return &from{
		l:   ((value * 0.508) + 0.76) / 1.3546,
		srm: value * 0.508,
		ebc: value,
	}
}

func FromSrm(value float64) *from {
	return &from{
		l:   (value + 0.76) / 1.3546,
		srm: value,
		ebc: value * 1.97,
	}
}

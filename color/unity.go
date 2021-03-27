package color

const (
	ebcInSrm = 0.508
	srmInEbc = 1.97

	srmToLovibondFactor     = 1.3546
	srmToLovibondCorrection = 0.76
)

func srmInLovibond(lovibond float64) float64 {
	return (srmToLovibondFactor * lovibond) - srmToLovibondCorrection
}

func lovibondInSrm(srm float64) float64 {
	return (srm + srmToLovibondCorrection) / srmToLovibondFactor
}

type Color struct {
	Srm      float64
	Ebc      float64
	Lovibond float64
}

func NewFromSrm(value float64) Color {
	return Color{
		Srm:      value,
		Ebc:      value * srmInEbc,
		Lovibond: lovibondInSrm(value),
	}
}

func NewFromEbc(value float64) Color {
	srm := value * ebcInSrm
	return Color{
		Srm:      srm,
		Ebc:      value,
		Lovibond: lovibondInSrm(srm),
	}
}

func FromLovibond(value float64) Color {
	srm := srmInLovibond(value)
	return Color{
		Srm:      srm,
		Ebc:      srm * srmInEbc,
		Lovibond: value,
	}
}

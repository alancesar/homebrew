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
	SRM      float64
	EBC      float64
	Lovibond float64
}

func NewFromSRM(value float64) Color {
	return Color{
		SRM:      value,
		EBC:      value * srmInEbc,
		Lovibond: lovibondInSrm(value),
	}
}

func NewFromEBC(value float64) Color {
	srm := value * ebcInSrm
	return Color{
		SRM:      srm,
		EBC:      value,
		Lovibond: lovibondInSrm(srm),
	}
}

func FromLovibond(value float64) Color {
	srm := srmInLovibond(value)
	return Color{
		SRM:      srm,
		EBC:      srm * srmInEbc,
		Lovibond: value,
	}
}

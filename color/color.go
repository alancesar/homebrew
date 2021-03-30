package color

import "math"

const (
	ebcInSrm = 0.508
	srmInEbc = 1.97

	srmToLovibondFactor     = 1.3546
	srmToLovibondCorrection = 0.76

	maxSRMValue = 40
	minSRMValue = 1
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

func (c Color) IsZero() bool {
	return c.SRM == 0 && c.EBC == 0 && c.Lovibond == 0
}

func (c *Color) RGB() string {
	srm := int(math.Round(c.SRM))

	if srm > maxSRMValue {
		return srmToRGB[maxSRMValue]
	} else if srm < minSRMValue {
		return srmToRGB[minSRMValue]
	}

	return srmToRGB[srm]
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

func NewFromLovibond(value float64) Color {
	srm := srmInLovibond(value)
	return Color{
		SRM:      srm,
		EBC:      srm * srmInEbc,
		Lovibond: value,
	}
}

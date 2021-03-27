package density

import "math"

func RefractometerCorrection(og, fg Density, wortCorrection float64) Density {
	correctedOgBrix := og.Brix * wortCorrection
	correctedFgBrix := fg.Brix * wortCorrection

	correction := 1 - 0.0044993*correctedOgBrix + 0.0117741*correctedFgBrix +
		0.000275806*math.Pow(correctedOgBrix, 2) - 0.00127169*math.Pow(correctedFgBrix, 2) -
		0.00000727999*math.Pow(correctedOgBrix, 3) + 0.0000632929*math.Pow(correctedFgBrix, 3)

	return NewFromSG(correction)
}

package density

import "math"

func RefractometerCorrection(og, fg Density) Density {
	correction := (1 - (0.0044993 * og.Brix)) +
		(0.011774 * fg.Brix) + ((0.00027581 * math.Pow(og.Brix, 2)) -
		(0.0012717 * math.Pow(fg.Brix, 2)) - (0.0000072800 * math.Pow(og.Brix, 3))) +
		(0.000063293 * math.Pow(fg.Brix, 3))

	return Sg(correction)
}

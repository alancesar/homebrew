package pressure

const (
	psiInKgfcm2 = 14.223
	barInKgfcm2 = 1.01972
)

type Pressure struct {
	PSI    float64
	Kgfcm2 float64
	Bar    float64
}

func NewFromPSI(value float64) Pressure {
	return Pressure{
		PSI:    value,
		Kgfcm2: value / psiInKgfcm2,
		Bar:    value / barInKgfcm2 / psiInKgfcm2,
	}
}

func NewFromKgfcm2(value float64) Pressure {
	return Pressure{
		PSI:    value * psiInKgfcm2,
		Kgfcm2: value,
		Bar:    value / barInKgfcm2,
	}
}

func NewFromBar(value float64) Pressure {
	return Pressure{
		PSI:    value * barInKgfcm2 * psiInKgfcm2,
		Kgfcm2: value * barInKgfcm2,
		Bar:    value,
	}
}

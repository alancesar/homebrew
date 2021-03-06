package pressure

import "github.com/alancesar/homebrew/measure"

const (
	psiInKgfcm2 = 14.223
	barInKgfcm2 = 1.01972
)

type (
	Pressure struct {
		PSI    float64
		Kgfcm2 float64
		Bar    float64
	}

	pressureConstructor func(value float64) Pressure
)

func (p Pressure) IsZero() bool {
	return p.PSI == 0 && p.Kgfcm2 == 0 && p.Bar == 0
}

var constructorsMap = map[string]pressureConstructor{
	"psi":     NewFromPSI,
	"bar":     NewFromBar,
	"kgfcm":   NewFromKgfcm2,
	"kgfcm²":  NewFromKgfcm2,
	"kgf/cm²": NewFromKgfcm2,
	"at":      NewFromKgfcm2,
}

func NewFrom(input string) Pressure {
	pressure := Pressure{}
	measure.NewFrom(input, func(symbol string, value float64) {
		if constructor, exists := constructorsMap[symbol]; exists {
			pressure = constructor(value)
		}
	})

	return pressure
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

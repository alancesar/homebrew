package volume

const (
	millilitersInLiters = 1000
	gallonsInLiters     = 0.26417205235815
)

type Volume struct {
	Milliliters float64
	Liters      float64
	Gallons     float64
}

func NewFromMilliliter(value float64) Volume {
	return createFromLiters(value / millilitersInLiters)
}

func NewFromLiter(value float64) Volume {
	return createFromLiters(value)
}

func NewFromGallon(value float64) Volume {
	return createFromGallons(value)
}

func createFromLiters(liters float64) Volume {
	return Volume{
		Milliliters: liters * millilitersInLiters,
		Liters:      liters,
		Gallons:     liters * gallonsInLiters,
	}
}

func createFromGallons(gallons float64) Volume {
	return Volume{
		Milliliters: gallons / gallonsInLiters * millilitersInLiters,
		Liters:      gallons / gallonsInLiters,
		Gallons:     gallons,
	}
}

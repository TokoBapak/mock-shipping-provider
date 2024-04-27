package provider

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository"
)

type Sicepat struct {
	Rate primitive.Rate
}

func NewSicepatCalculation(sicepatRate *primitive.Rate) repository.ProviderCalculation {
	return &Sicepat{
		Rate: primitive.Rate{
			PerKilogram:      sicepatRate.PerKilogram,
			PerKilometer:     sicepatRate.PerKilometer,
			PerCmCubic:       sicepatRate.PerCmCubic,
			KilometerPerHour: sicepatRate.KilometerPerHour,
		},
	}
}

func (sicepat *Sicepat) CalculatePrice(distance float64, dimension primitive.Dimension, weight float64) int64 {
	volume := dimension.Width * dimension.Height * dimension.Depth

	distanceCost := distance * float64(sicepat.Rate.PerKilometer)

	weightCost := weight * float64(sicepat.Rate.PerKilogram)

	volumeCost := volume * float64(sicepat.Rate.PerCmCubic)

	return int64(distanceCost + weightCost + volumeCost)
}

func (sicepat *Sicepat) CalculateTimeOfArrival(distance float64) int64 {
	if distance < float64(sicepat.Rate.KilometerPerHour) {
		return 1
	}
	return int64(distance) / sicepat.Rate.KilometerPerHour
}

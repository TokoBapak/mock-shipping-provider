package provider

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository"
)

type Anteraja struct {
	Rate primitive.Rate
}

func NewAnterajaCalculation(anterAjaRate *primitive.Rate) repository.ProviderCalculation {
	return &Anteraja{
		Rate: primitive.Rate{
			PerKilometer:     anterAjaRate.PerKilometer,
			PerKilogram:      anterAjaRate.PerKilogram,
			PerCmCubic:       anterAjaRate.PerCmCubic,
			KilometerPerHour: anterAjaRate.KilometerPerHour,
		},
	}
}

func (anterAja *Anteraja) CalculatePrice(distance float64, dimension primitive.Dimension, weight float64) int64 {
	volume := dimension.Width * dimension.Height * dimension.Depth

	distanceCost := distance * float64(anterAja.Rate.PerKilometer)

	weightCost := weight * float64(anterAja.Rate.PerKilogram)

	volumeCost := volume * float64(anterAja.Rate.PerCmCubic)

	return int64(distanceCost + weightCost + volumeCost)
}

func (anterAja *Anteraja) CalculateTimeOfArrival(distance float64) int64 {
	if distance < float64(anterAja.Rate.KilometerPerHour) {
		return 1
	}

	return int64(distance / float64(anterAja.Rate.KilometerPerHour))
}

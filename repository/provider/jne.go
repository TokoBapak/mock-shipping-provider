package provider

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository"
)

type JNE struct {
	Rate primitive.Rate
}

func NewJneCalculation(jneRate *primitive.Rate) repository.ProviderCalculation {
	return &JNE{
		Rate: primitive.Rate{
			PerKilogram:      jneRate.PerKilogram,
			PerKilometer:     jneRate.PerKilometer,
			PerCmCubic:       jneRate.PerCmCubic,
			KilometerPerHour: jneRate.KilometerPerHour,
		},
	}
}

func (jne *JNE) CalculatePrice(distance float64, dimension primitive.Dimension, weight float64) int64 {

	volume := dimension.Width * dimension.Height * dimension.Depth

	distanceCost := distance * float64(jne.Rate.PerKilometer)

	weightCost := weight * float64(jne.Rate.PerKilogram)

	volumeCost := volume * float64(jne.Rate.PerCmCubic)

	return int64(distanceCost + weightCost + volumeCost)
}

func (jne *JNE) CalculateTimeOfArrival(distance float64) int64 {
	if int64(distance) < jne.Rate.KilometerPerHour {
		return 1
	}

	return int64(distance) / jne.Rate.KilometerPerHour
}

package provider

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository"
)

type JNT struct {
	Rate primitive.Rate
}

func NewJntCalculation(jntRate *primitive.Rate) repository.ProviderCalculation {
	return &JNT{
		Rate: primitive.Rate{
			PerKilogram:      jntRate.PerKilogram,
			PerKilometer:     jntRate.PerKilometer,
			PerCmCubic:       jntRate.PerCmCubic,
			KilometerPerHour: jntRate.KilometerPerHour,
		},
	}
}

func (jnt *JNT) CalculatePrice(distance float64, dimension primitive.Dimension, weight float64) int64 {
	volume := dimension.Width * dimension.Height * dimension.Depth

	distanceCost := distance * float64(jnt.Rate.PerKilometer)

	weightCost := weight * float64(jnt.Rate.PerKilogram)

	volumeCost := volume * float64(jnt.Rate.PerCmCubic)

	return int64(distanceCost + weightCost + volumeCost)
}

func (jnt *JNT) CalculateTimeOfArrival(distance float64) int64 {
	if distance < float64(jnt.Rate.KilometerPerHour) {
		return 1
	}
	return int64(distance) / jnt.Rate.KilometerPerHour
}

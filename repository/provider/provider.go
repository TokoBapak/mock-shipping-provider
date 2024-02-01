package provider

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository"
)

type Calculation struct {
	JNE      repository.ProviderCalculation
	JNT      repository.ProviderCalculation
	SiCepat  repository.ProviderCalculation
	AnterAja repository.ProviderCalculation
}

func GetProviderCalculation() (*Calculation, error) {
	jneRate := primitive.Rate{
		PerKilogram:      700,
		PerKilometer:     600,
		PerCmCubic:       500,
		KilometerPerHour: 60,
	}
	if err := jneRate.Validate(); err != nil {
		return &Calculation{}, err
	}
	jneCalculation := NewJneCalculation(&jneRate)

	jntRate := primitive.Rate{
		PerKilogram:      700,
		PerKilometer:     600,
		PerCmCubic:       500,
		KilometerPerHour: 60,
	}
	if err := jntRate.Validate(); err != nil {
		return &Calculation{}, err
	}
	jntCalculation := NewJneCalculation(&jntRate)

	anterajaRate := primitive.Rate{
		PerKilogram:      700,
		PerKilometer:     600,
		PerCmCubic:       500,
		KilometerPerHour: 60,
	}
	if err := anterajaRate.Validate(); err != nil {
		return &Calculation{}, err
	}
	anterajaCalculation := NewJneCalculation(&anterajaRate)

	sicepatRate := primitive.Rate{
		PerKilogram:      700,
		PerKilometer:     600,
		PerCmCubic:       500,
		KilometerPerHour: 60,
	}
	if err := sicepatRate.Validate(); err != nil {
		return &Calculation{}, err
	}
	sicepatCalculation := NewJneCalculation(&sicepatRate)

	return &Calculation{
		JNE:      jneCalculation,
		JNT:      jntCalculation,
		SiCepat:  sicepatCalculation,
		AnterAja: anterajaCalculation,
	}, nil
}

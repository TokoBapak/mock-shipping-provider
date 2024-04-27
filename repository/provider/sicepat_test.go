package provider_test

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository/provider"
	"reflect"
	"testing"
)

func newSicepatRate() primitive.Rate {
	return primitive.Rate{
		PerKilogram:      1000,
		PerKilometer:     1200,
		PerCmCubic:       900,
		KilometerPerHour: 60,
	}
}

func createSicepatPriceCalculation(t *testing.T, distance float64, dimension primitive.Dimension, weight float64) int64 {

	sicepatRate := newSicepatRate()

	Sicepat := provider.NewSicepatCalculation(&sicepatRate)

	result := Sicepat.CalculatePrice(distance, dimension, weight)

	return result
}

func createSicepatTimeArrival(t *testing.T, distance float64) int64 {

	sicepatRate := newSicepatRate()

	Sicepat := provider.NewSicepatCalculation(&sicepatRate)

	result := Sicepat.CalculateTimeOfArrival(distance)

	return result
}

func TestProviderSicepat(t *testing.T) {
	sicepatRate := newSicepatRate()

	Sicepat := provider.NewSicepatCalculation(&sicepatRate)

	testCases := []struct {
		distance  float64
		dimension primitive.Dimension
		weight    float64
	}{
		{distance: 300},
		{dimension: primitive.Dimension{
			Width:  30,
			Height: 10,
			Depth:  11,
		}},
		{weight: 40},
	}

	for _, testCase := range testCases {
		t.Run("test calculate price", func(t *testing.T) {
			result := Sicepat.CalculatePrice(testCase.distance, testCase.dimension, testCase.weight)

			expectedResult := createSicepatPriceCalculation(t, testCase.distance, testCase.dimension, testCase.weight)
			if reflect.DeepEqual(result, expectedResult) == false {
				t.Errorf("must not error but get %v and %v different price meaning code error", result, expectedResult)
			}
		})
	}

	for _, testCase := range testCases {
		t.Run("test time of arrival", func(t *testing.T) {

			result := Sicepat.CalculateTimeOfArrival(testCase.distance)

			expectedResultTime := createSicepatTimeArrival(t, testCase.distance)

			if result != expectedResultTime {
				t.Errorf("must be not error but get result time: %v, end expected: %v", result, expectedResultTime)
			}
		})
	}

	t.Run("when distance return an hour", func(t *testing.T) {
		result := Sicepat.CalculateTimeOfArrival(90)

		if result != (1) {
			t.Errorf("error time is must one hour get %v", result)
		}
	})
}

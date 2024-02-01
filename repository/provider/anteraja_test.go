package provider_test

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository/provider"
	"reflect"
	"testing"
)

func newRateAnterAja() primitive.Rate {
	return primitive.Rate{
		PerKilogram:      700,
		PerKilometer:     500,
		PerCmCubic:       400,
		KilometerPerHour: 60,
	}
}

func createAnterAjaPrice(t *testing.T, distance float64, dimension primitive.Dimension, weight float64) int64 {
	anterAjaRate := newRateAnterAja()

	AnterAja := provider.NewAnterajaCalculation(&anterAjaRate)

	result := AnterAja.CalculatePrice(distance, dimension, weight)

	return result
}

func createAnterAjaTimeArrival(t *testing.T, distance float64) int64 {
	anterAjaRate := newRateAnterAja()

	AnterAja := provider.NewAnterajaCalculation(&anterAjaRate)

	result := AnterAja.CalculateTimeOfArrival(distance)

	return result
}

func TestAnterAjaCalculation(t *testing.T) {
	anterAjaRate := newRateAnterAja()

	AnterAja := provider.NewAnterajaCalculation(&anterAjaRate)

	testCases := []struct {
		distance  float64
		dimension primitive.Dimension
		weight    float64
	}{
		{distance: 3400},
		{dimension: primitive.Dimension{
			Height: 40,
			Depth:  50,
			Width:  9,
		}},
		{weight: 2},
	}

	for _, testCase := range testCases {
		t.Run("calculate the price", func(t *testing.T) {
			result := AnterAja.CalculatePrice(testCase.distance, testCase.dimension, testCase.weight)

			expectedPrice := createAnterAjaPrice(t, testCase.distance, testCase.dimension, testCase.weight)

			if reflect.DeepEqual(result, expectedPrice) == false {
				t.Errorf("get %v and expected %v", result, expectedPrice)
			}
		})
	}

	for _, testCase := range testCases {
		t.Run("calculate time an arrival", func(t *testing.T) {
			result := AnterAja.CalculateTimeOfArrival(testCase.distance)

			expectedTime := createAnterAjaTimeArrival(t, testCase.distance)

			if result != expectedTime {
				t.Errorf("error mean code err or paramter msiign argument get %v and expected %v", result, expectedTime)
			}

		})
	}

	t.Run("when distance return an hour", func(t *testing.T) {
		result := AnterAja.CalculateTimeOfArrival(30)

		if result != 1 {
			t.Errorf("error time is must one hour get %v", result)
		}
	})
}

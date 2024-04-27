package provider_test

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository/provider"
	"reflect"
	"testing"
)

func newRateForJnt() primitive.Rate {
	return primitive.Rate{
		PerKilogram:      800,
		PerKilometer:     600,
		PerCmCubic:       890,
		KilometerPerHour: 60,
	}
}

func createJntPriceCalculation(t *testing.T, distance float64, dimension primitive.Dimension, weight float64) int64 {
	rateJnt := newRateForJnt()

	JNT := provider.NewJntCalculation(&rateJnt)

	result := JNT.CalculatePrice(distance, dimension, weight)

	return result
}

func createJntTimeArrival(t *testing.T, distance float64) int64 {
	rateJnt := newRateForJnt()

	JNT := provider.NewJntCalculation(&rateJnt)

	result := JNT.CalculateTimeOfArrival(distance)

	return result

}

func TestJntCalculation(t *testing.T) {
	rateJnt := newRateForJnt()

	JNT := provider.NewJntCalculation(&rateJnt)

	testCases := []struct {
		distance  float64
		dimension primitive.Dimension
		weight    float64
	}{
		{distance: 2350},
		{dimension: primitive.Dimension{
			Height: 20,
			Width:  10,
			Depth:  11,
		}},
		{weight: 300},
	}

	for _, testCase := range testCases {
		t.Run("calculate price", func(t *testing.T) {
			result := JNT.CalculatePrice(testCase.distance, testCase.dimension, testCase.weight)

			expectedPrice := createJntPriceCalculation(t, testCase.distance, testCase.dimension, testCase.weight)

			if reflect.DeepEqual(result, expectedPrice) == false {
				t.Errorf("get differnce price %v expected %v", result, expectedPrice)
			}
		})
	}

	for _, testCase := range testCases {
		t.Run("time of arrival", func(t *testing.T) {
			result := JNT.CalculateTimeOfArrival(testCase.distance)

			expectedTime := createJntTimeArrival(t, testCase.distance)

			if reflect.DeepEqual(result, expectedTime) == false {
				t.Errorf("get time different %v expected %v", result, expectedTime)
			}
		})
	}

	t.Run("calculate time arrival an hour", func(t *testing.T) {
		result := JNT.CalculateTimeOfArrival(30)
		if result != (1) {
			t.Errorf("error estimate an arrival %v", result)
		}
		result = JNT.CalculateTimeOfArrival(100)
		if result != (1) {
			t.Errorf("error estimate an arrival %v", result)
		}
	})
}

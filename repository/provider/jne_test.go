package provider_test

import (
	"mock-shipping-provider/primitive"
	"mock-shipping-provider/repository/provider"
	"reflect"
	"testing"
)

func newJneRate() primitive.Rate {
	return primitive.Rate{
		PerKilogram:      900,
		PerKilometer:     800,
		PerCmCubic:       600,
		KilometerPerHour: 60,
	}
}

func createJneCalculatePrice(t *testing.T, distance float64, dimension primitive.Dimension, weight float64) int64 {
	jntRate := newJneRate()

	JNE := provider.NewJneCalculation(&jntRate)

	result := JNE.CalculatePrice(distance, dimension, weight)

	return result
}

func createJneTimeArrival(t *testing.T, distance float64) int64 {
	jntRate := newJneRate()

	JNE := provider.NewJneCalculation(&jntRate)

	result := JNE.CalculateTimeOfArrival(distance)

	return result
}

func TestProviderJne(t *testing.T) {
	testCases := []struct {
		distance  float64
		dimension primitive.Dimension
		weight    float64
	}{
		{distance: 4500.0},
		{dimension: primitive.Dimension{
			Width:  30,
			Height: 20,
			Depth:  30,
		}},
		{weight: 45},
	}

	jneRate := newJneRate()

	JNE := provider.NewJneCalculation(&jneRate)

	for _, testcase := range testCases {
		t.Run("test calculate price", func(t *testing.T) {
			result := JNE.CalculatePrice(float64(testcase.distance), testcase.dimension, testcase.weight)

			expectedPrice := createJneCalculatePrice(t, testcase.distance, testcase.dimension, testcase.weight)

			if reflect.DeepEqual(result, expectedPrice) == false {
				t.Errorf("must not error but get %v and %v different price meaning code error", result, expectedPrice)
			}
		})
	}
	for _, testCase := range testCases {
		t.Run("test time of arrival", func(t *testing.T) {
			time := JNE.CalculateTimeOfArrival(testCase.distance)

			expectedTime := createJneTimeArrival(t, testCase.distance)

			if reflect.DeepEqual(time, expectedTime) == false {
				t.Errorf("error get time arrival is : %v hours ", time)
			}
		})
	}

	t.Run("when distance return an hour", func(t *testing.T) {
		result := JNE.CalculateTimeOfArrival(30)

		if result != (1) {
			t.Errorf("error time is must one hour get %v", result)
		}
	})
}

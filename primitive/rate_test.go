package primitive_test

import (
	"mock-shipping-provider/primitive"
	"testing"
)

func TestRateInsuficient(t *testing.T) {
	var rate primitive.Rate

	err := rate.Validate()

	if err != nil {
		t.Error("get errorrrr")
	}
}

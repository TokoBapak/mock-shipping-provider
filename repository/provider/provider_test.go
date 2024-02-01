package provider_test

import (
	"mock-shipping-provider/repository/provider"
	"testing"
)

func TestProvider(t *testing.T) {
	ProviderCalculation, err := provider.GetProviderCalculation()
	if err != nil {
		t.Errorf("must not errorr but get: %v", err)
	}

	if ProviderCalculation == nil {
		t.Error("chech code in provider calculation")
	}
}

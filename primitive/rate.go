package primitive

import "errors"

type Rate struct {
	PerKilogram      int64
	PerKilometer     int64
	PerCmCubic       int64
	KilometerPerHour int64
}

func (rate Rate) Validate() error {
	if rate.PerKilogram < 0 {
		return errors.New("rate in kilogram lower than 0")
	}

	if rate.PerKilometer < 0 {
		return errors.New("rate in kilometer lower than 0")
	}

	if rate.PerCmCubic < 0 {
		return errors.New("rate cm cubic is lower than 0")
	}

	if rate.KilometerPerHour < 0 {
		return errors.New("kilometer per hour lower than 0")
	}

	return nil
}

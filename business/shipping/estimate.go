package shipping

import (
	"context"
	"mock-shipping-provider/business"
	"mock-shipping-provider/primitive"
)

func (d *Dependency) Estimate(ctx context.Context, request business.EstimateRequest) ([]business.EstimateResult, error) {
	// TODO implement me
	panic("implement me")
}

// ValidateEstimateRequest handle an action to
// validate the estimate request body
func ValidateEstimateRequest(request business.EstimateRequest) *business.RequestValidationError {
	var issues []business.RequestValidationIssue

	// Sender
	if request.Sender == (primitive.Coordinate{}) {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "sender",
			Message: "can not be empty",
		})
	} else {
		if request.Sender.Latitude < -90 || request.Sender.Latitude > 90 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeInvalidValue,
				Field:   "sender.latitude",
				Message: "must be between -90 and 90",
			})
		}

		if request.Sender.Longitude < -180 || request.Sender.Longitude > 180 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeInvalidValue,
				Field:   "sender.longitude",
				Message: "must be between -180 and 180",
			})
		}
	}

	// Recipient
	if request.Recipient == (primitive.Coordinate{}) {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "recipient",
			Message: "can not be empty",
		})
	} else {
		if request.Recipient.Latitude < -90 || request.Recipient.Latitude > 90 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeInvalidValue,
				Field:   "recipient.latitude",
				Message: "must be between -90 and 90",
			})
		}

		if request.Recipient.Longitude < -180 || request.Recipient.Longitude > 180 {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeInvalidValue,
				Field:   "recipient.longitude",
				Message: "must be between -180 and 180",
			})
		}
	}

	// Dimension
	if request.Dimension == (primitive.Dimension{}) {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeRequired,
			Field:   "dimension",
			Message: "can not be empty",
		})
	} else {
		if err := request.Dimension.Validate(); err != nil {
			issues = append(issues, business.RequestValidationIssue{
				Code:    business.RequestValidationCodeInvalidValue,
				Field:   "dimension",
				Message: err.Error(),
			})
		} else {
			// Dimension.Depth
			if request.Dimension.Depth <= 0 {
				issues = append(issues, business.RequestValidationIssue{
					Code:    business.RequestValidationCodeInvalidValue,
					Field:   "dimension.depth",
					Message: "must be greater than 0",
				})
			}

			// Dimension.Height
			if request.Dimension.Height <= 0 {
				issues = append(issues, business.RequestValidationIssue{
					Code:    business.RequestValidationCodeInvalidValue,
					Field:   "dimension.height",
					Message: "must be greater than 0",
				})
			}

			// Dimension.Width
			if request.Dimension.Width <= 0 {
				issues = append(issues, business.RequestValidationIssue{
					Code:    business.RequestValidationCodeInvalidValue,
					Field:   "width",
					Message: "must be greater than 0",
				})
			}
		}
	}

	// Weight
	if request.Weight <= 0 {
		issues = append(issues, business.RequestValidationIssue{
			Code:    business.RequestValidationCodeInvalidValue,
			Field:   "weight",
			Message: "must be greater than 0",
		})
	}

	if len(issues) > 0 {
		return &business.RequestValidationError{Issues: issues}
	}

	return nil
}

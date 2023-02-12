package business

import (
	"fmt"
	"strings"
)

// RequestValidationCode provides a typed string for validation error codes.
type RequestValidationCode string

const (
	RequestValidationCodeRequired        RequestValidationCode = "field_required"
	RequestValidationCodeTooShort        RequestValidationCode = "too_short"
	RequestValidationCodeTooLong         RequestValidationCode = "too_long"
	RequestValidationCodeProhibitedValue RequestValidationCode = "prohibited_value"
	RequestValidationCodeInvalidValue    RequestValidationCode = "invalid_value"
)

// RequestValidationIssue contains a specific validation issue for each field and rules.
// It should be embedded as array inside the RequestValidationError struct.
type RequestValidationIssue struct {
	// Code specifies the error code. You must not enter a custom code, instead
	// add another entry for the RequestValidationCode type.
	//
	// This should be aligned with the documentation on how the consumers (or users)
	// handle their validation errors from us.
	Code RequestValidationCode
	// Field specifies the field that the error happened. If the field is on a nested object,
	// you can separate it using a dot. For example:
	//
	// 		{ "customer": { "name": "string" } }
	// becomes
	// 		customer.name
	Field string
	// Message must contain helpful message that helps the user create proper request.
	// This should also be simple, and should not repeat what's on the code and field value.
	// For example: "maximum of 50 characters", "must be numeric", "must not empty".
	Message string
}

// RequestValidationError should be returned when a request validation error occurred.
// It contains array of issues that the validation encountered.
type RequestValidationError struct {
	Issues []RequestValidationIssue
}

func (r RequestValidationError) Error() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("%d issues found during validation: ", len(r.Issues)))

	for i, issue := range r.Issues {
		s.WriteString(fmt.Sprintf("%s for %s: %s", issue.Code, issue.Field, issue.Message))

		if i != len(r.Issues)-1 {
			s.WriteString(", ")
		}
	}

	return s.String()
}

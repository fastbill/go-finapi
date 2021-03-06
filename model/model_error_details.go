/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: 1.106.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Error details
type ErrorDetails struct {
	// Error message
	Message string `json:"message,omitempty"`
	// Error code. See the documentation of the individual services for details about what values may be returned.
	Code string `json:"code"`
	// Error type. BUSINESS errors depict error messages in the language of the bank (or the preferred language) for the user, e.g. from a bank server. TECHNICAL errors are meant to be read by developers and depict internal errors.
	Type_ string `json:"type"`
	// This field is set when a multi-step authentication is required, i.e. when you need to repeat the original service call and provide additional data. The field contains information about what additional data is required.
	MultiStepAuthentication *MultiStepAuthenticationChallenge `json:"multiStepAuthentication,omitempty"`
}

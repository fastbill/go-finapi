/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.79.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Error details
type ErrorDetails struct {
	// Error message
	Message string `json:"message,omitempty"`
	// Error code. See the documentation of the individual services for details about what values may be returned.
	Code string `json:"code"`
	// Error type. BUSINESS errors depict German error messages for the user, e.g. from a bank server. TECHNICAL errors depict internal errors.
	Type_ string `json:"type"`
	// This field is set when a multi-step authentication is required, i.e. when you need to repeat the original service call and provide additional data. The field contains information about what additional data is required.
	MultiStepAuthentication *MultiStepAuthenticationChallenge `json:"multiStepAuthentication,omitempty"`
}

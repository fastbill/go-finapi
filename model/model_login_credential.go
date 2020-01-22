/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.93.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Login credential
type LoginCredential struct {
	// The login field label, as defined by the respective bank interface.
	Label string `json:"label"`
	// The value for the login field
	Value string `json:"value"`
}

/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.93.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Container for securities resources
type SecurityList struct {
	// List of securities
	Securities []Security `json:"securities"`
}

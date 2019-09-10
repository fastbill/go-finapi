/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.79.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Container for IBAN rules
type IbanRuleList struct {
	// List of IBAN rules
	IbanRules []IbanRule `json:"ibanRules"`
}

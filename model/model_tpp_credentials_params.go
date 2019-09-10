/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.79.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// A container for new TPP client credentials data
type TppCredentialsParams struct {
	// The TPP Authentication Group Id for which the credentials can be used
	TppAuthenticationGroupId int64 `json:"tppAuthenticationGroupId"`
	// Optional label to credentials
	Label string `json:"label"`
	// ID of the TPP accessing the ASPSP API, as provided by the ASPSP as the result of registration
	TppClientId string `json:"tppClientId"`
	// Secret of the TPP accessing the ASPSP API, as provided by the ASPSP as the result of registration
	TppClientSecret string `json:"tppClientSecret"`
	// API Key provided by ASPSP  as the result of registration
	TppApiKey string `json:"tppApiKey,omitempty"`
	// Credentials \"valid from\" date in the format 'YYYY-MM-DD'. Default is today's date
	ValidFromDate string `json:"validFromDate,omitempty"`
	// Credentials \"valid until\" date in the format 'YYYY-MM-DD'. Default is null which means \"indefinite\" (no limit)
	ValidUntilDate string `json:"validUntilDate,omitempty"`
}

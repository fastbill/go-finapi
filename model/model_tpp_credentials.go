/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.93.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// A container for the TPP client credentials data
type TppCredentials struct {
	// The credential identifier.
	Id int64 `json:"id"`
	// Optional label of tpp client credentials set.
	Label string `json:"label,omitempty"`
	// TPP Authentication Group ID
	TppAuthenticationGroupId int64 `json:"tppAuthenticationGroupId,omitempty"`
	// Valid from date in the format 'YYYY-MM-DD'.
	ValidFrom string `json:"validFrom,omitempty"`
	// Valid until date in the format 'YYYY-MM-DD'.
	ValidUntil string `json:"validUntil,omitempty"`
}

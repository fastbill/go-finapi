/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.93.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// TPP Authentication groups with paging information
type PageableTppAuthenticationGroupResources struct {
	// List of received TPP authentication groups
	TppAuthenticationGroups []TppAuthenticationGroup `json:"tppAuthenticationGroups"`
	// Information for pagination
	Paging *Paging `json:"paging"`
}

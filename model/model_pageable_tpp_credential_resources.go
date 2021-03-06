/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: 1.106.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Container for TPP client credentials information with paging info
type PageableTppCredentialResources struct {
	// List of TPP client credentials
	TppCredentials []TppCredentials `json:"tppCredentials"`
	// Information for pagination
	Paging *Paging `json:"paging"`
}

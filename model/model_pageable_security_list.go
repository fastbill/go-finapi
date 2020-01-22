/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.93.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Container for page of securities
type PageableSecurityList struct {
	// List of securities
	Securities []Security `json:"securities"`
	// Information for pagination
	Paging *Paging `json:"paging"`
}

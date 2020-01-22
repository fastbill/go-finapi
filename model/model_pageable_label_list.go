/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.93.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Label resources with paging information
type PageableLabelList struct {
	// Labels
	Labels []Label `json:"labels"`
	// Information for pagination
	Paging *Paging `json:"paging"`
}

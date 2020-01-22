/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.93.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Container for keyword rule information with paging info
type PageableKeywordRuleList struct {
	// List of keyword rules
	KeywordRules []KeywordRule `json:"keywordRules"`
	// Information for pagination
	Paging *Paging `json:"paging"`
}

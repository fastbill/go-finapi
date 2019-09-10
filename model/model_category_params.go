/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.71.0.b1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Category parameters
type CategoryParams struct {
	// Name of the category. Maximum length is 128.
	Name string `json:"name"`
	// Identifier of the parent category, if the new category should be created as a sub-category. Must point to a main category in this case. If the new category should be a main category itself, this field must remain unset.
	ParentId int64 `json:"parentId,omitempty"`
}

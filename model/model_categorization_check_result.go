/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.71.0.b1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type CategorizationCheckResult struct {
	// Transaction identifier
	TransactionId string `json:"transactionId"`
	// Category
	Category *Category `json:"category,omitempty"`
}

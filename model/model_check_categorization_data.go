/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.93.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Transactions data for categorization check
type CheckCategorizationData struct {
	// Set of transaction data
	TransactionData []CheckCategorizationTransactionData `json:"transactionData"`
}

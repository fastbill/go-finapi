/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: 1.106.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Data for mock bank connection updates
type MockBatchUpdateParams struct {
	// List of mock bank connection updates
	MockBankConnectionUpdates []MockBankConnectionUpdate `json:"mockBankConnectionUpdates"`
	// Whether this call should trigger the dispatching of notifications. Default is 'false'.
	TriggerNotifications bool `json:"triggerNotifications,omitempty"`
}

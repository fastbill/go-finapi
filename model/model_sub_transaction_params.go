/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.79.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Data of a sub-transaction
type SubTransactionParams struct {
	// Amount
	Amount float32 `json:"amount"`
	// Category identifier. If not specified, the original transaction's category will be applied. If you explicitly want the sub-transaction to have no category, then pass this field with value '0' (zero).
	CategoryId int64 `json:"categoryId,omitempty"`
	// Purpose. Maximum length is 2000. If not specified, the original transaction's value will be applied.
	Purpose string `json:"purpose,omitempty"`
	// Counterpart. Maximum length is 80. If not specified, the original transaction's value will be applied.
	Counterpart string `json:"counterpart,omitempty"`
	// Counterpart account number. If not specified, the original transaction's value will be applied.
	CounterpartAccountNumber string `json:"counterpartAccountNumber,omitempty"`
	// Counterpart IBAN. If not specified, the original transaction's value will be applied.
	CounterpartIban string `json:"counterpartIban,omitempty"`
	// Counterpart BIC. If not specified, the original transaction's value will be applied.
	CounterpartBic string `json:"counterpartBic,omitempty"`
	// Counterpart BLZ. If not specified, the original transaction's value will be applied.
	CounterpartBlz string `json:"counterpartBlz,omitempty"`
	// List of connected labels. Note that when this field is not specified, then the labels of the original transaction will NOT get applied to the sub-transaction. Instead, the sub-transaction will have no labels assigned to it.
	LabelIds []int64 `json:"labelIds,omitempty"`
}

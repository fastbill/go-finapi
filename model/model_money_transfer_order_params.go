/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: 1.106.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Parameters for a money transfer order
type MoneyTransferOrderParams struct {
	// Name of the counterpart. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the counterpart name. Even if the name does not depict the actual registered account holder of the target account, the order might still be successful.
	CounterpartName string `json:"counterpartName"`
	// IBAN of the counterpart's account.
	CounterpartIban string `json:"counterpartIban"`
	// BIC of the counterpart's account. Only required when there is no 'IBAN_ONLY'-capability in the respective account/interface combination that is to be used when submitting the payment.
	CounterpartBic string `json:"counterpartBic,omitempty"`
	// The amount of the payment. Must be a positive decimal number with at most two decimal places.
	Amount float32 `json:"amount"`
	// The purpose of the transfer transaction
	Purpose string `json:"purpose,omitempty"`
	// SEPA purpose code, according to ISO 20022, external codes set.<br/>Please note that the SEPA purpose code may be ignored by some banks.
	SepaPurposeCode string `json:"sepaPurposeCode,omitempty"`
	// End-To-End ID for the transfer transaction
	EndToEndId string `json:"endToEndId,omitempty"`
}
/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.71.0.b1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Container for parameters for the execution of a submitted SEPA direct debit order
type ExecuteSepaDirectDebitParams struct {
	// Identifier of the bank account that you want to transfer money to
	AccountId int64 `json:"accountId"`
	// Banking TAN that the user received from the bank for executing the direct debit order. The field is required if you are licensed to perform SEPA direct debits yourself. Otherwise, i.e. when finAPI's web form flow is required, the web form will deal with executing the service itself.
	BankingTan string `json:"bankingTan,omitempty"`
}

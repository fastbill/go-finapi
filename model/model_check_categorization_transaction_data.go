/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.71.0.b1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Transaction data for categorization check
type CheckCategorizationTransactionData struct {
	// Identifier of transaction. This can be any arbitrary string that will be passed back in the response so that you can map the results to the given transactions. Note that the identifier must be unique within the given list of transactions.
	TransactionId string `json:"transactionId"`
	// Identifier of account type.<br/><br/>1 = Checking,<br/>2 = Savings,<br/>3 = CreditCard,<br/>4 = Security,<br/>5 = Loan,<br/>6 = Pocket (DEPRECATED; will not be returned for any account unless this type has explicitly been set via PATCH),<br/>7 = Membership,<br/>8 = Bausparen<br/><br/>
	AccountTypeId int64 `json:"accountTypeId"`
	// Amount
	Amount float32 `json:"amount"`
	// Purpose. Any symbols are allowed. Maximum length is 2000. Default value: null.
	Purpose string `json:"purpose,omitempty"`
	// Counterpart. Any symbols are allowed. Maximum length is 80. Default value: null.
	Counterpart string `json:"counterpart,omitempty"`
	// Counterpart IBAN. Default value: null.
	CounterpartIban string `json:"counterpartIban,omitempty"`
	// Counterpart account number. Default value: null.
	CounterpartAccountNumber string `json:"counterpartAccountNumber,omitempty"`
	// Counterpart BLZ. Default value: null.
	CounterpartBlz string `json:"counterpartBlz,omitempty"`
	// Counterpart BIC. Default value: null.
	CounterpartBic string `json:"counterpartBic,omitempty"`
	// Merchant category code (for credit card transactions only). May only contain up to 4 digits. Default value: null.
	McCode string `json:"mcCode,omitempty"`
	// ZKA business transaction code which relates to the transaction's type (Number from 0 through 999). Default value: null.
	TypeCodeZka string `json:"typeCodeZka,omitempty"`
}

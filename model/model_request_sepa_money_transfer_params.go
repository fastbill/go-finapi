/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.79.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Parameters for a single or collective SEPA money transfer order request
type RequestSepaMoneyTransferParams struct {
	// Name of the recipient. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the recipient name. Even if the recipient name does not depict the actual registered account holder of the specified recipient account, the money transfer request might still be successful. This field is optional only when you pass a clearing account as the recipient. Otherwise, this field is required.
	RecipientName string `json:"recipientName,omitempty"`
	// IBAN of the recipient's account. This field is optional only when you pass a clearing account as the recipient. Otherwise, this field is required.
	RecipientIban string `json:"recipientIban,omitempty"`
	// BIC of the recipient's account. Note: This field is optional when you pass a clearing account as the recipient or if the bank connection of the account that you want to transfer money from supports the IBAN-Only money transfer. You can find this out via GET /bankConnections/<id>. Also note that when a BIC is given, then this BIC will be used for the money transfer request independent of whether it is required or not (unless you pass a clearing account, in which case this field will always be ignored).
	RecipientBic string `json:"recipientBic,omitempty"`
	// Identifier of a clearing account. If this field is set, then the fields 'recipientName', 'recipientIban' and 'recipientBic' will be ignored and the recipient account will be the specified clearing account.
	ClearingAccountId string `json:"clearingAccountId,omitempty"`
	// The amount to transfer. Must be a positive decimal number with at most two decimal places (e.g. 99.99)
	Amount float32 `json:"amount"`
	// The purpose of the transfer transaction
	Purpose string `json:"purpose,omitempty"`
	// SEPA purpose code, according to ISO 20022, external codes set.
	SepaPurposeCode string `json:"sepaPurposeCode,omitempty"`
	// Identifier of the bank account that you want to transfer money from
	AccountId int64 `json:"accountId"`
	// Online banking PIN. Any symbols are allowed. Max length: 170. If a PIN is stored in the bank connection, then this field may remain unset. If finAPI's web form is not required and the field is set though then it will always be used (even if there is some other PIN stored in the bank connection). If you want the user to enter a PIN in finAPI's web form even when a PIN is stored, then just set the field to any value, so that the service recognizes that you wish to use the web form flow.
	BankingPin string `json:"bankingPin,omitempty"`
	// Whether to store the PIN. If the PIN is stored, it is not required to pass the PIN again when executing SEPA orders. Default value is 'false'. <br/><br/>NOTES:<br/> - before you set this field to true, please regard the 'pinsAreVolatile' flag of the bank connection that the account belongs to;<br/> - this field is ignored in case when the user will need to use finAPI's web form. The user will be able to decide whether to store the PIN or not in the web form, depending on the 'storeSecretsAvailableInWebForm' setting (see Client Configuration).
	StoreSecrets bool `json:"storeSecrets,omitempty"`
	// The bank-given ID of the two-step-procedure that should be used for the order. For a list of available two-step-procedures, see the corresponding bank connection (GET /bankConnections). If this field is not set, then the bank connection's default two-step-procedure will be used. Note that in this case, when the bank connection has no default two-step-procedure set, then the response of the service depends on whether you need to use finAPI's web form or not. If you need to use the web form, the user will be prompted to select the two-step-procedure within the web form. If you don't need to use the web form, then the service will return an error (passing a value for this field is required in this case).
	TwoStepProcedureId string `json:"twoStepProcedureId,omitempty"`
	// Execution date for the money transfer(s), in the format 'YYYY-MM-DD'. If not specified, then the current date will be used.
	ExecutionDate string `json:"executionDate,omitempty"`
	// This field is only regarded when you pass multiple orders. It determines whether the orders should be processed by the bank as one collective booking (in case of 'false'), or as single bookings (in case of 'true'). Default value is 'false'.
	SingleBooking bool `json:"singleBooking,omitempty"`
	// In case that you want to submit not just a single money transfer, but do a collective money transfer, use this field to pass a list of additional money transfer orders. The service will then pass a collective money transfer request to the bank, including both the money transfer specified on the top-level, as well as all money transfers specified in this list. The maximum count of money transfers that you can pass (in total) is 15000. Note that you should check the account's 'supportedOrders' field to find out whether or not it is supporting collective money transfers.
	AdditionalMoneyTransfers []SingleMoneyTransferRecipientData `json:"additionalMoneyTransfers,omitempty"`
	// NOTE: This field is DEPRECATED and will get removed at some point. Please refer to the 'multiStepAuthentication' field instead.<br/><br/>Challenge response. This field should be set only when the previous attempt to request a SEPA money transfer failed with HTTP code 510, i.e. the bank sent a challenge for the user for an additional authentication. In this case, this field must contain the response to the bank's challenge. Please note that in case of using finAPI's web form you have to leave this field unset and the application will automatically recognize that the user has to input challenge response and then a web form will be shown to the user.
	ChallengeResponse string `json:"challengeResponse,omitempty"`
	// Container for multi-step authentication data. Required when a previous service call initiated a multi-step authentication.
	MultiStepAuthentication *MultiStepAuthenticationCallback `json:"multiStepAuthentication,omitempty"`
	// Whether the finAPI web form should hide transaction details when prompting the caller for the second factor. Default value is false.
	HideTransactionDetailsInWebForm bool `json:"hideTransactionDetailsInWebForm,omitempty"`
	// Whether to store the PIN. If the PIN is stored, it is not required to pass the PIN again when executing SEPA orders. Default value is 'false'. <br/><br/>NOTES:<br/> - before you set this field to true, please regard the 'pinsAreVolatile' flag of the bank connection that the account belongs to;<br/> - this field is ignored in case when the user will need to use finAPI's web form. The user will be able to decide whether to store the PIN or not in the web form, depending on the 'storeSecretsAvailableInWebForm' setting (see Client Configuration).<br><br>NOTE: This field is deprecated and will be removed at some point. Use 'storeSecrets' instead.
	StorePin bool `json:"storePin,omitempty"`
}

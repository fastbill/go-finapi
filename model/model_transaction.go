/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: 1.106.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Container for a transaction's data
type Transaction struct {
	// Transaction identifier
	Id int64 `json:"id"`
	// Parent transaction identifier
	ParentId int64 `json:"parentId,omitempty"`
	// Account identifier
	AccountId int64 `json:"accountId"`
	// Value date in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	ValueDate string `json:"valueDate"`
	// Bank booking date in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	BankBookingDate string `json:"bankBookingDate"`
	// finAPI Booking date in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time). NOTE: In some cases, banks may deliver transactions that are booked in future, but already included in the current account balance. To keep the account balance consistent with the set of transactions, such \"future transactions\" will be imported with their finapiBookingDate set to the current date (i.e.: date of import). The finapiBookingDate will automatically get adjusted towards the bankBookingDate each time the associated bank account is updated. Example: A transaction is imported on July, 3rd, with a bank reported booking date of July, 6th. The transaction will be imported with its finapiBookingDate set to July, 3rd. Then, on July 4th, the associated account is updated. During this update, the transaction's finapiBookingDate will be automatically adjusted to July 4th. This adjustment of the finapiBookingDate takes place on each update until the bank account is updated on July 6th or later, in which case the transaction's finapiBookingDate will be adjusted to its final value, July 6th.<br/> The finapiBookingDate is the date that is used by the finAPI PFM services. E.g. when you calculate the spendings of an account for the current month, and have a transaction with finapiBookingDate in the current month but bankBookingDate at the beginning of the next month, then this transaction is included in the calculations (as the bank has this transaction's amount included in the current account balance as well).
	FinapiBookingDate string `json:"finapiBookingDate"`
	// Transaction amount
	Amount float64 `json:"amount"`
	// Transaction purpose. Maximum length: 2000
	Purpose string `json:"purpose,omitempty"`
	// Counterpart name. Maximum length: 80
	CounterpartName string `json:"counterpartName,omitempty"`
	// Counterpart account number
	CounterpartAccountNumber string `json:"counterpartAccountNumber,omitempty"`
	// Counterpart IBAN
	CounterpartIban string `json:"counterpartIban,omitempty"`
	// Counterpart BLZ
	CounterpartBlz string `json:"counterpartBlz,omitempty"`
	// Counterpart BIC
	CounterpartBic string `json:"counterpartBic,omitempty"`
	// Counterpart Bank name
	CounterpartBankName string `json:"counterpartBankName,omitempty"`
	// The mandate reference of the counterpart
	CounterpartMandateReference string `json:"counterpartMandateReference,omitempty"`
	// The customer reference of the counterpart
	CounterpartCustomerReference string `json:"counterpartCustomerReference,omitempty"`
	// The creditor ID of the counterpart. Exists only for SEPA direct debit transactions (\"Lastschrift\").
	CounterpartCreditorId string `json:"counterpartCreditorId,omitempty"`
	// The originator's identification code. Exists only for SEPA money transfer transactions (\"Überweisung\").
	CounterpartDebitorId string `json:"counterpartDebitorId,omitempty"`
	// Transaction type, according to the bank. If set, this will contain a German term that you can display to the user. Some examples of common values are: \"Lastschrift\", \"Auslands&uuml;berweisung\", \"Geb&uuml;hren\", \"Zinsen\". The maximum possible length of this field is 255 characters.
	Type_ string `json:"type,omitempty"`
	// ZKA business transaction code which relates to the transaction's type. Possible values range from 1 through 999. If no information about the ZKA type code is available, then this field will be null.
	TypeCodeZka string `json:"typeCodeZka,omitempty"`
	// SWIFT transaction type code. If no information about the SWIFT code is available, then this field will be null.
	TypeCodeSwift string `json:"typeCodeSwift,omitempty"`
	// SEPA purpose code, according to ISO 20022
	SepaPurposeCode string `json:"sepaPurposeCode,omitempty"`
	// Transaction primanota (bank side identification number)
	Primanota string `json:"primanota,omitempty"`
	// Transaction category, if any is assigned. Note: Recently imported transactions that have currently no category assigned might still get categorized by the background categorization process. To check the status of the background categorization, see GET /bankConnections. Manual category assignments to a transaction will remove the transaction from the background categorization process (i.e. the background categorization process will never overwrite a manual category assignment).
	Category *Category `json:"category,omitempty"`
	// Array of assigned labels
	Labels []Label `json:"labels,omitempty"`
	// While finAPI uses a well-elaborated algorithm for uniquely identifying transactions, there is still the possibility that during an account update, a transaction that was imported previously may be imported a second time as a new transaction. For example, this can happen if some transaction data changes on the bank server side. However, finAPI also includes an algorithm of identifying such \"potential duplicate\" transactions. If this field is set to true, it means that finAPI detected a similar transaction that might actually be the same. It is recommended to communicate this information to the end user, and give him an option to delete the transaction in case he confirms that it really is a duplicate.
	IsPotentialDuplicate bool `json:"isPotentialDuplicate"`
	// Indicating whether this transaction is an adjusting entry ('Zwischensaldo').<br/><br/>Adjusting entries do not originate from the bank, but are added by finAPI during an account update when the bank reported account balance does not add up to the set of transactions that finAPI receives for the account. In this case, the adjusting entry will fix the deviation between the balance and the received transactions so that both adds up again.<br/><br/>Possible causes for such deviations are:<br/>- Inconsistencies in how the bank calculates the balance, for instance when not yet booked transactions are already included in the balance, but not included in the set of transactions<br/>- Gaps in the transaction history that finAPI receives, for instance because the account has not been updated for a while and older transactions are no longer available
	IsAdjustingEntry bool `json:"isAdjustingEntry"`
	// Indicating whether this transaction is 'new' or not. Any newly imported transaction will have this flag initially set to true. How you use this field is up to your interpretation. For example, you might want to set it to false once a user has clicked on/seen the transaction. You can change this flag to 'false' with the PATCH method.
	IsNew bool `json:"isNew"`
	// Date of transaction import, in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	ImportDate string `json:"importDate"`
	// Sub-transactions identifiers (if this transaction is split)
	Children []int64 `json:"children,omitempty"`
	// Additional, PayPal-specific transaction data. This field is only set for transactions that belong to an account of the 'PayPal' bank (BLZ 'PAYPAL').<br/>NOTE: This field is deprecated as the bank with blz 'PAYPAL' is no longer supported. Do not use this field, as it will be removed at some point.
	PaypalData *PaypalTransactionData `json:"paypalData,omitempty"`
	// End-To-End reference
	EndToEndReference string `json:"endToEndReference,omitempty"`
	// Compensation Amount. Sum of reimbursement of out-of-pocket expenses plus processing brokerage in case of a national return / refund debit as well as an optional interest equalisation. Exists predominantly for SEPA direct debit returns.
	CompensationAmount float64 `json:"compensationAmount,omitempty"`
	// Original Amount of the original direct debit. Exists predominantly for SEPA direct debit returns.
	OriginalAmount float64 `json:"originalAmount,omitempty"`
	// Payer's/debtor's reference party (in the case of a credit transfer) or payee's/creditor's reference party (in the case of a direct debit)
	DifferentDebitor string `json:"differentDebitor,omitempty"`
	// Payee's/creditor's reference party (in the case of a credit transfer) or payer's/debtor's reference party (in the case of a direct debit)
	DifferentCreditor string `json:"differentCreditor,omitempty"`
}

/*
 * finAPI Access (with deprecation)
 *
 * <strong>RESTful API for Account Information Services (AIS) and Payment Initiation Services (PIS)</strong>  The following pages give you some general information on how to use our APIs.<br/> The actual API services documentation then follows further below. You can use the menu to jump between API sections. <br/> <br/> This page has a built-in HTTP(S) client, so you can test the services directly from within this page, by filling in the request parameters and/or body in the respective services, and then hitting the TRY button. Note that you need to be authorized to make a successful API call. To authorize, refer to the 'Authorization' section of the API, or just use the OAUTH button that can be found near the TRY button. <br/>  <h2 id=\"general-information\">General information</h2>  <h3 id=\"general-error-responses\"><strong>Error Responses</strong></h3> When an API call returns with an error, then in general it has the structure shown in the following example:  <pre> {   \"errors\": [     {       \"message\": \"Interface 'FINTS_SERVER' is not supported for this operation.\",       \"code\": \"BAD_REQUEST\",       \"type\": \"TECHNICAL\"     }   ],   \"date\": \"2020-11-19 16:54:06.854\",   \"requestId\": \"selfgen-312042e7-df55-47e4-bffd-956a68ef37b5\",   \"endpoint\": \"POST /api/v1/bankConnections/import\",   \"authContext\": \"1/21\",   \"bank\": \"DEMO0002 - finAPI Test Redirect Bank\" } </pre>  If an API call requires an additional authentication by the user, HTTP code 510 is returned and the error response contains the additional \"multiStepAuthentication\" object, see the following example:  <pre> {   \"errors\": [     {       \"message\": \"Es ist eine zusätzliche Authentifizierung erforderlich. Bitte geben Sie folgenden Code an: 123456\",       \"code\": \"ADDITIONAL_AUTHENTICATION_REQUIRED\",       \"type\": \"BUSINESS\",       \"multiStepAuthentication\": {         \"hash\": \"678b13f4be9ed7d981a840af8131223a\",         \"status\": \"CHALLENGE_RESPONSE_REQUIRED\",         \"challengeMessage\": \"Es ist eine zusätzliche Authentifizierung erforderlich. Bitte geben Sie folgenden Code an: 123456\",         \"answerFieldLabel\": \"TAN\",         \"redirectUrl\": null,         \"redirectContext\": null,         \"redirectContextField\": null,         \"twoStepProcedures\": null,         \"photoTanMimeType\": null,         \"photoTanData\": null,         \"opticalData\": null       }     }   ],   \"date\": \"2019-11-29 09:51:55.931\",   \"requestId\": \"selfgen-45059c99-1b14-4df7-9bd3-9d5f126df294\",   \"endpoint\": \"POST /api/v1/bankConnections/import\",   \"authContext\": \"1/18\",   \"bank\": \"DEMO0001 - finAPI Test Bank\" } </pre>  An exception to this error format are API authentication errors, where the following structure is returned:  <pre> {   \"error\": \"invalid_token\",   \"error_description\": \"Invalid access token: cccbce46-xxxx-xxxx-xxxx-xxxxxxxxxx\" } </pre>  <h3 id=\"general-paging\"><strong>Paging</strong></h3> API services that may potentially return a lot of data implement paging. They return a limited number of entries within a \"page\". Further entries must be fetched with subsequent calls. <br/><br/> Any API service that implements paging provides the following input parameters:<br/> &bull; \"page\": the number of the page to be retrieved (starting with 1).<br/> &bull; \"perPage\": the number of entries within a page. The default and maximum value is stated in the documentation of the respective services.  A paged response contains an additional \"paging\" object with the following structure:  <pre> {   ...   ,   \"paging\": {     \"page\": 1,     \"perPage\": 20,     \"pageCount\": 234,     \"totalCount\": 4662   } } </pre>  <h3 id=\"general-internationalization\"><strong>Internationalization</strong></h3> The finAPI services support internationalization which means you can define the language you prefer for API service responses. <br/><br/> The following languages are available: German, English, Czech, Slovak. <br/><br/> The preferred language can be defined by providing the official HTTP <strong>Accept-Language</strong> header. <br/><br/> finAPI reacts on the official iso language codes &quot;de&quot;, &quot;en&quot;, &quot;cs&quot; and &quot;sk&quot; for the named languages. Additional subtags supported by the Accept-Language header may be provided, e.g. &quot;en-US&quot;, but are ignored. <br/> If no Accept-Language header is given, German is used as the default language. <br/><br/> Exceptions:<br/> &bull; Bank login hints and login fields are only available in the language of the bank and not being translated.<br/> &bull; Direct messages from the bank systems typically returned as BUSINESS errors will not be translated.<br/> &bull; BUSINESS errors created by finAPI directly are available in German and English.<br/> &bull; TECHNICAL errors messages meant for developers are mostly in English, but also may be translated.  <h3 id=\"general-request-ids\"><strong>Request IDs</strong></h3> With any API call, you can pass a request ID via a header with name \"X-Request-Id\". The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. <br/><br/> If you don't pass a request ID for a call, finAPI will generate a random ID internally. <br/><br/> The request ID is always returned back in the response of a service, as a header with name \"X-Request-Id\". <br/><br/> We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster.  <h3 id=\"general-overriding-http-methods\"><strong>Overriding HTTP methods</strong></h3> Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with a special HTTP header indicating the originally intended HTTP method. <br/><br/> The header's name is <strong>X-HTTP-Method-Override</strong>. Set its value to either <strong>PATCH</strong> or <strong>DELETE</strong>. POST Requests having this header set will be treated either as PATCH or DELETE by the finAPI servers. <br/><br/> Example: <br/><br/> <strong>X-HTTP-Method-Override: PATCH</strong><br/> POST /api/v1/label/51<br/> {\"name\": \"changed label\"}<br/><br/> will be interpreted by finAPI as:<br/><br/> PATCH /api/v1/label/51<br/> {\"name\": \"changed label\"}<br/>  <h3 id=\"general-user-metadata\"><strong>User metadata</strong></h3> With the migration to PSD2 APIs, a new term called \"User metadata\" (also known as \"PSU metadata\") has been introduced to the API. This user metadata aims to inform the banking API if there was a real end-user behind an HTTP request or if the request was triggered by a system (e.g. by an automatic batch update). In the latter case, the bank may apply some restrictions such as limiting the number of HTTP requests for a single consent. Also, some operations may be forbidden entirely by the banking API. For example, some banks do not allow issuing a new consent without the end-user being involved. Therefore, it is certainly necessary and obligatory for the customer to provide the PSU metadata for such operations. <br/><br/> As finAPI does not have direct interaction with the end-user, it is the client application's responsibility to provide all the necessary information about the end-user. This must be done by sending additional headers with every request triggered on behalf of the end-user. <br/><br/> At the moment, the following headers are supported by the API:<br/> &bull; \"PSU-IP-Address\" - the IP address of the user's device.<br/> &bull; \"PSU-Device-OS\" - the user's device and/or operating system identification.<br/> &bull; \"PSU-User-Agent\" - the user's web browser or other client device identification.  <h3 id=\"general-faq\"><strong>FAQ</strong></h3> <strong>Is there a finAPI SDK?</strong> <br/> Currently we do not offer a native SDK, but there is the option to generate an SDK for almost any target language via OpenAPI. Use the 'Download SDK' button on this page for SDK generation. <br/> <br/> <strong>How can I enable finAPI's automatic batch update?</strong> <br/> Currently there is no way to set up the batch update via the API. Please contact support@finapi.io for this. <br/> <br/> <strong>Why do I need to keep authorizing when calling services on this page?</strong> <br/> This page is a \"one-page-app\". Reloading the page resets the OAuth authorization context. There is generally no need to reload the page, so just don't do it and your authorization will persist. 
 *
 * API version: 1.151.1
 * Contact: kontakt@finapi.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finapi

import (
	"encoding/json"
)

// Transaction Container for a transaction's data
type Transaction struct {
	// Transaction identifier
	Id int64 `json:"id"`
	// Parent transaction identifier
	ParentId NullableInt64 `json:"parentId"`
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
	// <strong>Type:</strong> Currency<br/> Transaction currency in ISO 4217 format.This field can be null if not explicitly provided the bank. In this case it can be assumed as account’s currency.
	Currency NullableCurrency `json:"currency"`
	// Transaction purpose. Maximum length: 2000
	Purpose NullableString `json:"purpose"`
	// Counterpart name. Maximum length: 80
	CounterpartName NullableString `json:"counterpartName"`
	// Counterpart account number
	CounterpartAccountNumber NullableString `json:"counterpartAccountNumber"`
	// Counterpart IBAN
	CounterpartIban NullableString `json:"counterpartIban"`
	// Counterpart BLZ
	CounterpartBlz NullableString `json:"counterpartBlz"`
	// Counterpart BIC
	CounterpartBic NullableString `json:"counterpartBic"`
	// Counterpart Bank name
	CounterpartBankName NullableString `json:"counterpartBankName"`
	// The mandate reference of the counterpart
	CounterpartMandateReference NullableString `json:"counterpartMandateReference"`
	// The customer reference of the counterpart
	CounterpartCustomerReference NullableString `json:"counterpartCustomerReference"`
	// The creditor ID of the counterpart. Exists only for SEPA direct debit transactions (\"Lastschrift\").
	CounterpartCreditorId NullableString `json:"counterpartCreditorId"`
	// The originator's identification code. Exists only for SEPA money transfer transactions (\"Überweisung\").
	CounterpartDebitorId NullableString `json:"counterpartDebitorId"`
	// Transaction type, according to the bank. If set, this will contain a German term that you can display to the user. Some examples of common values are: \"Lastschrift\", \"Auslands&uuml;berweisung\", \"Geb&uuml;hren\", \"Zinsen\". The maximum possible length of this field is 255 characters.
	Type NullableString `json:"type"`
	// ZKA business transaction code which relates to the transaction's type. Possible values range from 1 through 999. If no information about the ZKA type code is available, then this field will be null.
	TypeCodeZka NullableString `json:"typeCodeZka"`
	// SWIFT transaction type code. If no information about the SWIFT code is available, then this field will be null.
	TypeCodeSwift NullableString `json:"typeCodeSwift"`
	// SEPA purpose code, according to ISO 20022
	SepaPurposeCode NullableString `json:"sepaPurposeCode"`
	// Bank transaction code, according to ISO 20022
	BankTransactionCode NullableString `json:"bankTransactionCode"`
	// Transaction primanota (bank side identification number)
	Primanota NullableString `json:"primanota"`
	// <strong>Type:</strong> Category<br/> Transaction category, if any is assigned. Note: Recently imported transactions that have currently no category assigned might still get categorized by the background categorization process. To check the status of the background categorization, see GET /bankConnections. Manual category assignments to a transaction will remove the transaction from the background categorization process (i.e. the background categorization process will never overwrite a manual category assignment).
	Category NullableCategory `json:"category"`
	// <strong>Type:</strong> Label<br/> Array of assigned labels
	Labels []Label `json:"labels"`
	// While finAPI uses a well-elaborated algorithm for uniquely identifying transactions, there is still the possibility that during an account update, a transaction that was imported previously may be imported a second time as a new transaction. For example, this can happen if some transaction data changes on the bank server side. However, finAPI also includes an algorithm of identifying such \"potential duplicate\" transactions. If this field is set to true, it means that finAPI detected a similar transaction that might actually be the same. It is recommended to communicate this information to the end user, and give him an option to delete the transaction in case he confirms that it really is a duplicate.
	IsPotentialDuplicate bool `json:"isPotentialDuplicate"`
	// Indicating whether this transaction is an adjusting entry ('Zwischensaldo').<br/><br/>Adjusting entries do not originate from the bank, but are added by finAPI during an account update when the bank reported account balance does not add up to the set of transactions that finAPI receives for the account. In this case, the adjusting entry will fix the deviation between the balance and the received transactions so that both adds up again.<br/><br/>Possible causes for such deviations are:<br/>- Inconsistencies in how the bank calculates the balance, for instance when not yet booked transactions are already included in the balance, but not included in the set of transactions<br/>- Gaps in the transaction history that finAPI receives, for instance because the account has not been updated for a while and older transactions are no longer available
	IsAdjustingEntry bool `json:"isAdjustingEntry"`
	// Indicating whether this transaction is 'new' or not. Any newly imported transaction will have this flag initially set to true. How you use this field is up to your interpretation. For example, you might want to set it to false once a user has clicked on/seen the transaction. You can change this flag to 'false' with the PATCH method.
	IsNew bool `json:"isNew"`
	// Date of transaction import, in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	ImportDate string `json:"importDate"`
	// Sub-transactions identifiers (if this transaction is split)
	Children []int64 `json:"children"`
	// <strong>Type:</strong> PaypalTransactionData<br/> Additional, PayPal-specific transaction data.
	PaypalData NullablePaypalTransactionData `json:"paypalData"`
	// End-To-End reference
	EndToEndReference NullableString `json:"endToEndReference"`
	// Compensation Amount. Sum of reimbursement of out-of-pocket expenses plus processing brokerage in case of a national return / refund debit as well as an optional interest equalisation. Exists predominantly for SEPA direct debit returns.
	CompensationAmount NullableFloat64 `json:"compensationAmount"`
	// Original Amount of the original direct debit. Exists predominantly for SEPA direct debit returns.
	OriginalAmount NullableFloat64 `json:"originalAmount"`
	// <strong>Type:</strong> Currency<br/> Currency of the original amount in ISO 4217 format. This field can be null if not explicitly provided the bank. In this case it can be assumed as account’s currency.
	OriginalCurrency NullableCurrency `json:"originalCurrency"`
	// Payer's/debtor's reference party (in the case of a credit transfer) or payee's/creditor's reference party (in the case of a direct debit)
	DifferentDebitor NullableString `json:"differentDebitor"`
	// Payee's/creditor's reference party (in the case of a credit transfer) or payer's/debtor's reference party (in the case of a direct debit)
	DifferentCreditor NullableString `json:"differentCreditor"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(id int64, parentId NullableInt64, accountId int64, valueDate string, bankBookingDate string, finapiBookingDate string, amount float64, currency NullableCurrency, purpose NullableString, counterpartName NullableString, counterpartAccountNumber NullableString, counterpartIban NullableString, counterpartBlz NullableString, counterpartBic NullableString, counterpartBankName NullableString, counterpartMandateReference NullableString, counterpartCustomerReference NullableString, counterpartCreditorId NullableString, counterpartDebitorId NullableString, type_ NullableString, typeCodeZka NullableString, typeCodeSwift NullableString, sepaPurposeCode NullableString, bankTransactionCode NullableString, primanota NullableString, category NullableCategory, labels []Label, isPotentialDuplicate bool, isAdjustingEntry bool, isNew bool, importDate string, children []int64, paypalData NullablePaypalTransactionData, endToEndReference NullableString, compensationAmount NullableFloat64, originalAmount NullableFloat64, originalCurrency NullableCurrency, differentDebitor NullableString, differentCreditor NullableString, ) *Transaction {
	this := Transaction{}
	this.Id = id
	this.ParentId = parentId
	this.AccountId = accountId
	this.ValueDate = valueDate
	this.BankBookingDate = bankBookingDate
	this.FinapiBookingDate = finapiBookingDate
	this.Amount = amount
	this.Currency = currency
	this.Purpose = purpose
	this.CounterpartName = counterpartName
	this.CounterpartAccountNumber = counterpartAccountNumber
	this.CounterpartIban = counterpartIban
	this.CounterpartBlz = counterpartBlz
	this.CounterpartBic = counterpartBic
	this.CounterpartBankName = counterpartBankName
	this.CounterpartMandateReference = counterpartMandateReference
	this.CounterpartCustomerReference = counterpartCustomerReference
	this.CounterpartCreditorId = counterpartCreditorId
	this.CounterpartDebitorId = counterpartDebitorId
	this.Type = type_
	this.TypeCodeZka = typeCodeZka
	this.TypeCodeSwift = typeCodeSwift
	this.SepaPurposeCode = sepaPurposeCode
	this.BankTransactionCode = bankTransactionCode
	this.Primanota = primanota
	this.Category = category
	this.Labels = labels
	this.IsPotentialDuplicate = isPotentialDuplicate
	this.IsAdjustingEntry = isAdjustingEntry
	this.IsNew = isNew
	this.ImportDate = importDate
	this.Children = children
	this.PaypalData = paypalData
	this.EndToEndReference = endToEndReference
	this.CompensationAmount = compensationAmount
	this.OriginalAmount = originalAmount
	this.OriginalCurrency = originalCurrency
	this.DifferentDebitor = differentDebitor
	this.DifferentCreditor = differentCreditor
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetId returns the Id field value
func (o *Transaction) GetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Transaction) SetId(v int64) {
	o.Id = v
}

// GetParentId returns the ParentId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *Transaction) GetParentId() int64 {
	if o == nil || o.ParentId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetParentIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// SetParentId sets field value
func (o *Transaction) SetParentId(v int64) {
	o.ParentId.Set(&v)
}

// GetAccountId returns the AccountId field value
func (o *Transaction) GetAccountId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetAccountIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Transaction) SetAccountId(v int64) {
	o.AccountId = v
}

// GetValueDate returns the ValueDate field value
func (o *Transaction) GetValueDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ValueDate
}

// GetValueDateOk returns a tuple with the ValueDate field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetValueDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ValueDate, true
}

// SetValueDate sets field value
func (o *Transaction) SetValueDate(v string) {
	o.ValueDate = v
}

// GetBankBookingDate returns the BankBookingDate field value
func (o *Transaction) GetBankBookingDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BankBookingDate
}

// GetBankBookingDateOk returns a tuple with the BankBookingDate field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetBankBookingDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankBookingDate, true
}

// SetBankBookingDate sets field value
func (o *Transaction) SetBankBookingDate(v string) {
	o.BankBookingDate = v
}

// GetFinapiBookingDate returns the FinapiBookingDate field value
func (o *Transaction) GetFinapiBookingDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.FinapiBookingDate
}

// GetFinapiBookingDateOk returns a tuple with the FinapiBookingDate field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetFinapiBookingDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FinapiBookingDate, true
}

// SetFinapiBookingDate sets field value
func (o *Transaction) SetFinapiBookingDate(v string) {
	o.FinapiBookingDate = v
}

// GetAmount returns the Amount field value
func (o *Transaction) GetAmount() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Transaction) SetAmount(v float64) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for Currency will be returned
func (o *Transaction) GetCurrency() Currency {
	if o == nil || o.Currency.Get() == nil {
		var ret Currency
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCurrencyOk() (*Currency, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *Transaction) SetCurrency(v Currency) {
	o.Currency.Set(&v)
}

// GetPurpose returns the Purpose field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetPurpose() string {
	if o == nil || o.Purpose.Get() == nil {
		var ret string
		return ret
	}

	return *o.Purpose.Get()
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetPurposeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Purpose.Get(), o.Purpose.IsSet()
}

// SetPurpose sets field value
func (o *Transaction) SetPurpose(v string) {
	o.Purpose.Set(&v)
}

// GetCounterpartName returns the CounterpartName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetCounterpartName() string {
	if o == nil || o.CounterpartName.Get() == nil {
		var ret string
		return ret
	}

	return *o.CounterpartName.Get()
}

// GetCounterpartNameOk returns a tuple with the CounterpartName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCounterpartNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CounterpartName.Get(), o.CounterpartName.IsSet()
}

// SetCounterpartName sets field value
func (o *Transaction) SetCounterpartName(v string) {
	o.CounterpartName.Set(&v)
}

// GetCounterpartAccountNumber returns the CounterpartAccountNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetCounterpartAccountNumber() string {
	if o == nil || o.CounterpartAccountNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.CounterpartAccountNumber.Get()
}

// GetCounterpartAccountNumberOk returns a tuple with the CounterpartAccountNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCounterpartAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CounterpartAccountNumber.Get(), o.CounterpartAccountNumber.IsSet()
}

// SetCounterpartAccountNumber sets field value
func (o *Transaction) SetCounterpartAccountNumber(v string) {
	o.CounterpartAccountNumber.Set(&v)
}

// GetCounterpartIban returns the CounterpartIban field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetCounterpartIban() string {
	if o == nil || o.CounterpartIban.Get() == nil {
		var ret string
		return ret
	}

	return *o.CounterpartIban.Get()
}

// GetCounterpartIbanOk returns a tuple with the CounterpartIban field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCounterpartIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CounterpartIban.Get(), o.CounterpartIban.IsSet()
}

// SetCounterpartIban sets field value
func (o *Transaction) SetCounterpartIban(v string) {
	o.CounterpartIban.Set(&v)
}

// GetCounterpartBlz returns the CounterpartBlz field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetCounterpartBlz() string {
	if o == nil || o.CounterpartBlz.Get() == nil {
		var ret string
		return ret
	}

	return *o.CounterpartBlz.Get()
}

// GetCounterpartBlzOk returns a tuple with the CounterpartBlz field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCounterpartBlzOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CounterpartBlz.Get(), o.CounterpartBlz.IsSet()
}

// SetCounterpartBlz sets field value
func (o *Transaction) SetCounterpartBlz(v string) {
	o.CounterpartBlz.Set(&v)
}

// GetCounterpartBic returns the CounterpartBic field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetCounterpartBic() string {
	if o == nil || o.CounterpartBic.Get() == nil {
		var ret string
		return ret
	}

	return *o.CounterpartBic.Get()
}

// GetCounterpartBicOk returns a tuple with the CounterpartBic field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCounterpartBicOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CounterpartBic.Get(), o.CounterpartBic.IsSet()
}

// SetCounterpartBic sets field value
func (o *Transaction) SetCounterpartBic(v string) {
	o.CounterpartBic.Set(&v)
}

// GetCounterpartBankName returns the CounterpartBankName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetCounterpartBankName() string {
	if o == nil || o.CounterpartBankName.Get() == nil {
		var ret string
		return ret
	}

	return *o.CounterpartBankName.Get()
}

// GetCounterpartBankNameOk returns a tuple with the CounterpartBankName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCounterpartBankNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CounterpartBankName.Get(), o.CounterpartBankName.IsSet()
}

// SetCounterpartBankName sets field value
func (o *Transaction) SetCounterpartBankName(v string) {
	o.CounterpartBankName.Set(&v)
}

// GetCounterpartMandateReference returns the CounterpartMandateReference field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetCounterpartMandateReference() string {
	if o == nil || o.CounterpartMandateReference.Get() == nil {
		var ret string
		return ret
	}

	return *o.CounterpartMandateReference.Get()
}

// GetCounterpartMandateReferenceOk returns a tuple with the CounterpartMandateReference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCounterpartMandateReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CounterpartMandateReference.Get(), o.CounterpartMandateReference.IsSet()
}

// SetCounterpartMandateReference sets field value
func (o *Transaction) SetCounterpartMandateReference(v string) {
	o.CounterpartMandateReference.Set(&v)
}

// GetCounterpartCustomerReference returns the CounterpartCustomerReference field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetCounterpartCustomerReference() string {
	if o == nil || o.CounterpartCustomerReference.Get() == nil {
		var ret string
		return ret
	}

	return *o.CounterpartCustomerReference.Get()
}

// GetCounterpartCustomerReferenceOk returns a tuple with the CounterpartCustomerReference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCounterpartCustomerReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CounterpartCustomerReference.Get(), o.CounterpartCustomerReference.IsSet()
}

// SetCounterpartCustomerReference sets field value
func (o *Transaction) SetCounterpartCustomerReference(v string) {
	o.CounterpartCustomerReference.Set(&v)
}

// GetCounterpartCreditorId returns the CounterpartCreditorId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetCounterpartCreditorId() string {
	if o == nil || o.CounterpartCreditorId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CounterpartCreditorId.Get()
}

// GetCounterpartCreditorIdOk returns a tuple with the CounterpartCreditorId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCounterpartCreditorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CounterpartCreditorId.Get(), o.CounterpartCreditorId.IsSet()
}

// SetCounterpartCreditorId sets field value
func (o *Transaction) SetCounterpartCreditorId(v string) {
	o.CounterpartCreditorId.Set(&v)
}

// GetCounterpartDebitorId returns the CounterpartDebitorId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetCounterpartDebitorId() string {
	if o == nil || o.CounterpartDebitorId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CounterpartDebitorId.Get()
}

// GetCounterpartDebitorIdOk returns a tuple with the CounterpartDebitorId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCounterpartDebitorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CounterpartDebitorId.Get(), o.CounterpartDebitorId.IsSet()
}

// SetCounterpartDebitorId sets field value
func (o *Transaction) SetCounterpartDebitorId(v string) {
	o.CounterpartDebitorId.Set(&v)
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *Transaction) SetType(v string) {
	o.Type.Set(&v)
}

// GetTypeCodeZka returns the TypeCodeZka field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetTypeCodeZka() string {
	if o == nil || o.TypeCodeZka.Get() == nil {
		var ret string
		return ret
	}

	return *o.TypeCodeZka.Get()
}

// GetTypeCodeZkaOk returns a tuple with the TypeCodeZka field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetTypeCodeZkaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TypeCodeZka.Get(), o.TypeCodeZka.IsSet()
}

// SetTypeCodeZka sets field value
func (o *Transaction) SetTypeCodeZka(v string) {
	o.TypeCodeZka.Set(&v)
}

// GetTypeCodeSwift returns the TypeCodeSwift field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetTypeCodeSwift() string {
	if o == nil || o.TypeCodeSwift.Get() == nil {
		var ret string
		return ret
	}

	return *o.TypeCodeSwift.Get()
}

// GetTypeCodeSwiftOk returns a tuple with the TypeCodeSwift field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetTypeCodeSwiftOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TypeCodeSwift.Get(), o.TypeCodeSwift.IsSet()
}

// SetTypeCodeSwift sets field value
func (o *Transaction) SetTypeCodeSwift(v string) {
	o.TypeCodeSwift.Set(&v)
}

// GetSepaPurposeCode returns the SepaPurposeCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetSepaPurposeCode() string {
	if o == nil || o.SepaPurposeCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.SepaPurposeCode.Get()
}

// GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetSepaPurposeCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SepaPurposeCode.Get(), o.SepaPurposeCode.IsSet()
}

// SetSepaPurposeCode sets field value
func (o *Transaction) SetSepaPurposeCode(v string) {
	o.SepaPurposeCode.Set(&v)
}

// GetBankTransactionCode returns the BankTransactionCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetBankTransactionCode() string {
	if o == nil || o.BankTransactionCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.BankTransactionCode.Get()
}

// GetBankTransactionCodeOk returns a tuple with the BankTransactionCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetBankTransactionCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankTransactionCode.Get(), o.BankTransactionCode.IsSet()
}

// SetBankTransactionCode sets field value
func (o *Transaction) SetBankTransactionCode(v string) {
	o.BankTransactionCode.Set(&v)
}

// GetPrimanota returns the Primanota field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetPrimanota() string {
	if o == nil || o.Primanota.Get() == nil {
		var ret string
		return ret
	}

	return *o.Primanota.Get()
}

// GetPrimanotaOk returns a tuple with the Primanota field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetPrimanotaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Primanota.Get(), o.Primanota.IsSet()
}

// SetPrimanota sets field value
func (o *Transaction) SetPrimanota(v string) {
	o.Primanota.Set(&v)
}

// GetCategory returns the Category field value
// If the value is explicit nil, the zero value for Category will be returned
func (o *Transaction) GetCategory() Category {
	if o == nil || o.Category.Get() == nil {
		var ret Category
		return ret
	}

	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCategoryOk() (*Category, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// SetCategory sets field value
func (o *Transaction) SetCategory(v Category) {
	o.Category.Set(&v)
}

// GetLabels returns the Labels field value
func (o *Transaction) GetLabels() []Label {
	if o == nil  {
		var ret []Label
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetLabelsOk() (*[]Label, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *Transaction) SetLabels(v []Label) {
	o.Labels = v
}

// GetIsPotentialDuplicate returns the IsPotentialDuplicate field value
func (o *Transaction) GetIsPotentialDuplicate() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsPotentialDuplicate
}

// GetIsPotentialDuplicateOk returns a tuple with the IsPotentialDuplicate field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetIsPotentialDuplicateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsPotentialDuplicate, true
}

// SetIsPotentialDuplicate sets field value
func (o *Transaction) SetIsPotentialDuplicate(v bool) {
	o.IsPotentialDuplicate = v
}

// GetIsAdjustingEntry returns the IsAdjustingEntry field value
func (o *Transaction) GetIsAdjustingEntry() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsAdjustingEntry
}

// GetIsAdjustingEntryOk returns a tuple with the IsAdjustingEntry field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetIsAdjustingEntryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsAdjustingEntry, true
}

// SetIsAdjustingEntry sets field value
func (o *Transaction) SetIsAdjustingEntry(v bool) {
	o.IsAdjustingEntry = v
}

// GetIsNew returns the IsNew field value
func (o *Transaction) GetIsNew() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsNew
}

// GetIsNewOk returns a tuple with the IsNew field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetIsNewOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsNew, true
}

// SetIsNew sets field value
func (o *Transaction) SetIsNew(v bool) {
	o.IsNew = v
}

// GetImportDate returns the ImportDate field value
func (o *Transaction) GetImportDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ImportDate
}

// GetImportDateOk returns a tuple with the ImportDate field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetImportDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImportDate, true
}

// SetImportDate sets field value
func (o *Transaction) SetImportDate(v string) {
	o.ImportDate = v
}

// GetChildren returns the Children field value
// If the value is explicit nil, the zero value for []int64 will be returned
func (o *Transaction) GetChildren() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetChildrenOk() (*[]int64, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return &o.Children, true
}

// SetChildren sets field value
func (o *Transaction) SetChildren(v []int64) {
	o.Children = v
}

// GetPaypalData returns the PaypalData field value
// If the value is explicit nil, the zero value for PaypalTransactionData will be returned
func (o *Transaction) GetPaypalData() PaypalTransactionData {
	if o == nil || o.PaypalData.Get() == nil {
		var ret PaypalTransactionData
		return ret
	}

	return *o.PaypalData.Get()
}

// GetPaypalDataOk returns a tuple with the PaypalData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetPaypalDataOk() (*PaypalTransactionData, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaypalData.Get(), o.PaypalData.IsSet()
}

// SetPaypalData sets field value
func (o *Transaction) SetPaypalData(v PaypalTransactionData) {
	o.PaypalData.Set(&v)
}

// GetEndToEndReference returns the EndToEndReference field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetEndToEndReference() string {
	if o == nil || o.EndToEndReference.Get() == nil {
		var ret string
		return ret
	}

	return *o.EndToEndReference.Get()
}

// GetEndToEndReferenceOk returns a tuple with the EndToEndReference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetEndToEndReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndToEndReference.Get(), o.EndToEndReference.IsSet()
}

// SetEndToEndReference sets field value
func (o *Transaction) SetEndToEndReference(v string) {
	o.EndToEndReference.Set(&v)
}

// GetCompensationAmount returns the CompensationAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Transaction) GetCompensationAmount() float64 {
	if o == nil || o.CompensationAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.CompensationAmount.Get()
}

// GetCompensationAmountOk returns a tuple with the CompensationAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCompensationAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompensationAmount.Get(), o.CompensationAmount.IsSet()
}

// SetCompensationAmount sets field value
func (o *Transaction) SetCompensationAmount(v float64) {
	o.CompensationAmount.Set(&v)
}

// GetOriginalAmount returns the OriginalAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Transaction) GetOriginalAmount() float64 {
	if o == nil || o.OriginalAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.OriginalAmount.Get()
}

// GetOriginalAmountOk returns a tuple with the OriginalAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetOriginalAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalAmount.Get(), o.OriginalAmount.IsSet()
}

// SetOriginalAmount sets field value
func (o *Transaction) SetOriginalAmount(v float64) {
	o.OriginalAmount.Set(&v)
}

// GetOriginalCurrency returns the OriginalCurrency field value
// If the value is explicit nil, the zero value for Currency will be returned
func (o *Transaction) GetOriginalCurrency() Currency {
	if o == nil || o.OriginalCurrency.Get() == nil {
		var ret Currency
		return ret
	}

	return *o.OriginalCurrency.Get()
}

// GetOriginalCurrencyOk returns a tuple with the OriginalCurrency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetOriginalCurrencyOk() (*Currency, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalCurrency.Get(), o.OriginalCurrency.IsSet()
}

// SetOriginalCurrency sets field value
func (o *Transaction) SetOriginalCurrency(v Currency) {
	o.OriginalCurrency.Set(&v)
}

// GetDifferentDebitor returns the DifferentDebitor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetDifferentDebitor() string {
	if o == nil || o.DifferentDebitor.Get() == nil {
		var ret string
		return ret
	}

	return *o.DifferentDebitor.Get()
}

// GetDifferentDebitorOk returns a tuple with the DifferentDebitor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetDifferentDebitorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DifferentDebitor.Get(), o.DifferentDebitor.IsSet()
}

// SetDifferentDebitor sets field value
func (o *Transaction) SetDifferentDebitor(v string) {
	o.DifferentDebitor.Set(&v)
}

// GetDifferentCreditor returns the DifferentCreditor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetDifferentCreditor() string {
	if o == nil || o.DifferentCreditor.Get() == nil {
		var ret string
		return ret
	}

	return *o.DifferentCreditor.Get()
}

// GetDifferentCreditorOk returns a tuple with the DifferentCreditor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetDifferentCreditorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DifferentCreditor.Get(), o.DifferentCreditor.IsSet()
}

// SetDifferentCreditor sets field value
func (o *Transaction) SetDifferentCreditor(v string) {
	o.DifferentCreditor.Set(&v)
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["valueDate"] = o.ValueDate
	}
	if true {
		toSerialize["bankBookingDate"] = o.BankBookingDate
	}
	if true {
		toSerialize["finapiBookingDate"] = o.FinapiBookingDate
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency.Get()
	}
	if true {
		toSerialize["purpose"] = o.Purpose.Get()
	}
	if true {
		toSerialize["counterpartName"] = o.CounterpartName.Get()
	}
	if true {
		toSerialize["counterpartAccountNumber"] = o.CounterpartAccountNumber.Get()
	}
	if true {
		toSerialize["counterpartIban"] = o.CounterpartIban.Get()
	}
	if true {
		toSerialize["counterpartBlz"] = o.CounterpartBlz.Get()
	}
	if true {
		toSerialize["counterpartBic"] = o.CounterpartBic.Get()
	}
	if true {
		toSerialize["counterpartBankName"] = o.CounterpartBankName.Get()
	}
	if true {
		toSerialize["counterpartMandateReference"] = o.CounterpartMandateReference.Get()
	}
	if true {
		toSerialize["counterpartCustomerReference"] = o.CounterpartCustomerReference.Get()
	}
	if true {
		toSerialize["counterpartCreditorId"] = o.CounterpartCreditorId.Get()
	}
	if true {
		toSerialize["counterpartDebitorId"] = o.CounterpartDebitorId.Get()
	}
	if true {
		toSerialize["type"] = o.Type.Get()
	}
	if true {
		toSerialize["typeCodeZka"] = o.TypeCodeZka.Get()
	}
	if true {
		toSerialize["typeCodeSwift"] = o.TypeCodeSwift.Get()
	}
	if true {
		toSerialize["sepaPurposeCode"] = o.SepaPurposeCode.Get()
	}
	if true {
		toSerialize["bankTransactionCode"] = o.BankTransactionCode.Get()
	}
	if true {
		toSerialize["primanota"] = o.Primanota.Get()
	}
	if true {
		toSerialize["category"] = o.Category.Get()
	}
	if true {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["isPotentialDuplicate"] = o.IsPotentialDuplicate
	}
	if true {
		toSerialize["isAdjustingEntry"] = o.IsAdjustingEntry
	}
	if true {
		toSerialize["isNew"] = o.IsNew
	}
	if true {
		toSerialize["importDate"] = o.ImportDate
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if true {
		toSerialize["paypalData"] = o.PaypalData.Get()
	}
	if true {
		toSerialize["endToEndReference"] = o.EndToEndReference.Get()
	}
	if true {
		toSerialize["compensationAmount"] = o.CompensationAmount.Get()
	}
	if true {
		toSerialize["originalAmount"] = o.OriginalAmount.Get()
	}
	if true {
		toSerialize["originalCurrency"] = o.OriginalCurrency.Get()
	}
	if true {
		toSerialize["differentDebitor"] = o.DifferentDebitor.Get()
	}
	if true {
		toSerialize["differentCreditor"] = o.DifferentCreditor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



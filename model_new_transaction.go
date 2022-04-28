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

// ExtendedTransaction Mock transaction data
type ExtendedTransaction struct {
	// Amount. Required.
	Amount float64 `json:"amount"`
	// <strong>Type:</strong> Currency<br/> Transaction currency.
	Currency *Currency `json:"currency,omitempty"`
	// Original amount
	OriginalAmount *float64 `json:"originalAmount,omitempty"`
	// <strong>Type:</strong> Currency<br/> Currency of the original amount.
	OriginalCurrency *Currency `json:"originalCurrency,omitempty"`
	// Purpose. Any symbols are allowed. Maximum length is 2000. Optional. Default value: null.
	Purpose *string `json:"purpose,omitempty"`
	// Counterpart. Any symbols are allowed. Maximum length is 80. Optional. Default value: null.
	Counterpart *string `json:"counterpart,omitempty"`
	// Counterpart IBAN. Optional. Default value: null.
	CounterpartIban *string `json:"counterpartIban,omitempty"`
	// Counterpart BLZ. Optional. Default value: null.
	CounterpartBlz *string `json:"counterpartBlz,omitempty"`
	// Counterpart BIC. Optional. Default value: null.
	CounterpartBic *string `json:"counterpartBic,omitempty"`
	// Counterpart account number. Maximum length is 34. Optional. Default value: null.
	CounterpartAccountNumber *string `json:"counterpartAccountNumber,omitempty"`
	// Booking date in the format 'YYYY-MM-DD'.<br/><br/>If the date lies back more than 10 days from the booking date of the latest transaction that currently exists in the account, then this transaction will be ignored and not imported. If the date depicts a date in the future, then finAPI will deal with it the same way as it does with real transactions during a real update (see fields 'bankBookingDate' and 'finapiBookingDate' in the Transaction Resource for explanation).<br/><br/>This field is optional, default value is the current date.
	BookingDate *string `json:"bookingDate,omitempty"`
	// Value date in the format 'YYYY-MM-DD'. Optional. Default value: Same as the booking date.
	ValueDate *string `json:"valueDate,omitempty"`
	// The transaction type id. It's usually a number between 1 and 999. You can look up valid transaction in the following document on page 198: <a href='https://www.hbci-zka.de/dokumente/spezifikation_deutsch/fintsv4/FinTS_4.1_Messages_Finanzdatenformate_2014-01-20-FV.pdf' target='_blank'>FinTS Financial Transaction Services</a>.<br/> For numbers not listed here, the service call might fail.
	TypeId *int32 `json:"typeId,omitempty"`
	// The mandate reference of the counterpart. The maximum possible length of this field is 270 characters.
	CounterpartMandateReference *string `json:"counterpartMandateReference,omitempty"`
	// The creditor ID of the counterpart. Exists only for SEPA direct debit transactions (\"Lastschrift\"). The maximum possible length of this field is 270 characters.
	CounterpartCreditorId *string `json:"counterpartCreditorId,omitempty"`
	// The customer reference of the counterpart. The maximum possible length of this field is 270 characters.
	CounterpartCustomerReference *string `json:"counterpartCustomerReference,omitempty"`
	// The originator's identification code. Exists only for SEPA money transfer transactions (\"Überweisung\"). The maximum possible length of this field is 100 characters.
	CounterpartDebitorId *string `json:"counterpartDebitorId,omitempty"`
	// Transaction type, according to the bank. If set, this will contain a German term that you can display to the user. Some examples of common values are: \"Lastschrift\", \"Auslands&uuml;berweisung\", \"Geb&uuml;hren\", \"Zinsen\". The maximum possible length of this field is 270 characters.
	Type *string `json:"type,omitempty"`
	// SWIFT transaction type code.
	TypeCodeSwift *string `json:"typeCodeSwift,omitempty"`
	// SEPA purpose code.
	SepaPurposeCode *string `json:"sepaPurposeCode,omitempty"`
}

// NewExtendedTransaction instantiates a new ExtendedTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedTransaction(amount float64, ) *ExtendedTransaction {
	this := ExtendedTransaction{}
	this.Amount = amount
	return &this
}

// NewExtendedTransactionWithDefaults instantiates a new ExtendedTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedTransactionWithDefaults() *ExtendedTransaction {
	this := ExtendedTransaction{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ExtendedTransaction) GetAmount() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ExtendedTransaction) SetAmount(v float64) {
	o.Amount = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetCurrency() Currency {
	if o == nil || o.Currency == nil {
		var ret Currency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetCurrencyOk() (*Currency, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given Currency and assigns it to the Currency field.
func (o *ExtendedTransaction) SetCurrency(v Currency) {
	o.Currency = &v
}

// GetOriginalAmount returns the OriginalAmount field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetOriginalAmount() float64 {
	if o == nil || o.OriginalAmount == nil {
		var ret float64
		return ret
	}
	return *o.OriginalAmount
}

// GetOriginalAmountOk returns a tuple with the OriginalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetOriginalAmountOk() (*float64, bool) {
	if o == nil || o.OriginalAmount == nil {
		return nil, false
	}
	return o.OriginalAmount, true
}

// HasOriginalAmount returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasOriginalAmount() bool {
	if o != nil && o.OriginalAmount != nil {
		return true
	}

	return false
}

// SetOriginalAmount gets a reference to the given float64 and assigns it to the OriginalAmount field.
func (o *ExtendedTransaction) SetOriginalAmount(v float64) {
	o.OriginalAmount = &v
}

// GetOriginalCurrency returns the OriginalCurrency field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetOriginalCurrency() Currency {
	if o == nil || o.OriginalCurrency == nil {
		var ret Currency
		return ret
	}
	return *o.OriginalCurrency
}

// GetOriginalCurrencyOk returns a tuple with the OriginalCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetOriginalCurrencyOk() (*Currency, bool) {
	if o == nil || o.OriginalCurrency == nil {
		return nil, false
	}
	return o.OriginalCurrency, true
}

// HasOriginalCurrency returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasOriginalCurrency() bool {
	if o != nil && o.OriginalCurrency != nil {
		return true
	}

	return false
}

// SetOriginalCurrency gets a reference to the given Currency and assigns it to the OriginalCurrency field.
func (o *ExtendedTransaction) SetOriginalCurrency(v Currency) {
	o.OriginalCurrency = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *ExtendedTransaction) SetPurpose(v string) {
	o.Purpose = &v
}

// GetCounterpart returns the Counterpart field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetCounterpart() string {
	if o == nil || o.Counterpart == nil {
		var ret string
		return ret
	}
	return *o.Counterpart
}

// GetCounterpartOk returns a tuple with the Counterpart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetCounterpartOk() (*string, bool) {
	if o == nil || o.Counterpart == nil {
		return nil, false
	}
	return o.Counterpart, true
}

// HasCounterpart returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasCounterpart() bool {
	if o != nil && o.Counterpart != nil {
		return true
	}

	return false
}

// SetCounterpart gets a reference to the given string and assigns it to the Counterpart field.
func (o *ExtendedTransaction) SetCounterpart(v string) {
	o.Counterpart = &v
}

// GetCounterpartIban returns the CounterpartIban field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetCounterpartIban() string {
	if o == nil || o.CounterpartIban == nil {
		var ret string
		return ret
	}
	return *o.CounterpartIban
}

// GetCounterpartIbanOk returns a tuple with the CounterpartIban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetCounterpartIbanOk() (*string, bool) {
	if o == nil || o.CounterpartIban == nil {
		return nil, false
	}
	return o.CounterpartIban, true
}

// HasCounterpartIban returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasCounterpartIban() bool {
	if o != nil && o.CounterpartIban != nil {
		return true
	}

	return false
}

// SetCounterpartIban gets a reference to the given string and assigns it to the CounterpartIban field.
func (o *ExtendedTransaction) SetCounterpartIban(v string) {
	o.CounterpartIban = &v
}

// GetCounterpartBlz returns the CounterpartBlz field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetCounterpartBlz() string {
	if o == nil || o.CounterpartBlz == nil {
		var ret string
		return ret
	}
	return *o.CounterpartBlz
}

// GetCounterpartBlzOk returns a tuple with the CounterpartBlz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetCounterpartBlzOk() (*string, bool) {
	if o == nil || o.CounterpartBlz == nil {
		return nil, false
	}
	return o.CounterpartBlz, true
}

// HasCounterpartBlz returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasCounterpartBlz() bool {
	if o != nil && o.CounterpartBlz != nil {
		return true
	}

	return false
}

// SetCounterpartBlz gets a reference to the given string and assigns it to the CounterpartBlz field.
func (o *ExtendedTransaction) SetCounterpartBlz(v string) {
	o.CounterpartBlz = &v
}

// GetCounterpartBic returns the CounterpartBic field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetCounterpartBic() string {
	if o == nil || o.CounterpartBic == nil {
		var ret string
		return ret
	}
	return *o.CounterpartBic
}

// GetCounterpartBicOk returns a tuple with the CounterpartBic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetCounterpartBicOk() (*string, bool) {
	if o == nil || o.CounterpartBic == nil {
		return nil, false
	}
	return o.CounterpartBic, true
}

// HasCounterpartBic returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasCounterpartBic() bool {
	if o != nil && o.CounterpartBic != nil {
		return true
	}

	return false
}

// SetCounterpartBic gets a reference to the given string and assigns it to the CounterpartBic field.
func (o *ExtendedTransaction) SetCounterpartBic(v string) {
	o.CounterpartBic = &v
}

// GetCounterpartAccountNumber returns the CounterpartAccountNumber field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetCounterpartAccountNumber() string {
	if o == nil || o.CounterpartAccountNumber == nil {
		var ret string
		return ret
	}
	return *o.CounterpartAccountNumber
}

// GetCounterpartAccountNumberOk returns a tuple with the CounterpartAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetCounterpartAccountNumberOk() (*string, bool) {
	if o == nil || o.CounterpartAccountNumber == nil {
		return nil, false
	}
	return o.CounterpartAccountNumber, true
}

// HasCounterpartAccountNumber returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasCounterpartAccountNumber() bool {
	if o != nil && o.CounterpartAccountNumber != nil {
		return true
	}

	return false
}

// SetCounterpartAccountNumber gets a reference to the given string and assigns it to the CounterpartAccountNumber field.
func (o *ExtendedTransaction) SetCounterpartAccountNumber(v string) {
	o.CounterpartAccountNumber = &v
}

// GetBookingDate returns the BookingDate field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetBookingDate() string {
	if o == nil || o.BookingDate == nil {
		var ret string
		return ret
	}
	return *o.BookingDate
}

// GetBookingDateOk returns a tuple with the BookingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetBookingDateOk() (*string, bool) {
	if o == nil || o.BookingDate == nil {
		return nil, false
	}
	return o.BookingDate, true
}

// HasBookingDate returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasBookingDate() bool {
	if o != nil && o.BookingDate != nil {
		return true
	}

	return false
}

// SetBookingDate gets a reference to the given string and assigns it to the BookingDate field.
func (o *ExtendedTransaction) SetBookingDate(v string) {
	o.BookingDate = &v
}

// GetValueDate returns the ValueDate field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetValueDate() string {
	if o == nil || o.ValueDate == nil {
		var ret string
		return ret
	}
	return *o.ValueDate
}

// GetValueDateOk returns a tuple with the ValueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetValueDateOk() (*string, bool) {
	if o == nil || o.ValueDate == nil {
		return nil, false
	}
	return o.ValueDate, true
}

// HasValueDate returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasValueDate() bool {
	if o != nil && o.ValueDate != nil {
		return true
	}

	return false
}

// SetValueDate gets a reference to the given string and assigns it to the ValueDate field.
func (o *ExtendedTransaction) SetValueDate(v string) {
	o.ValueDate = &v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetTypeId() int32 {
	if o == nil || o.TypeId == nil {
		var ret int32
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetTypeIdOk() (*int32, bool) {
	if o == nil || o.TypeId == nil {
		return nil, false
	}
	return o.TypeId, true
}

// HasTypeId returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasTypeId() bool {
	if o != nil && o.TypeId != nil {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given int32 and assigns it to the TypeId field.
func (o *ExtendedTransaction) SetTypeId(v int32) {
	o.TypeId = &v
}

// GetCounterpartMandateReference returns the CounterpartMandateReference field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetCounterpartMandateReference() string {
	if o == nil || o.CounterpartMandateReference == nil {
		var ret string
		return ret
	}
	return *o.CounterpartMandateReference
}

// GetCounterpartMandateReferenceOk returns a tuple with the CounterpartMandateReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetCounterpartMandateReferenceOk() (*string, bool) {
	if o == nil || o.CounterpartMandateReference == nil {
		return nil, false
	}
	return o.CounterpartMandateReference, true
}

// HasCounterpartMandateReference returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasCounterpartMandateReference() bool {
	if o != nil && o.CounterpartMandateReference != nil {
		return true
	}

	return false
}

// SetCounterpartMandateReference gets a reference to the given string and assigns it to the CounterpartMandateReference field.
func (o *ExtendedTransaction) SetCounterpartMandateReference(v string) {
	o.CounterpartMandateReference = &v
}

// GetCounterpartCreditorId returns the CounterpartCreditorId field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetCounterpartCreditorId() string {
	if o == nil || o.CounterpartCreditorId == nil {
		var ret string
		return ret
	}
	return *o.CounterpartCreditorId
}

// GetCounterpartCreditorIdOk returns a tuple with the CounterpartCreditorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetCounterpartCreditorIdOk() (*string, bool) {
	if o == nil || o.CounterpartCreditorId == nil {
		return nil, false
	}
	return o.CounterpartCreditorId, true
}

// HasCounterpartCreditorId returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasCounterpartCreditorId() bool {
	if o != nil && o.CounterpartCreditorId != nil {
		return true
	}

	return false
}

// SetCounterpartCreditorId gets a reference to the given string and assigns it to the CounterpartCreditorId field.
func (o *ExtendedTransaction) SetCounterpartCreditorId(v string) {
	o.CounterpartCreditorId = &v
}

// GetCounterpartCustomerReference returns the CounterpartCustomerReference field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetCounterpartCustomerReference() string {
	if o == nil || o.CounterpartCustomerReference == nil {
		var ret string
		return ret
	}
	return *o.CounterpartCustomerReference
}

// GetCounterpartCustomerReferenceOk returns a tuple with the CounterpartCustomerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetCounterpartCustomerReferenceOk() (*string, bool) {
	if o == nil || o.CounterpartCustomerReference == nil {
		return nil, false
	}
	return o.CounterpartCustomerReference, true
}

// HasCounterpartCustomerReference returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasCounterpartCustomerReference() bool {
	if o != nil && o.CounterpartCustomerReference != nil {
		return true
	}

	return false
}

// SetCounterpartCustomerReference gets a reference to the given string and assigns it to the CounterpartCustomerReference field.
func (o *ExtendedTransaction) SetCounterpartCustomerReference(v string) {
	o.CounterpartCustomerReference = &v
}

// GetCounterpartDebitorId returns the CounterpartDebitorId field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetCounterpartDebitorId() string {
	if o == nil || o.CounterpartDebitorId == nil {
		var ret string
		return ret
	}
	return *o.CounterpartDebitorId
}

// GetCounterpartDebitorIdOk returns a tuple with the CounterpartDebitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetCounterpartDebitorIdOk() (*string, bool) {
	if o == nil || o.CounterpartDebitorId == nil {
		return nil, false
	}
	return o.CounterpartDebitorId, true
}

// HasCounterpartDebitorId returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasCounterpartDebitorId() bool {
	if o != nil && o.CounterpartDebitorId != nil {
		return true
	}

	return false
}

// SetCounterpartDebitorId gets a reference to the given string and assigns it to the CounterpartDebitorId field.
func (o *ExtendedTransaction) SetCounterpartDebitorId(v string) {
	o.CounterpartDebitorId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ExtendedTransaction) SetType(v string) {
	o.Type = &v
}

// GetTypeCodeSwift returns the TypeCodeSwift field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetTypeCodeSwift() string {
	if o == nil || o.TypeCodeSwift == nil {
		var ret string
		return ret
	}
	return *o.TypeCodeSwift
}

// GetTypeCodeSwiftOk returns a tuple with the TypeCodeSwift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetTypeCodeSwiftOk() (*string, bool) {
	if o == nil || o.TypeCodeSwift == nil {
		return nil, false
	}
	return o.TypeCodeSwift, true
}

// HasTypeCodeSwift returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasTypeCodeSwift() bool {
	if o != nil && o.TypeCodeSwift != nil {
		return true
	}

	return false
}

// SetTypeCodeSwift gets a reference to the given string and assigns it to the TypeCodeSwift field.
func (o *ExtendedTransaction) SetTypeCodeSwift(v string) {
	o.TypeCodeSwift = &v
}

// GetSepaPurposeCode returns the SepaPurposeCode field value if set, zero value otherwise.
func (o *ExtendedTransaction) GetSepaPurposeCode() string {
	if o == nil || o.SepaPurposeCode == nil {
		var ret string
		return ret
	}
	return *o.SepaPurposeCode
}

// GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTransaction) GetSepaPurposeCodeOk() (*string, bool) {
	if o == nil || o.SepaPurposeCode == nil {
		return nil, false
	}
	return o.SepaPurposeCode, true
}

// HasSepaPurposeCode returns a boolean if a field has been set.
func (o *ExtendedTransaction) HasSepaPurposeCode() bool {
	if o != nil && o.SepaPurposeCode != nil {
		return true
	}

	return false
}

// SetSepaPurposeCode gets a reference to the given string and assigns it to the SepaPurposeCode field.
func (o *ExtendedTransaction) SetSepaPurposeCode(v string) {
	o.SepaPurposeCode = &v
}

func (o ExtendedTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.OriginalAmount != nil {
		toSerialize["originalAmount"] = o.OriginalAmount
	}
	if o.OriginalCurrency != nil {
		toSerialize["originalCurrency"] = o.OriginalCurrency
	}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	if o.Counterpart != nil {
		toSerialize["counterpart"] = o.Counterpart
	}
	if o.CounterpartIban != nil {
		toSerialize["counterpartIban"] = o.CounterpartIban
	}
	if o.CounterpartBlz != nil {
		toSerialize["counterpartBlz"] = o.CounterpartBlz
	}
	if o.CounterpartBic != nil {
		toSerialize["counterpartBic"] = o.CounterpartBic
	}
	if o.CounterpartAccountNumber != nil {
		toSerialize["counterpartAccountNumber"] = o.CounterpartAccountNumber
	}
	if o.BookingDate != nil {
		toSerialize["bookingDate"] = o.BookingDate
	}
	if o.ValueDate != nil {
		toSerialize["valueDate"] = o.ValueDate
	}
	if o.TypeId != nil {
		toSerialize["typeId"] = o.TypeId
	}
	if o.CounterpartMandateReference != nil {
		toSerialize["counterpartMandateReference"] = o.CounterpartMandateReference
	}
	if o.CounterpartCreditorId != nil {
		toSerialize["counterpartCreditorId"] = o.CounterpartCreditorId
	}
	if o.CounterpartCustomerReference != nil {
		toSerialize["counterpartCustomerReference"] = o.CounterpartCustomerReference
	}
	if o.CounterpartDebitorId != nil {
		toSerialize["counterpartDebitorId"] = o.CounterpartDebitorId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.TypeCodeSwift != nil {
		toSerialize["typeCodeSwift"] = o.TypeCodeSwift
	}
	if o.SepaPurposeCode != nil {
		toSerialize["sepaPurposeCode"] = o.SepaPurposeCode
	}
	return json.Marshal(toSerialize)
}

type NullableExtendedTransaction struct {
	value *ExtendedTransaction
	isSet bool
}

func (v NullableExtendedTransaction) Get() *ExtendedTransaction {
	return v.value
}

func (v *NullableExtendedTransaction) Set(val *ExtendedTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedTransaction(val *ExtendedTransaction) *NullableExtendedTransaction {
	return &NullableExtendedTransaction{value: val, isSet: true}
}

func (v NullableExtendedTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



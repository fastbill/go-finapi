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

// StandingOrder Container for a standing order's data
type StandingOrder struct {
	// Standing order identifier
	Id int64 `json:"id"`
	// Identifier of the account to which this standing order relates. This field is only set if it was specified upon creation of the standing order.
	AccountId NullableInt64 `json:"accountId"`
	// IBAN of the account to which this standing order relates. This field is only set if it was specified upon creation of the standing order.
	Iban NullableString `json:"iban"`
	// Amount of the standing order, as absolute value.
	Amount float64 `json:"amount"`
	// <strong>Type:</strong> Currency<br/> The currency code of the 'amount', in the ISO 4217 Alpha 3 format.
	Currency Currency `json:"currency"`
	// Start date of the standing order, in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	StartDate string `json:"startDate"`
	// Termination date of the standing order, in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time). If this field is not set, then the standing order has no termination date.
	EndDate NullableString `json:"endDate"`
	// <strong>Type:</strong> StandingOrderFrequency<br/> The frequency of the standing order.
	Frequency StandingOrderFrequency `json:"frequency"`
	// Scheduled day of execution. '31' depicts ultimo.
	DayOfExecution NullableInt32 `json:"dayOfExecution"`
	// Time of when finAPI submitted this standing order to the bank, in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time)
	RequestDate NullableString `json:"requestDate"`
	// Time of when the submission of this standing order was finalized, in the format ‘YYYY-MM-DD HH:MM:SS.SSS' (german time). Note: When the submission of a standing order is finalized, it does not necessarily mean that the bank accepted the standing order. Please refer to the standing order’s 'status' for its final status.
	RequestCompletionDate NullableString `json:"requestCompletionDate"`
	// <strong>Type:</strong> OrderInitiationStatus<br/> Current standing order status:<br/> &bull; OPEN: means that this standing order has been created in finAPI, but not yet submitted to the bank.<br/> &bull; PENDING: means that this standing order has been submitted to the bank,  but has not been confirmed yet.<br/> &bull; SUCCESSFUL: means that this standing order has been successfully initiated.<br/> &bull; NOT_SUCCESSFUL: means that this standing order could not be initiated successfully.<br/> &bull; DISCARDED: means that this standing order was discarded, either because another standing order was requested for the same account before this standing order was initiated and the bank does not support this, or because the user has aborted the initiation (when using finAPI's Web Form).
	Status OrderInitiationStatus `json:"status"`
	// The bank's response to the most recent request for this standing order. Note that this field may not always (or never) be set. Also, as long as the standing order has not reached its final status, this field can always change.
	BankMessage NullableString `json:"bankMessage"`
}

// NewStandingOrder instantiates a new StandingOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandingOrder(id int64, accountId NullableInt64, iban NullableString, amount float64, currency Currency, startDate string, endDate NullableString, frequency StandingOrderFrequency, dayOfExecution NullableInt32, requestDate NullableString, requestCompletionDate NullableString, status OrderInitiationStatus, bankMessage NullableString, ) *StandingOrder {
	this := StandingOrder{}
	this.Id = id
	this.AccountId = accountId
	this.Iban = iban
	this.Amount = amount
	this.Currency = currency
	this.StartDate = startDate
	this.EndDate = endDate
	this.Frequency = frequency
	this.DayOfExecution = dayOfExecution
	this.RequestDate = requestDate
	this.RequestCompletionDate = requestCompletionDate
	this.Status = status
	this.BankMessage = bankMessage
	return &this
}

// NewStandingOrderWithDefaults instantiates a new StandingOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandingOrderWithDefaults() *StandingOrder {
	this := StandingOrder{}
	return &this
}

// GetId returns the Id field value
func (o *StandingOrder) GetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StandingOrder) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StandingOrder) SetId(v int64) {
	o.Id = v
}

// GetAccountId returns the AccountId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *StandingOrder) GetAccountId() int64 {
	if o == nil || o.AccountId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandingOrder) GetAccountIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// SetAccountId sets field value
func (o *StandingOrder) SetAccountId(v int64) {
	o.AccountId.Set(&v)
}

// GetIban returns the Iban field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StandingOrder) GetIban() string {
	if o == nil || o.Iban.Get() == nil {
		var ret string
		return ret
	}

	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandingOrder) GetIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// SetIban sets field value
func (o *StandingOrder) SetIban(v string) {
	o.Iban.Set(&v)
}

// GetAmount returns the Amount field value
func (o *StandingOrder) GetAmount() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *StandingOrder) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *StandingOrder) SetAmount(v float64) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *StandingOrder) GetCurrency() Currency {
	if o == nil  {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *StandingOrder) GetCurrencyOk() (*Currency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *StandingOrder) SetCurrency(v Currency) {
	o.Currency = v
}

// GetStartDate returns the StartDate field value
func (o *StandingOrder) GetStartDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *StandingOrder) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *StandingOrder) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StandingOrder) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandingOrder) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// SetEndDate sets field value
func (o *StandingOrder) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// GetFrequency returns the Frequency field value
func (o *StandingOrder) GetFrequency() StandingOrderFrequency {
	if o == nil  {
		var ret StandingOrderFrequency
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *StandingOrder) GetFrequencyOk() (*StandingOrderFrequency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *StandingOrder) SetFrequency(v StandingOrderFrequency) {
	o.Frequency = v
}

// GetDayOfExecution returns the DayOfExecution field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *StandingOrder) GetDayOfExecution() int32 {
	if o == nil || o.DayOfExecution.Get() == nil {
		var ret int32
		return ret
	}

	return *o.DayOfExecution.Get()
}

// GetDayOfExecutionOk returns a tuple with the DayOfExecution field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandingOrder) GetDayOfExecutionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DayOfExecution.Get(), o.DayOfExecution.IsSet()
}

// SetDayOfExecution sets field value
func (o *StandingOrder) SetDayOfExecution(v int32) {
	o.DayOfExecution.Set(&v)
}

// GetRequestDate returns the RequestDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StandingOrder) GetRequestDate() string {
	if o == nil || o.RequestDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.RequestDate.Get()
}

// GetRequestDateOk returns a tuple with the RequestDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandingOrder) GetRequestDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestDate.Get(), o.RequestDate.IsSet()
}

// SetRequestDate sets field value
func (o *StandingOrder) SetRequestDate(v string) {
	o.RequestDate.Set(&v)
}

// GetRequestCompletionDate returns the RequestCompletionDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StandingOrder) GetRequestCompletionDate() string {
	if o == nil || o.RequestCompletionDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.RequestCompletionDate.Get()
}

// GetRequestCompletionDateOk returns a tuple with the RequestCompletionDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandingOrder) GetRequestCompletionDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestCompletionDate.Get(), o.RequestCompletionDate.IsSet()
}

// SetRequestCompletionDate sets field value
func (o *StandingOrder) SetRequestCompletionDate(v string) {
	o.RequestCompletionDate.Set(&v)
}

// GetStatus returns the Status field value
func (o *StandingOrder) GetStatus() OrderInitiationStatus {
	if o == nil  {
		var ret OrderInitiationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StandingOrder) GetStatusOk() (*OrderInitiationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StandingOrder) SetStatus(v OrderInitiationStatus) {
	o.Status = v
}

// GetBankMessage returns the BankMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StandingOrder) GetBankMessage() string {
	if o == nil || o.BankMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.BankMessage.Get()
}

// GetBankMessageOk returns a tuple with the BankMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandingOrder) GetBankMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankMessage.Get(), o.BankMessage.IsSet()
}

// SetBankMessage sets field value
func (o *StandingOrder) SetBankMessage(v string) {
	o.BankMessage.Set(&v)
}

func (o StandingOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["accountId"] = o.AccountId.Get()
	}
	if true {
		toSerialize["iban"] = o.Iban.Get()
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["startDate"] = o.StartDate
	}
	if true {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if true {
		toSerialize["frequency"] = o.Frequency
	}
	if true {
		toSerialize["dayOfExecution"] = o.DayOfExecution.Get()
	}
	if true {
		toSerialize["requestDate"] = o.RequestDate.Get()
	}
	if true {
		toSerialize["requestCompletionDate"] = o.RequestCompletionDate.Get()
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["bankMessage"] = o.BankMessage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableStandingOrder struct {
	value *StandingOrder
	isSet bool
}

func (v NullableStandingOrder) Get() *StandingOrder {
	return v.value
}

func (v *NullableStandingOrder) Set(val *StandingOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableStandingOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableStandingOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandingOrder(val *StandingOrder) *NullableStandingOrder {
	return &NullableStandingOrder{value: val, isSet: true}
}

func (v NullableStandingOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandingOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



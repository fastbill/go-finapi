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

// CreateStandingOrderParams Container for standing order creation parameters
type CreateStandingOrderParams struct {
	// Identifier of the account that should be used for the standing order. If you want to do a standalone standing order (i.e. for an account that is not imported in finAPI) leave this field unset and instead use the field 'iban'.
	AccountId *int64 `json:"accountId,omitempty"`
	// IBAN of the account that should be used for the standing order. Use this field only if you want to do a standalone standing order (i.e. for an account that is not imported in finAPI) otherwise, use the 'accountId' field and leave this field unset.
	Iban *string `json:"iban,omitempty"`
	// Name of the counterpart. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the counterpart name. Even if the name does not depict the actual registered account holder of the target account, the order might still be successful.
	CounterpartName string `json:"counterpartName"`
	// IBAN of the counterpart's account.
	CounterpartIban string `json:"counterpartIban"`
	// The amount of the standing order. Must be a positive decimal number with at most two decimal places.
	Amount float64 `json:"amount"`
	// <strong>Type:</strong> Currency<br/> The currency code of the 'amount'. To be provided in the ISO 4217 Alpha 3 format.
	Currency Currency `json:"currency"`
	// The purpose of the standing order.
	Purpose *string `json:"purpose,omitempty"`
	// SEPA purpose code, according to ISO 20022, external codes set.<br/>Please note that the SEPA purpose code may be ignored by some banks.
	SepaPurposeCode *string `json:"sepaPurposeCode,omitempty"`
	// End-To-End ID of the standing order
	EndToEndId *string `json:"endToEndId,omitempty"`
	// Start date of the standing order. Date must be in the future (at least tomorrow). To be provided in the format 'YYYY-MM-DD'.
	StartDate string `json:"startDate"`
	// Termination date of the standing order. If provided, it must be after the 'startDate'. If not provided, then the standing order will have no termination. To be provided in the format 'YYYY-MM-DD'.
	EndDate *string `json:"endDate,omitempty"`
	// <strong>Type:</strong> StandingOrderFrequency<br/> The frequency of the standing order
	Frequency StandingOrderFrequency `json:"frequency"`
	// Desired day of execution. If not provided, it will be derived from the 'startDate'. Use '31' for ultimo. Allowed values: 1-31.
	DayOfExecution *int32 `json:"dayOfExecution,omitempty"`
}

// NewCreateStandingOrderParams instantiates a new CreateStandingOrderParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStandingOrderParams(counterpartName string, counterpartIban string, amount float64, currency Currency, startDate string, frequency StandingOrderFrequency, ) *CreateStandingOrderParams {
	this := CreateStandingOrderParams{}
	this.CounterpartName = counterpartName
	this.CounterpartIban = counterpartIban
	this.Amount = amount
	this.Currency = currency
	this.StartDate = startDate
	this.Frequency = frequency
	return &this
}

// NewCreateStandingOrderParamsWithDefaults instantiates a new CreateStandingOrderParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStandingOrderParamsWithDefaults() *CreateStandingOrderParams {
	this := CreateStandingOrderParams{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *CreateStandingOrderParams) GetAccountId() int64 {
	if o == nil || o.AccountId == nil {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetAccountIdOk() (*int64, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *CreateStandingOrderParams) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *CreateStandingOrderParams) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *CreateStandingOrderParams) GetIban() string {
	if o == nil || o.Iban == nil {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetIbanOk() (*string, bool) {
	if o == nil || o.Iban == nil {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *CreateStandingOrderParams) HasIban() bool {
	if o != nil && o.Iban != nil {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *CreateStandingOrderParams) SetIban(v string) {
	o.Iban = &v
}

// GetCounterpartName returns the CounterpartName field value
func (o *CreateStandingOrderParams) GetCounterpartName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CounterpartName
}

// GetCounterpartNameOk returns a tuple with the CounterpartName field value
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetCounterpartNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CounterpartName, true
}

// SetCounterpartName sets field value
func (o *CreateStandingOrderParams) SetCounterpartName(v string) {
	o.CounterpartName = v
}

// GetCounterpartIban returns the CounterpartIban field value
func (o *CreateStandingOrderParams) GetCounterpartIban() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CounterpartIban
}

// GetCounterpartIbanOk returns a tuple with the CounterpartIban field value
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetCounterpartIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CounterpartIban, true
}

// SetCounterpartIban sets field value
func (o *CreateStandingOrderParams) SetCounterpartIban(v string) {
	o.CounterpartIban = v
}

// GetAmount returns the Amount field value
func (o *CreateStandingOrderParams) GetAmount() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateStandingOrderParams) SetAmount(v float64) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *CreateStandingOrderParams) GetCurrency() Currency {
	if o == nil  {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetCurrencyOk() (*Currency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreateStandingOrderParams) SetCurrency(v Currency) {
	o.Currency = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *CreateStandingOrderParams) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *CreateStandingOrderParams) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *CreateStandingOrderParams) SetPurpose(v string) {
	o.Purpose = &v
}

// GetSepaPurposeCode returns the SepaPurposeCode field value if set, zero value otherwise.
func (o *CreateStandingOrderParams) GetSepaPurposeCode() string {
	if o == nil || o.SepaPurposeCode == nil {
		var ret string
		return ret
	}
	return *o.SepaPurposeCode
}

// GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetSepaPurposeCodeOk() (*string, bool) {
	if o == nil || o.SepaPurposeCode == nil {
		return nil, false
	}
	return o.SepaPurposeCode, true
}

// HasSepaPurposeCode returns a boolean if a field has been set.
func (o *CreateStandingOrderParams) HasSepaPurposeCode() bool {
	if o != nil && o.SepaPurposeCode != nil {
		return true
	}

	return false
}

// SetSepaPurposeCode gets a reference to the given string and assigns it to the SepaPurposeCode field.
func (o *CreateStandingOrderParams) SetSepaPurposeCode(v string) {
	o.SepaPurposeCode = &v
}

// GetEndToEndId returns the EndToEndId field value if set, zero value otherwise.
func (o *CreateStandingOrderParams) GetEndToEndId() string {
	if o == nil || o.EndToEndId == nil {
		var ret string
		return ret
	}
	return *o.EndToEndId
}

// GetEndToEndIdOk returns a tuple with the EndToEndId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetEndToEndIdOk() (*string, bool) {
	if o == nil || o.EndToEndId == nil {
		return nil, false
	}
	return o.EndToEndId, true
}

// HasEndToEndId returns a boolean if a field has been set.
func (o *CreateStandingOrderParams) HasEndToEndId() bool {
	if o != nil && o.EndToEndId != nil {
		return true
	}

	return false
}

// SetEndToEndId gets a reference to the given string and assigns it to the EndToEndId field.
func (o *CreateStandingOrderParams) SetEndToEndId(v string) {
	o.EndToEndId = &v
}

// GetStartDate returns the StartDate field value
func (o *CreateStandingOrderParams) GetStartDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *CreateStandingOrderParams) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CreateStandingOrderParams) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CreateStandingOrderParams) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *CreateStandingOrderParams) SetEndDate(v string) {
	o.EndDate = &v
}

// GetFrequency returns the Frequency field value
func (o *CreateStandingOrderParams) GetFrequency() StandingOrderFrequency {
	if o == nil  {
		var ret StandingOrderFrequency
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetFrequencyOk() (*StandingOrderFrequency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *CreateStandingOrderParams) SetFrequency(v StandingOrderFrequency) {
	o.Frequency = v
}

// GetDayOfExecution returns the DayOfExecution field value if set, zero value otherwise.
func (o *CreateStandingOrderParams) GetDayOfExecution() int32 {
	if o == nil || o.DayOfExecution == nil {
		var ret int32
		return ret
	}
	return *o.DayOfExecution
}

// GetDayOfExecutionOk returns a tuple with the DayOfExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStandingOrderParams) GetDayOfExecutionOk() (*int32, bool) {
	if o == nil || o.DayOfExecution == nil {
		return nil, false
	}
	return o.DayOfExecution, true
}

// HasDayOfExecution returns a boolean if a field has been set.
func (o *CreateStandingOrderParams) HasDayOfExecution() bool {
	if o != nil && o.DayOfExecution != nil {
		return true
	}

	return false
}

// SetDayOfExecution gets a reference to the given int32 and assigns it to the DayOfExecution field.
func (o *CreateStandingOrderParams) SetDayOfExecution(v int32) {
	o.DayOfExecution = &v
}

func (o CreateStandingOrderParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.Iban != nil {
		toSerialize["iban"] = o.Iban
	}
	if true {
		toSerialize["counterpartName"] = o.CounterpartName
	}
	if true {
		toSerialize["counterpartIban"] = o.CounterpartIban
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	if o.SepaPurposeCode != nil {
		toSerialize["sepaPurposeCode"] = o.SepaPurposeCode
	}
	if o.EndToEndId != nil {
		toSerialize["endToEndId"] = o.EndToEndId
	}
	if true {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if true {
		toSerialize["frequency"] = o.Frequency
	}
	if o.DayOfExecution != nil {
		toSerialize["dayOfExecution"] = o.DayOfExecution
	}
	return json.Marshal(toSerialize)
}

type NullableCreateStandingOrderParams struct {
	value *CreateStandingOrderParams
	isSet bool
}

func (v NullableCreateStandingOrderParams) Get() *CreateStandingOrderParams {
	return v.value
}

func (v *NullableCreateStandingOrderParams) Set(val *CreateStandingOrderParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStandingOrderParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStandingOrderParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStandingOrderParams(val *CreateStandingOrderParams) *NullableCreateStandingOrderParams {
	return &NullableCreateStandingOrderParams{value: val, isSet: true}
}

func (v NullableCreateStandingOrderParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStandingOrderParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



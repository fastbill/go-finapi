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

// TrainCategorizationTransactionData Transaction data for categorization training
type TrainCategorizationTransactionData struct {
	// Identifier of account type.<br/><br/>1 = Checking,<br/>2 = Savings,<br/>3 = CreditCard,<br/>4 = Security,<br/>5 = Loan,<br/>6 = Pocket (DEPRECATED; will not be returned for any account unless this type has explicitly been set via PATCH),<br/>7 = Membership,<br/>8 = Bausparen<br/><br/>
	AccountTypeId int64 `json:"accountTypeId"`
	// Amount
	Amount float64 `json:"amount"`
	// Purpose. Any symbols are allowed. Maximum length is 2000. Default value: null.
	Purpose *string `json:"purpose,omitempty"`
	// Counterpart. Any symbols are allowed. Maximum length is 80. Default value: null.
	Counterpart *string `json:"counterpart,omitempty"`
	// Counterpart IBAN. Default value: null.
	CounterpartIban *string `json:"counterpartIban,omitempty"`
	// Counterpart account number. Default value: null.
	CounterpartAccountNumber *string `json:"counterpartAccountNumber,omitempty"`
	// Counterpart BLZ. Default value: null.
	CounterpartBlz *string `json:"counterpartBlz,omitempty"`
	// Counterpart BIC. Default value: null.
	CounterpartBic *string `json:"counterpartBic,omitempty"`
	// Merchant category code (for credit card transactions only). Default value: null. NOTE: This field is currently not regarded.
	McCode *string `json:"mcCode,omitempty"`
}

// NewTrainCategorizationTransactionData instantiates a new TrainCategorizationTransactionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrainCategorizationTransactionData(accountTypeId int64, amount float64, ) *TrainCategorizationTransactionData {
	this := TrainCategorizationTransactionData{}
	this.AccountTypeId = accountTypeId
	this.Amount = amount
	return &this
}

// NewTrainCategorizationTransactionDataWithDefaults instantiates a new TrainCategorizationTransactionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrainCategorizationTransactionDataWithDefaults() *TrainCategorizationTransactionData {
	this := TrainCategorizationTransactionData{}
	return &this
}

// GetAccountTypeId returns the AccountTypeId field value
func (o *TrainCategorizationTransactionData) GetAccountTypeId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.AccountTypeId
}

// GetAccountTypeIdOk returns a tuple with the AccountTypeId field value
// and a boolean to check if the value has been set.
func (o *TrainCategorizationTransactionData) GetAccountTypeIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountTypeId, true
}

// SetAccountTypeId sets field value
func (o *TrainCategorizationTransactionData) SetAccountTypeId(v int64) {
	o.AccountTypeId = v
}

// GetAmount returns the Amount field value
func (o *TrainCategorizationTransactionData) GetAmount() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TrainCategorizationTransactionData) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TrainCategorizationTransactionData) SetAmount(v float64) {
	o.Amount = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *TrainCategorizationTransactionData) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainCategorizationTransactionData) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *TrainCategorizationTransactionData) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *TrainCategorizationTransactionData) SetPurpose(v string) {
	o.Purpose = &v
}

// GetCounterpart returns the Counterpart field value if set, zero value otherwise.
func (o *TrainCategorizationTransactionData) GetCounterpart() string {
	if o == nil || o.Counterpart == nil {
		var ret string
		return ret
	}
	return *o.Counterpart
}

// GetCounterpartOk returns a tuple with the Counterpart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainCategorizationTransactionData) GetCounterpartOk() (*string, bool) {
	if o == nil || o.Counterpart == nil {
		return nil, false
	}
	return o.Counterpart, true
}

// HasCounterpart returns a boolean if a field has been set.
func (o *TrainCategorizationTransactionData) HasCounterpart() bool {
	if o != nil && o.Counterpart != nil {
		return true
	}

	return false
}

// SetCounterpart gets a reference to the given string and assigns it to the Counterpart field.
func (o *TrainCategorizationTransactionData) SetCounterpart(v string) {
	o.Counterpart = &v
}

// GetCounterpartIban returns the CounterpartIban field value if set, zero value otherwise.
func (o *TrainCategorizationTransactionData) GetCounterpartIban() string {
	if o == nil || o.CounterpartIban == nil {
		var ret string
		return ret
	}
	return *o.CounterpartIban
}

// GetCounterpartIbanOk returns a tuple with the CounterpartIban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainCategorizationTransactionData) GetCounterpartIbanOk() (*string, bool) {
	if o == nil || o.CounterpartIban == nil {
		return nil, false
	}
	return o.CounterpartIban, true
}

// HasCounterpartIban returns a boolean if a field has been set.
func (o *TrainCategorizationTransactionData) HasCounterpartIban() bool {
	if o != nil && o.CounterpartIban != nil {
		return true
	}

	return false
}

// SetCounterpartIban gets a reference to the given string and assigns it to the CounterpartIban field.
func (o *TrainCategorizationTransactionData) SetCounterpartIban(v string) {
	o.CounterpartIban = &v
}

// GetCounterpartAccountNumber returns the CounterpartAccountNumber field value if set, zero value otherwise.
func (o *TrainCategorizationTransactionData) GetCounterpartAccountNumber() string {
	if o == nil || o.CounterpartAccountNumber == nil {
		var ret string
		return ret
	}
	return *o.CounterpartAccountNumber
}

// GetCounterpartAccountNumberOk returns a tuple with the CounterpartAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainCategorizationTransactionData) GetCounterpartAccountNumberOk() (*string, bool) {
	if o == nil || o.CounterpartAccountNumber == nil {
		return nil, false
	}
	return o.CounterpartAccountNumber, true
}

// HasCounterpartAccountNumber returns a boolean if a field has been set.
func (o *TrainCategorizationTransactionData) HasCounterpartAccountNumber() bool {
	if o != nil && o.CounterpartAccountNumber != nil {
		return true
	}

	return false
}

// SetCounterpartAccountNumber gets a reference to the given string and assigns it to the CounterpartAccountNumber field.
func (o *TrainCategorizationTransactionData) SetCounterpartAccountNumber(v string) {
	o.CounterpartAccountNumber = &v
}

// GetCounterpartBlz returns the CounterpartBlz field value if set, zero value otherwise.
func (o *TrainCategorizationTransactionData) GetCounterpartBlz() string {
	if o == nil || o.CounterpartBlz == nil {
		var ret string
		return ret
	}
	return *o.CounterpartBlz
}

// GetCounterpartBlzOk returns a tuple with the CounterpartBlz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainCategorizationTransactionData) GetCounterpartBlzOk() (*string, bool) {
	if o == nil || o.CounterpartBlz == nil {
		return nil, false
	}
	return o.CounterpartBlz, true
}

// HasCounterpartBlz returns a boolean if a field has been set.
func (o *TrainCategorizationTransactionData) HasCounterpartBlz() bool {
	if o != nil && o.CounterpartBlz != nil {
		return true
	}

	return false
}

// SetCounterpartBlz gets a reference to the given string and assigns it to the CounterpartBlz field.
func (o *TrainCategorizationTransactionData) SetCounterpartBlz(v string) {
	o.CounterpartBlz = &v
}

// GetCounterpartBic returns the CounterpartBic field value if set, zero value otherwise.
func (o *TrainCategorizationTransactionData) GetCounterpartBic() string {
	if o == nil || o.CounterpartBic == nil {
		var ret string
		return ret
	}
	return *o.CounterpartBic
}

// GetCounterpartBicOk returns a tuple with the CounterpartBic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainCategorizationTransactionData) GetCounterpartBicOk() (*string, bool) {
	if o == nil || o.CounterpartBic == nil {
		return nil, false
	}
	return o.CounterpartBic, true
}

// HasCounterpartBic returns a boolean if a field has been set.
func (o *TrainCategorizationTransactionData) HasCounterpartBic() bool {
	if o != nil && o.CounterpartBic != nil {
		return true
	}

	return false
}

// SetCounterpartBic gets a reference to the given string and assigns it to the CounterpartBic field.
func (o *TrainCategorizationTransactionData) SetCounterpartBic(v string) {
	o.CounterpartBic = &v
}

// GetMcCode returns the McCode field value if set, zero value otherwise.
func (o *TrainCategorizationTransactionData) GetMcCode() string {
	if o == nil || o.McCode == nil {
		var ret string
		return ret
	}
	return *o.McCode
}

// GetMcCodeOk returns a tuple with the McCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrainCategorizationTransactionData) GetMcCodeOk() (*string, bool) {
	if o == nil || o.McCode == nil {
		return nil, false
	}
	return o.McCode, true
}

// HasMcCode returns a boolean if a field has been set.
func (o *TrainCategorizationTransactionData) HasMcCode() bool {
	if o != nil && o.McCode != nil {
		return true
	}

	return false
}

// SetMcCode gets a reference to the given string and assigns it to the McCode field.
func (o *TrainCategorizationTransactionData) SetMcCode(v string) {
	o.McCode = &v
}

func (o TrainCategorizationTransactionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountTypeId"] = o.AccountTypeId
	}
	if true {
		toSerialize["amount"] = o.Amount
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
	if o.CounterpartAccountNumber != nil {
		toSerialize["counterpartAccountNumber"] = o.CounterpartAccountNumber
	}
	if o.CounterpartBlz != nil {
		toSerialize["counterpartBlz"] = o.CounterpartBlz
	}
	if o.CounterpartBic != nil {
		toSerialize["counterpartBic"] = o.CounterpartBic
	}
	if o.McCode != nil {
		toSerialize["mcCode"] = o.McCode
	}
	return json.Marshal(toSerialize)
}

type NullableTrainCategorizationTransactionData struct {
	value *TrainCategorizationTransactionData
	isSet bool
}

func (v NullableTrainCategorizationTransactionData) Get() *TrainCategorizationTransactionData {
	return v.value
}

func (v *NullableTrainCategorizationTransactionData) Set(val *TrainCategorizationTransactionData) {
	v.value = val
	v.isSet = true
}

func (v NullableTrainCategorizationTransactionData) IsSet() bool {
	return v.isSet
}

func (v *NullableTrainCategorizationTransactionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrainCategorizationTransactionData(val *TrainCategorizationTransactionData) *NullableTrainCategorizationTransactionData {
	return &NullableTrainCategorizationTransactionData{value: val, isSet: true}
}

func (v NullableTrainCategorizationTransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrainCategorizationTransactionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



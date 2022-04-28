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

// Payment Container for a payment's data
type Payment struct {
	// Payment identifier
	Id int64 `json:"id"`
	// Identifier of the account to which this payment relates. This field is only set if it was specified upon creation of the payment.
	AccountId NullableInt64 `json:"accountId"`
	// IBAN of the account to which this payment relates. This field is only set if it was specified upon creation of the payment.
	Iban NullableString `json:"iban"`
	// <strong>Type:</strong> PaymentType<br/> Payment type
	Type PaymentType `json:"type"`
	// Total money amount of the payment order(s), as absolute value
	Amount float64 `json:"amount"`
	// Total count of orders included in this payment
	OrderCount int32 `json:"orderCount"`
	// <strong>Type:</strong> OrderInitiationStatus<br/> Current payment status:<br/> &bull; OPEN: means that this payment has been created in finAPI, but not yet submitted to the bank.<br/> &bull; PENDING: means that this payment has been requested at the bank,  but has not been confirmed yet.<br/> &bull; SUCCESSFUL: means that this payment has been successfully initiated.<br/> &bull; NOT_SUCCESSFUL: means that this payment could not be initiated successfully.<br/> &bull; DISCARDED: means that this payment was discarded, either because another payment was requested for the same account before this payment was initiated and the bank does not support this, or because the user has aborted the initiation (when using finAPI's Web Form).
	Status OrderInitiationStatus `json:"status"`
	// The bank's response to the most recent request for this payment. Possible requests are: Initial submission of the payment, execution request or subsequent status checks. Note that this field may not always (or never) be set. Also, as long as the payment has not reached its final status, this field can always change.
	BankMessage NullableString `json:"bankMessage"`
	// Time of when finAPI submitted this payment to the bank, in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time)
	RequestDate NullableString `json:"requestDate"`
	// Time of when the execution of this payment has completed, in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).<br/><br/>Note:<br/>&bull; When the execution of a payment has completed, it does not necessarily mean that the payment was successful. Please refer to the payment 'status' for its final status.<br/>&bull; The execution date may deviate from the date when the bank will actually book the payment (for example if the 'instructedExecutionDate' is in the future).
	ExecutionDate NullableString `json:"executionDate"`
	// The date that was specified as 'executionDate' upon creation of the payment, in the format 'YYYY-MM-DD'. This field may not be set if no 'executionDate' was specified upon payment creation.
	InstructedExecutionDate NullableString `json:"instructedExecutionDate"`
}

// NewPayment instantiates a new Payment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayment(id int64, accountId NullableInt64, iban NullableString, type_ PaymentType, amount float64, orderCount int32, status OrderInitiationStatus, bankMessage NullableString, requestDate NullableString, executionDate NullableString, instructedExecutionDate NullableString, ) *Payment {
	this := Payment{}
	this.Id = id
	this.AccountId = accountId
	this.Iban = iban
	this.Type = type_
	this.Amount = amount
	this.OrderCount = orderCount
	this.Status = status
	this.BankMessage = bankMessage
	this.RequestDate = requestDate
	this.ExecutionDate = executionDate
	this.InstructedExecutionDate = instructedExecutionDate
	return &this
}

// NewPaymentWithDefaults instantiates a new Payment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentWithDefaults() *Payment {
	this := Payment{}
	return &this
}

// GetId returns the Id field value
func (o *Payment) GetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Payment) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Payment) SetId(v int64) {
	o.Id = v
}

// GetAccountId returns the AccountId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *Payment) GetAccountId() int64 {
	if o == nil || o.AccountId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetAccountIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// SetAccountId sets field value
func (o *Payment) SetAccountId(v int64) {
	o.AccountId.Set(&v)
}

// GetIban returns the Iban field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Payment) GetIban() string {
	if o == nil || o.Iban.Get() == nil {
		var ret string
		return ret
	}

	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// SetIban sets field value
func (o *Payment) SetIban(v string) {
	o.Iban.Set(&v)
}

// GetType returns the Type field value
func (o *Payment) GetType() PaymentType {
	if o == nil  {
		var ret PaymentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Payment) GetTypeOk() (*PaymentType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Payment) SetType(v PaymentType) {
	o.Type = v
}

// GetAmount returns the Amount field value
func (o *Payment) GetAmount() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Payment) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Payment) SetAmount(v float64) {
	o.Amount = v
}

// GetOrderCount returns the OrderCount field value
func (o *Payment) GetOrderCount() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.OrderCount
}

// GetOrderCountOk returns a tuple with the OrderCount field value
// and a boolean to check if the value has been set.
func (o *Payment) GetOrderCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrderCount, true
}

// SetOrderCount sets field value
func (o *Payment) SetOrderCount(v int32) {
	o.OrderCount = v
}

// GetStatus returns the Status field value
func (o *Payment) GetStatus() OrderInitiationStatus {
	if o == nil  {
		var ret OrderInitiationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Payment) GetStatusOk() (*OrderInitiationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Payment) SetStatus(v OrderInitiationStatus) {
	o.Status = v
}

// GetBankMessage returns the BankMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Payment) GetBankMessage() string {
	if o == nil || o.BankMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.BankMessage.Get()
}

// GetBankMessageOk returns a tuple with the BankMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetBankMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankMessage.Get(), o.BankMessage.IsSet()
}

// SetBankMessage sets field value
func (o *Payment) SetBankMessage(v string) {
	o.BankMessage.Set(&v)
}

// GetRequestDate returns the RequestDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Payment) GetRequestDate() string {
	if o == nil || o.RequestDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.RequestDate.Get()
}

// GetRequestDateOk returns a tuple with the RequestDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetRequestDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestDate.Get(), o.RequestDate.IsSet()
}

// SetRequestDate sets field value
func (o *Payment) SetRequestDate(v string) {
	o.RequestDate.Set(&v)
}

// GetExecutionDate returns the ExecutionDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Payment) GetExecutionDate() string {
	if o == nil || o.ExecutionDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExecutionDate.Get()
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetExecutionDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExecutionDate.Get(), o.ExecutionDate.IsSet()
}

// SetExecutionDate sets field value
func (o *Payment) SetExecutionDate(v string) {
	o.ExecutionDate.Set(&v)
}

// GetInstructedExecutionDate returns the InstructedExecutionDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Payment) GetInstructedExecutionDate() string {
	if o == nil || o.InstructedExecutionDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.InstructedExecutionDate.Get()
}

// GetInstructedExecutionDateOk returns a tuple with the InstructedExecutionDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetInstructedExecutionDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstructedExecutionDate.Get(), o.InstructedExecutionDate.IsSet()
}

// SetInstructedExecutionDate sets field value
func (o *Payment) SetInstructedExecutionDate(v string) {
	o.InstructedExecutionDate.Set(&v)
}

func (o Payment) MarshalJSON() ([]byte, error) {
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
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["orderCount"] = o.OrderCount
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["bankMessage"] = o.BankMessage.Get()
	}
	if true {
		toSerialize["requestDate"] = o.RequestDate.Get()
	}
	if true {
		toSerialize["executionDate"] = o.ExecutionDate.Get()
	}
	if true {
		toSerialize["instructedExecutionDate"] = o.InstructedExecutionDate.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePayment struct {
	value *Payment
	isSet bool
}

func (v NullablePayment) Get() *Payment {
	return v.value
}

func (v *NullablePayment) Set(val *Payment) {
	v.value = val
	v.isSet = true
}

func (v NullablePayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayment(val *Payment) *NullablePayment {
	return &NullablePayment{value: val, isSet: true}
}

func (v NullablePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



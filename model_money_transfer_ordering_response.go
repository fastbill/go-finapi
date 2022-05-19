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

// MoneyTransferOrderingResponse Bank server's response to a money transfer order request
type MoneyTransferOrderingResponse struct {
	// Technical message from the bank server, confirming the success of the request. Typically, you would not want to present this message to the user. Note that this field may not be set. However if it is not set, it does not necessarily mean that there was an error in processing the request.
	SuccessMessage NullableString `json:"successMessage"`
	// In some cases, a bank server may accept the requested order, but return a warn message. This message may be of technical nature, but could also be of interest to the user.
	WarnMessage NullableString `json:"warnMessage"`
	// Payment identifier. Can be used to retrieve the status of the payment (see 'Get payments' service).
	PaymentId int64 `json:"paymentId"`
	// Message from the bank server containing information or instructions on how to retrieve the TAN that is needed to execute the requested order. This message should be presented to the user. Note that some bank servers may limit the message to just the most crucial information, e.g. the message may contain just a single number that depicts the target TAN number on a user's TAN list. You may want to parse the challenge message for such cases and extend it with more detailed information before showing it to the user.
	ChallengeMessage NullableString `json:"challengeMessage"`
	// Suggestion from the bank server on how you can label your input field where the user must enter his TAN. A typical value that many bank servers give is 'TAN-Nummer'.
	AnswerFieldLabel NullableString `json:"answerFieldLabel"`
	// In case that the bank server has instructed the user to look up a TAN from a TAN list, this field may contain the identification number of the TAN list. However in most cases, this field is only set (i.e. not null) when the user has multiple active TAN lists.
	TanListNumber NullableString `json:"tanListNumber"`
	// In case that the bank server has instructed the user to scan a flicker code, then this field will contain the raw data for the flicker animation as a BASE-64 string. Otherwise, this field will be not set (i.e. null). For more information on how to process the flicker code data, please refer to <a href='https://documentation.finapi.io/access/Flicker-Code-Template.2807824454.html' target='_blank'>this</a> article.
	OpticalData NullableString `json:"opticalData"`
	// In case that the 'photoTanData' field is set (i.e. not null), this field contains the MIME type to use for interpreting the photo data (e.g.: 'image/png')
	PhotoTanMimeType NullableString `json:"photoTanMimeType"`
	// In case that the bank server has instructed the user to scan a photo (or more generally speaking, any kind of QR-code-like data), then this field will contain the raw data of the photo as a BASE-64 string. Otherwise, this field will be not set (i.e. null).
	PhotoTanData NullableString `json:"photoTanData"`
}

// NewMoneyTransferOrderingResponse instantiates a new MoneyTransferOrderingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyTransferOrderingResponse(successMessage NullableString, warnMessage NullableString, paymentId int64, challengeMessage NullableString, answerFieldLabel NullableString, tanListNumber NullableString, opticalData NullableString, photoTanMimeType NullableString, photoTanData NullableString, ) *MoneyTransferOrderingResponse {
	this := MoneyTransferOrderingResponse{}
	this.SuccessMessage = successMessage
	this.WarnMessage = warnMessage
	this.PaymentId = paymentId
	this.ChallengeMessage = challengeMessage
	this.AnswerFieldLabel = answerFieldLabel
	this.TanListNumber = tanListNumber
	this.OpticalData = opticalData
	this.PhotoTanMimeType = photoTanMimeType
	this.PhotoTanData = photoTanData
	return &this
}

// NewMoneyTransferOrderingResponseWithDefaults instantiates a new MoneyTransferOrderingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyTransferOrderingResponseWithDefaults() *MoneyTransferOrderingResponse {
	this := MoneyTransferOrderingResponse{}
	return &this
}

// GetSuccessMessage returns the SuccessMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MoneyTransferOrderingResponse) GetSuccessMessage() string {
	if o == nil || o.SuccessMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.SuccessMessage.Get()
}

// GetSuccessMessageOk returns a tuple with the SuccessMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoneyTransferOrderingResponse) GetSuccessMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SuccessMessage.Get(), o.SuccessMessage.IsSet()
}

// SetSuccessMessage sets field value
func (o *MoneyTransferOrderingResponse) SetSuccessMessage(v string) {
	o.SuccessMessage.Set(&v)
}

// GetWarnMessage returns the WarnMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MoneyTransferOrderingResponse) GetWarnMessage() string {
	if o == nil || o.WarnMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.WarnMessage.Get()
}

// GetWarnMessageOk returns a tuple with the WarnMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoneyTransferOrderingResponse) GetWarnMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WarnMessage.Get(), o.WarnMessage.IsSet()
}

// SetWarnMessage sets field value
func (o *MoneyTransferOrderingResponse) SetWarnMessage(v string) {
	o.WarnMessage.Set(&v)
}

// GetPaymentId returns the PaymentId field value
func (o *MoneyTransferOrderingResponse) GetPaymentId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *MoneyTransferOrderingResponse) GetPaymentIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *MoneyTransferOrderingResponse) SetPaymentId(v int64) {
	o.PaymentId = v
}

// GetChallengeMessage returns the ChallengeMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MoneyTransferOrderingResponse) GetChallengeMessage() string {
	if o == nil || o.ChallengeMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.ChallengeMessage.Get()
}

// GetChallengeMessageOk returns a tuple with the ChallengeMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoneyTransferOrderingResponse) GetChallengeMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ChallengeMessage.Get(), o.ChallengeMessage.IsSet()
}

// SetChallengeMessage sets field value
func (o *MoneyTransferOrderingResponse) SetChallengeMessage(v string) {
	o.ChallengeMessage.Set(&v)
}

// GetAnswerFieldLabel returns the AnswerFieldLabel field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MoneyTransferOrderingResponse) GetAnswerFieldLabel() string {
	if o == nil || o.AnswerFieldLabel.Get() == nil {
		var ret string
		return ret
	}

	return *o.AnswerFieldLabel.Get()
}

// GetAnswerFieldLabelOk returns a tuple with the AnswerFieldLabel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoneyTransferOrderingResponse) GetAnswerFieldLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AnswerFieldLabel.Get(), o.AnswerFieldLabel.IsSet()
}

// SetAnswerFieldLabel sets field value
func (o *MoneyTransferOrderingResponse) SetAnswerFieldLabel(v string) {
	o.AnswerFieldLabel.Set(&v)
}

// GetTanListNumber returns the TanListNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MoneyTransferOrderingResponse) GetTanListNumber() string {
	if o == nil || o.TanListNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.TanListNumber.Get()
}

// GetTanListNumberOk returns a tuple with the TanListNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoneyTransferOrderingResponse) GetTanListNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TanListNumber.Get(), o.TanListNumber.IsSet()
}

// SetTanListNumber sets field value
func (o *MoneyTransferOrderingResponse) SetTanListNumber(v string) {
	o.TanListNumber.Set(&v)
}

// GetOpticalData returns the OpticalData field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MoneyTransferOrderingResponse) GetOpticalData() string {
	if o == nil || o.OpticalData.Get() == nil {
		var ret string
		return ret
	}

	return *o.OpticalData.Get()
}

// GetOpticalDataOk returns a tuple with the OpticalData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoneyTransferOrderingResponse) GetOpticalDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OpticalData.Get(), o.OpticalData.IsSet()
}

// SetOpticalData sets field value
func (o *MoneyTransferOrderingResponse) SetOpticalData(v string) {
	o.OpticalData.Set(&v)
}

// GetPhotoTanMimeType returns the PhotoTanMimeType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MoneyTransferOrderingResponse) GetPhotoTanMimeType() string {
	if o == nil || o.PhotoTanMimeType.Get() == nil {
		var ret string
		return ret
	}

	return *o.PhotoTanMimeType.Get()
}

// GetPhotoTanMimeTypeOk returns a tuple with the PhotoTanMimeType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoneyTransferOrderingResponse) GetPhotoTanMimeTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhotoTanMimeType.Get(), o.PhotoTanMimeType.IsSet()
}

// SetPhotoTanMimeType sets field value
func (o *MoneyTransferOrderingResponse) SetPhotoTanMimeType(v string) {
	o.PhotoTanMimeType.Set(&v)
}

// GetPhotoTanData returns the PhotoTanData field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MoneyTransferOrderingResponse) GetPhotoTanData() string {
	if o == nil || o.PhotoTanData.Get() == nil {
		var ret string
		return ret
	}

	return *o.PhotoTanData.Get()
}

// GetPhotoTanDataOk returns a tuple with the PhotoTanData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoneyTransferOrderingResponse) GetPhotoTanDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhotoTanData.Get(), o.PhotoTanData.IsSet()
}

// SetPhotoTanData sets field value
func (o *MoneyTransferOrderingResponse) SetPhotoTanData(v string) {
	o.PhotoTanData.Set(&v)
}

func (o MoneyTransferOrderingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["successMessage"] = o.SuccessMessage.Get()
	}
	if true {
		toSerialize["warnMessage"] = o.WarnMessage.Get()
	}
	if true {
		toSerialize["paymentId"] = o.PaymentId
	}
	if true {
		toSerialize["challengeMessage"] = o.ChallengeMessage.Get()
	}
	if true {
		toSerialize["answerFieldLabel"] = o.AnswerFieldLabel.Get()
	}
	if true {
		toSerialize["tanListNumber"] = o.TanListNumber.Get()
	}
	if true {
		toSerialize["opticalData"] = o.OpticalData.Get()
	}
	if true {
		toSerialize["photoTanMimeType"] = o.PhotoTanMimeType.Get()
	}
	if true {
		toSerialize["photoTanData"] = o.PhotoTanData.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMoneyTransferOrderingResponse struct {
	value *MoneyTransferOrderingResponse
	isSet bool
}

func (v NullableMoneyTransferOrderingResponse) Get() *MoneyTransferOrderingResponse {
	return v.value
}

func (v *NullableMoneyTransferOrderingResponse) Set(val *MoneyTransferOrderingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyTransferOrderingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyTransferOrderingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyTransferOrderingResponse(val *MoneyTransferOrderingResponse) *NullableMoneyTransferOrderingResponse {
	return &NullableMoneyTransferOrderingResponse{value: val, isSet: true}
}

func (v NullableMoneyTransferOrderingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyTransferOrderingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



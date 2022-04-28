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

// WebForm Container for a Web Form's data
type WebForm struct {
	// Web Form identifier, as returned in the 451 response of the REST service call that initiated the Web Form flow.
	Id int64 `json:"id"`
	// Token for the finAPI Web Form page, as contained in the 451 response of the REST service call that initiated the Web Form flow (in the 'Location' header).
	Token string `json:"token"`
	// <strong>Type:</strong> WebFormStatus<br/> Status of a Web Form. Possible values are:<br/>&bull; NOT_YET_OPENED - the Web Form URL was not yet called;<br/>&bull; IN_PROGRESS - the Web Form has been opened but not yet submitted by the user;<br/>&bull; COMPLETED - the user has opened and submitted the Web Form;<br/>&bull; ABORTED - the user has opened but then aborted the Web Form, or the Web Form was aborted by the finAPI system because it has expired (this is the case when a Web Form is opened and then not submitted within 10 minutes)
	Status WebFormStatus `json:"status"`
	// HTTP response code of the REST service call that initiated the Web Form flow. This field can be queried as soon as the status becomes COMPLETED or ABORTED. Note that it is still not guaranteed in this case that the field has a value, i.e. it might be null.
	ServiceResponseCode NullableInt32 `json:"serviceResponseCode"`
	// HTTP response body of the REST service call that initiated the Web Form flow. This field can be queried as soon as the status becomes COMPLETED or ABORTED. Note that it is still not guaranteed in this case that the field has a value, i.e. it might be null.
	ServiceResponseBody NullableString `json:"serviceResponseBody"`
}

// NewWebForm instantiates a new WebForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebForm(id int64, token string, status WebFormStatus, serviceResponseCode NullableInt32, serviceResponseBody NullableString, ) *WebForm {
	this := WebForm{}
	this.Id = id
	this.Token = token
	this.Status = status
	this.ServiceResponseCode = serviceResponseCode
	this.ServiceResponseBody = serviceResponseBody
	return &this
}

// NewWebFormWithDefaults instantiates a new WebForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebFormWithDefaults() *WebForm {
	this := WebForm{}
	return &this
}

// GetId returns the Id field value
func (o *WebForm) GetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebForm) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebForm) SetId(v int64) {
	o.Id = v
}

// GetToken returns the Token field value
func (o *WebForm) GetToken() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *WebForm) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *WebForm) SetToken(v string) {
	o.Token = v
}

// GetStatus returns the Status field value
func (o *WebForm) GetStatus() WebFormStatus {
	if o == nil  {
		var ret WebFormStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WebForm) GetStatusOk() (*WebFormStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WebForm) SetStatus(v WebFormStatus) {
	o.Status = v
}

// GetServiceResponseCode returns the ServiceResponseCode field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *WebForm) GetServiceResponseCode() int32 {
	if o == nil || o.ServiceResponseCode.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ServiceResponseCode.Get()
}

// GetServiceResponseCodeOk returns a tuple with the ServiceResponseCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebForm) GetServiceResponseCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceResponseCode.Get(), o.ServiceResponseCode.IsSet()
}

// SetServiceResponseCode sets field value
func (o *WebForm) SetServiceResponseCode(v int32) {
	o.ServiceResponseCode.Set(&v)
}

// GetServiceResponseBody returns the ServiceResponseBody field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebForm) GetServiceResponseBody() string {
	if o == nil || o.ServiceResponseBody.Get() == nil {
		var ret string
		return ret
	}

	return *o.ServiceResponseBody.Get()
}

// GetServiceResponseBodyOk returns a tuple with the ServiceResponseBody field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebForm) GetServiceResponseBodyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceResponseBody.Get(), o.ServiceResponseBody.IsSet()
}

// SetServiceResponseBody sets field value
func (o *WebForm) SetServiceResponseBody(v string) {
	o.ServiceResponseBody.Set(&v)
}

func (o WebForm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["serviceResponseCode"] = o.ServiceResponseCode.Get()
	}
	if true {
		toSerialize["serviceResponseBody"] = o.ServiceResponseBody.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableWebForm struct {
	value *WebForm
	isSet bool
}

func (v NullableWebForm) Get() *WebForm {
	return v.value
}

func (v *NullableWebForm) Set(val *WebForm) {
	v.value = val
	v.isSet = true
}

func (v NullableWebForm) IsSet() bool {
	return v.isSet
}

func (v *NullableWebForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebForm(val *WebForm) *NullableWebForm {
	return &NullableWebForm{value: val, isSet: true}
}

func (v NullableWebForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



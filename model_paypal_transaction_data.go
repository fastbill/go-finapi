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

// PaypalTransactionData Additional, PayPal-specific transaction data. This field is only set for transactions that belong to an account of the PayPal bank.
type PaypalTransactionData struct {
	// Invoice Number.
	InvoiceNumber NullableString `json:"invoiceNumber"`
	// Fee value.
	Fee NullableFloat64 `json:"fee"`
	// Net value.
	Net NullableFloat64 `json:"net"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/><br/> Auction Site.
	AuctionSite NullableString `json:"auctionSite"`
}

// NewPaypalTransactionData instantiates a new PaypalTransactionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalTransactionData(invoiceNumber NullableString, fee NullableFloat64, net NullableFloat64, auctionSite NullableString, ) *PaypalTransactionData {
	this := PaypalTransactionData{}
	this.InvoiceNumber = invoiceNumber
	this.Fee = fee
	this.Net = net
	this.AuctionSite = auctionSite
	return &this
}

// NewPaypalTransactionDataWithDefaults instantiates a new PaypalTransactionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalTransactionDataWithDefaults() *PaypalTransactionData {
	this := PaypalTransactionData{}
	return &this
}

// GetInvoiceNumber returns the InvoiceNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaypalTransactionData) GetInvoiceNumber() string {
	if o == nil || o.InvoiceNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.InvoiceNumber.Get()
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaypalTransactionData) GetInvoiceNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InvoiceNumber.Get(), o.InvoiceNumber.IsSet()
}

// SetInvoiceNumber sets field value
func (o *PaypalTransactionData) SetInvoiceNumber(v string) {
	o.InvoiceNumber.Set(&v)
}

// GetFee returns the Fee field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PaypalTransactionData) GetFee() float64 {
	if o == nil || o.Fee.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Fee.Get()
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaypalTransactionData) GetFeeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Fee.Get(), o.Fee.IsSet()
}

// SetFee sets field value
func (o *PaypalTransactionData) SetFee(v float64) {
	o.Fee.Set(&v)
}

// GetNet returns the Net field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PaypalTransactionData) GetNet() float64 {
	if o == nil || o.Net.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Net.Get()
}

// GetNetOk returns a tuple with the Net field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaypalTransactionData) GetNetOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Net.Get(), o.Net.IsSet()
}

// SetNet sets field value
func (o *PaypalTransactionData) SetNet(v float64) {
	o.Net.Set(&v)
}

// GetAuctionSite returns the AuctionSite field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaypalTransactionData) GetAuctionSite() string {
	if o == nil || o.AuctionSite.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuctionSite.Get()
}

// GetAuctionSiteOk returns a tuple with the AuctionSite field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaypalTransactionData) GetAuctionSiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuctionSite.Get(), o.AuctionSite.IsSet()
}

// SetAuctionSite sets field value
func (o *PaypalTransactionData) SetAuctionSite(v string) {
	o.AuctionSite.Set(&v)
}

func (o PaypalTransactionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["invoiceNumber"] = o.InvoiceNumber.Get()
	}
	if true {
		toSerialize["fee"] = o.Fee.Get()
	}
	if true {
		toSerialize["net"] = o.Net.Get()
	}
	if true {
		toSerialize["auctionSite"] = o.AuctionSite.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaypalTransactionData struct {
	value *PaypalTransactionData
	isSet bool
}

func (v NullablePaypalTransactionData) Get() *PaypalTransactionData {
	return v.value
}

func (v *NullablePaypalTransactionData) Set(val *PaypalTransactionData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalTransactionData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalTransactionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalTransactionData(val *PaypalTransactionData) *NullablePaypalTransactionData {
	return &NullablePaypalTransactionData{value: val, isSet: true}
}

func (v NullablePaypalTransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalTransactionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



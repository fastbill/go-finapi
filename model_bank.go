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

// Bank Container for a bank's data
type Bank struct {
	// Bank identifier.<br/><br/>NOTE: Do NOT assume that the identifiers of banks are the same across different finAPI environments. In fact, the identifiers may change whenever a new finAPI version is released, even within the same environment. The identifiers are meant to be used for references within the finAPI services only, but not for hard-coding them in your application. If you need to hard-code the usage of a certain bank within your application, please instead refer to the BLZ.
	Id int64 `json:"id"`
	// Name of bank
	Name string `json:"name"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginHint' field in the 'interfaces' instead.<br/><br/>Login hint. Contains a German message for the user that explains what kind of credentials are expected.<br/><br/>Please note that it is strongly recommended to always show the login hint to the user if there is one, as the credentials that finAPI requires for the bank might be different to the credentials that the user knows from the bank's website.<br/><br/>Also note that the contents of this field should always be interpreted as HTML, as the text might contain HTML tags for highlighted words, paragraphs, etc.
	LoginHint NullableString `json:"loginHint"`
	// BIC of bank
	Bic NullableString `json:"bic"`
	Blzs []string `json:"blzs"`
	// BLZ of bank
	Blz string `json:"blz"`
	// Bank location (two-letter country code; ISO 3166 ALPHA-2). Note that when this field is not set, it means that this bank depicts an international institute which is not bound to any specific country.
	Location NullableString `json:"location"`
	// City that this bank is located in. Note that this field may not be set for some banks.
	City NullableString `json:"city"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'isMoneyTransferSupported' and 'isAisSupported' fields in the 'interfaces' instead.<br/><br/>Whether this bank is supported by finAPI, i.e. whether you can import/update a bank connection of this bank.<br/><br/>NOTE: this field only regards FINTS_SERVER and WEB_SCRAPER interfaces. XS2A is not considered.
	IsSupported bool `json:"isSupported"`
	// If true, then this bank does not depict a real bank, but rather a testing endpoint provided by a bank or by finAPI. You probably want to regard these banks only during the development of your application, but not in production. You can filter out these banks in production by making sure that the 'isTestBank' parameter is always set to 'false' whenever your application is calling the 'Get and search all banks' service.
	IsTestBank bool `json:"isTestBank"`
	// Popularity of this bank with your users (mandator-wide, i.e. across all of your clients). The value equals the number of bank connections that are currently imported for this bank across all of your users (which means it is a constantly adjusting value). You can use this field for statistical evaluation, and also for ordering bank search results (see service 'Get and search all banks').
	Popularity int32 `json:"popularity"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'health' field in the 'interfaces' instead. <br/><br/>The health status of this bank. This is a value between 0 and 100, depicting the percentage of successful communication attempts with this bank during the latest couple of bank connection imports or updates (across the entire finAPI system). Note that 'successful' means that there was no technical error trying to establish a communication with the bank. Non-technical errors (like incorrect credentials) are regarded successful communication attempts.
	Health int32 `json:"health"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' in the 'interfaces' instead.<br/><br/>Label for the user ID login field, as it is called on the bank's website (e.g. \"Nutzerkennung\"). If this field is set (i.e. not null) then you should prompt your users to enter the required data in a text field which you can label with this field's value.
	LoginFieldUserId NullableString `json:"loginFieldUserId"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' in the 'interfaces' instead.<br/><br/>Label for the customer ID login field, as it is called on the bank's website (e.g. \"Kundennummer\"). If this field is set (i.e. not null) then you should prompt your users to enter the required data in a text field which you can label with this field's value.
	LoginFieldCustomerId NullableString `json:"loginFieldCustomerId"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' in the 'interfaces' instead.<br/><br/>Label for the PIN field, as it is called on the bank's website (mostly \"PIN\"). If this field is set (i.e. not null) then you should prompt your users to enter the required data in a text field which you can label with this field's value.
	LoginFieldPin NullableString `json:"loginFieldPin"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'isVolatile' field of the 'loginCredentials' in the 'interfaces' instead.<br/><br/>Whether the PINs that are used for authentication with the bank are volatile. If the PINs are volatile, it means that the value for the field, which is provided by the user, is valid only for a single authentication and then gets invalidated on bank-side. If pinsAreVolatile is true, it is not possible to store the PIN in finAPI, as a stored PIN will never work for future authentications.
	PinsAreVolatile bool `json:"pinsAreVolatile"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'isSecret' field of the 'loginCredentials' in 'interfaces' instead.<br/><br/>Whether the banking customer ID has to be treated like a password. Certain banks require a second password (besides the PIN) for the user to login. In this case your application should use a password input field when prompting users for their credentials.
	IsCustomerIdPassword bool `json:"isCustomerIdPassword"`
	SupportedDataSources []SupportedDataSource `json:"supportedDataSources"`
	// <strong>Type:</strong> BankInterface<br/> Set of interfaces that finAPI can use to connect to the bank. Note that this set will be empty for non-supported banks. Note also that the WEB_SCRAPER interface might be disabled for your client (see GET /clientConfiguration). When this is the case, then finAPI will not use the web scraper for data download, and if the web scraper is the only supported interface of this bank, then finAPI will not allow to download any data for this bank at all (for details, see POST /bankConnections/import and POST /bankConnections/update).
	Interfaces []BankInterface `json:"interfaces"`
	// <strong>Type:</strong> BankGroup<br/> Bank group
	BankGroup NullableBankGroup `json:"bankGroup"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'lastCommunicationAttempt' field in the 'interfaces' instead. <br/><br/>Time of the last communication attempt with this bank during a bank connection import or update (across the entire finAPI system). The value is returned in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	LastCommunicationAttempt NullableString `json:"lastCommunicationAttempt"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'lastSuccessfulCommunication' field in the 'interfaces' instead. <br/><br/>Time of the last successful communication with this bank during a bank connection import or update (across the entire finAPI system). The value is returned in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	LastSuccessfulCommunication NullableString `json:"lastSuccessfulCommunication"`
	// Whether this bank is in beta phase. For more details, please refer to the field ClientConfiguration.betaBanksEnabled.
	IsBeta bool `json:"isBeta"`
}

// NewBank instantiates a new Bank object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBank(id int64, name string, loginHint NullableString, bic NullableString, blzs []string, blz string, location NullableString, city NullableString, isSupported bool, isTestBank bool, popularity int32, health int32, loginFieldUserId NullableString, loginFieldCustomerId NullableString, loginFieldPin NullableString, pinsAreVolatile bool, isCustomerIdPassword bool, supportedDataSources []SupportedDataSource, interfaces []BankInterface, bankGroup NullableBankGroup, lastCommunicationAttempt NullableString, lastSuccessfulCommunication NullableString, isBeta bool, ) *Bank {
	this := Bank{}
	this.Id = id
	this.Name = name
	this.LoginHint = loginHint
	this.Bic = bic
	this.Blzs = blzs
	this.Blz = blz
	this.Location = location
	this.City = city
	this.IsSupported = isSupported
	this.IsTestBank = isTestBank
	this.Popularity = popularity
	this.Health = health
	this.LoginFieldUserId = loginFieldUserId
	this.LoginFieldCustomerId = loginFieldCustomerId
	this.LoginFieldPin = loginFieldPin
	this.PinsAreVolatile = pinsAreVolatile
	this.IsCustomerIdPassword = isCustomerIdPassword
	this.SupportedDataSources = supportedDataSources
	this.Interfaces = interfaces
	this.BankGroup = bankGroup
	this.LastCommunicationAttempt = lastCommunicationAttempt
	this.LastSuccessfulCommunication = lastSuccessfulCommunication
	this.IsBeta = isBeta
	return &this
}

// NewBankWithDefaults instantiates a new Bank object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankWithDefaults() *Bank {
	this := Bank{}
	return &this
}

// GetId returns the Id field value
func (o *Bank) GetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Bank) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Bank) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Bank) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Bank) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Bank) SetName(v string) {
	o.Name = v
}

// GetLoginHint returns the LoginHint field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bank) GetLoginHint() string {
	if o == nil || o.LoginHint.Get() == nil {
		var ret string
		return ret
	}

	return *o.LoginHint.Get()
}

// GetLoginHintOk returns a tuple with the LoginHint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bank) GetLoginHintOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoginHint.Get(), o.LoginHint.IsSet()
}

// SetLoginHint sets field value
func (o *Bank) SetLoginHint(v string) {
	o.LoginHint.Set(&v)
}

// GetBic returns the Bic field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bank) GetBic() string {
	if o == nil || o.Bic.Get() == nil {
		var ret string
		return ret
	}

	return *o.Bic.Get()
}

// GetBicOk returns a tuple with the Bic field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bank) GetBicOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bic.Get(), o.Bic.IsSet()
}

// SetBic sets field value
func (o *Bank) SetBic(v string) {
	o.Bic.Set(&v)
}

// GetBlzs returns the Blzs field value
func (o *Bank) GetBlzs() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Blzs
}

// GetBlzsOk returns a tuple with the Blzs field value
// and a boolean to check if the value has been set.
func (o *Bank) GetBlzsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Blzs, true
}

// SetBlzs sets field value
func (o *Bank) SetBlzs(v []string) {
	o.Blzs = v
}

// GetBlz returns the Blz field value
func (o *Bank) GetBlz() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Blz
}

// GetBlzOk returns a tuple with the Blz field value
// and a boolean to check if the value has been set.
func (o *Bank) GetBlzOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Blz, true
}

// SetBlz sets field value
func (o *Bank) SetBlz(v string) {
	o.Blz = v
}

// GetLocation returns the Location field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bank) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}

	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bank) GetLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// SetLocation sets field value
func (o *Bank) SetLocation(v string) {
	o.Location.Set(&v)
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bank) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bank) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *Bank) SetCity(v string) {
	o.City.Set(&v)
}

// GetIsSupported returns the IsSupported field value
func (o *Bank) GetIsSupported() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsSupported
}

// GetIsSupportedOk returns a tuple with the IsSupported field value
// and a boolean to check if the value has been set.
func (o *Bank) GetIsSupportedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSupported, true
}

// SetIsSupported sets field value
func (o *Bank) SetIsSupported(v bool) {
	o.IsSupported = v
}

// GetIsTestBank returns the IsTestBank field value
func (o *Bank) GetIsTestBank() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsTestBank
}

// GetIsTestBankOk returns a tuple with the IsTestBank field value
// and a boolean to check if the value has been set.
func (o *Bank) GetIsTestBankOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsTestBank, true
}

// SetIsTestBank sets field value
func (o *Bank) SetIsTestBank(v bool) {
	o.IsTestBank = v
}

// GetPopularity returns the Popularity field value
func (o *Bank) GetPopularity() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value
// and a boolean to check if the value has been set.
func (o *Bank) GetPopularityOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Popularity, true
}

// SetPopularity sets field value
func (o *Bank) SetPopularity(v int32) {
	o.Popularity = v
}

// GetHealth returns the Health field value
func (o *Bank) GetHealth() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Health
}

// GetHealthOk returns a tuple with the Health field value
// and a boolean to check if the value has been set.
func (o *Bank) GetHealthOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Health, true
}

// SetHealth sets field value
func (o *Bank) SetHealth(v int32) {
	o.Health = v
}

// GetLoginFieldUserId returns the LoginFieldUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bank) GetLoginFieldUserId() string {
	if o == nil || o.LoginFieldUserId.Get() == nil {
		var ret string
		return ret
	}

	return *o.LoginFieldUserId.Get()
}

// GetLoginFieldUserIdOk returns a tuple with the LoginFieldUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bank) GetLoginFieldUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoginFieldUserId.Get(), o.LoginFieldUserId.IsSet()
}

// SetLoginFieldUserId sets field value
func (o *Bank) SetLoginFieldUserId(v string) {
	o.LoginFieldUserId.Set(&v)
}

// GetLoginFieldCustomerId returns the LoginFieldCustomerId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bank) GetLoginFieldCustomerId() string {
	if o == nil || o.LoginFieldCustomerId.Get() == nil {
		var ret string
		return ret
	}

	return *o.LoginFieldCustomerId.Get()
}

// GetLoginFieldCustomerIdOk returns a tuple with the LoginFieldCustomerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bank) GetLoginFieldCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoginFieldCustomerId.Get(), o.LoginFieldCustomerId.IsSet()
}

// SetLoginFieldCustomerId sets field value
func (o *Bank) SetLoginFieldCustomerId(v string) {
	o.LoginFieldCustomerId.Set(&v)
}

// GetLoginFieldPin returns the LoginFieldPin field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bank) GetLoginFieldPin() string {
	if o == nil || o.LoginFieldPin.Get() == nil {
		var ret string
		return ret
	}

	return *o.LoginFieldPin.Get()
}

// GetLoginFieldPinOk returns a tuple with the LoginFieldPin field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bank) GetLoginFieldPinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoginFieldPin.Get(), o.LoginFieldPin.IsSet()
}

// SetLoginFieldPin sets field value
func (o *Bank) SetLoginFieldPin(v string) {
	o.LoginFieldPin.Set(&v)
}

// GetPinsAreVolatile returns the PinsAreVolatile field value
func (o *Bank) GetPinsAreVolatile() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.PinsAreVolatile
}

// GetPinsAreVolatileOk returns a tuple with the PinsAreVolatile field value
// and a boolean to check if the value has been set.
func (o *Bank) GetPinsAreVolatileOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PinsAreVolatile, true
}

// SetPinsAreVolatile sets field value
func (o *Bank) SetPinsAreVolatile(v bool) {
	o.PinsAreVolatile = v
}

// GetIsCustomerIdPassword returns the IsCustomerIdPassword field value
func (o *Bank) GetIsCustomerIdPassword() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsCustomerIdPassword
}

// GetIsCustomerIdPasswordOk returns a tuple with the IsCustomerIdPassword field value
// and a boolean to check if the value has been set.
func (o *Bank) GetIsCustomerIdPasswordOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsCustomerIdPassword, true
}

// SetIsCustomerIdPassword sets field value
func (o *Bank) SetIsCustomerIdPassword(v bool) {
	o.IsCustomerIdPassword = v
}

// GetSupportedDataSources returns the SupportedDataSources field value
func (o *Bank) GetSupportedDataSources() []SupportedDataSource {
	if o == nil  {
		var ret []SupportedDataSource
		return ret
	}

	return o.SupportedDataSources
}

// GetSupportedDataSourcesOk returns a tuple with the SupportedDataSources field value
// and a boolean to check if the value has been set.
func (o *Bank) GetSupportedDataSourcesOk() (*[]SupportedDataSource, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupportedDataSources, true
}

// SetSupportedDataSources sets field value
func (o *Bank) SetSupportedDataSources(v []SupportedDataSource) {
	o.SupportedDataSources = v
}

// GetInterfaces returns the Interfaces field value
func (o *Bank) GetInterfaces() []BankInterface {
	if o == nil  {
		var ret []BankInterface
		return ret
	}

	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value
// and a boolean to check if the value has been set.
func (o *Bank) GetInterfacesOk() (*[]BankInterface, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Interfaces, true
}

// SetInterfaces sets field value
func (o *Bank) SetInterfaces(v []BankInterface) {
	o.Interfaces = v
}

// GetBankGroup returns the BankGroup field value
// If the value is explicit nil, the zero value for BankGroup will be returned
func (o *Bank) GetBankGroup() BankGroup {
	if o == nil || o.BankGroup.Get() == nil {
		var ret BankGroup
		return ret
	}

	return *o.BankGroup.Get()
}

// GetBankGroupOk returns a tuple with the BankGroup field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bank) GetBankGroupOk() (*BankGroup, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankGroup.Get(), o.BankGroup.IsSet()
}

// SetBankGroup sets field value
func (o *Bank) SetBankGroup(v BankGroup) {
	o.BankGroup.Set(&v)
}

// GetLastCommunicationAttempt returns the LastCommunicationAttempt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bank) GetLastCommunicationAttempt() string {
	if o == nil || o.LastCommunicationAttempt.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastCommunicationAttempt.Get()
}

// GetLastCommunicationAttemptOk returns a tuple with the LastCommunicationAttempt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bank) GetLastCommunicationAttemptOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastCommunicationAttempt.Get(), o.LastCommunicationAttempt.IsSet()
}

// SetLastCommunicationAttempt sets field value
func (o *Bank) SetLastCommunicationAttempt(v string) {
	o.LastCommunicationAttempt.Set(&v)
}

// GetLastSuccessfulCommunication returns the LastSuccessfulCommunication field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Bank) GetLastSuccessfulCommunication() string {
	if o == nil || o.LastSuccessfulCommunication.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastSuccessfulCommunication.Get()
}

// GetLastSuccessfulCommunicationOk returns a tuple with the LastSuccessfulCommunication field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bank) GetLastSuccessfulCommunicationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastSuccessfulCommunication.Get(), o.LastSuccessfulCommunication.IsSet()
}

// SetLastSuccessfulCommunication sets field value
func (o *Bank) SetLastSuccessfulCommunication(v string) {
	o.LastSuccessfulCommunication.Set(&v)
}

// GetIsBeta returns the IsBeta field value
func (o *Bank) GetIsBeta() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsBeta
}

// GetIsBetaOk returns a tuple with the IsBeta field value
// and a boolean to check if the value has been set.
func (o *Bank) GetIsBetaOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsBeta, true
}

// SetIsBeta sets field value
func (o *Bank) SetIsBeta(v bool) {
	o.IsBeta = v
}

func (o Bank) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["loginHint"] = o.LoginHint.Get()
	}
	if true {
		toSerialize["bic"] = o.Bic.Get()
	}
	if true {
		toSerialize["blzs"] = o.Blzs
	}
	if true {
		toSerialize["blz"] = o.Blz
	}
	if true {
		toSerialize["location"] = o.Location.Get()
	}
	if true {
		toSerialize["city"] = o.City.Get()
	}
	if true {
		toSerialize["isSupported"] = o.IsSupported
	}
	if true {
		toSerialize["isTestBank"] = o.IsTestBank
	}
	if true {
		toSerialize["popularity"] = o.Popularity
	}
	if true {
		toSerialize["health"] = o.Health
	}
	if true {
		toSerialize["loginFieldUserId"] = o.LoginFieldUserId.Get()
	}
	if true {
		toSerialize["loginFieldCustomerId"] = o.LoginFieldCustomerId.Get()
	}
	if true {
		toSerialize["loginFieldPin"] = o.LoginFieldPin.Get()
	}
	if true {
		toSerialize["pinsAreVolatile"] = o.PinsAreVolatile
	}
	if true {
		toSerialize["isCustomerIdPassword"] = o.IsCustomerIdPassword
	}
	if true {
		toSerialize["supportedDataSources"] = o.SupportedDataSources
	}
	if true {
		toSerialize["interfaces"] = o.Interfaces
	}
	if true {
		toSerialize["bankGroup"] = o.BankGroup.Get()
	}
	if true {
		toSerialize["lastCommunicationAttempt"] = o.LastCommunicationAttempt.Get()
	}
	if true {
		toSerialize["lastSuccessfulCommunication"] = o.LastSuccessfulCommunication.Get()
	}
	if true {
		toSerialize["isBeta"] = o.IsBeta
	}
	return json.Marshal(toSerialize)
}

type NullableBank struct {
	value *Bank
	isSet bool
}

func (v NullableBank) Get() *Bank {
	return v.value
}

func (v *NullableBank) Set(val *Bank) {
	v.value = val
	v.isSet = true
}

func (v NullableBank) IsSet() bool {
	return v.isSet
}

func (v *NullableBank) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBank(val *Bank) *NullableBank {
	return &NullableBank{value: val, isSet: true}
}

func (v NullableBank) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBank) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



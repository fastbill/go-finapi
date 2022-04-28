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

// BankConnectionInterface Resource representing a bank connection interface
type BankConnectionInterface struct {
	// <strong>Type:</strong> BankingInterface<br/> Bank interface. Possible values:<br><br>&bull; <code>WEB_SCRAPER</code> - means that finAPI will parse data from the bank's online banking website.<br>&bull; <code>FINTS_SERVER</code> - means that finAPI will download data via the bank's FinTS interface.<br>&bull; <code>XS2A</code> - means that finAPI will download data via the bank's XS2A interface.<br>
	Interface BankingInterface `json:"interface"`
	// <strong>Type:</strong> LoginCredentialResource<br/> Login fields for this interface (in the order that we suggest to show them to the user), with their currently stored values. Note that this list always contains all existing login fields for this interface, even when there is no stored value for a field (value will be null in such a case).
	LoginCredentials []LoginCredentialResource `json:"loginCredentials"`
	// The default two-step-procedure for this interface. Must match one of the available 'procedureId's from the 'twoStepProcedures' list. When this field is set, then finAPI will automatically try to select the procedure wherever applicable. Note that the list of available procedures of a bank connection may change as a result of an update of the connection, and if this field references a procedure that is no longer available after an update, finAPI will automatically clear the default procedure (set it to null).
	DefaultTwoStepProcedureId NullableString `json:"defaultTwoStepProcedureId"`
	// <strong>Type:</strong> TwoStepProcedure<br/> Available two-step-procedures in this interface, used for submitting a money transfer or direct debit request (see /accounts/requestSepaMoneyTransfer or /requestSepaDirectDebit),or for multi-step-authentication during bank connection import or update. The available two-step-procedures mya be re-evaluated each time this bank connection is updated (/bankConnections/update). This means that this list may change as a result of an update.
	TwoStepProcedures []TwoStepProcedure `json:"twoStepProcedures"`
	// <strong>Type:</strong> BankConsent<br/> If this field is set, it means that this interface is handing out a consent to finAPI in exchange for the login credentials. finAPI needs to use this consent to get access to the account list and account data (i.e. Account Information Services, AIS). If this field is not set, it means that this interface does not use such consents.
	AisConsent NullableBankConsent `json:"aisConsent"`
	// <strong>Type:</strong> UpdateResult<br/> Result of the last manual update of the associated bank connection using this interface. If no manual update has ever been done so far with this interface, then this field will not be set.
	LastManualUpdate NullableUpdateResult `json:"lastManualUpdate"`
	// <strong>Type:</strong> UpdateResult<br/> Result of the last auto update of the associated bank connection using this interface (ran by finAPI's automatic batch update process). If no auto update has ever been done so far with this interface, then this field will not be set.
	LastAutoUpdate NullableUpdateResult `json:"lastAutoUpdate"`
	// This field indicates whether the user's attention is required for the next update of the given bank connection interface.<br/>If the field is true, finAPI stops auto-updates of this bank connection interface to mitigate the risk of locking the user's bank account and also of triggering a multi-step authentication that might lead to a notification being sent to the end-user.<br/>If the field is false, the user's attention might still be required for the next bank update, e.g. because of new Terms and Conditions that have to get approved by the user.(this only applies to users whose mandator doesn't have an AIS license)<br/>Every communication with the bank (e.g. updating a bank connection, submitting a money transfer or a direct debit, etc.) can change the value of this flag. If the field is true, we recommend to ask the end-user to trigger a manual update of the bank connection interface (using the 'Update a bank connection' service). If the update completes successfully without triggering a strong customer authentication or results in storing a valid XS2A consent, this flag will switch to false. The logic about determination of the user's attention being required might change in time. Please use this as a convenience function to know, when you have to involve the user in the next communication with the bank. Once the flag switches to false, the bank connection interface will be enabled again for the auto-update (if it is configured).
	UserActionRequired bool `json:"userActionRequired"`
	// This setting defines how much of an account's transactions history will be downloaded whenever a new account is imported. More technically, it depicts the number of days to download transactions for, starting from - and including - the date of the account import. For example, on an account import that happens today, the value 30 would instruct finAPI to download transactions from the past 30 days (including today). The minimum allowed value is 14, the maximum value is 3650. Also possible is the value 0 (which is the default value), in which case there will be no limit to the transactions download and finAPI will try to get all transactions that it can. <br/><br/>Note:<br/>&bull; There is no guarantee that finAPI will actually download transactions for the entire defined date range, as there may be limitations to the download range (set by the bank or by finAPI, e.g. see ClientConfiguration.transactionImportLimitation). <br/>&bull; This parameter only applies to transactions, not to security positions; For security accounts, finAPI will always download all security positions that it can. <br/>&bull; This setting is stored for each interface individually.<br/>&bull; After an interface has been connected with this setting, there is no way to change the setting for that interface afterwards.<br/>&bull; <b>If you do not limit the download range to a value less than 90 days, the bank is more likely to trigger a strong customer authentication request for the user when finAPI is attempting to download the transactions.</b>
	MaxDaysForDownload int32 `json:"maxDaysForDownload"`
}

// NewBankConnectionInterface instantiates a new BankConnectionInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankConnectionInterface(interface_ BankingInterface, loginCredentials []LoginCredentialResource, defaultTwoStepProcedureId NullableString, twoStepProcedures []TwoStepProcedure, aisConsent NullableBankConsent, lastManualUpdate NullableUpdateResult, lastAutoUpdate NullableUpdateResult, userActionRequired bool, maxDaysForDownload int32, ) *BankConnectionInterface {
	this := BankConnectionInterface{}
	this.Interface = interface_
	this.LoginCredentials = loginCredentials
	this.DefaultTwoStepProcedureId = defaultTwoStepProcedureId
	this.TwoStepProcedures = twoStepProcedures
	this.AisConsent = aisConsent
	this.LastManualUpdate = lastManualUpdate
	this.LastAutoUpdate = lastAutoUpdate
	this.UserActionRequired = userActionRequired
	this.MaxDaysForDownload = maxDaysForDownload
	return &this
}

// NewBankConnectionInterfaceWithDefaults instantiates a new BankConnectionInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankConnectionInterfaceWithDefaults() *BankConnectionInterface {
	this := BankConnectionInterface{}
	return &this
}

// GetInterface returns the Interface field value
func (o *BankConnectionInterface) GetInterface() BankingInterface {
	if o == nil  {
		var ret BankingInterface
		return ret
	}

	return o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value
// and a boolean to check if the value has been set.
func (o *BankConnectionInterface) GetInterfaceOk() (*BankingInterface, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Interface, true
}

// SetInterface sets field value
func (o *BankConnectionInterface) SetInterface(v BankingInterface) {
	o.Interface = v
}

// GetLoginCredentials returns the LoginCredentials field value
func (o *BankConnectionInterface) GetLoginCredentials() []LoginCredentialResource {
	if o == nil  {
		var ret []LoginCredentialResource
		return ret
	}

	return o.LoginCredentials
}

// GetLoginCredentialsOk returns a tuple with the LoginCredentials field value
// and a boolean to check if the value has been set.
func (o *BankConnectionInterface) GetLoginCredentialsOk() (*[]LoginCredentialResource, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LoginCredentials, true
}

// SetLoginCredentials sets field value
func (o *BankConnectionInterface) SetLoginCredentials(v []LoginCredentialResource) {
	o.LoginCredentials = v
}

// GetDefaultTwoStepProcedureId returns the DefaultTwoStepProcedureId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BankConnectionInterface) GetDefaultTwoStepProcedureId() string {
	if o == nil || o.DefaultTwoStepProcedureId.Get() == nil {
		var ret string
		return ret
	}

	return *o.DefaultTwoStepProcedureId.Get()
}

// GetDefaultTwoStepProcedureIdOk returns a tuple with the DefaultTwoStepProcedureId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionInterface) GetDefaultTwoStepProcedureIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DefaultTwoStepProcedureId.Get(), o.DefaultTwoStepProcedureId.IsSet()
}

// SetDefaultTwoStepProcedureId sets field value
func (o *BankConnectionInterface) SetDefaultTwoStepProcedureId(v string) {
	o.DefaultTwoStepProcedureId.Set(&v)
}

// GetTwoStepProcedures returns the TwoStepProcedures field value
func (o *BankConnectionInterface) GetTwoStepProcedures() []TwoStepProcedure {
	if o == nil  {
		var ret []TwoStepProcedure
		return ret
	}

	return o.TwoStepProcedures
}

// GetTwoStepProceduresOk returns a tuple with the TwoStepProcedures field value
// and a boolean to check if the value has been set.
func (o *BankConnectionInterface) GetTwoStepProceduresOk() (*[]TwoStepProcedure, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TwoStepProcedures, true
}

// SetTwoStepProcedures sets field value
func (o *BankConnectionInterface) SetTwoStepProcedures(v []TwoStepProcedure) {
	o.TwoStepProcedures = v
}

// GetAisConsent returns the AisConsent field value
// If the value is explicit nil, the zero value for BankConsent will be returned
func (o *BankConnectionInterface) GetAisConsent() BankConsent {
	if o == nil || o.AisConsent.Get() == nil {
		var ret BankConsent
		return ret
	}

	return *o.AisConsent.Get()
}

// GetAisConsentOk returns a tuple with the AisConsent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionInterface) GetAisConsentOk() (*BankConsent, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AisConsent.Get(), o.AisConsent.IsSet()
}

// SetAisConsent sets field value
func (o *BankConnectionInterface) SetAisConsent(v BankConsent) {
	o.AisConsent.Set(&v)
}

// GetLastManualUpdate returns the LastManualUpdate field value
// If the value is explicit nil, the zero value for UpdateResult will be returned
func (o *BankConnectionInterface) GetLastManualUpdate() UpdateResult {
	if o == nil || o.LastManualUpdate.Get() == nil {
		var ret UpdateResult
		return ret
	}

	return *o.LastManualUpdate.Get()
}

// GetLastManualUpdateOk returns a tuple with the LastManualUpdate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionInterface) GetLastManualUpdateOk() (*UpdateResult, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastManualUpdate.Get(), o.LastManualUpdate.IsSet()
}

// SetLastManualUpdate sets field value
func (o *BankConnectionInterface) SetLastManualUpdate(v UpdateResult) {
	o.LastManualUpdate.Set(&v)
}

// GetLastAutoUpdate returns the LastAutoUpdate field value
// If the value is explicit nil, the zero value for UpdateResult will be returned
func (o *BankConnectionInterface) GetLastAutoUpdate() UpdateResult {
	if o == nil || o.LastAutoUpdate.Get() == nil {
		var ret UpdateResult
		return ret
	}

	return *o.LastAutoUpdate.Get()
}

// GetLastAutoUpdateOk returns a tuple with the LastAutoUpdate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnectionInterface) GetLastAutoUpdateOk() (*UpdateResult, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastAutoUpdate.Get(), o.LastAutoUpdate.IsSet()
}

// SetLastAutoUpdate sets field value
func (o *BankConnectionInterface) SetLastAutoUpdate(v UpdateResult) {
	o.LastAutoUpdate.Set(&v)
}

// GetUserActionRequired returns the UserActionRequired field value
func (o *BankConnectionInterface) GetUserActionRequired() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.UserActionRequired
}

// GetUserActionRequiredOk returns a tuple with the UserActionRequired field value
// and a boolean to check if the value has been set.
func (o *BankConnectionInterface) GetUserActionRequiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserActionRequired, true
}

// SetUserActionRequired sets field value
func (o *BankConnectionInterface) SetUserActionRequired(v bool) {
	o.UserActionRequired = v
}

// GetMaxDaysForDownload returns the MaxDaysForDownload field value
func (o *BankConnectionInterface) GetMaxDaysForDownload() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.MaxDaysForDownload
}

// GetMaxDaysForDownloadOk returns a tuple with the MaxDaysForDownload field value
// and a boolean to check if the value has been set.
func (o *BankConnectionInterface) GetMaxDaysForDownloadOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxDaysForDownload, true
}

// SetMaxDaysForDownload sets field value
func (o *BankConnectionInterface) SetMaxDaysForDownload(v int32) {
	o.MaxDaysForDownload = v
}

func (o BankConnectionInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interface"] = o.Interface
	}
	if true {
		toSerialize["loginCredentials"] = o.LoginCredentials
	}
	if true {
		toSerialize["defaultTwoStepProcedureId"] = o.DefaultTwoStepProcedureId.Get()
	}
	if true {
		toSerialize["twoStepProcedures"] = o.TwoStepProcedures
	}
	if true {
		toSerialize["aisConsent"] = o.AisConsent.Get()
	}
	if true {
		toSerialize["lastManualUpdate"] = o.LastManualUpdate.Get()
	}
	if true {
		toSerialize["lastAutoUpdate"] = o.LastAutoUpdate.Get()
	}
	if true {
		toSerialize["userActionRequired"] = o.UserActionRequired
	}
	if true {
		toSerialize["maxDaysForDownload"] = o.MaxDaysForDownload
	}
	return json.Marshal(toSerialize)
}

type NullableBankConnectionInterface struct {
	value *BankConnectionInterface
	isSet bool
}

func (v NullableBankConnectionInterface) Get() *BankConnectionInterface {
	return v.value
}

func (v *NullableBankConnectionInterface) Set(val *BankConnectionInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableBankConnectionInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableBankConnectionInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankConnectionInterface(val *BankConnectionInterface) *NullableBankConnectionInterface {
	return &NullableBankConnectionInterface{value: val, isSet: true}
}

func (v NullableBankConnectionInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankConnectionInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



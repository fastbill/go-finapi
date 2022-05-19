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

// UserInfo Container for user information
type UserInfo struct {
	// User's identifier
	UserId string `json:"userId"`
	// User's registration date, in the format 'YYYY-MM-DD'
	RegistrationDate string `json:"registrationDate"`
	// User's deletion date, in the format 'YYYY-MM-DD'. May be null if the user has not been deleted.
	DeletionDate NullableString `json:"deletionDate"`
	// User's last active date, in the format 'YYYY-MM-DD'. May be null if the user has not yet logged in.
	LastActiveDate NullableString `json:"lastActiveDate"`
	// Number of bank connections that currently exist for this user.
	BankConnectionCount int32 `json:"bankConnectionCount"`
	// Latest date of when a bank connection was imported for this user, in the format 'YYYY-MM-DD'. This field is null when there has never been a bank connection import.
	LatestBankConnectionImportDate NullableString `json:"latestBankConnectionImportDate"`
	// Latest date of when a bank connection was deleted for this user, in the format 'YYYY-MM-DD'. This field is null when there has never been a bank connection deletion.
	LatestBankConnectionDeletionDate NullableString `json:"latestBankConnectionDeletionDate"`
	// <strong>Type:</strong> MonthlyUserStats<br/> Additional information about the user's data or activities, broken down in months. The list will by default contain an entry for each month starting with the month of when the user was registered, up to the current month. The date range may vary when you have limited it in the request. <br/><br/>Please note:<br/>&bull; this field is only set when 'includeMonthlyStats' = true, otherwise it will be null.<br/>&bull; the list is always ordered from the latest month first, to the oldest month last.<br/>&bull; the list will never contain an entry for a month that was prior to the month of when the user was registered, or after the month of when the user was deleted, even when you have explicitly set a respective date range. This means that the list may be empty if you are requesting a date range where the user didn't exist yet, or didn't exist any longer.
	MonthlyStats []MonthlyUserStats `json:"monthlyStats"`
	// Whether the user is currently locked (for further information, see the 'maxUserLoginAttempts' setting in your client's configuration). Note that deleted users will always have this field set to 'false'.
	IsLocked bool `json:"isLocked"`
}

// NewUserInfo instantiates a new UserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInfo(userId string, registrationDate string, deletionDate NullableString, lastActiveDate NullableString, bankConnectionCount int32, latestBankConnectionImportDate NullableString, latestBankConnectionDeletionDate NullableString, monthlyStats []MonthlyUserStats, isLocked bool, ) *UserInfo {
	this := UserInfo{}
	this.UserId = userId
	this.RegistrationDate = registrationDate
	this.DeletionDate = deletionDate
	this.LastActiveDate = lastActiveDate
	this.BankConnectionCount = bankConnectionCount
	this.LatestBankConnectionImportDate = latestBankConnectionImportDate
	this.LatestBankConnectionDeletionDate = latestBankConnectionDeletionDate
	this.MonthlyStats = monthlyStats
	this.IsLocked = isLocked
	return &this
}

// NewUserInfoWithDefaults instantiates a new UserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInfoWithDefaults() *UserInfo {
	this := UserInfo{}
	return &this
}

// GetUserId returns the UserId field value
func (o *UserInfo) GetUserId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserInfo) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserInfo) SetUserId(v string) {
	o.UserId = v
}

// GetRegistrationDate returns the RegistrationDate field value
func (o *UserInfo) GetRegistrationDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RegistrationDate
}

// GetRegistrationDateOk returns a tuple with the RegistrationDate field value
// and a boolean to check if the value has been set.
func (o *UserInfo) GetRegistrationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegistrationDate, true
}

// SetRegistrationDate sets field value
func (o *UserInfo) SetRegistrationDate(v string) {
	o.RegistrationDate = v
}

// GetDeletionDate returns the DeletionDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserInfo) GetDeletionDate() string {
	if o == nil || o.DeletionDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.DeletionDate.Get()
}

// GetDeletionDateOk returns a tuple with the DeletionDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetDeletionDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletionDate.Get(), o.DeletionDate.IsSet()
}

// SetDeletionDate sets field value
func (o *UserInfo) SetDeletionDate(v string) {
	o.DeletionDate.Set(&v)
}

// GetLastActiveDate returns the LastActiveDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserInfo) GetLastActiveDate() string {
	if o == nil || o.LastActiveDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastActiveDate.Get()
}

// GetLastActiveDateOk returns a tuple with the LastActiveDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetLastActiveDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastActiveDate.Get(), o.LastActiveDate.IsSet()
}

// SetLastActiveDate sets field value
func (o *UserInfo) SetLastActiveDate(v string) {
	o.LastActiveDate.Set(&v)
}

// GetBankConnectionCount returns the BankConnectionCount field value
func (o *UserInfo) GetBankConnectionCount() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.BankConnectionCount
}

// GetBankConnectionCountOk returns a tuple with the BankConnectionCount field value
// and a boolean to check if the value has been set.
func (o *UserInfo) GetBankConnectionCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankConnectionCount, true
}

// SetBankConnectionCount sets field value
func (o *UserInfo) SetBankConnectionCount(v int32) {
	o.BankConnectionCount = v
}

// GetLatestBankConnectionImportDate returns the LatestBankConnectionImportDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserInfo) GetLatestBankConnectionImportDate() string {
	if o == nil || o.LatestBankConnectionImportDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.LatestBankConnectionImportDate.Get()
}

// GetLatestBankConnectionImportDateOk returns a tuple with the LatestBankConnectionImportDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetLatestBankConnectionImportDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LatestBankConnectionImportDate.Get(), o.LatestBankConnectionImportDate.IsSet()
}

// SetLatestBankConnectionImportDate sets field value
func (o *UserInfo) SetLatestBankConnectionImportDate(v string) {
	o.LatestBankConnectionImportDate.Set(&v)
}

// GetLatestBankConnectionDeletionDate returns the LatestBankConnectionDeletionDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserInfo) GetLatestBankConnectionDeletionDate() string {
	if o == nil || o.LatestBankConnectionDeletionDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.LatestBankConnectionDeletionDate.Get()
}

// GetLatestBankConnectionDeletionDateOk returns a tuple with the LatestBankConnectionDeletionDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetLatestBankConnectionDeletionDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LatestBankConnectionDeletionDate.Get(), o.LatestBankConnectionDeletionDate.IsSet()
}

// SetLatestBankConnectionDeletionDate sets field value
func (o *UserInfo) SetLatestBankConnectionDeletionDate(v string) {
	o.LatestBankConnectionDeletionDate.Set(&v)
}

// GetMonthlyStats returns the MonthlyStats field value
// If the value is explicit nil, the zero value for []MonthlyUserStats will be returned
func (o *UserInfo) GetMonthlyStats() []MonthlyUserStats {
	if o == nil  {
		var ret []MonthlyUserStats
		return ret
	}

	return o.MonthlyStats
}

// GetMonthlyStatsOk returns a tuple with the MonthlyStats field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetMonthlyStatsOk() (*[]MonthlyUserStats, bool) {
	if o == nil || o.MonthlyStats == nil {
		return nil, false
	}
	return &o.MonthlyStats, true
}

// SetMonthlyStats sets field value
func (o *UserInfo) SetMonthlyStats(v []MonthlyUserStats) {
	o.MonthlyStats = v
}

// GetIsLocked returns the IsLocked field value
func (o *UserInfo) GetIsLocked() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value
// and a boolean to check if the value has been set.
func (o *UserInfo) GetIsLockedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsLocked, true
}

// SetIsLocked sets field value
func (o *UserInfo) SetIsLocked(v bool) {
	o.IsLocked = v
}

func (o UserInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["userId"] = o.UserId
	}
	if true {
		toSerialize["registrationDate"] = o.RegistrationDate
	}
	if true {
		toSerialize["deletionDate"] = o.DeletionDate.Get()
	}
	if true {
		toSerialize["lastActiveDate"] = o.LastActiveDate.Get()
	}
	if true {
		toSerialize["bankConnectionCount"] = o.BankConnectionCount
	}
	if true {
		toSerialize["latestBankConnectionImportDate"] = o.LatestBankConnectionImportDate.Get()
	}
	if true {
		toSerialize["latestBankConnectionDeletionDate"] = o.LatestBankConnectionDeletionDate.Get()
	}
	if o.MonthlyStats != nil {
		toSerialize["monthlyStats"] = o.MonthlyStats
	}
	if true {
		toSerialize["isLocked"] = o.IsLocked
	}
	return json.Marshal(toSerialize)
}

type NullableUserInfo struct {
	value *UserInfo
	isSet bool
}

func (v NullableUserInfo) Get() *UserInfo {
	return v.value
}

func (v *NullableUserInfo) Set(val *UserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInfo(val *UserInfo) *NullableUserInfo {
	return &NullableUserInfo{value: val, isSet: true}
}

func (v NullableUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



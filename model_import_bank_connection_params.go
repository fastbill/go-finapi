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

// ImportBankConnectionParams Container for bank connection import parameters
type ImportBankConnectionParams struct {
	// Bank Identifier
	BankId int64 `json:"bankId"`
	// Custom name for the bank connection. Maximum length is 64. If you do not want to set a name, you can leave this field unset.
	Name *string `json:"name,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' and 'interface' fields instead. If any of those two fields is used, then the value of this field will be ignored.<br/><br/>Online banking user ID credential. Max length: 170. NOTES:<br/>- if you import the 'demo connection', this field can be left unset;<br/> - if the user will need to enter his credentials in finAPI's Web Form, this field can contain any value. It will be ignored.
	BankingUserId *string `json:"bankingUserId,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' and 'interface' fields instead. If any of those two fields is used, then the value of this field will be ignored.<br/><br/>Online banking customer ID credential (for most banks this field can remain unset). Max length: 170. NOTES:<br/>- if the user will need to enter his credentials in finAPI's Web Form, this field can contain any value. It will be ignored.
	BankingCustomerId *string `json:"bankingCustomerId,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' and 'interface' fields instead. If any of those two fields is used, then the value of this field will be ignored.<br/><br/>Online banking PIN. Max length: 170. Any symbols are allowed. NOTES:<br/>- if you import the 'demo connection', this field can be left unset;<br/> - if the user will need to enter his credentials in finAPI's Web Form, this field can be left unset or contain any value. It will be ignored.
	BankingPin *string `json:"bankingPin,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'storeSecrets' field instead.<br/><br/>Whether to store the PIN. If the PIN is stored, it is not required to pass the PIN again when updating this bank connection or when executing orders (like money transfers). Default is false. <br/><br/>NOTES:<br/> - before you set this field to true, please regard the 'pinsAreVolatile' flag of this connection's bank. Please note that volatile credentials will not be stored, even if provided, to enforce user involvement in the next communication with the bank;<br/> - this field is ignored in case when the user will need to use finAPI's Web Form. The user will be able to decide whether to store the PIN or not in the Web Form, depending on the 'storeSecretsAvailableInWebForm' setting (see Client Configuration).
	StorePin *bool `json:"storePin,omitempty"`
	// <strong>Type:</strong> BankingInterface<br/> The interface to use for connecting with the bank.
	Interface *BankingInterface `json:"interface,omitempty"`
	// <strong>Type:</strong> LoginCredential<br/> Set of login credentials. Must be passed in combination with the 'interface' field.
	LoginCredentials *[]LoginCredential `json:"loginCredentials,omitempty"`
	// Whether to store the secret login fields. If the secret fields are stored, then updates can be triggered without the involvement of the users, as long as the credentials remain valid and the bank consent has not expired. Note that bank consent will be stored regardless of the field value. Default value is false.
	StoreSecrets *bool `json:"storeSecrets,omitempty"`
	// Whether to skip the download of transactions and securities or not. If set to true, then finAPI will download just the accounts list with the accounts' information (like account name, number, holder, etc), as well as the accounts' balances (if possible), but skip the download of transactions and securities. Default is false.<br/><br/>NOTES:<br/>&bull; Setting this flag to true is only meant to be used if A) you generally never download positions, because you are only interested in the accounts list and/or balances, or B) you want to get just the list of accounts in a first step, and then delete unwanted accounts from the bank connection, before you trigger another update with transactions download. This approach allows you to download transactions only for the accounts that you want.<br/>&bull; If you skip the download of transactions and securities during an import or update, you can still download them on a later update (though you might not get all positions at a later point, because the date range in which the bank servers provide this data is usually limited).<br/>&bull; If an account already had any positions imported before an update, and you skip the positions download in the update, then the account's updated balance might not add up to the set of transactions / security positions. Be aware that certain services (like GET /accounts/dailyBalances) may give incorrect results for accounts in such a state.<br/>&bull; If this bank connection is updated via finAPI's automatic batch update, then transactions and security positions (of already imported accounts) <u>will</u> be downloaded in any case!<br/>&bull; For security accounts, skipping the downloading of the securities might result in the account's balance also not being downloaded.<br/>&bull; For Bausparen accounts, this field is ignored. finAPI will always download transactions for Bausparen accounts.<br/>
	SkipPositionsDownload *bool `json:"skipPositionsDownload,omitempty"`
	// Whether to load information about the bank connection owner(s) - see field 'owners'. Default value is 'false'.<br/><br/>NOTE: This feature is supported only by the WEB_SCRAPER interface.
	LoadOwnerData *bool `json:"loadOwnerData,omitempty"`
	// This setting defines how much of an account's transactions history will be downloaded whenever a new account is imported. More technically, it depicts the number of days to download transactions for, starting from - and including - the date of the account import. For example, on an account import that happens today, the value 30 would instruct finAPI to download transactions from the past 30 days (including today). The minimum allowed value is 14, the maximum value is 3650. Also possible is the value 0 (which is the default value), in which case there will be no limit to the transactions download and finAPI will try to get all transactions that it can. <br/><br/>Note:<br/>&bull; There is no guarantee that finAPI will actually download transactions for the entire defined date range, as there may be limitations to the download range (set by the bank or by finAPI, e.g. see ClientConfiguration.transactionImportLimitation). <br/>&bull; This parameter only applies to transactions, not to security positions; For security accounts, finAPI will always download all security positions that it can. <br/>&bull; This setting is stored for each interface individually.<br/>&bull; After an interface has been connected with this setting, there is no way to change the setting for that interface afterwards.<br/>&bull; <b>If you do not limit the download range to a value less than 90 days, the bank is more likely to trigger a strong customer authentication request for the user when finAPI is attempting to download the transactions.</b>
	MaxDaysForDownload *int32 `json:"maxDaysForDownload,omitempty"`
	AccountTypes *[]AccountType `json:"accountTypes,omitempty"`
	AccountTypeIds *[]int64 `json:"accountTypeIds,omitempty"`
	// <strong>Type:</strong> AccountReference<br/> List of accounts for which access is requested from the bank. It must only be passed if the bank interface has the DETAILED_CONSENT property set.
	AccountReferences *[]AccountReference `json:"accountReferences,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'multiStepAuthentication' field instead.<br/><br/>Challenge response. This field should be set only when the previous attempt of importing the bank connection failed with HTTP code 510, i.e. the bank sent a challenge for the user for an additional authentication. In this case, this field must contain the response to the bank's challenge. Note that in the context of finAPI's Web Form flow, finAPI will automatically deal with getting the challenge response from the user via the Web Form.
	ChallengeResponse *string `json:"challengeResponse,omitempty"`
	// <strong>Type:</strong> MultiStepAuthenticationCallback<br/> Container for multi-step authentication data. Required when a previous service call initiated a multi-step authentication.
	MultiStepAuthentication *MultiStepAuthenticationCallback `json:"multiStepAuthentication,omitempty"`
	// Must only be passed when the used interface has the property REDIRECT_APPROACH. The user will be redirected to the given URL from the bank's website after completing the bank login and (possibly) the SCA.
	RedirectUrl *string `json:"redirectUrl,omitempty"`
}

// NewImportBankConnectionParams instantiates a new ImportBankConnectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportBankConnectionParams(bankId int64, ) *ImportBankConnectionParams {
	this := ImportBankConnectionParams{}
	this.BankId = bankId
	var storePin bool = false
	this.StorePin = &storePin
	var storeSecrets bool = false
	this.StoreSecrets = &storeSecrets
	var skipPositionsDownload bool = false
	this.SkipPositionsDownload = &skipPositionsDownload
	var loadOwnerData bool = false
	this.LoadOwnerData = &loadOwnerData
	var maxDaysForDownload int32 = 0
	this.MaxDaysForDownload = &maxDaysForDownload
	return &this
}

// NewImportBankConnectionParamsWithDefaults instantiates a new ImportBankConnectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportBankConnectionParamsWithDefaults() *ImportBankConnectionParams {
	this := ImportBankConnectionParams{}
	var storePin bool = false
	this.StorePin = &storePin
	var storeSecrets bool = false
	this.StoreSecrets = &storeSecrets
	var skipPositionsDownload bool = false
	this.SkipPositionsDownload = &skipPositionsDownload
	var loadOwnerData bool = false
	this.LoadOwnerData = &loadOwnerData
	var maxDaysForDownload int32 = 0
	this.MaxDaysForDownload = &maxDaysForDownload
	return &this
}

// GetBankId returns the BankId field value
func (o *ImportBankConnectionParams) GetBankId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetBankIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankId, true
}

// SetBankId sets field value
func (o *ImportBankConnectionParams) SetBankId(v int64) {
	o.BankId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImportBankConnectionParams) SetName(v string) {
	o.Name = &v
}

// GetBankingUserId returns the BankingUserId field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetBankingUserId() string {
	if o == nil || o.BankingUserId == nil {
		var ret string
		return ret
	}
	return *o.BankingUserId
}

// GetBankingUserIdOk returns a tuple with the BankingUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetBankingUserIdOk() (*string, bool) {
	if o == nil || o.BankingUserId == nil {
		return nil, false
	}
	return o.BankingUserId, true
}

// HasBankingUserId returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasBankingUserId() bool {
	if o != nil && o.BankingUserId != nil {
		return true
	}

	return false
}

// SetBankingUserId gets a reference to the given string and assigns it to the BankingUserId field.
func (o *ImportBankConnectionParams) SetBankingUserId(v string) {
	o.BankingUserId = &v
}

// GetBankingCustomerId returns the BankingCustomerId field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetBankingCustomerId() string {
	if o == nil || o.BankingCustomerId == nil {
		var ret string
		return ret
	}
	return *o.BankingCustomerId
}

// GetBankingCustomerIdOk returns a tuple with the BankingCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetBankingCustomerIdOk() (*string, bool) {
	if o == nil || o.BankingCustomerId == nil {
		return nil, false
	}
	return o.BankingCustomerId, true
}

// HasBankingCustomerId returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasBankingCustomerId() bool {
	if o != nil && o.BankingCustomerId != nil {
		return true
	}

	return false
}

// SetBankingCustomerId gets a reference to the given string and assigns it to the BankingCustomerId field.
func (o *ImportBankConnectionParams) SetBankingCustomerId(v string) {
	o.BankingCustomerId = &v
}

// GetBankingPin returns the BankingPin field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetBankingPin() string {
	if o == nil || o.BankingPin == nil {
		var ret string
		return ret
	}
	return *o.BankingPin
}

// GetBankingPinOk returns a tuple with the BankingPin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetBankingPinOk() (*string, bool) {
	if o == nil || o.BankingPin == nil {
		return nil, false
	}
	return o.BankingPin, true
}

// HasBankingPin returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasBankingPin() bool {
	if o != nil && o.BankingPin != nil {
		return true
	}

	return false
}

// SetBankingPin gets a reference to the given string and assigns it to the BankingPin field.
func (o *ImportBankConnectionParams) SetBankingPin(v string) {
	o.BankingPin = &v
}

// GetStorePin returns the StorePin field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetStorePin() bool {
	if o == nil || o.StorePin == nil {
		var ret bool
		return ret
	}
	return *o.StorePin
}

// GetStorePinOk returns a tuple with the StorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetStorePinOk() (*bool, bool) {
	if o == nil || o.StorePin == nil {
		return nil, false
	}
	return o.StorePin, true
}

// HasStorePin returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasStorePin() bool {
	if o != nil && o.StorePin != nil {
		return true
	}

	return false
}

// SetStorePin gets a reference to the given bool and assigns it to the StorePin field.
func (o *ImportBankConnectionParams) SetStorePin(v bool) {
	o.StorePin = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetInterface() BankingInterface {
	if o == nil || o.Interface == nil {
		var ret BankingInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetInterfaceOk() (*BankingInterface, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasInterface() bool {
	if o != nil && o.Interface != nil {
		return true
	}

	return false
}

// SetInterface gets a reference to the given BankingInterface and assigns it to the Interface field.
func (o *ImportBankConnectionParams) SetInterface(v BankingInterface) {
	o.Interface = &v
}

// GetLoginCredentials returns the LoginCredentials field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetLoginCredentials() []LoginCredential {
	if o == nil || o.LoginCredentials == nil {
		var ret []LoginCredential
		return ret
	}
	return *o.LoginCredentials
}

// GetLoginCredentialsOk returns a tuple with the LoginCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetLoginCredentialsOk() (*[]LoginCredential, bool) {
	if o == nil || o.LoginCredentials == nil {
		return nil, false
	}
	return o.LoginCredentials, true
}

// HasLoginCredentials returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasLoginCredentials() bool {
	if o != nil && o.LoginCredentials != nil {
		return true
	}

	return false
}

// SetLoginCredentials gets a reference to the given []LoginCredential and assigns it to the LoginCredentials field.
func (o *ImportBankConnectionParams) SetLoginCredentials(v []LoginCredential) {
	o.LoginCredentials = &v
}

// GetStoreSecrets returns the StoreSecrets field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetStoreSecrets() bool {
	if o == nil || o.StoreSecrets == nil {
		var ret bool
		return ret
	}
	return *o.StoreSecrets
}

// GetStoreSecretsOk returns a tuple with the StoreSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetStoreSecretsOk() (*bool, bool) {
	if o == nil || o.StoreSecrets == nil {
		return nil, false
	}
	return o.StoreSecrets, true
}

// HasStoreSecrets returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasStoreSecrets() bool {
	if o != nil && o.StoreSecrets != nil {
		return true
	}

	return false
}

// SetStoreSecrets gets a reference to the given bool and assigns it to the StoreSecrets field.
func (o *ImportBankConnectionParams) SetStoreSecrets(v bool) {
	o.StoreSecrets = &v
}

// GetSkipPositionsDownload returns the SkipPositionsDownload field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetSkipPositionsDownload() bool {
	if o == nil || o.SkipPositionsDownload == nil {
		var ret bool
		return ret
	}
	return *o.SkipPositionsDownload
}

// GetSkipPositionsDownloadOk returns a tuple with the SkipPositionsDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetSkipPositionsDownloadOk() (*bool, bool) {
	if o == nil || o.SkipPositionsDownload == nil {
		return nil, false
	}
	return o.SkipPositionsDownload, true
}

// HasSkipPositionsDownload returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasSkipPositionsDownload() bool {
	if o != nil && o.SkipPositionsDownload != nil {
		return true
	}

	return false
}

// SetSkipPositionsDownload gets a reference to the given bool and assigns it to the SkipPositionsDownload field.
func (o *ImportBankConnectionParams) SetSkipPositionsDownload(v bool) {
	o.SkipPositionsDownload = &v
}

// GetLoadOwnerData returns the LoadOwnerData field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetLoadOwnerData() bool {
	if o == nil || o.LoadOwnerData == nil {
		var ret bool
		return ret
	}
	return *o.LoadOwnerData
}

// GetLoadOwnerDataOk returns a tuple with the LoadOwnerData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetLoadOwnerDataOk() (*bool, bool) {
	if o == nil || o.LoadOwnerData == nil {
		return nil, false
	}
	return o.LoadOwnerData, true
}

// HasLoadOwnerData returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasLoadOwnerData() bool {
	if o != nil && o.LoadOwnerData != nil {
		return true
	}

	return false
}

// SetLoadOwnerData gets a reference to the given bool and assigns it to the LoadOwnerData field.
func (o *ImportBankConnectionParams) SetLoadOwnerData(v bool) {
	o.LoadOwnerData = &v
}

// GetMaxDaysForDownload returns the MaxDaysForDownload field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetMaxDaysForDownload() int32 {
	if o == nil || o.MaxDaysForDownload == nil {
		var ret int32
		return ret
	}
	return *o.MaxDaysForDownload
}

// GetMaxDaysForDownloadOk returns a tuple with the MaxDaysForDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetMaxDaysForDownloadOk() (*int32, bool) {
	if o == nil || o.MaxDaysForDownload == nil {
		return nil, false
	}
	return o.MaxDaysForDownload, true
}

// HasMaxDaysForDownload returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasMaxDaysForDownload() bool {
	if o != nil && o.MaxDaysForDownload != nil {
		return true
	}

	return false
}

// SetMaxDaysForDownload gets a reference to the given int32 and assigns it to the MaxDaysForDownload field.
func (o *ImportBankConnectionParams) SetMaxDaysForDownload(v int32) {
	o.MaxDaysForDownload = &v
}

// GetAccountTypes returns the AccountTypes field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetAccountTypes() []AccountType {
	if o == nil || o.AccountTypes == nil {
		var ret []AccountType
		return ret
	}
	return *o.AccountTypes
}

// GetAccountTypesOk returns a tuple with the AccountTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetAccountTypesOk() (*[]AccountType, bool) {
	if o == nil || o.AccountTypes == nil {
		return nil, false
	}
	return o.AccountTypes, true
}

// HasAccountTypes returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasAccountTypes() bool {
	if o != nil && o.AccountTypes != nil {
		return true
	}

	return false
}

// SetAccountTypes gets a reference to the given []AccountType and assigns it to the AccountTypes field.
func (o *ImportBankConnectionParams) SetAccountTypes(v []AccountType) {
	o.AccountTypes = &v
}

// GetAccountTypeIds returns the AccountTypeIds field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetAccountTypeIds() []int64 {
	if o == nil || o.AccountTypeIds == nil {
		var ret []int64
		return ret
	}
	return *o.AccountTypeIds
}

// GetAccountTypeIdsOk returns a tuple with the AccountTypeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetAccountTypeIdsOk() (*[]int64, bool) {
	if o == nil || o.AccountTypeIds == nil {
		return nil, false
	}
	return o.AccountTypeIds, true
}

// HasAccountTypeIds returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasAccountTypeIds() bool {
	if o != nil && o.AccountTypeIds != nil {
		return true
	}

	return false
}

// SetAccountTypeIds gets a reference to the given []int64 and assigns it to the AccountTypeIds field.
func (o *ImportBankConnectionParams) SetAccountTypeIds(v []int64) {
	o.AccountTypeIds = &v
}

// GetAccountReferences returns the AccountReferences field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetAccountReferences() []AccountReference {
	if o == nil || o.AccountReferences == nil {
		var ret []AccountReference
		return ret
	}
	return *o.AccountReferences
}

// GetAccountReferencesOk returns a tuple with the AccountReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetAccountReferencesOk() (*[]AccountReference, bool) {
	if o == nil || o.AccountReferences == nil {
		return nil, false
	}
	return o.AccountReferences, true
}

// HasAccountReferences returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasAccountReferences() bool {
	if o != nil && o.AccountReferences != nil {
		return true
	}

	return false
}

// SetAccountReferences gets a reference to the given []AccountReference and assigns it to the AccountReferences field.
func (o *ImportBankConnectionParams) SetAccountReferences(v []AccountReference) {
	o.AccountReferences = &v
}

// GetChallengeResponse returns the ChallengeResponse field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetChallengeResponse() string {
	if o == nil || o.ChallengeResponse == nil {
		var ret string
		return ret
	}
	return *o.ChallengeResponse
}

// GetChallengeResponseOk returns a tuple with the ChallengeResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetChallengeResponseOk() (*string, bool) {
	if o == nil || o.ChallengeResponse == nil {
		return nil, false
	}
	return o.ChallengeResponse, true
}

// HasChallengeResponse returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasChallengeResponse() bool {
	if o != nil && o.ChallengeResponse != nil {
		return true
	}

	return false
}

// SetChallengeResponse gets a reference to the given string and assigns it to the ChallengeResponse field.
func (o *ImportBankConnectionParams) SetChallengeResponse(v string) {
	o.ChallengeResponse = &v
}

// GetMultiStepAuthentication returns the MultiStepAuthentication field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetMultiStepAuthentication() MultiStepAuthenticationCallback {
	if o == nil || o.MultiStepAuthentication == nil {
		var ret MultiStepAuthenticationCallback
		return ret
	}
	return *o.MultiStepAuthentication
}

// GetMultiStepAuthenticationOk returns a tuple with the MultiStepAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetMultiStepAuthenticationOk() (*MultiStepAuthenticationCallback, bool) {
	if o == nil || o.MultiStepAuthentication == nil {
		return nil, false
	}
	return o.MultiStepAuthentication, true
}

// HasMultiStepAuthentication returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasMultiStepAuthentication() bool {
	if o != nil && o.MultiStepAuthentication != nil {
		return true
	}

	return false
}

// SetMultiStepAuthentication gets a reference to the given MultiStepAuthenticationCallback and assigns it to the MultiStepAuthentication field.
func (o *ImportBankConnectionParams) SetMultiStepAuthentication(v MultiStepAuthenticationCallback) {
	o.MultiStepAuthentication = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *ImportBankConnectionParams) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBankConnectionParams) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *ImportBankConnectionParams) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl != nil {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *ImportBankConnectionParams) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

func (o ImportBankConnectionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bankId"] = o.BankId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.BankingUserId != nil {
		toSerialize["bankingUserId"] = o.BankingUserId
	}
	if o.BankingCustomerId != nil {
		toSerialize["bankingCustomerId"] = o.BankingCustomerId
	}
	if o.BankingPin != nil {
		toSerialize["bankingPin"] = o.BankingPin
	}
	if o.StorePin != nil {
		toSerialize["storePin"] = o.StorePin
	}
	if o.Interface != nil {
		toSerialize["interface"] = o.Interface
	}
	if o.LoginCredentials != nil {
		toSerialize["loginCredentials"] = o.LoginCredentials
	}
	if o.StoreSecrets != nil {
		toSerialize["storeSecrets"] = o.StoreSecrets
	}
	if o.SkipPositionsDownload != nil {
		toSerialize["skipPositionsDownload"] = o.SkipPositionsDownload
	}
	if o.LoadOwnerData != nil {
		toSerialize["loadOwnerData"] = o.LoadOwnerData
	}
	if o.MaxDaysForDownload != nil {
		toSerialize["maxDaysForDownload"] = o.MaxDaysForDownload
	}
	if o.AccountTypes != nil {
		toSerialize["accountTypes"] = o.AccountTypes
	}
	if o.AccountTypeIds != nil {
		toSerialize["accountTypeIds"] = o.AccountTypeIds
	}
	if o.AccountReferences != nil {
		toSerialize["accountReferences"] = o.AccountReferences
	}
	if o.ChallengeResponse != nil {
		toSerialize["challengeResponse"] = o.ChallengeResponse
	}
	if o.MultiStepAuthentication != nil {
		toSerialize["multiStepAuthentication"] = o.MultiStepAuthentication
	}
	if o.RedirectUrl != nil {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	return json.Marshal(toSerialize)
}

type NullableImportBankConnectionParams struct {
	value *ImportBankConnectionParams
	isSet bool
}

func (v NullableImportBankConnectionParams) Get() *ImportBankConnectionParams {
	return v.value
}

func (v *NullableImportBankConnectionParams) Set(val *ImportBankConnectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableImportBankConnectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableImportBankConnectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportBankConnectionParams(val *ImportBankConnectionParams) *NullableImportBankConnectionParams {
	return &NullableImportBankConnectionParams{value: val, isSet: true}
}

func (v NullableImportBankConnectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportBankConnectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



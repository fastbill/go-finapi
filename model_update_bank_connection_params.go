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

// UpdateBankConnectionParams Container for bank connection update parameters
type UpdateBankConnectionParams struct {
	// Bank connection identifier
	BankConnectionId int64 `json:"bankConnectionId"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' and 'interface' fields instead. If any of those two fields is used, then the value of this field will be ignored.<br><br>Online banking PIN. Any symbols are allowed. Max length: 170. If a PIN is stored in the bank connection, then this field may remain unset. If finAPI's Web Form is not required and the field is set though then it will always be used (even if there is some other PIN stored in the bank connection). If you want the user to enter a PIN in finAPI's Web Form even when a PIN is stored, then just set the field to any value, so that the service recognizes that you wish to use the Web Form flow.
	BankingPin *string `json:"bankingPin,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'storeSecrets' field instead.<br><br>Whether to store the PIN. If the PIN is stored, it is not required to pass the PIN again when updating this bank connection or when executing orders (like money transfers). Default is false. <br/><br/>NOTES:<br/> - before you set this field to true, please regard the 'pinsAreVolatile' flag of this connection's bank. Please note that volatile credentials will not be stored, even if provided, to enforce user involvement in the next communication with the bank;<br/> - this field is ignored in case when the user will need to use finAPI's Web Form. The user will be able to decide whether to store the PIN or not in the Web Form, depending on the 'storeSecretsAvailableInWebForm' setting (see Client Configuration).
	StorePin *bool `json:"storePin,omitempty"`
	// <strong>Type:</strong> BankingInterface<br/> The interface to use for connecting with the bank.
	Interface *BankingInterface `json:"interface,omitempty"`
	// <strong>Type:</strong> LoginCredential<br/> Set of login credentials. Must be passed in combination with the 'interface' field, if the credentials have not been previously stored. The labels that you pass must match with the login credential labels that the respective interface defines. finAPI will combine the given credentials with any credentials that it has stored. You can leave this field unset in case finAPI has stored all required credentials.
	LoginCredentials *[]LoginCredential `json:"loginCredentials,omitempty"`
	// Whether to store the secret login fields. If the secret fields are stored, then updates can be triggered without the involvement of the users, as long as the credentials remain valid and the bank consent has not expired. Note that bank consent will be stored regardless of the field value. Default value is false.
	StoreSecrets *bool `json:"storeSecrets,omitempty"`
	// Whether new accounts that have not yet been imported will be imported or not. Default is false. <br/><br/>NOTES:<br/>&bull; For best performance of the bank connection update, you should not enable this flag unless you really expect new accounts to be available in the connection. It is recommended to let your users tell you through your application when they want the service to look for new accounts.<br/>&bull; If you have imported an interface using a limited set of account types, you would import all other accounts (e.g. security accounts or credit cards) by setting <code>importNewAccounts</code> to <code>true</code>. To avoid importing account types that you are not interested in, make sure to keep this field unset (or set to false) for the update.
	ImportNewAccounts *bool `json:"importNewAccounts,omitempty"`
	// Whether to skip the download of transactions and securities or not. If set to true, then finAPI will download just the accounts list with the accounts' information (like account name, number, holder, etc), as well as the accounts' balances (if possible), but skip the download of transactions and securities. Default is false.<br/><br/>NOTES:<br/>&bull; Setting this flag to true is only meant to be used if A) you generally never download positions, because you are only interested in the accounts list and/or balances, or B) you want to get just the list of accounts in a first step, and then delete unwanted accounts from the bank connection, before you trigger another update with transactions download. This approach allows you to download transactions only for the accounts that you want.<br/>&bull; If you skip the download of transactions and securities during an import or update, you can still download them on a later update (though you might not get all positions at a later point, because the date range in which the bank servers provide this data is usually limited).<br/>&bull; If an account already had any positions imported before an update, and you skip the positions download in the update, then the account's updated balance might not add up to the set of transactions / security positions. Be aware that certain services (like GET /accounts/dailyBalances) may give incorrect results for accounts in such a state.<br/>&bull; If this bank connection is updated via finAPI's automatic batch update, then transactions and security positions (of already imported accounts) <u>will</u> be downloaded in any case!<br/>&bull; For security accounts, skipping the downloading of the securities might result in the account's balance also not being downloaded.<br/>&bull; For Bausparen accounts, this field is ignored. finAPI will always download transactions for Bausparen accounts.<br/>
	SkipPositionsDownload *bool `json:"skipPositionsDownload,omitempty"`
	// Whether to load/refresh information about the bank connection owner(s) - see field 'owners'. Default value is 'false'. Note that owner data is NOT loaded/refreshed during finAPI's automatic bank connection update.
	LoadOwnerData *bool `json:"loadOwnerData,omitempty"`
	AccountTypes *[]AccountType `json:"accountTypes,omitempty"`
	// <strong>Type:</strong> AccountReference<br/> List of accounts for which access is requested from the bank. It may only be passed if the bank interface has the DETAILED_CONSENT property set. if omitted, finAPI will use the list of existing accounts. Note that the parameter is still required if you want to import new accounts (i.e. call with importNewAccounts=true).
	AccountReferences *[]AccountReference `json:"accountReferences,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'multiStepAuthentication' field instead.<br/><br/>Challenge response. This field should be set only when the previous attempt of update the bank connection failed with HTTP code 510, i.e. the bank sent a challenge for the user for an additional authentication. In this case, this field must contain the response to the bank's challenge. Note that in the context of finAPI's Web Form flow, finAPI will automatically deal with getting the challenge response from the user via the Web Form.
	ChallengeResponse *string `json:"challengeResponse,omitempty"`
	// Must only be passed when the used interface has the property REDIRECT_APPROACH. The user will be redirected to the given URL from the bank's website after completing the bank login and (possibly) the SCA.
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer <a href='?product=web_form_2.0#post-/api/tasks/backgroundUpdate' target='_blank'>here</a> to the 'editSavedSettings' field instead.<br/><br/>If the user has stored credentials in finAPI for the selected bank connection, then the finAPI Web Form will only be shown when the user must be involved for a second authentication, or when the previous connection to the bank via the selected interface had failed. However if you want to provide the Web Form to the user in any case, you can set this field to true. It will force the Web Form flow for the user and allow him to make changes to the data that he has stored in finAPI. Default value is 'false'.
	ForceWebForm *bool `json:"forceWebForm,omitempty"`
	// <strong>Type:</strong> MultiStepAuthenticationCallback<br/> Container for multi-step authentication data. Required when a previous service call initiated a multi-step authentication.
	MultiStepAuthentication *MultiStepAuthenticationCallback `json:"multiStepAuthentication,omitempty"`
}

// NewUpdateBankConnectionParams instantiates a new UpdateBankConnectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBankConnectionParams(bankConnectionId int64, ) *UpdateBankConnectionParams {
	this := UpdateBankConnectionParams{}
	this.BankConnectionId = bankConnectionId
	var storePin bool = false
	this.StorePin = &storePin
	var storeSecrets bool = false
	this.StoreSecrets = &storeSecrets
	var importNewAccounts bool = false
	this.ImportNewAccounts = &importNewAccounts
	var skipPositionsDownload bool = false
	this.SkipPositionsDownload = &skipPositionsDownload
	var loadOwnerData bool = false
	this.LoadOwnerData = &loadOwnerData
	var forceWebForm bool = false
	this.ForceWebForm = &forceWebForm
	return &this
}

// NewUpdateBankConnectionParamsWithDefaults instantiates a new UpdateBankConnectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBankConnectionParamsWithDefaults() *UpdateBankConnectionParams {
	this := UpdateBankConnectionParams{}
	var storePin bool = false
	this.StorePin = &storePin
	var storeSecrets bool = false
	this.StoreSecrets = &storeSecrets
	var importNewAccounts bool = false
	this.ImportNewAccounts = &importNewAccounts
	var skipPositionsDownload bool = false
	this.SkipPositionsDownload = &skipPositionsDownload
	var loadOwnerData bool = false
	this.LoadOwnerData = &loadOwnerData
	var forceWebForm bool = false
	this.ForceWebForm = &forceWebForm
	return &this
}

// GetBankConnectionId returns the BankConnectionId field value
func (o *UpdateBankConnectionParams) GetBankConnectionId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.BankConnectionId
}

// GetBankConnectionIdOk returns a tuple with the BankConnectionId field value
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetBankConnectionIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankConnectionId, true
}

// SetBankConnectionId sets field value
func (o *UpdateBankConnectionParams) SetBankConnectionId(v int64) {
	o.BankConnectionId = v
}

// GetBankingPin returns the BankingPin field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetBankingPin() string {
	if o == nil || o.BankingPin == nil {
		var ret string
		return ret
	}
	return *o.BankingPin
}

// GetBankingPinOk returns a tuple with the BankingPin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetBankingPinOk() (*string, bool) {
	if o == nil || o.BankingPin == nil {
		return nil, false
	}
	return o.BankingPin, true
}

// HasBankingPin returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasBankingPin() bool {
	if o != nil && o.BankingPin != nil {
		return true
	}

	return false
}

// SetBankingPin gets a reference to the given string and assigns it to the BankingPin field.
func (o *UpdateBankConnectionParams) SetBankingPin(v string) {
	o.BankingPin = &v
}

// GetStorePin returns the StorePin field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetStorePin() bool {
	if o == nil || o.StorePin == nil {
		var ret bool
		return ret
	}
	return *o.StorePin
}

// GetStorePinOk returns a tuple with the StorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetStorePinOk() (*bool, bool) {
	if o == nil || o.StorePin == nil {
		return nil, false
	}
	return o.StorePin, true
}

// HasStorePin returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasStorePin() bool {
	if o != nil && o.StorePin != nil {
		return true
	}

	return false
}

// SetStorePin gets a reference to the given bool and assigns it to the StorePin field.
func (o *UpdateBankConnectionParams) SetStorePin(v bool) {
	o.StorePin = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetInterface() BankingInterface {
	if o == nil || o.Interface == nil {
		var ret BankingInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetInterfaceOk() (*BankingInterface, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasInterface() bool {
	if o != nil && o.Interface != nil {
		return true
	}

	return false
}

// SetInterface gets a reference to the given BankingInterface and assigns it to the Interface field.
func (o *UpdateBankConnectionParams) SetInterface(v BankingInterface) {
	o.Interface = &v
}

// GetLoginCredentials returns the LoginCredentials field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetLoginCredentials() []LoginCredential {
	if o == nil || o.LoginCredentials == nil {
		var ret []LoginCredential
		return ret
	}
	return *o.LoginCredentials
}

// GetLoginCredentialsOk returns a tuple with the LoginCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetLoginCredentialsOk() (*[]LoginCredential, bool) {
	if o == nil || o.LoginCredentials == nil {
		return nil, false
	}
	return o.LoginCredentials, true
}

// HasLoginCredentials returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasLoginCredentials() bool {
	if o != nil && o.LoginCredentials != nil {
		return true
	}

	return false
}

// SetLoginCredentials gets a reference to the given []LoginCredential and assigns it to the LoginCredentials field.
func (o *UpdateBankConnectionParams) SetLoginCredentials(v []LoginCredential) {
	o.LoginCredentials = &v
}

// GetStoreSecrets returns the StoreSecrets field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetStoreSecrets() bool {
	if o == nil || o.StoreSecrets == nil {
		var ret bool
		return ret
	}
	return *o.StoreSecrets
}

// GetStoreSecretsOk returns a tuple with the StoreSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetStoreSecretsOk() (*bool, bool) {
	if o == nil || o.StoreSecrets == nil {
		return nil, false
	}
	return o.StoreSecrets, true
}

// HasStoreSecrets returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasStoreSecrets() bool {
	if o != nil && o.StoreSecrets != nil {
		return true
	}

	return false
}

// SetStoreSecrets gets a reference to the given bool and assigns it to the StoreSecrets field.
func (o *UpdateBankConnectionParams) SetStoreSecrets(v bool) {
	o.StoreSecrets = &v
}

// GetImportNewAccounts returns the ImportNewAccounts field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetImportNewAccounts() bool {
	if o == nil || o.ImportNewAccounts == nil {
		var ret bool
		return ret
	}
	return *o.ImportNewAccounts
}

// GetImportNewAccountsOk returns a tuple with the ImportNewAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetImportNewAccountsOk() (*bool, bool) {
	if o == nil || o.ImportNewAccounts == nil {
		return nil, false
	}
	return o.ImportNewAccounts, true
}

// HasImportNewAccounts returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasImportNewAccounts() bool {
	if o != nil && o.ImportNewAccounts != nil {
		return true
	}

	return false
}

// SetImportNewAccounts gets a reference to the given bool and assigns it to the ImportNewAccounts field.
func (o *UpdateBankConnectionParams) SetImportNewAccounts(v bool) {
	o.ImportNewAccounts = &v
}

// GetSkipPositionsDownload returns the SkipPositionsDownload field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetSkipPositionsDownload() bool {
	if o == nil || o.SkipPositionsDownload == nil {
		var ret bool
		return ret
	}
	return *o.SkipPositionsDownload
}

// GetSkipPositionsDownloadOk returns a tuple with the SkipPositionsDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetSkipPositionsDownloadOk() (*bool, bool) {
	if o == nil || o.SkipPositionsDownload == nil {
		return nil, false
	}
	return o.SkipPositionsDownload, true
}

// HasSkipPositionsDownload returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasSkipPositionsDownload() bool {
	if o != nil && o.SkipPositionsDownload != nil {
		return true
	}

	return false
}

// SetSkipPositionsDownload gets a reference to the given bool and assigns it to the SkipPositionsDownload field.
func (o *UpdateBankConnectionParams) SetSkipPositionsDownload(v bool) {
	o.SkipPositionsDownload = &v
}

// GetLoadOwnerData returns the LoadOwnerData field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetLoadOwnerData() bool {
	if o == nil || o.LoadOwnerData == nil {
		var ret bool
		return ret
	}
	return *o.LoadOwnerData
}

// GetLoadOwnerDataOk returns a tuple with the LoadOwnerData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetLoadOwnerDataOk() (*bool, bool) {
	if o == nil || o.LoadOwnerData == nil {
		return nil, false
	}
	return o.LoadOwnerData, true
}

// HasLoadOwnerData returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasLoadOwnerData() bool {
	if o != nil && o.LoadOwnerData != nil {
		return true
	}

	return false
}

// SetLoadOwnerData gets a reference to the given bool and assigns it to the LoadOwnerData field.
func (o *UpdateBankConnectionParams) SetLoadOwnerData(v bool) {
	o.LoadOwnerData = &v
}

// GetAccountTypes returns the AccountTypes field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetAccountTypes() []AccountType {
	if o == nil || o.AccountTypes == nil {
		var ret []AccountType
		return ret
	}
	return *o.AccountTypes
}

// GetAccountTypesOk returns a tuple with the AccountTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetAccountTypesOk() (*[]AccountType, bool) {
	if o == nil || o.AccountTypes == nil {
		return nil, false
	}
	return o.AccountTypes, true
}

// HasAccountTypes returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasAccountTypes() bool {
	if o != nil && o.AccountTypes != nil {
		return true
	}

	return false
}

// SetAccountTypes gets a reference to the given []AccountType and assigns it to the AccountTypes field.
func (o *UpdateBankConnectionParams) SetAccountTypes(v []AccountType) {
	o.AccountTypes = &v
}

// GetAccountReferences returns the AccountReferences field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetAccountReferences() []AccountReference {
	if o == nil || o.AccountReferences == nil {
		var ret []AccountReference
		return ret
	}
	return *o.AccountReferences
}

// GetAccountReferencesOk returns a tuple with the AccountReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetAccountReferencesOk() (*[]AccountReference, bool) {
	if o == nil || o.AccountReferences == nil {
		return nil, false
	}
	return o.AccountReferences, true
}

// HasAccountReferences returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasAccountReferences() bool {
	if o != nil && o.AccountReferences != nil {
		return true
	}

	return false
}

// SetAccountReferences gets a reference to the given []AccountReference and assigns it to the AccountReferences field.
func (o *UpdateBankConnectionParams) SetAccountReferences(v []AccountReference) {
	o.AccountReferences = &v
}

// GetChallengeResponse returns the ChallengeResponse field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetChallengeResponse() string {
	if o == nil || o.ChallengeResponse == nil {
		var ret string
		return ret
	}
	return *o.ChallengeResponse
}

// GetChallengeResponseOk returns a tuple with the ChallengeResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetChallengeResponseOk() (*string, bool) {
	if o == nil || o.ChallengeResponse == nil {
		return nil, false
	}
	return o.ChallengeResponse, true
}

// HasChallengeResponse returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasChallengeResponse() bool {
	if o != nil && o.ChallengeResponse != nil {
		return true
	}

	return false
}

// SetChallengeResponse gets a reference to the given string and assigns it to the ChallengeResponse field.
func (o *UpdateBankConnectionParams) SetChallengeResponse(v string) {
	o.ChallengeResponse = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl != nil {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *UpdateBankConnectionParams) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetForceWebForm returns the ForceWebForm field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetForceWebForm() bool {
	if o == nil || o.ForceWebForm == nil {
		var ret bool
		return ret
	}
	return *o.ForceWebForm
}

// GetForceWebFormOk returns a tuple with the ForceWebForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetForceWebFormOk() (*bool, bool) {
	if o == nil || o.ForceWebForm == nil {
		return nil, false
	}
	return o.ForceWebForm, true
}

// HasForceWebForm returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasForceWebForm() bool {
	if o != nil && o.ForceWebForm != nil {
		return true
	}

	return false
}

// SetForceWebForm gets a reference to the given bool and assigns it to the ForceWebForm field.
func (o *UpdateBankConnectionParams) SetForceWebForm(v bool) {
	o.ForceWebForm = &v
}

// GetMultiStepAuthentication returns the MultiStepAuthentication field value if set, zero value otherwise.
func (o *UpdateBankConnectionParams) GetMultiStepAuthentication() MultiStepAuthenticationCallback {
	if o == nil || o.MultiStepAuthentication == nil {
		var ret MultiStepAuthenticationCallback
		return ret
	}
	return *o.MultiStepAuthentication
}

// GetMultiStepAuthenticationOk returns a tuple with the MultiStepAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBankConnectionParams) GetMultiStepAuthenticationOk() (*MultiStepAuthenticationCallback, bool) {
	if o == nil || o.MultiStepAuthentication == nil {
		return nil, false
	}
	return o.MultiStepAuthentication, true
}

// HasMultiStepAuthentication returns a boolean if a field has been set.
func (o *UpdateBankConnectionParams) HasMultiStepAuthentication() bool {
	if o != nil && o.MultiStepAuthentication != nil {
		return true
	}

	return false
}

// SetMultiStepAuthentication gets a reference to the given MultiStepAuthenticationCallback and assigns it to the MultiStepAuthentication field.
func (o *UpdateBankConnectionParams) SetMultiStepAuthentication(v MultiStepAuthenticationCallback) {
	o.MultiStepAuthentication = &v
}

func (o UpdateBankConnectionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bankConnectionId"] = o.BankConnectionId
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
	if o.ImportNewAccounts != nil {
		toSerialize["importNewAccounts"] = o.ImportNewAccounts
	}
	if o.SkipPositionsDownload != nil {
		toSerialize["skipPositionsDownload"] = o.SkipPositionsDownload
	}
	if o.LoadOwnerData != nil {
		toSerialize["loadOwnerData"] = o.LoadOwnerData
	}
	if o.AccountTypes != nil {
		toSerialize["accountTypes"] = o.AccountTypes
	}
	if o.AccountReferences != nil {
		toSerialize["accountReferences"] = o.AccountReferences
	}
	if o.ChallengeResponse != nil {
		toSerialize["challengeResponse"] = o.ChallengeResponse
	}
	if o.RedirectUrl != nil {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if o.ForceWebForm != nil {
		toSerialize["forceWebForm"] = o.ForceWebForm
	}
	if o.MultiStepAuthentication != nil {
		toSerialize["multiStepAuthentication"] = o.MultiStepAuthentication
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateBankConnectionParams struct {
	value *UpdateBankConnectionParams
	isSet bool
}

func (v NullableUpdateBankConnectionParams) Get() *UpdateBankConnectionParams {
	return v.value
}

func (v *NullableUpdateBankConnectionParams) Set(val *UpdateBankConnectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBankConnectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBankConnectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBankConnectionParams(val *UpdateBankConnectionParams) *NullableUpdateBankConnectionParams {
	return &NullableUpdateBankConnectionParams{value: val, isSet: true}
}

func (v NullableUpdateBankConnectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBankConnectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



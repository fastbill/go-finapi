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

// BankConnection Container for a bank connection's data
type BankConnection struct {
	// Bank connection identifier
	Id int64 `json:"id"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'bank' field instead.<br/><br/>Identifier of the bank that this connection belongs to. 
	BankId int64 `json:"bankId"`
	// Custom name for the bank connection. You can set this field with the 'Edit a bank connection' service, as well as during the initial import of the bank connection. Maximum length is 64.
	Name NullableString `json:"name"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' in the 'interfaces' instead.<br/><br/>Stored online banking user ID credential. This field may be null for the 'demo connection'. If your client has no license for processing banking credentials then a banking user ID will always be 'XXXXX'
	BankingUserId NullableString `json:"bankingUserId"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' in the 'interfaces' instead.<br/><br/>Stored online banking customer ID credential. If your client has no license for processing banking credentials or if this field contains a value that requires password protection (see field 'isCustomerIdPassword' in Bank Resource) then the banking customer ID will always be 'XXXXX
	BankingCustomerId NullableString `json:"bankingCustomerId"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' in the 'interfaces' instead.<br/><br/>Stored online banking PIN. If a PIN is stored, this will always be 'XXXXX'
	BankingPin NullableString `json:"bankingPin"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Bank connection type
	Type string `json:"type"`
	// Current status of data download (account balances and transactions/securities). The POST /bankConnections/import and POST /bankConnections/<id>/update services will set this flag to IN_PROGRESS before they return. Once the import or update has finished, the status will be changed to READY.
	UpdateStatus string `json:"updateStatus"`
	// <strong>Type:</strong> CategorizationStatus<br/> Current status of transaction categorization. The asynchronous download process that is triggered by a call of the POST /bankConnections/import and POST /bankConnections/<id>/update services (and also by finAPI's auto update, if enabled) will set this flag to PENDING once the download has finished and a categorization is scheduled for the imported transactions. A separate categorization thread will then start to categorize the transactions (during this process, the status is IN_PROGRESS). When categorization has finished, the status will be (re-)set to READY. Note that the current categorization status should only be queried after the download has finished, i.e. once the download status has switched from IN_PROGRESS to READY.
	CategorizationStatus CategorizationStatus `json:"categorizationStatus"`
	// <strong>Type:</strong> UpdateResult<br/> THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the corresponding field in 'interfaces' instead.<br/><br/>Result of the last manual update of this bank connection. If no manual update has ever been done so far, then this field will not be set.
	LastManualUpdate NullableUpdateResult `json:"lastManualUpdate"`
	// <strong>Type:</strong> UpdateResult<br/> THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the corresponding field in 'interfaces' instead.<br/><br/>Result of the last auto update of this bank connection (ran by finAPI's automatic batch update process). If no auto update has ever been done so far, then this field will not be set.
	LastAutoUpdate NullableUpdateResult `json:"lastAutoUpdate"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'capabilities' field of the 'interfaces' in the Account resource instead.<br/><br/>Whether this bank connection accepts money transfer requests where the recipient's account is defined just by the IBAN (without an additional BIC). This field is re-evaluated each time this bank connection is updated. <br/>See also: /accounts/requestSepaMoneyTransfer
	IbanOnlyMoneyTransferSupported bool `json:"ibanOnlyMoneyTransferSupported"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'capabilities' field of the 'interfaces' in the Account resource instead.<br/><br/>Whether this bank connection accepts direct debit requests where the debitor's account is defined just by the IBAN (without an additional BIC). This field is re-evaluated each time this bank connection is updated. <br/>See also: /accounts/requestSepaDirectDebit
	IbanOnlyDirectDebitSupported bool `json:"ibanOnlyDirectDebitSupported"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'capabilities' field of the 'interfaces' in the Account resource instead.<br/><br/>Whether this bank connection supports submitting collective money transfers. This field is re-evaluated each time this bank connection is updated. <br/>See also: /accounts/requestSepaMoneyTransfer
	CollectiveMoneyTransferSupported bool `json:"collectiveMoneyTransferSupported"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the corresponding field in 'interfaces' instead.<br/><br/>The default two-step-procedure. Must match one of the available 'procedureId's from the 'twoStepProcedures' list. When this field is set, then finAPI will automatically try to select the procedure wherever applicable. Note that the list of available procedures of a bank connection may change as a result of an update of the connection, and if this field references a procedure that is no longer available after an update, finAPI will automatically clear the default procedure (set it to null).
	DefaultTwoStepProcedureId NullableString `json:"defaultTwoStepProcedureId"`
	// <strong>Type:</strong> TwoStepProcedure<br/> THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the corresponding field in 'interfaces' instead.<br/><br/>Available two-step-procedures for this bank connection, used for submitting a money transfer or direct debit request (see /accounts/requestSepaMoneyTransfer or /requestSepaDirectDebit). The available two-step-procedures are re-evaluated each time this bank connection is updated (/bankConnections/update). This means that this list may change as a result of an update.
	TwoStepProcedures []TwoStepProcedure `json:"twoStepProcedures"`
	// <strong>Type:</strong> BankConnectionInterface<br/> Set of interfaces that are connected for this bank connection.
	Interfaces []BankConnectionInterface `json:"interfaces"`
	// Identifiers of the accounts that belong to this bank connection
	AccountIds []int64 `json:"accountIds"`
	// <strong>Type:</strong> BankConnectionOwner<br/> Information about the owner(s) of the bank connection
	Owners []BankConnectionOwner `json:"owners"`
	// <strong>Type:</strong> Bank<br/> Bank that this connection belongs to
	Bank Bank `json:"bank"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'userActionRequired' field of the BankConnectionInterface resource instead.<br/><br/>This field indicates whether the last communication with the bank failed with an error that requires the user's attention or multi-step authentication error. If 'furtherLoginNotRecommended' is true, finAPI will stop auto updates of this bank connection to mitigate the risk of the user's bank account getting locked by the bank. Every communication with the bank (via updates, money_transfers, direct debits. etc.) can change the value of this flag. If this field is true, we recommend the user to check his credentials and try a manual update of the bank connection. If the update is successful without any multi-step authentication error, the 'furtherLoginNotRecommended' field will be set to false and the bank connection will be reincluded in the next batch update process. A manual update of the bank connection with incorrect credentials or if multi-step authentication error happens will set this field to true and lead to the exclusion of the bank connection from the following batch updates.
	FurtherLoginNotRecommended bool `json:"furtherLoginNotRecommended"`
}

// NewBankConnection instantiates a new BankConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankConnection(id int64, bankId int64, name NullableString, bankingUserId NullableString, bankingCustomerId NullableString, bankingPin NullableString, type_ string, updateStatus string, categorizationStatus CategorizationStatus, lastManualUpdate NullableUpdateResult, lastAutoUpdate NullableUpdateResult, ibanOnlyMoneyTransferSupported bool, ibanOnlyDirectDebitSupported bool, collectiveMoneyTransferSupported bool, defaultTwoStepProcedureId NullableString, twoStepProcedures []TwoStepProcedure, interfaces []BankConnectionInterface, accountIds []int64, owners []BankConnectionOwner, bank Bank, furtherLoginNotRecommended bool, ) *BankConnection {
	this := BankConnection{}
	this.Id = id
	this.BankId = bankId
	this.Name = name
	this.BankingUserId = bankingUserId
	this.BankingCustomerId = bankingCustomerId
	this.BankingPin = bankingPin
	this.Type = type_
	this.UpdateStatus = updateStatus
	this.CategorizationStatus = categorizationStatus
	this.LastManualUpdate = lastManualUpdate
	this.LastAutoUpdate = lastAutoUpdate
	this.IbanOnlyMoneyTransferSupported = ibanOnlyMoneyTransferSupported
	this.IbanOnlyDirectDebitSupported = ibanOnlyDirectDebitSupported
	this.CollectiveMoneyTransferSupported = collectiveMoneyTransferSupported
	this.DefaultTwoStepProcedureId = defaultTwoStepProcedureId
	this.TwoStepProcedures = twoStepProcedures
	this.Interfaces = interfaces
	this.AccountIds = accountIds
	this.Owners = owners
	this.Bank = bank
	this.FurtherLoginNotRecommended = furtherLoginNotRecommended
	return &this
}

// NewBankConnectionWithDefaults instantiates a new BankConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankConnectionWithDefaults() *BankConnection {
	this := BankConnection{}
	return &this
}

// GetId returns the Id field value
func (o *BankConnection) GetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BankConnection) SetId(v int64) {
	o.Id = v
}

// GetBankId returns the BankId field value
func (o *BankConnection) GetBankId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetBankIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankId, true
}

// SetBankId sets field value
func (o *BankConnection) SetBankId(v int64) {
	o.BankId = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BankConnection) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnection) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *BankConnection) SetName(v string) {
	o.Name.Set(&v)
}

// GetBankingUserId returns the BankingUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BankConnection) GetBankingUserId() string {
	if o == nil || o.BankingUserId.Get() == nil {
		var ret string
		return ret
	}

	return *o.BankingUserId.Get()
}

// GetBankingUserIdOk returns a tuple with the BankingUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnection) GetBankingUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankingUserId.Get(), o.BankingUserId.IsSet()
}

// SetBankingUserId sets field value
func (o *BankConnection) SetBankingUserId(v string) {
	o.BankingUserId.Set(&v)
}

// GetBankingCustomerId returns the BankingCustomerId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BankConnection) GetBankingCustomerId() string {
	if o == nil || o.BankingCustomerId.Get() == nil {
		var ret string
		return ret
	}

	return *o.BankingCustomerId.Get()
}

// GetBankingCustomerIdOk returns a tuple with the BankingCustomerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnection) GetBankingCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankingCustomerId.Get(), o.BankingCustomerId.IsSet()
}

// SetBankingCustomerId sets field value
func (o *BankConnection) SetBankingCustomerId(v string) {
	o.BankingCustomerId.Set(&v)
}

// GetBankingPin returns the BankingPin field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BankConnection) GetBankingPin() string {
	if o == nil || o.BankingPin.Get() == nil {
		var ret string
		return ret
	}

	return *o.BankingPin.Get()
}

// GetBankingPinOk returns a tuple with the BankingPin field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnection) GetBankingPinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankingPin.Get(), o.BankingPin.IsSet()
}

// SetBankingPin sets field value
func (o *BankConnection) SetBankingPin(v string) {
	o.BankingPin.Set(&v)
}

// GetType returns the Type field value
func (o *BankConnection) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BankConnection) SetType(v string) {
	o.Type = v
}

// GetUpdateStatus returns the UpdateStatus field value
func (o *BankConnection) GetUpdateStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.UpdateStatus
}

// GetUpdateStatusOk returns a tuple with the UpdateStatus field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetUpdateStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateStatus, true
}

// SetUpdateStatus sets field value
func (o *BankConnection) SetUpdateStatus(v string) {
	o.UpdateStatus = v
}

// GetCategorizationStatus returns the CategorizationStatus field value
func (o *BankConnection) GetCategorizationStatus() CategorizationStatus {
	if o == nil  {
		var ret CategorizationStatus
		return ret
	}

	return o.CategorizationStatus
}

// GetCategorizationStatusOk returns a tuple with the CategorizationStatus field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetCategorizationStatusOk() (*CategorizationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CategorizationStatus, true
}

// SetCategorizationStatus sets field value
func (o *BankConnection) SetCategorizationStatus(v CategorizationStatus) {
	o.CategorizationStatus = v
}

// GetLastManualUpdate returns the LastManualUpdate field value
// If the value is explicit nil, the zero value for UpdateResult will be returned
func (o *BankConnection) GetLastManualUpdate() UpdateResult {
	if o == nil || o.LastManualUpdate.Get() == nil {
		var ret UpdateResult
		return ret
	}

	return *o.LastManualUpdate.Get()
}

// GetLastManualUpdateOk returns a tuple with the LastManualUpdate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnection) GetLastManualUpdateOk() (*UpdateResult, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastManualUpdate.Get(), o.LastManualUpdate.IsSet()
}

// SetLastManualUpdate sets field value
func (o *BankConnection) SetLastManualUpdate(v UpdateResult) {
	o.LastManualUpdate.Set(&v)
}

// GetLastAutoUpdate returns the LastAutoUpdate field value
// If the value is explicit nil, the zero value for UpdateResult will be returned
func (o *BankConnection) GetLastAutoUpdate() UpdateResult {
	if o == nil || o.LastAutoUpdate.Get() == nil {
		var ret UpdateResult
		return ret
	}

	return *o.LastAutoUpdate.Get()
}

// GetLastAutoUpdateOk returns a tuple with the LastAutoUpdate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnection) GetLastAutoUpdateOk() (*UpdateResult, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastAutoUpdate.Get(), o.LastAutoUpdate.IsSet()
}

// SetLastAutoUpdate sets field value
func (o *BankConnection) SetLastAutoUpdate(v UpdateResult) {
	o.LastAutoUpdate.Set(&v)
}

// GetIbanOnlyMoneyTransferSupported returns the IbanOnlyMoneyTransferSupported field value
func (o *BankConnection) GetIbanOnlyMoneyTransferSupported() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IbanOnlyMoneyTransferSupported
}

// GetIbanOnlyMoneyTransferSupportedOk returns a tuple with the IbanOnlyMoneyTransferSupported field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetIbanOnlyMoneyTransferSupportedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IbanOnlyMoneyTransferSupported, true
}

// SetIbanOnlyMoneyTransferSupported sets field value
func (o *BankConnection) SetIbanOnlyMoneyTransferSupported(v bool) {
	o.IbanOnlyMoneyTransferSupported = v
}

// GetIbanOnlyDirectDebitSupported returns the IbanOnlyDirectDebitSupported field value
func (o *BankConnection) GetIbanOnlyDirectDebitSupported() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IbanOnlyDirectDebitSupported
}

// GetIbanOnlyDirectDebitSupportedOk returns a tuple with the IbanOnlyDirectDebitSupported field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetIbanOnlyDirectDebitSupportedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IbanOnlyDirectDebitSupported, true
}

// SetIbanOnlyDirectDebitSupported sets field value
func (o *BankConnection) SetIbanOnlyDirectDebitSupported(v bool) {
	o.IbanOnlyDirectDebitSupported = v
}

// GetCollectiveMoneyTransferSupported returns the CollectiveMoneyTransferSupported field value
func (o *BankConnection) GetCollectiveMoneyTransferSupported() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.CollectiveMoneyTransferSupported
}

// GetCollectiveMoneyTransferSupportedOk returns a tuple with the CollectiveMoneyTransferSupported field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetCollectiveMoneyTransferSupportedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CollectiveMoneyTransferSupported, true
}

// SetCollectiveMoneyTransferSupported sets field value
func (o *BankConnection) SetCollectiveMoneyTransferSupported(v bool) {
	o.CollectiveMoneyTransferSupported = v
}

// GetDefaultTwoStepProcedureId returns the DefaultTwoStepProcedureId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BankConnection) GetDefaultTwoStepProcedureId() string {
	if o == nil || o.DefaultTwoStepProcedureId.Get() == nil {
		var ret string
		return ret
	}

	return *o.DefaultTwoStepProcedureId.Get()
}

// GetDefaultTwoStepProcedureIdOk returns a tuple with the DefaultTwoStepProcedureId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnection) GetDefaultTwoStepProcedureIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DefaultTwoStepProcedureId.Get(), o.DefaultTwoStepProcedureId.IsSet()
}

// SetDefaultTwoStepProcedureId sets field value
func (o *BankConnection) SetDefaultTwoStepProcedureId(v string) {
	o.DefaultTwoStepProcedureId.Set(&v)
}

// GetTwoStepProcedures returns the TwoStepProcedures field value
func (o *BankConnection) GetTwoStepProcedures() []TwoStepProcedure {
	if o == nil  {
		var ret []TwoStepProcedure
		return ret
	}

	return o.TwoStepProcedures
}

// GetTwoStepProceduresOk returns a tuple with the TwoStepProcedures field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetTwoStepProceduresOk() (*[]TwoStepProcedure, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TwoStepProcedures, true
}

// SetTwoStepProcedures sets field value
func (o *BankConnection) SetTwoStepProcedures(v []TwoStepProcedure) {
	o.TwoStepProcedures = v
}

// GetInterfaces returns the Interfaces field value
func (o *BankConnection) GetInterfaces() []BankConnectionInterface {
	if o == nil  {
		var ret []BankConnectionInterface
		return ret
	}

	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetInterfacesOk() (*[]BankConnectionInterface, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Interfaces, true
}

// SetInterfaces sets field value
func (o *BankConnection) SetInterfaces(v []BankConnectionInterface) {
	o.Interfaces = v
}

// GetAccountIds returns the AccountIds field value
func (o *BankConnection) GetAccountIds() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}

	return o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetAccountIdsOk() (*[]int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountIds, true
}

// SetAccountIds sets field value
func (o *BankConnection) SetAccountIds(v []int64) {
	o.AccountIds = v
}

// GetOwners returns the Owners field value
// If the value is explicit nil, the zero value for []BankConnectionOwner will be returned
func (o *BankConnection) GetOwners() []BankConnectionOwner {
	if o == nil  {
		var ret []BankConnectionOwner
		return ret
	}

	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankConnection) GetOwnersOk() (*[]BankConnectionOwner, bool) {
	if o == nil || o.Owners == nil {
		return nil, false
	}
	return &o.Owners, true
}

// SetOwners sets field value
func (o *BankConnection) SetOwners(v []BankConnectionOwner) {
	o.Owners = v
}

// GetBank returns the Bank field value
func (o *BankConnection) GetBank() Bank {
	if o == nil  {
		var ret Bank
		return ret
	}

	return o.Bank
}

// GetBankOk returns a tuple with the Bank field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetBankOk() (*Bank, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bank, true
}

// SetBank sets field value
func (o *BankConnection) SetBank(v Bank) {
	o.Bank = v
}

// GetFurtherLoginNotRecommended returns the FurtherLoginNotRecommended field value
func (o *BankConnection) GetFurtherLoginNotRecommended() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.FurtherLoginNotRecommended
}

// GetFurtherLoginNotRecommendedOk returns a tuple with the FurtherLoginNotRecommended field value
// and a boolean to check if the value has been set.
func (o *BankConnection) GetFurtherLoginNotRecommendedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FurtherLoginNotRecommended, true
}

// SetFurtherLoginNotRecommended sets field value
func (o *BankConnection) SetFurtherLoginNotRecommended(v bool) {
	o.FurtherLoginNotRecommended = v
}

func (o BankConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["bankId"] = o.BankId
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["bankingUserId"] = o.BankingUserId.Get()
	}
	if true {
		toSerialize["bankingCustomerId"] = o.BankingCustomerId.Get()
	}
	if true {
		toSerialize["bankingPin"] = o.BankingPin.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["updateStatus"] = o.UpdateStatus
	}
	if true {
		toSerialize["categorizationStatus"] = o.CategorizationStatus
	}
	if true {
		toSerialize["lastManualUpdate"] = o.LastManualUpdate.Get()
	}
	if true {
		toSerialize["lastAutoUpdate"] = o.LastAutoUpdate.Get()
	}
	if true {
		toSerialize["ibanOnlyMoneyTransferSupported"] = o.IbanOnlyMoneyTransferSupported
	}
	if true {
		toSerialize["ibanOnlyDirectDebitSupported"] = o.IbanOnlyDirectDebitSupported
	}
	if true {
		toSerialize["collectiveMoneyTransferSupported"] = o.CollectiveMoneyTransferSupported
	}
	if true {
		toSerialize["defaultTwoStepProcedureId"] = o.DefaultTwoStepProcedureId.Get()
	}
	if true {
		toSerialize["twoStepProcedures"] = o.TwoStepProcedures
	}
	if true {
		toSerialize["interfaces"] = o.Interfaces
	}
	if true {
		toSerialize["accountIds"] = o.AccountIds
	}
	if o.Owners != nil {
		toSerialize["owners"] = o.Owners
	}
	if true {
		toSerialize["bank"] = o.Bank
	}
	if true {
		toSerialize["furtherLoginNotRecommended"] = o.FurtherLoginNotRecommended
	}
	return json.Marshal(toSerialize)
}

type NullableBankConnection struct {
	value *BankConnection
	isSet bool
}

func (v NullableBankConnection) Get() *BankConnection {
	return v.value
}

func (v *NullableBankConnection) Set(val *BankConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableBankConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableBankConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankConnection(val *BankConnection) *NullableBankConnection {
	return &NullableBankConnection{value: val, isSet: true}
}

func (v NullableBankConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



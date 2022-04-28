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

// Account Container for a bank account's data
type Account struct {
	// Account identifier
	Id int64 `json:"id"`
	// Identifier of the bank connection that this account belongs to
	BankConnectionId int64 `json:"bankConnectionId"`
	// Account name
	AccountName NullableString `json:"accountName"`
	// Account's IBAN. Note that this field can change from 'null' to a value - or vice versa - any time when the account is being updated. This is subject to changes within the bank's internal account management.
	Iban NullableString `json:"iban"`
	// (National) account number. Note that this value might change whenever the account is updated (for example, leading zeros might be added or removed).
	AccountNumber string `json:"accountNumber"`
	// Account's sub-account-number. Note that this field can change from 'null' to a value - or vice versa - any time when the account is being updated. This is subject to changes within the bank's internal account management.
	SubAccountNumber NullableString `json:"subAccountNumber"`
	// Name of the account holder
	AccountHolderName NullableString `json:"accountHolderName"`
	// Bank's internal identification of the account holder. Note that if your client has no license for processing this field, it will always be 'XXXXX'
	AccountHolderId NullableString `json:"accountHolderId"`
	// Account's currency
	AccountCurrency NullableString `json:"accountCurrency"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'accountType' field instead.<br/><br/>Identifier of the account's type. Note that, in general, the type of an account can change any time when the account is being updated. This is subject to changes within the bank's internal account management. However, if the account's type has previously been changed explicitly (via the PATCH method), then the explicitly set type will NOT be automatically changed anymore, even if the type has changed on the bank side. <br/>1 = Checking,<br/>2 = Savings,<br/>3 = CreditCard,<br/>4 = Security,<br/>5 = Loan,<br/>6 = Pocket (DEPRECATED; will not be returned for any account unless this type has explicitly been set via PATCH),<br/>7 = Membership,<br/>8 = Bausparen<br/>
	AccountTypeId int64 `json:"accountTypeId"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'accountType' field instead.<br/><br/>Name of the account's type.
	AccountTypeName string `json:"accountTypeName"`
	// <strong>Type:</strong> AccountType<br/> An account type.<br/><br/>Checking,<br/>Savings,<br/>CreditCard,<br/>Security,<br/>Loan,<br/>Pocket (DEPRECATED; will not be returned for any account unless this type has explicitly been set via PATCH),<br/>Membership,<br/>Bausparen<br/><br/>
	AccountType AccountType `json:"accountType"`
	// Current account balance
	Balance NullableFloat64 `json:"balance"`
	// Current overdraft
	Overdraft NullableFloat64 `json:"overdraft"`
	// Overdraft limit
	OverdraftLimit NullableFloat64 `json:"overdraftLimit"`
	// Current available funds. Note that this field is only set if finAPI can make a definite statement about the current available funds. This might not always be the case, for example if there is not enough information available about the overdraft limit and current overdraft.
	AvailableFunds NullableFloat64 `json:"availableFunds"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the corresponding field in 'interfaces' instead.<br/><br/>Timestamp of when the account was last successfully updated (or initially imported); more precisely: time when the account data (balance and positions) has been stored into the finAPI databases. The value is returned in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	LastSuccessfulUpdate NullableString `json:"lastSuccessfulUpdate"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the corresponding field in 'interfaces' instead.<br/><br/>Timestamp of when the account was last tried to be updated (or initially imported); more precisely: time when the update (or initial import) was triggered. The value is returned in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	LastUpdateAttempt NullableString `json:"lastUpdateAttempt"`
	// Indicating whether this account is 'new' or not. Any newly imported account will have this flag initially set to true, and remain so until you set it to false (see PATCH /accounts/<id>). How you use this field is up to your interpretation, however it is recommended to set the flag to false for all accounts right after the initial import of the bank connection. This way, you will be able recognize accounts that get newly imported during a later update of the bank connection, by checking for any accounts with the flag set to true right after an update.
	IsNew bool `json:"isNew"`
	// <strong>Type:</strong> AccountStatus<br/> THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'status' field in 'interfaces' instead.<br/><br/>The current status of the account. Possible values are:<br/>&bull; <code>UPDATED</code> means that the account is up to date from finAPI's point of view. This means that no current import/update is running, and the previous import/update could successfully update the account's data (e.g. transactions and securities), and the bank given balance matched the transaction's calculated sum, meaning that no adjusting entry ('Zwischensaldo' transaction) was inserted.<br/>&bull; <code>UPDATED_FIXED</code> means that the account is up to date from finAPI's point of view (no current import/update is running, and the previous import/update could successfully update the account's data), BUT there was a deviation in the bank given balance which was fixed by adding an adjusting entry ('Zwischensaldo' transaction).<br/>&bull; <code>DOWNLOAD_IN_PROGRESS</code> means that the account's data is currently being imported/updated.<br/>&bull; <code>DOWNLOAD_FAILED</code> means that the account data could not get successfully imported or updated. Possible reasons: finAPI could not get the account's balance, or it could not parse all transactions/securities, or some internal error has occurred. Also, it could mean that finAPI could not even get to the point of receiving the account data from the bank server, for example because of incorrect login credentials or a network problem. Note however that when we get a balance and just an empty list of transactions or securities, then this is regarded as valid and successful download. The reason for this is that for some accounts that have little activity, we may actually get no recent transactions but only a balance.<br/>&bull; <code>DEPRECATED</code> means that the account could no longer get matched with any account from the bank server. This can mean either that the account was terminated by the user and is no longer sent by the bank server, or that finAPI could no longer match it because the account's data (name, type, iban, account number, etc.) has been changed by the bank.
	Status AccountStatus `json:"status"`
	SupportedOrders []SupportedOrder `json:"supportedOrders"`
	// <strong>Type:</strong> AccountInterface<br/> Set of interfaces to which this account is connected
	Interfaces []AccountInterface `json:"interfaces"`
	// <strong>Type:</strong> ClearingAccountData<br/> THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/><br/>List of clearing accounts that relate to this account. Clearing accounts can be used for money transfers (see field 'clearingAccountId' of the 'Request SEPA Money Transfer' service).
	ClearingAccounts []ClearingAccountData `json:"clearingAccounts"`
	// Whether this account is seized. Note that this information is not received from the bank, but determined by finAPI based on the available account information.
	IsSeized bool `json:"isSeized"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(id int64, bankConnectionId int64, accountName NullableString, iban NullableString, accountNumber string, subAccountNumber NullableString, accountHolderName NullableString, accountHolderId NullableString, accountCurrency NullableString, accountTypeId int64, accountTypeName string, accountType AccountType, balance NullableFloat64, overdraft NullableFloat64, overdraftLimit NullableFloat64, availableFunds NullableFloat64, lastSuccessfulUpdate NullableString, lastUpdateAttempt NullableString, isNew bool, status AccountStatus, supportedOrders []SupportedOrder, interfaces []AccountInterface, clearingAccounts []ClearingAccountData, isSeized bool, ) *Account {
	this := Account{}
	this.Id = id
	this.BankConnectionId = bankConnectionId
	this.AccountName = accountName
	this.Iban = iban
	this.AccountNumber = accountNumber
	this.SubAccountNumber = subAccountNumber
	this.AccountHolderName = accountHolderName
	this.AccountHolderId = accountHolderId
	this.AccountCurrency = accountCurrency
	this.AccountTypeId = accountTypeId
	this.AccountTypeName = accountTypeName
	this.AccountType = accountType
	this.Balance = balance
	this.Overdraft = overdraft
	this.OverdraftLimit = overdraftLimit
	this.AvailableFunds = availableFunds
	this.LastSuccessfulUpdate = lastSuccessfulUpdate
	this.LastUpdateAttempt = lastUpdateAttempt
	this.IsNew = isNew
	this.Status = status
	this.SupportedOrders = supportedOrders
	this.Interfaces = interfaces
	this.ClearingAccounts = clearingAccounts
	this.IsSeized = isSeized
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetId returns the Id field value
func (o *Account) GetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Account) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Account) SetId(v int64) {
	o.Id = v
}

// GetBankConnectionId returns the BankConnectionId field value
func (o *Account) GetBankConnectionId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.BankConnectionId
}

// GetBankConnectionIdOk returns a tuple with the BankConnectionId field value
// and a boolean to check if the value has been set.
func (o *Account) GetBankConnectionIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankConnectionId, true
}

// SetBankConnectionId sets field value
func (o *Account) SetBankConnectionId(v int64) {
	o.BankConnectionId = v
}

// GetAccountName returns the AccountName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Account) GetAccountName() string {
	if o == nil || o.AccountName.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountName.Get()
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountName.Get(), o.AccountName.IsSet()
}

// SetAccountName sets field value
func (o *Account) SetAccountName(v string) {
	o.AccountName.Set(&v)
}

// GetIban returns the Iban field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Account) GetIban() string {
	if o == nil || o.Iban.Get() == nil {
		var ret string
		return ret
	}

	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// SetIban sets field value
func (o *Account) SetIban(v string) {
	o.Iban.Set(&v)
}

// GetAccountNumber returns the AccountNumber field value
func (o *Account) GetAccountNumber() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *Account) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *Account) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetSubAccountNumber returns the SubAccountNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Account) GetSubAccountNumber() string {
	if o == nil || o.SubAccountNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubAccountNumber.Get()
}

// GetSubAccountNumberOk returns a tuple with the SubAccountNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetSubAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubAccountNumber.Get(), o.SubAccountNumber.IsSet()
}

// SetSubAccountNumber sets field value
func (o *Account) SetSubAccountNumber(v string) {
	o.SubAccountNumber.Set(&v)
}

// GetAccountHolderName returns the AccountHolderName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Account) GetAccountHolderName() string {
	if o == nil || o.AccountHolderName.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountHolderName.Get()
}

// GetAccountHolderNameOk returns a tuple with the AccountHolderName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetAccountHolderNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountHolderName.Get(), o.AccountHolderName.IsSet()
}

// SetAccountHolderName sets field value
func (o *Account) SetAccountHolderName(v string) {
	o.AccountHolderName.Set(&v)
}

// GetAccountHolderId returns the AccountHolderId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Account) GetAccountHolderId() string {
	if o == nil || o.AccountHolderId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountHolderId.Get()
}

// GetAccountHolderIdOk returns a tuple with the AccountHolderId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetAccountHolderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountHolderId.Get(), o.AccountHolderId.IsSet()
}

// SetAccountHolderId sets field value
func (o *Account) SetAccountHolderId(v string) {
	o.AccountHolderId.Set(&v)
}

// GetAccountCurrency returns the AccountCurrency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Account) GetAccountCurrency() string {
	if o == nil || o.AccountCurrency.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountCurrency.Get()
}

// GetAccountCurrencyOk returns a tuple with the AccountCurrency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetAccountCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountCurrency.Get(), o.AccountCurrency.IsSet()
}

// SetAccountCurrency sets field value
func (o *Account) SetAccountCurrency(v string) {
	o.AccountCurrency.Set(&v)
}

// GetAccountTypeId returns the AccountTypeId field value
func (o *Account) GetAccountTypeId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.AccountTypeId
}

// GetAccountTypeIdOk returns a tuple with the AccountTypeId field value
// and a boolean to check if the value has been set.
func (o *Account) GetAccountTypeIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountTypeId, true
}

// SetAccountTypeId sets field value
func (o *Account) SetAccountTypeId(v int64) {
	o.AccountTypeId = v
}

// GetAccountTypeName returns the AccountTypeName field value
func (o *Account) GetAccountTypeName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountTypeName
}

// GetAccountTypeNameOk returns a tuple with the AccountTypeName field value
// and a boolean to check if the value has been set.
func (o *Account) GetAccountTypeNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountTypeName, true
}

// SetAccountTypeName sets field value
func (o *Account) SetAccountTypeName(v string) {
	o.AccountTypeName = v
}

// GetAccountType returns the AccountType field value
func (o *Account) GetAccountType() AccountType {
	if o == nil  {
		var ret AccountType
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *Account) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *Account) SetAccountType(v AccountType) {
	o.AccountType = v
}

// GetBalance returns the Balance field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Account) GetBalance() float64 {
	if o == nil || o.Balance.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetBalanceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// SetBalance sets field value
func (o *Account) SetBalance(v float64) {
	o.Balance.Set(&v)
}

// GetOverdraft returns the Overdraft field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Account) GetOverdraft() float64 {
	if o == nil || o.Overdraft.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Overdraft.Get()
}

// GetOverdraftOk returns a tuple with the Overdraft field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetOverdraftOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Overdraft.Get(), o.Overdraft.IsSet()
}

// SetOverdraft sets field value
func (o *Account) SetOverdraft(v float64) {
	o.Overdraft.Set(&v)
}

// GetOverdraftLimit returns the OverdraftLimit field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Account) GetOverdraftLimit() float64 {
	if o == nil || o.OverdraftLimit.Get() == nil {
		var ret float64
		return ret
	}

	return *o.OverdraftLimit.Get()
}

// GetOverdraftLimitOk returns a tuple with the OverdraftLimit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetOverdraftLimitOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverdraftLimit.Get(), o.OverdraftLimit.IsSet()
}

// SetOverdraftLimit sets field value
func (o *Account) SetOverdraftLimit(v float64) {
	o.OverdraftLimit.Set(&v)
}

// GetAvailableFunds returns the AvailableFunds field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Account) GetAvailableFunds() float64 {
	if o == nil || o.AvailableFunds.Get() == nil {
		var ret float64
		return ret
	}

	return *o.AvailableFunds.Get()
}

// GetAvailableFundsOk returns a tuple with the AvailableFunds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetAvailableFundsOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AvailableFunds.Get(), o.AvailableFunds.IsSet()
}

// SetAvailableFunds sets field value
func (o *Account) SetAvailableFunds(v float64) {
	o.AvailableFunds.Set(&v)
}

// GetLastSuccessfulUpdate returns the LastSuccessfulUpdate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Account) GetLastSuccessfulUpdate() string {
	if o == nil || o.LastSuccessfulUpdate.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastSuccessfulUpdate.Get()
}

// GetLastSuccessfulUpdateOk returns a tuple with the LastSuccessfulUpdate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetLastSuccessfulUpdateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastSuccessfulUpdate.Get(), o.LastSuccessfulUpdate.IsSet()
}

// SetLastSuccessfulUpdate sets field value
func (o *Account) SetLastSuccessfulUpdate(v string) {
	o.LastSuccessfulUpdate.Set(&v)
}

// GetLastUpdateAttempt returns the LastUpdateAttempt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Account) GetLastUpdateAttempt() string {
	if o == nil || o.LastUpdateAttempt.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastUpdateAttempt.Get()
}

// GetLastUpdateAttemptOk returns a tuple with the LastUpdateAttempt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetLastUpdateAttemptOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUpdateAttempt.Get(), o.LastUpdateAttempt.IsSet()
}

// SetLastUpdateAttempt sets field value
func (o *Account) SetLastUpdateAttempt(v string) {
	o.LastUpdateAttempt.Set(&v)
}

// GetIsNew returns the IsNew field value
func (o *Account) GetIsNew() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsNew
}

// GetIsNewOk returns a tuple with the IsNew field value
// and a boolean to check if the value has been set.
func (o *Account) GetIsNewOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsNew, true
}

// SetIsNew sets field value
func (o *Account) SetIsNew(v bool) {
	o.IsNew = v
}

// GetStatus returns the Status field value
func (o *Account) GetStatus() AccountStatus {
	if o == nil  {
		var ret AccountStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Account) GetStatusOk() (*AccountStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Account) SetStatus(v AccountStatus) {
	o.Status = v
}

// GetSupportedOrders returns the SupportedOrders field value
func (o *Account) GetSupportedOrders() []SupportedOrder {
	if o == nil  {
		var ret []SupportedOrder
		return ret
	}

	return o.SupportedOrders
}

// GetSupportedOrdersOk returns a tuple with the SupportedOrders field value
// and a boolean to check if the value has been set.
func (o *Account) GetSupportedOrdersOk() (*[]SupportedOrder, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupportedOrders, true
}

// SetSupportedOrders sets field value
func (o *Account) SetSupportedOrders(v []SupportedOrder) {
	o.SupportedOrders = v
}

// GetInterfaces returns the Interfaces field value
func (o *Account) GetInterfaces() []AccountInterface {
	if o == nil  {
		var ret []AccountInterface
		return ret
	}

	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value
// and a boolean to check if the value has been set.
func (o *Account) GetInterfacesOk() (*[]AccountInterface, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Interfaces, true
}

// SetInterfaces sets field value
func (o *Account) SetInterfaces(v []AccountInterface) {
	o.Interfaces = v
}

// GetClearingAccounts returns the ClearingAccounts field value
func (o *Account) GetClearingAccounts() []ClearingAccountData {
	if o == nil  {
		var ret []ClearingAccountData
		return ret
	}

	return o.ClearingAccounts
}

// GetClearingAccountsOk returns a tuple with the ClearingAccounts field value
// and a boolean to check if the value has been set.
func (o *Account) GetClearingAccountsOk() (*[]ClearingAccountData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClearingAccounts, true
}

// SetClearingAccounts sets field value
func (o *Account) SetClearingAccounts(v []ClearingAccountData) {
	o.ClearingAccounts = v
}

// GetIsSeized returns the IsSeized field value
func (o *Account) GetIsSeized() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsSeized
}

// GetIsSeizedOk returns a tuple with the IsSeized field value
// and a boolean to check if the value has been set.
func (o *Account) GetIsSeizedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSeized, true
}

// SetIsSeized sets field value
func (o *Account) SetIsSeized(v bool) {
	o.IsSeized = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["bankConnectionId"] = o.BankConnectionId
	}
	if true {
		toSerialize["accountName"] = o.AccountName.Get()
	}
	if true {
		toSerialize["iban"] = o.Iban.Get()
	}
	if true {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if true {
		toSerialize["subAccountNumber"] = o.SubAccountNumber.Get()
	}
	if true {
		toSerialize["accountHolderName"] = o.AccountHolderName.Get()
	}
	if true {
		toSerialize["accountHolderId"] = o.AccountHolderId.Get()
	}
	if true {
		toSerialize["accountCurrency"] = o.AccountCurrency.Get()
	}
	if true {
		toSerialize["accountTypeId"] = o.AccountTypeId
	}
	if true {
		toSerialize["accountTypeName"] = o.AccountTypeName
	}
	if true {
		toSerialize["accountType"] = o.AccountType
	}
	if true {
		toSerialize["balance"] = o.Balance.Get()
	}
	if true {
		toSerialize["overdraft"] = o.Overdraft.Get()
	}
	if true {
		toSerialize["overdraftLimit"] = o.OverdraftLimit.Get()
	}
	if true {
		toSerialize["availableFunds"] = o.AvailableFunds.Get()
	}
	if true {
		toSerialize["lastSuccessfulUpdate"] = o.LastSuccessfulUpdate.Get()
	}
	if true {
		toSerialize["lastUpdateAttempt"] = o.LastUpdateAttempt.Get()
	}
	if true {
		toSerialize["isNew"] = o.IsNew
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["supportedOrders"] = o.SupportedOrders
	}
	if true {
		toSerialize["interfaces"] = o.Interfaces
	}
	if true {
		toSerialize["clearingAccounts"] = o.ClearingAccounts
	}
	if true {
		toSerialize["isSeized"] = o.IsSeized
	}
	return json.Marshal(toSerialize)
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



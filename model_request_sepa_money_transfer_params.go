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

// RequestSepaMoneyTransferParams Parameters for a single or collective SEPA money transfer order request
type RequestSepaMoneyTransferParams struct {
	// Name of the recipient. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the recipient name. Even if the recipient name does not depict the actual registered account holder of the specified recipient account, the money transfer request might still be successful. This field is optional only when you pass a clearing account as the recipient. Otherwise, this field is required.
	RecipientName *string `json:"recipientName,omitempty"`
	// IBAN of the recipient's account. This field is optional only when you pass a clearing account as the recipient. Otherwise, this field is required.
	RecipientIban *string `json:"recipientIban,omitempty"`
	// BIC of the recipient's account. Note: This field is optional when you pass a clearing account as the recipient or if the bank connection of the account that you want to transfer money from supports the IBAN-Only money transfer. You can find this out via GET /bankConnections/<id>. If no BIC is given, finAPI will try to recognize it using the given recipientIban value (if it's given). And then if the result value is not empty, it will be used for the money transfer request independent of whether it is required or not (unless you pass a clearing account, in which case the value will always be ignored).
	RecipientBic *string `json:"recipientBic,omitempty"`
	// Identifier of a clearing account. If this field is set, then the fields 'recipientName', 'recipientIban' and 'recipientBic' will be ignored and the recipient account will be the specified clearing account.
	ClearingAccountId *string `json:"clearingAccountId,omitempty"`
	// End-To-End ID for the transfer transaction
	EndToEndId *string `json:"endToEndId,omitempty"`
	// The amount to transfer. Must be a positive decimal number with at most two decimal places (e.g. 99.99)
	Amount float64 `json:"amount"`
	// The purpose of the transfer transaction
	Purpose *string `json:"purpose,omitempty"`
	// SEPA purpose code, according to ISO 20022, external codes set.
	SepaPurposeCode *string `json:"sepaPurposeCode,omitempty"`
	// Identifier of the bank account that you want to transfer money from
	AccountId int64 `json:"accountId"`
	// Online banking PIN. Any symbols are allowed. Max length: 170. If a PIN is stored in the bank connection, then this field may remain unset. If finAPI's Web Form is not required and the field is set though then it will always be used (even if there is some other PIN stored in the bank connection). If you want the user to enter a PIN in finAPI's Web Form even when a PIN is stored, then just set the field to any value, so that the service recognizes that you wish to use the Web Form flow.
	BankingPin *string `json:"bankingPin,omitempty"`
	// Whether to store the PIN. If the PIN is stored, it is not required to pass the PIN again when executing SEPA orders. Default value is 'false'. <br/><br/>NOTES:<br/> - before you set this field to true, please regard the 'pinsAreVolatile' flag of the bank connection that the account belongs to. Please note that volatile credentials will not be stored, even if provided, to enforce user involvement in the next communication with the bank;<br/> - this field is ignored in case when the user will need to use finAPI's Web Form. The user will be able to decide whether to store the PIN or not in the Web Form, depending on the 'storeSecretsAvailableInWebForm' setting (see Client Configuration).
	StoreSecrets *bool `json:"storeSecrets,omitempty"`
	// The bank-given ID of the two-step-procedure that should be used for the order. For a list of available two-step-procedures, see the corresponding bank connection (GET /bankConnections). If this field is not set, then the bank connection's default two-step-procedure will be used. Note that in this case, when the bank connection has no default two-step-procedure set, then the response of the service depends on whether you need to use finAPI's Web Form or not. If you need to use the Web Form, the user will be prompted to select the two-step-procedure within the Web Form. If you don't need to use the Web Form, then the service will return an error (passing a value for this field is required in this case).
	TwoStepProcedureId *string `json:"twoStepProcedureId,omitempty"`
	// Execution date for the money transfer(s), in the format 'YYYY-MM-DD'. If not specified, then the current date will be used.
	ExecutionDate *string `json:"executionDate,omitempty"`
	// This field is only regarded when you pass multiple orders. It determines whether the orders should be processed by the bank as one collective booking (in case of 'false'), or as single bookings (in case of 'true'). Default value is 'false'.
	SingleBooking *bool `json:"singleBooking,omitempty"`
	// <strong>Type:</strong> SingleMoneyTransferRecipientData<br/> In case that you want to submit not just a single money transfer, but do a collective money transfer, use this field to pass a list of additional money transfer orders. The service will then pass a collective money transfer request to the bank, including both the money transfer specified on the top-level, as well as all money transfers specified in this list. The maximum count of money transfers that you can pass (in total) is 15000. Note that you should check the account's 'supportedOrders' field to find out whether or not it is supporting collective money transfers.
	AdditionalMoneyTransfers *[]SingleMoneyTransferRecipientData `json:"additionalMoneyTransfers,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'multiStepAuthentication' field instead.<br/><br/>Challenge response. This field should be set only when the previous attempt to request a SEPA money transfer failed with HTTP code 510, i.e. the bank sent a challenge for the user for an additional authentication. In this case, this field must contain the response to the bank's challenge. Please note that in case of using finAPI's Web Form you have to leave this field unset and the application will automatically recognize that the user has to input challenge response and then a Web Form will be shown to the user.
	ChallengeResponse *string `json:"challengeResponse,omitempty"`
	// Whether the finAPI Web Form should hide transaction details when prompting the caller for the second factor. Default value is false.
	HideTransactionDetailsInWebForm *bool `json:"hideTransactionDetailsInWebForm,omitempty"`
	// <strong>Type:</strong> MultiStepAuthenticationCallback<br/> Container for multi-step authentication data. Required when a previous service call initiated a multi-step authentication.
	MultiStepAuthentication *MultiStepAuthenticationCallback `json:"multiStepAuthentication,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'storeSecrets' field instead.<br/><br/>Whether to store the PIN. If the PIN is stored, it is not required to pass the PIN again when executing SEPA orders. Default value is 'false'. <br/><br/>NOTES:<br/> - before you set this field to true, please regard the 'pinsAreVolatile' flag of the bank connection that the account belongs to. Please note that volatile credentials will not be stored, even if provided, to enforce user involvement in the next communication with the bank;<br/> - this field is ignored in case when the user will need to use finAPI's Web Form. The user will be able to decide whether to store the PIN or not in the Web Form, depending on the 'storeSecretsAvailableInWebForm' setting (see Client Configuration).
	StorePin *bool `json:"storePin,omitempty"`
}

// NewRequestSepaMoneyTransferParams instantiates a new RequestSepaMoneyTransferParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSepaMoneyTransferParams(amount float64, accountId int64, ) *RequestSepaMoneyTransferParams {
	this := RequestSepaMoneyTransferParams{}
	this.Amount = amount
	this.AccountId = accountId
	var storeSecrets bool = false
	this.StoreSecrets = &storeSecrets
	var singleBooking bool = false
	this.SingleBooking = &singleBooking
	var hideTransactionDetailsInWebForm bool = false
	this.HideTransactionDetailsInWebForm = &hideTransactionDetailsInWebForm
	var storePin bool = false
	this.StorePin = &storePin
	return &this
}

// NewRequestSepaMoneyTransferParamsWithDefaults instantiates a new RequestSepaMoneyTransferParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSepaMoneyTransferParamsWithDefaults() *RequestSepaMoneyTransferParams {
	this := RequestSepaMoneyTransferParams{}
	var storeSecrets bool = false
	this.StoreSecrets = &storeSecrets
	var singleBooking bool = false
	this.SingleBooking = &singleBooking
	var hideTransactionDetailsInWebForm bool = false
	this.HideTransactionDetailsInWebForm = &hideTransactionDetailsInWebForm
	var storePin bool = false
	this.StorePin = &storePin
	return &this
}

// GetRecipientName returns the RecipientName field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetRecipientName() string {
	if o == nil || o.RecipientName == nil {
		var ret string
		return ret
	}
	return *o.RecipientName
}

// GetRecipientNameOk returns a tuple with the RecipientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetRecipientNameOk() (*string, bool) {
	if o == nil || o.RecipientName == nil {
		return nil, false
	}
	return o.RecipientName, true
}

// HasRecipientName returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasRecipientName() bool {
	if o != nil && o.RecipientName != nil {
		return true
	}

	return false
}

// SetRecipientName gets a reference to the given string and assigns it to the RecipientName field.
func (o *RequestSepaMoneyTransferParams) SetRecipientName(v string) {
	o.RecipientName = &v
}

// GetRecipientIban returns the RecipientIban field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetRecipientIban() string {
	if o == nil || o.RecipientIban == nil {
		var ret string
		return ret
	}
	return *o.RecipientIban
}

// GetRecipientIbanOk returns a tuple with the RecipientIban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetRecipientIbanOk() (*string, bool) {
	if o == nil || o.RecipientIban == nil {
		return nil, false
	}
	return o.RecipientIban, true
}

// HasRecipientIban returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasRecipientIban() bool {
	if o != nil && o.RecipientIban != nil {
		return true
	}

	return false
}

// SetRecipientIban gets a reference to the given string and assigns it to the RecipientIban field.
func (o *RequestSepaMoneyTransferParams) SetRecipientIban(v string) {
	o.RecipientIban = &v
}

// GetRecipientBic returns the RecipientBic field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetRecipientBic() string {
	if o == nil || o.RecipientBic == nil {
		var ret string
		return ret
	}
	return *o.RecipientBic
}

// GetRecipientBicOk returns a tuple with the RecipientBic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetRecipientBicOk() (*string, bool) {
	if o == nil || o.RecipientBic == nil {
		return nil, false
	}
	return o.RecipientBic, true
}

// HasRecipientBic returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasRecipientBic() bool {
	if o != nil && o.RecipientBic != nil {
		return true
	}

	return false
}

// SetRecipientBic gets a reference to the given string and assigns it to the RecipientBic field.
func (o *RequestSepaMoneyTransferParams) SetRecipientBic(v string) {
	o.RecipientBic = &v
}

// GetClearingAccountId returns the ClearingAccountId field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetClearingAccountId() string {
	if o == nil || o.ClearingAccountId == nil {
		var ret string
		return ret
	}
	return *o.ClearingAccountId
}

// GetClearingAccountIdOk returns a tuple with the ClearingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetClearingAccountIdOk() (*string, bool) {
	if o == nil || o.ClearingAccountId == nil {
		return nil, false
	}
	return o.ClearingAccountId, true
}

// HasClearingAccountId returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasClearingAccountId() bool {
	if o != nil && o.ClearingAccountId != nil {
		return true
	}

	return false
}

// SetClearingAccountId gets a reference to the given string and assigns it to the ClearingAccountId field.
func (o *RequestSepaMoneyTransferParams) SetClearingAccountId(v string) {
	o.ClearingAccountId = &v
}

// GetEndToEndId returns the EndToEndId field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetEndToEndId() string {
	if o == nil || o.EndToEndId == nil {
		var ret string
		return ret
	}
	return *o.EndToEndId
}

// GetEndToEndIdOk returns a tuple with the EndToEndId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetEndToEndIdOk() (*string, bool) {
	if o == nil || o.EndToEndId == nil {
		return nil, false
	}
	return o.EndToEndId, true
}

// HasEndToEndId returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasEndToEndId() bool {
	if o != nil && o.EndToEndId != nil {
		return true
	}

	return false
}

// SetEndToEndId gets a reference to the given string and assigns it to the EndToEndId field.
func (o *RequestSepaMoneyTransferParams) SetEndToEndId(v string) {
	o.EndToEndId = &v
}

// GetAmount returns the Amount field value
func (o *RequestSepaMoneyTransferParams) GetAmount() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *RequestSepaMoneyTransferParams) SetAmount(v float64) {
	o.Amount = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *RequestSepaMoneyTransferParams) SetPurpose(v string) {
	o.Purpose = &v
}

// GetSepaPurposeCode returns the SepaPurposeCode field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetSepaPurposeCode() string {
	if o == nil || o.SepaPurposeCode == nil {
		var ret string
		return ret
	}
	return *o.SepaPurposeCode
}

// GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetSepaPurposeCodeOk() (*string, bool) {
	if o == nil || o.SepaPurposeCode == nil {
		return nil, false
	}
	return o.SepaPurposeCode, true
}

// HasSepaPurposeCode returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasSepaPurposeCode() bool {
	if o != nil && o.SepaPurposeCode != nil {
		return true
	}

	return false
}

// SetSepaPurposeCode gets a reference to the given string and assigns it to the SepaPurposeCode field.
func (o *RequestSepaMoneyTransferParams) SetSepaPurposeCode(v string) {
	o.SepaPurposeCode = &v
}

// GetAccountId returns the AccountId field value
func (o *RequestSepaMoneyTransferParams) GetAccountId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetAccountIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *RequestSepaMoneyTransferParams) SetAccountId(v int64) {
	o.AccountId = v
}

// GetBankingPin returns the BankingPin field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetBankingPin() string {
	if o == nil || o.BankingPin == nil {
		var ret string
		return ret
	}
	return *o.BankingPin
}

// GetBankingPinOk returns a tuple with the BankingPin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetBankingPinOk() (*string, bool) {
	if o == nil || o.BankingPin == nil {
		return nil, false
	}
	return o.BankingPin, true
}

// HasBankingPin returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasBankingPin() bool {
	if o != nil && o.BankingPin != nil {
		return true
	}

	return false
}

// SetBankingPin gets a reference to the given string and assigns it to the BankingPin field.
func (o *RequestSepaMoneyTransferParams) SetBankingPin(v string) {
	o.BankingPin = &v
}

// GetStoreSecrets returns the StoreSecrets field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetStoreSecrets() bool {
	if o == nil || o.StoreSecrets == nil {
		var ret bool
		return ret
	}
	return *o.StoreSecrets
}

// GetStoreSecretsOk returns a tuple with the StoreSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetStoreSecretsOk() (*bool, bool) {
	if o == nil || o.StoreSecrets == nil {
		return nil, false
	}
	return o.StoreSecrets, true
}

// HasStoreSecrets returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasStoreSecrets() bool {
	if o != nil && o.StoreSecrets != nil {
		return true
	}

	return false
}

// SetStoreSecrets gets a reference to the given bool and assigns it to the StoreSecrets field.
func (o *RequestSepaMoneyTransferParams) SetStoreSecrets(v bool) {
	o.StoreSecrets = &v
}

// GetTwoStepProcedureId returns the TwoStepProcedureId field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetTwoStepProcedureId() string {
	if o == nil || o.TwoStepProcedureId == nil {
		var ret string
		return ret
	}
	return *o.TwoStepProcedureId
}

// GetTwoStepProcedureIdOk returns a tuple with the TwoStepProcedureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetTwoStepProcedureIdOk() (*string, bool) {
	if o == nil || o.TwoStepProcedureId == nil {
		return nil, false
	}
	return o.TwoStepProcedureId, true
}

// HasTwoStepProcedureId returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasTwoStepProcedureId() bool {
	if o != nil && o.TwoStepProcedureId != nil {
		return true
	}

	return false
}

// SetTwoStepProcedureId gets a reference to the given string and assigns it to the TwoStepProcedureId field.
func (o *RequestSepaMoneyTransferParams) SetTwoStepProcedureId(v string) {
	o.TwoStepProcedureId = &v
}

// GetExecutionDate returns the ExecutionDate field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetExecutionDate() string {
	if o == nil || o.ExecutionDate == nil {
		var ret string
		return ret
	}
	return *o.ExecutionDate
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetExecutionDateOk() (*string, bool) {
	if o == nil || o.ExecutionDate == nil {
		return nil, false
	}
	return o.ExecutionDate, true
}

// HasExecutionDate returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasExecutionDate() bool {
	if o != nil && o.ExecutionDate != nil {
		return true
	}

	return false
}

// SetExecutionDate gets a reference to the given string and assigns it to the ExecutionDate field.
func (o *RequestSepaMoneyTransferParams) SetExecutionDate(v string) {
	o.ExecutionDate = &v
}

// GetSingleBooking returns the SingleBooking field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetSingleBooking() bool {
	if o == nil || o.SingleBooking == nil {
		var ret bool
		return ret
	}
	return *o.SingleBooking
}

// GetSingleBookingOk returns a tuple with the SingleBooking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetSingleBookingOk() (*bool, bool) {
	if o == nil || o.SingleBooking == nil {
		return nil, false
	}
	return o.SingleBooking, true
}

// HasSingleBooking returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasSingleBooking() bool {
	if o != nil && o.SingleBooking != nil {
		return true
	}

	return false
}

// SetSingleBooking gets a reference to the given bool and assigns it to the SingleBooking field.
func (o *RequestSepaMoneyTransferParams) SetSingleBooking(v bool) {
	o.SingleBooking = &v
}

// GetAdditionalMoneyTransfers returns the AdditionalMoneyTransfers field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetAdditionalMoneyTransfers() []SingleMoneyTransferRecipientData {
	if o == nil || o.AdditionalMoneyTransfers == nil {
		var ret []SingleMoneyTransferRecipientData
		return ret
	}
	return *o.AdditionalMoneyTransfers
}

// GetAdditionalMoneyTransfersOk returns a tuple with the AdditionalMoneyTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetAdditionalMoneyTransfersOk() (*[]SingleMoneyTransferRecipientData, bool) {
	if o == nil || o.AdditionalMoneyTransfers == nil {
		return nil, false
	}
	return o.AdditionalMoneyTransfers, true
}

// HasAdditionalMoneyTransfers returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasAdditionalMoneyTransfers() bool {
	if o != nil && o.AdditionalMoneyTransfers != nil {
		return true
	}

	return false
}

// SetAdditionalMoneyTransfers gets a reference to the given []SingleMoneyTransferRecipientData and assigns it to the AdditionalMoneyTransfers field.
func (o *RequestSepaMoneyTransferParams) SetAdditionalMoneyTransfers(v []SingleMoneyTransferRecipientData) {
	o.AdditionalMoneyTransfers = &v
}

// GetChallengeResponse returns the ChallengeResponse field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetChallengeResponse() string {
	if o == nil || o.ChallengeResponse == nil {
		var ret string
		return ret
	}
	return *o.ChallengeResponse
}

// GetChallengeResponseOk returns a tuple with the ChallengeResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetChallengeResponseOk() (*string, bool) {
	if o == nil || o.ChallengeResponse == nil {
		return nil, false
	}
	return o.ChallengeResponse, true
}

// HasChallengeResponse returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasChallengeResponse() bool {
	if o != nil && o.ChallengeResponse != nil {
		return true
	}

	return false
}

// SetChallengeResponse gets a reference to the given string and assigns it to the ChallengeResponse field.
func (o *RequestSepaMoneyTransferParams) SetChallengeResponse(v string) {
	o.ChallengeResponse = &v
}

// GetHideTransactionDetailsInWebForm returns the HideTransactionDetailsInWebForm field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetHideTransactionDetailsInWebForm() bool {
	if o == nil || o.HideTransactionDetailsInWebForm == nil {
		var ret bool
		return ret
	}
	return *o.HideTransactionDetailsInWebForm
}

// GetHideTransactionDetailsInWebFormOk returns a tuple with the HideTransactionDetailsInWebForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetHideTransactionDetailsInWebFormOk() (*bool, bool) {
	if o == nil || o.HideTransactionDetailsInWebForm == nil {
		return nil, false
	}
	return o.HideTransactionDetailsInWebForm, true
}

// HasHideTransactionDetailsInWebForm returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasHideTransactionDetailsInWebForm() bool {
	if o != nil && o.HideTransactionDetailsInWebForm != nil {
		return true
	}

	return false
}

// SetHideTransactionDetailsInWebForm gets a reference to the given bool and assigns it to the HideTransactionDetailsInWebForm field.
func (o *RequestSepaMoneyTransferParams) SetHideTransactionDetailsInWebForm(v bool) {
	o.HideTransactionDetailsInWebForm = &v
}

// GetMultiStepAuthentication returns the MultiStepAuthentication field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetMultiStepAuthentication() MultiStepAuthenticationCallback {
	if o == nil || o.MultiStepAuthentication == nil {
		var ret MultiStepAuthenticationCallback
		return ret
	}
	return *o.MultiStepAuthentication
}

// GetMultiStepAuthenticationOk returns a tuple with the MultiStepAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetMultiStepAuthenticationOk() (*MultiStepAuthenticationCallback, bool) {
	if o == nil || o.MultiStepAuthentication == nil {
		return nil, false
	}
	return o.MultiStepAuthentication, true
}

// HasMultiStepAuthentication returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasMultiStepAuthentication() bool {
	if o != nil && o.MultiStepAuthentication != nil {
		return true
	}

	return false
}

// SetMultiStepAuthentication gets a reference to the given MultiStepAuthenticationCallback and assigns it to the MultiStepAuthentication field.
func (o *RequestSepaMoneyTransferParams) SetMultiStepAuthentication(v MultiStepAuthenticationCallback) {
	o.MultiStepAuthentication = &v
}

// GetStorePin returns the StorePin field value if set, zero value otherwise.
func (o *RequestSepaMoneyTransferParams) GetStorePin() bool {
	if o == nil || o.StorePin == nil {
		var ret bool
		return ret
	}
	return *o.StorePin
}

// GetStorePinOk returns a tuple with the StorePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSepaMoneyTransferParams) GetStorePinOk() (*bool, bool) {
	if o == nil || o.StorePin == nil {
		return nil, false
	}
	return o.StorePin, true
}

// HasStorePin returns a boolean if a field has been set.
func (o *RequestSepaMoneyTransferParams) HasStorePin() bool {
	if o != nil && o.StorePin != nil {
		return true
	}

	return false
}

// SetStorePin gets a reference to the given bool and assigns it to the StorePin field.
func (o *RequestSepaMoneyTransferParams) SetStorePin(v bool) {
	o.StorePin = &v
}

func (o RequestSepaMoneyTransferParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecipientName != nil {
		toSerialize["recipientName"] = o.RecipientName
	}
	if o.RecipientIban != nil {
		toSerialize["recipientIban"] = o.RecipientIban
	}
	if o.RecipientBic != nil {
		toSerialize["recipientBic"] = o.RecipientBic
	}
	if o.ClearingAccountId != nil {
		toSerialize["clearingAccountId"] = o.ClearingAccountId
	}
	if o.EndToEndId != nil {
		toSerialize["endToEndId"] = o.EndToEndId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	if o.SepaPurposeCode != nil {
		toSerialize["sepaPurposeCode"] = o.SepaPurposeCode
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if o.BankingPin != nil {
		toSerialize["bankingPin"] = o.BankingPin
	}
	if o.StoreSecrets != nil {
		toSerialize["storeSecrets"] = o.StoreSecrets
	}
	if o.TwoStepProcedureId != nil {
		toSerialize["twoStepProcedureId"] = o.TwoStepProcedureId
	}
	if o.ExecutionDate != nil {
		toSerialize["executionDate"] = o.ExecutionDate
	}
	if o.SingleBooking != nil {
		toSerialize["singleBooking"] = o.SingleBooking
	}
	if o.AdditionalMoneyTransfers != nil {
		toSerialize["additionalMoneyTransfers"] = o.AdditionalMoneyTransfers
	}
	if o.ChallengeResponse != nil {
		toSerialize["challengeResponse"] = o.ChallengeResponse
	}
	if o.HideTransactionDetailsInWebForm != nil {
		toSerialize["hideTransactionDetailsInWebForm"] = o.HideTransactionDetailsInWebForm
	}
	if o.MultiStepAuthentication != nil {
		toSerialize["multiStepAuthentication"] = o.MultiStepAuthentication
	}
	if o.StorePin != nil {
		toSerialize["storePin"] = o.StorePin
	}
	return json.Marshal(toSerialize)
}

type NullableRequestSepaMoneyTransferParams struct {
	value *RequestSepaMoneyTransferParams
	isSet bool
}

func (v NullableRequestSepaMoneyTransferParams) Get() *RequestSepaMoneyTransferParams {
	return v.value
}

func (v *NullableRequestSepaMoneyTransferParams) Set(val *RequestSepaMoneyTransferParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSepaMoneyTransferParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSepaMoneyTransferParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSepaMoneyTransferParams(val *RequestSepaMoneyTransferParams) *NullableRequestSepaMoneyTransferParams {
	return &NullableRequestSepaMoneyTransferParams{value: val, isSet: true}
}

func (v NullableRequestSepaMoneyTransferParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSepaMoneyTransferParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



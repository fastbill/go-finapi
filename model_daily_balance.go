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

// DailyBalance Balance data for a single day
type DailyBalance struct {
	// Date in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	Date string `json:"date"`
	// Calculated balance at the end of day, across all regarded accounts. Note that the balance may not always add up to the income/spending of the day. This happens when a bank reports a balance that includes transactions which the bank didn't (yet) deliver. In any case, it is recommended to rely on the balance rather than on calculations based on the income/spending.
	Balance float64 `json:"balance"`
	// The sum of income of the given day, based on the 'transactions', across all regarded accounts.
	Income float64 `json:"income"`
	// The sum of spending of the given day, based on the 'transactions', across all regarded accounts. Note that this is an absolute (i.e. positive) value.
	Spending float64 `json:"spending"`
	// Sometimes finAPI may detect deviations between the bank reported account balance and the set of transactions received from the bank. This is an expected behaviour when an account has not been updated for a while, as banks provide only a limited history of transactions. In such cases, finAPI adds an adjusting entry (see the field Transaction.isAdjustingEntry), which will be contained in the 'transactions' list, just as any other transaction.<br/><br/>However, if an account was regularly updated and gaps in the transaction history are not expected, then finAPI will fix such deviations by adding an internal adjusting entry. These internal entries are not visible in the API and will not be contained in the 'transactions' list, and thus also not regarded for the calculations of 'income' and 'spending'. They are however regarded for the calculation of the 'balance'.<br/><br/>As long as you don't do your own balance calculations, you do not need to regard this field here; The 'balance' will always be correct. But if you do your own calculations, then you should not only regard the 'income' and 'spending', but this field as well.<br/><br/>Note that unlike the 'income' and 'spending', this field can have a positive or negative value.
	InternalAdjustingEntries float64 `json:"internalAdjustingEntries"`
	// Identifiers of the transactions for the given day
	Transactions []int64 `json:"transactions"`
}

// NewDailyBalance instantiates a new DailyBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyBalance(date string, balance float64, income float64, spending float64, internalAdjustingEntries float64, transactions []int64, ) *DailyBalance {
	this := DailyBalance{}
	this.Date = date
	this.Balance = balance
	this.Income = income
	this.Spending = spending
	this.InternalAdjustingEntries = internalAdjustingEntries
	this.Transactions = transactions
	return &this
}

// NewDailyBalanceWithDefaults instantiates a new DailyBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyBalanceWithDefaults() *DailyBalance {
	this := DailyBalance{}
	return &this
}

// GetDate returns the Date field value
func (o *DailyBalance) GetDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *DailyBalance) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *DailyBalance) SetDate(v string) {
	o.Date = v
}

// GetBalance returns the Balance field value
func (o *DailyBalance) GetBalance() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *DailyBalance) GetBalanceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *DailyBalance) SetBalance(v float64) {
	o.Balance = v
}

// GetIncome returns the Income field value
func (o *DailyBalance) GetIncome() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Income
}

// GetIncomeOk returns a tuple with the Income field value
// and a boolean to check if the value has been set.
func (o *DailyBalance) GetIncomeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Income, true
}

// SetIncome sets field value
func (o *DailyBalance) SetIncome(v float64) {
	o.Income = v
}

// GetSpending returns the Spending field value
func (o *DailyBalance) GetSpending() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.Spending
}

// GetSpendingOk returns a tuple with the Spending field value
// and a boolean to check if the value has been set.
func (o *DailyBalance) GetSpendingOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Spending, true
}

// SetSpending sets field value
func (o *DailyBalance) SetSpending(v float64) {
	o.Spending = v
}

// GetInternalAdjustingEntries returns the InternalAdjustingEntries field value
func (o *DailyBalance) GetInternalAdjustingEntries() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.InternalAdjustingEntries
}

// GetInternalAdjustingEntriesOk returns a tuple with the InternalAdjustingEntries field value
// and a boolean to check if the value has been set.
func (o *DailyBalance) GetInternalAdjustingEntriesOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InternalAdjustingEntries, true
}

// SetInternalAdjustingEntries sets field value
func (o *DailyBalance) SetInternalAdjustingEntries(v float64) {
	o.InternalAdjustingEntries = v
}

// GetTransactions returns the Transactions field value
func (o *DailyBalance) GetTransactions() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *DailyBalance) GetTransactionsOk() (*[]int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *DailyBalance) SetTransactions(v []int64) {
	o.Transactions = v
}

func (o DailyBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if true {
		toSerialize["income"] = o.Income
	}
	if true {
		toSerialize["spending"] = o.Spending
	}
	if true {
		toSerialize["internalAdjustingEntries"] = o.InternalAdjustingEntries
	}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableDailyBalance struct {
	value *DailyBalance
	isSet bool
}

func (v NullableDailyBalance) Get() *DailyBalance {
	return v.value
}

func (v *NullableDailyBalance) Set(val *DailyBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyBalance(val *DailyBalance) *NullableDailyBalance {
	return &NullableDailyBalance{value: val, isSet: true}
}

func (v NullableDailyBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



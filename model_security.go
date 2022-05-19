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

// Security Container for a security position's data
type Security struct {
	// Identifier. Note: Whenever a security account is being updated, its security positions will be internally re-created, meaning that the identifier of a security position might change over time.
	Id int64 `json:"id"`
	// Security account identifier
	AccountId int64 `json:"accountId"`
	// Name
	Name NullableString `json:"name"`
	// ISIN
	Isin NullableString `json:"isin"`
	// WKN
	Wkn NullableString `json:"wkn"`
	// Quote
	Quote NullableFloat64 `json:"quote"`
	// Currency of quote
	QuoteCurrency NullableString `json:"quoteCurrency"`
	// <strong>Type:</strong> SecurityPositionQuoteType<br/> Type of quote. 'PERC' if quote is a percentage value, 'ACTU' if quote is the actual amount
	QuoteType NullableSecurityPositionQuoteType `json:"quoteType"`
	// Quote date in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	QuoteDate NullableString `json:"quoteDate"`
	// Value of quantity or nominal
	QuantityNominal NullableFloat64 `json:"quantityNominal"`
	// <strong>Type:</strong> SecurityPositionQuantityNominalType<br/> Type of quantity or nominal value. 'UNIT' if value is a quantity, 'FAMT' if value is the nominal amount
	QuantityNominalType NullableSecurityPositionQuantityNominalType `json:"quantityNominalType"`
	// Market value
	MarketValue NullableFloat64 `json:"marketValue"`
	// Currency of market value
	MarketValueCurrency NullableString `json:"marketValueCurrency"`
	// Entry quote
	EntryQuote NullableFloat64 `json:"entryQuote"`
	// Currency of entry quote
	EntryQuoteCurrency NullableString `json:"entryQuoteCurrency"`
	// Current profit or loss
	ProfitOrLoss NullableFloat64 `json:"profitOrLoss"`
}

// NewSecurity instantiates a new Security object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurity(id int64, accountId int64, name NullableString, isin NullableString, wkn NullableString, quote NullableFloat64, quoteCurrency NullableString, quoteType NullableSecurityPositionQuoteType, quoteDate NullableString, quantityNominal NullableFloat64, quantityNominalType NullableSecurityPositionQuantityNominalType, marketValue NullableFloat64, marketValueCurrency NullableString, entryQuote NullableFloat64, entryQuoteCurrency NullableString, profitOrLoss NullableFloat64, ) *Security {
	this := Security{}
	this.Id = id
	this.AccountId = accountId
	this.Name = name
	this.Isin = isin
	this.Wkn = wkn
	this.Quote = quote
	this.QuoteCurrency = quoteCurrency
	this.QuoteType = quoteType
	this.QuoteDate = quoteDate
	this.QuantityNominal = quantityNominal
	this.QuantityNominalType = quantityNominalType
	this.MarketValue = marketValue
	this.MarketValueCurrency = marketValueCurrency
	this.EntryQuote = entryQuote
	this.EntryQuoteCurrency = entryQuoteCurrency
	this.ProfitOrLoss = profitOrLoss
	return &this
}

// NewSecurityWithDefaults instantiates a new Security object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityWithDefaults() *Security {
	this := Security{}
	return &this
}

// GetId returns the Id field value
func (o *Security) GetId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Security) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Security) SetId(v int64) {
	o.Id = v
}

// GetAccountId returns the AccountId field value
func (o *Security) GetAccountId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Security) GetAccountIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Security) SetAccountId(v int64) {
	o.AccountId = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Security) SetName(v string) {
	o.Name.Set(&v)
}

// GetIsin returns the Isin field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetIsin() string {
	if o == nil || o.Isin.Get() == nil {
		var ret string
		return ret
	}

	return *o.Isin.Get()
}

// GetIsinOk returns a tuple with the Isin field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetIsinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Isin.Get(), o.Isin.IsSet()
}

// SetIsin sets field value
func (o *Security) SetIsin(v string) {
	o.Isin.Set(&v)
}

// GetWkn returns the Wkn field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetWkn() string {
	if o == nil || o.Wkn.Get() == nil {
		var ret string
		return ret
	}

	return *o.Wkn.Get()
}

// GetWknOk returns a tuple with the Wkn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetWknOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Wkn.Get(), o.Wkn.IsSet()
}

// SetWkn sets field value
func (o *Security) SetWkn(v string) {
	o.Wkn.Set(&v)
}

// GetQuote returns the Quote field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Security) GetQuote() float64 {
	if o == nil || o.Quote.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Quote.Get()
}

// GetQuoteOk returns a tuple with the Quote field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetQuoteOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Quote.Get(), o.Quote.IsSet()
}

// SetQuote sets field value
func (o *Security) SetQuote(v float64) {
	o.Quote.Set(&v)
}

// GetQuoteCurrency returns the QuoteCurrency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetQuoteCurrency() string {
	if o == nil || o.QuoteCurrency.Get() == nil {
		var ret string
		return ret
	}

	return *o.QuoteCurrency.Get()
}

// GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetQuoteCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.QuoteCurrency.Get(), o.QuoteCurrency.IsSet()
}

// SetQuoteCurrency sets field value
func (o *Security) SetQuoteCurrency(v string) {
	o.QuoteCurrency.Set(&v)
}

// GetQuoteType returns the QuoteType field value
// If the value is explicit nil, the zero value for SecurityPositionQuoteType will be returned
func (o *Security) GetQuoteType() SecurityPositionQuoteType {
	if o == nil || o.QuoteType.Get() == nil {
		var ret SecurityPositionQuoteType
		return ret
	}

	return *o.QuoteType.Get()
}

// GetQuoteTypeOk returns a tuple with the QuoteType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetQuoteTypeOk() (*SecurityPositionQuoteType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.QuoteType.Get(), o.QuoteType.IsSet()
}

// SetQuoteType sets field value
func (o *Security) SetQuoteType(v SecurityPositionQuoteType) {
	o.QuoteType.Set(&v)
}

// GetQuoteDate returns the QuoteDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetQuoteDate() string {
	if o == nil || o.QuoteDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.QuoteDate.Get()
}

// GetQuoteDateOk returns a tuple with the QuoteDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetQuoteDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.QuoteDate.Get(), o.QuoteDate.IsSet()
}

// SetQuoteDate sets field value
func (o *Security) SetQuoteDate(v string) {
	o.QuoteDate.Set(&v)
}

// GetQuantityNominal returns the QuantityNominal field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Security) GetQuantityNominal() float64 {
	if o == nil || o.QuantityNominal.Get() == nil {
		var ret float64
		return ret
	}

	return *o.QuantityNominal.Get()
}

// GetQuantityNominalOk returns a tuple with the QuantityNominal field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetQuantityNominalOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.QuantityNominal.Get(), o.QuantityNominal.IsSet()
}

// SetQuantityNominal sets field value
func (o *Security) SetQuantityNominal(v float64) {
	o.QuantityNominal.Set(&v)
}

// GetQuantityNominalType returns the QuantityNominalType field value
// If the value is explicit nil, the zero value for SecurityPositionQuantityNominalType will be returned
func (o *Security) GetQuantityNominalType() SecurityPositionQuantityNominalType {
	if o == nil || o.QuantityNominalType.Get() == nil {
		var ret SecurityPositionQuantityNominalType
		return ret
	}

	return *o.QuantityNominalType.Get()
}

// GetQuantityNominalTypeOk returns a tuple with the QuantityNominalType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetQuantityNominalTypeOk() (*SecurityPositionQuantityNominalType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.QuantityNominalType.Get(), o.QuantityNominalType.IsSet()
}

// SetQuantityNominalType sets field value
func (o *Security) SetQuantityNominalType(v SecurityPositionQuantityNominalType) {
	o.QuantityNominalType.Set(&v)
}

// GetMarketValue returns the MarketValue field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Security) GetMarketValue() float64 {
	if o == nil || o.MarketValue.Get() == nil {
		var ret float64
		return ret
	}

	return *o.MarketValue.Get()
}

// GetMarketValueOk returns a tuple with the MarketValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetMarketValueOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MarketValue.Get(), o.MarketValue.IsSet()
}

// SetMarketValue sets field value
func (o *Security) SetMarketValue(v float64) {
	o.MarketValue.Set(&v)
}

// GetMarketValueCurrency returns the MarketValueCurrency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetMarketValueCurrency() string {
	if o == nil || o.MarketValueCurrency.Get() == nil {
		var ret string
		return ret
	}

	return *o.MarketValueCurrency.Get()
}

// GetMarketValueCurrencyOk returns a tuple with the MarketValueCurrency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetMarketValueCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MarketValueCurrency.Get(), o.MarketValueCurrency.IsSet()
}

// SetMarketValueCurrency sets field value
func (o *Security) SetMarketValueCurrency(v string) {
	o.MarketValueCurrency.Set(&v)
}

// GetEntryQuote returns the EntryQuote field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Security) GetEntryQuote() float64 {
	if o == nil || o.EntryQuote.Get() == nil {
		var ret float64
		return ret
	}

	return *o.EntryQuote.Get()
}

// GetEntryQuoteOk returns a tuple with the EntryQuote field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetEntryQuoteOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EntryQuote.Get(), o.EntryQuote.IsSet()
}

// SetEntryQuote sets field value
func (o *Security) SetEntryQuote(v float64) {
	o.EntryQuote.Set(&v)
}

// GetEntryQuoteCurrency returns the EntryQuoteCurrency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Security) GetEntryQuoteCurrency() string {
	if o == nil || o.EntryQuoteCurrency.Get() == nil {
		var ret string
		return ret
	}

	return *o.EntryQuoteCurrency.Get()
}

// GetEntryQuoteCurrencyOk returns a tuple with the EntryQuoteCurrency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetEntryQuoteCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EntryQuoteCurrency.Get(), o.EntryQuoteCurrency.IsSet()
}

// SetEntryQuoteCurrency sets field value
func (o *Security) SetEntryQuoteCurrency(v string) {
	o.EntryQuoteCurrency.Set(&v)
}

// GetProfitOrLoss returns the ProfitOrLoss field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *Security) GetProfitOrLoss() float64 {
	if o == nil || o.ProfitOrLoss.Get() == nil {
		var ret float64
		return ret
	}

	return *o.ProfitOrLoss.Get()
}

// GetProfitOrLossOk returns a tuple with the ProfitOrLoss field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Security) GetProfitOrLossOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProfitOrLoss.Get(), o.ProfitOrLoss.IsSet()
}

// SetProfitOrLoss sets field value
func (o *Security) SetProfitOrLoss(v float64) {
	o.ProfitOrLoss.Set(&v)
}

func (o Security) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["isin"] = o.Isin.Get()
	}
	if true {
		toSerialize["wkn"] = o.Wkn.Get()
	}
	if true {
		toSerialize["quote"] = o.Quote.Get()
	}
	if true {
		toSerialize["quoteCurrency"] = o.QuoteCurrency.Get()
	}
	if true {
		toSerialize["quoteType"] = o.QuoteType.Get()
	}
	if true {
		toSerialize["quoteDate"] = o.QuoteDate.Get()
	}
	if true {
		toSerialize["quantityNominal"] = o.QuantityNominal.Get()
	}
	if true {
		toSerialize["quantityNominalType"] = o.QuantityNominalType.Get()
	}
	if true {
		toSerialize["marketValue"] = o.MarketValue.Get()
	}
	if true {
		toSerialize["marketValueCurrency"] = o.MarketValueCurrency.Get()
	}
	if true {
		toSerialize["entryQuote"] = o.EntryQuote.Get()
	}
	if true {
		toSerialize["entryQuoteCurrency"] = o.EntryQuoteCurrency.Get()
	}
	if true {
		toSerialize["profitOrLoss"] = o.ProfitOrLoss.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSecurity struct {
	value *Security
	isSet bool
}

func (v NullableSecurity) Get() *Security {
	return v.value
}

func (v *NullableSecurity) Set(val *Security) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurity(val *Security) *NullableSecurity {
	return &NullableSecurity{value: val, isSet: true}
}

func (v NullableSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



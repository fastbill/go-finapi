# DirectDebitOrderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CounterpartName** | **string** | Name of the counterpart. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the counterpart name. Even if the name does not depict the actual registered account holder of the target account, the order might still be successful. | 
**CounterpartIban** | **string** | IBAN of the counterpart&#39;s account. | 
**CounterpartBic** | Pointer to **string** | BIC of the counterpart&#39;s account. Only required when there is no &#39;IBAN_ONLY&#39;-capability in the respective account/interface combination that is to be used when submitting the payment. | [optional] 
**Amount** | **float64** | The amount of the payment. Must be a positive decimal number with at most two decimal places. When debiting money using the FINTS_SERVER or WEB_SCRAPER interface, the currency is always EUR. | 
**Purpose** | Pointer to **string** | The purpose of the transfer transaction | [optional] 
**SepaPurposeCode** | Pointer to **string** | SEPA purpose code, according to ISO 20022, external codes set.&lt;br/&gt;Please note that the SEPA purpose code may be ignored by some banks. | [optional] 
**EndToEndId** | Pointer to **string** | End-To-End ID for the transfer transaction | [optional] 
**MandateId** | **string** | Mandate ID that this direct debit order is based on. | 
**MandateDate** | **string** | Date of the mandate that this direct debit order is based on, in the format &#39;YYYY-MM-DD&#39; | 
**CreditorId** | **string** | Creditor ID of the source account&#39;s holder | 
**CounterpartAddress** | Pointer to **string** | The postal address of the debtor. This should be defined for direct debits created for debtors outside of the european union. | [optional] 
**CounterpartCountry** | Pointer to [**ISO3166Alpha2Codes**](ISO3166Alpha2Codes.md) | &lt;strong&gt;Type:&lt;/strong&gt; ISO3166Alpha2Codes&lt;br/&gt; The ISO 3166 ALPHA-2 country code of the debtor’s address. Examples: &#39;GB&#39; for the United Kingdom or &#39;CH&#39; for Switzerland. This should be defined for direct debits created for debtors outside of the european union. | [optional] 

## Methods

### NewDirectDebitOrderParams

`func NewDirectDebitOrderParams(counterpartName string, counterpartIban string, amount float64, mandateId string, mandateDate string, creditorId string, ) *DirectDebitOrderParams`

NewDirectDebitOrderParams instantiates a new DirectDebitOrderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDebitOrderParamsWithDefaults

`func NewDirectDebitOrderParamsWithDefaults() *DirectDebitOrderParams`

NewDirectDebitOrderParamsWithDefaults instantiates a new DirectDebitOrderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounterpartName

`func (o *DirectDebitOrderParams) GetCounterpartName() string`

GetCounterpartName returns the CounterpartName field if non-nil, zero value otherwise.

### GetCounterpartNameOk

`func (o *DirectDebitOrderParams) GetCounterpartNameOk() (*string, bool)`

GetCounterpartNameOk returns a tuple with the CounterpartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartName

`func (o *DirectDebitOrderParams) SetCounterpartName(v string)`

SetCounterpartName sets CounterpartName field to given value.


### GetCounterpartIban

`func (o *DirectDebitOrderParams) GetCounterpartIban() string`

GetCounterpartIban returns the CounterpartIban field if non-nil, zero value otherwise.

### GetCounterpartIbanOk

`func (o *DirectDebitOrderParams) GetCounterpartIbanOk() (*string, bool)`

GetCounterpartIbanOk returns a tuple with the CounterpartIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartIban

`func (o *DirectDebitOrderParams) SetCounterpartIban(v string)`

SetCounterpartIban sets CounterpartIban field to given value.


### GetCounterpartBic

`func (o *DirectDebitOrderParams) GetCounterpartBic() string`

GetCounterpartBic returns the CounterpartBic field if non-nil, zero value otherwise.

### GetCounterpartBicOk

`func (o *DirectDebitOrderParams) GetCounterpartBicOk() (*string, bool)`

GetCounterpartBicOk returns a tuple with the CounterpartBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBic

`func (o *DirectDebitOrderParams) SetCounterpartBic(v string)`

SetCounterpartBic sets CounterpartBic field to given value.

### HasCounterpartBic

`func (o *DirectDebitOrderParams) HasCounterpartBic() bool`

HasCounterpartBic returns a boolean if a field has been set.

### GetAmount

`func (o *DirectDebitOrderParams) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DirectDebitOrderParams) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DirectDebitOrderParams) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPurpose

`func (o *DirectDebitOrderParams) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *DirectDebitOrderParams) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *DirectDebitOrderParams) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *DirectDebitOrderParams) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSepaPurposeCode

`func (o *DirectDebitOrderParams) GetSepaPurposeCode() string`

GetSepaPurposeCode returns the SepaPurposeCode field if non-nil, zero value otherwise.

### GetSepaPurposeCodeOk

`func (o *DirectDebitOrderParams) GetSepaPurposeCodeOk() (*string, bool)`

GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaPurposeCode

`func (o *DirectDebitOrderParams) SetSepaPurposeCode(v string)`

SetSepaPurposeCode sets SepaPurposeCode field to given value.

### HasSepaPurposeCode

`func (o *DirectDebitOrderParams) HasSepaPurposeCode() bool`

HasSepaPurposeCode returns a boolean if a field has been set.

### GetEndToEndId

`func (o *DirectDebitOrderParams) GetEndToEndId() string`

GetEndToEndId returns the EndToEndId field if non-nil, zero value otherwise.

### GetEndToEndIdOk

`func (o *DirectDebitOrderParams) GetEndToEndIdOk() (*string, bool)`

GetEndToEndIdOk returns a tuple with the EndToEndId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndToEndId

`func (o *DirectDebitOrderParams) SetEndToEndId(v string)`

SetEndToEndId sets EndToEndId field to given value.

### HasEndToEndId

`func (o *DirectDebitOrderParams) HasEndToEndId() bool`

HasEndToEndId returns a boolean if a field has been set.

### GetMandateId

`func (o *DirectDebitOrderParams) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *DirectDebitOrderParams) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *DirectDebitOrderParams) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.


### GetMandateDate

`func (o *DirectDebitOrderParams) GetMandateDate() string`

GetMandateDate returns the MandateDate field if non-nil, zero value otherwise.

### GetMandateDateOk

`func (o *DirectDebitOrderParams) GetMandateDateOk() (*string, bool)`

GetMandateDateOk returns a tuple with the MandateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateDate

`func (o *DirectDebitOrderParams) SetMandateDate(v string)`

SetMandateDate sets MandateDate field to given value.


### GetCreditorId

`func (o *DirectDebitOrderParams) GetCreditorId() string`

GetCreditorId returns the CreditorId field if non-nil, zero value otherwise.

### GetCreditorIdOk

`func (o *DirectDebitOrderParams) GetCreditorIdOk() (*string, bool)`

GetCreditorIdOk returns a tuple with the CreditorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditorId

`func (o *DirectDebitOrderParams) SetCreditorId(v string)`

SetCreditorId sets CreditorId field to given value.


### GetCounterpartAddress

`func (o *DirectDebitOrderParams) GetCounterpartAddress() string`

GetCounterpartAddress returns the CounterpartAddress field if non-nil, zero value otherwise.

### GetCounterpartAddressOk

`func (o *DirectDebitOrderParams) GetCounterpartAddressOk() (*string, bool)`

GetCounterpartAddressOk returns a tuple with the CounterpartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartAddress

`func (o *DirectDebitOrderParams) SetCounterpartAddress(v string)`

SetCounterpartAddress sets CounterpartAddress field to given value.

### HasCounterpartAddress

`func (o *DirectDebitOrderParams) HasCounterpartAddress() bool`

HasCounterpartAddress returns a boolean if a field has been set.

### GetCounterpartCountry

`func (o *DirectDebitOrderParams) GetCounterpartCountry() ISO3166Alpha2Codes`

GetCounterpartCountry returns the CounterpartCountry field if non-nil, zero value otherwise.

### GetCounterpartCountryOk

`func (o *DirectDebitOrderParams) GetCounterpartCountryOk() (*ISO3166Alpha2Codes, bool)`

GetCounterpartCountryOk returns a tuple with the CounterpartCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartCountry

`func (o *DirectDebitOrderParams) SetCounterpartCountry(v ISO3166Alpha2Codes)`

SetCounterpartCountry sets CounterpartCountry field to given value.

### HasCounterpartCountry

`func (o *DirectDebitOrderParams) HasCounterpartCountry() bool`

HasCounterpartCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



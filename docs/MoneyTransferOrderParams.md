# MoneyTransferOrderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CounterpartName** | **string** | Name of the counterpart. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the counterpart name. Even if the name does not depict the actual registered account holder of the target account, the order might still be successful. | 
**CounterpartIban** | **string** | IBAN of the counterpart&#39;s account. | 
**CounterpartBic** | Pointer to **string** | BIC of the counterpart&#39;s account. Only required when there is no &#39;IBAN_ONLY&#39;-capability in the respective account/interface combination that is to be used when submitting the payment. | [optional] 
**Amount** | **float64** | The amount of the payment. Must be a positive decimal number with at most two decimal places. For money transfers over the XS2A interface, finAPI will interpret the amount to be in the currency of the related account. For money transfers over other interfaces (FINTS_SERVER, WEB_SCRAPER), as well as for standalone money transfers (finAPI Payment product) over all interfaces (FINTS_SERVER, WEB_SCRAPER, XS2A), finAPI will consider the amount to be in EUR. | 
**Purpose** | Pointer to **string** | The purpose of the transfer transaction | [optional] 
**SepaPurposeCode** | Pointer to **string** | SEPA purpose code, according to ISO 20022, external codes set.&lt;br/&gt;Please note that the SEPA purpose code may be ignored by some banks. | [optional] 
**EndToEndId** | Pointer to **string** | End-To-End ID for the transfer transaction | [optional] 

## Methods

### NewMoneyTransferOrderParams

`func NewMoneyTransferOrderParams(counterpartName string, counterpartIban string, amount float64, ) *MoneyTransferOrderParams`

NewMoneyTransferOrderParams instantiates a new MoneyTransferOrderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyTransferOrderParamsWithDefaults

`func NewMoneyTransferOrderParamsWithDefaults() *MoneyTransferOrderParams`

NewMoneyTransferOrderParamsWithDefaults instantiates a new MoneyTransferOrderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounterpartName

`func (o *MoneyTransferOrderParams) GetCounterpartName() string`

GetCounterpartName returns the CounterpartName field if non-nil, zero value otherwise.

### GetCounterpartNameOk

`func (o *MoneyTransferOrderParams) GetCounterpartNameOk() (*string, bool)`

GetCounterpartNameOk returns a tuple with the CounterpartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartName

`func (o *MoneyTransferOrderParams) SetCounterpartName(v string)`

SetCounterpartName sets CounterpartName field to given value.


### GetCounterpartIban

`func (o *MoneyTransferOrderParams) GetCounterpartIban() string`

GetCounterpartIban returns the CounterpartIban field if non-nil, zero value otherwise.

### GetCounterpartIbanOk

`func (o *MoneyTransferOrderParams) GetCounterpartIbanOk() (*string, bool)`

GetCounterpartIbanOk returns a tuple with the CounterpartIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartIban

`func (o *MoneyTransferOrderParams) SetCounterpartIban(v string)`

SetCounterpartIban sets CounterpartIban field to given value.


### GetCounterpartBic

`func (o *MoneyTransferOrderParams) GetCounterpartBic() string`

GetCounterpartBic returns the CounterpartBic field if non-nil, zero value otherwise.

### GetCounterpartBicOk

`func (o *MoneyTransferOrderParams) GetCounterpartBicOk() (*string, bool)`

GetCounterpartBicOk returns a tuple with the CounterpartBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBic

`func (o *MoneyTransferOrderParams) SetCounterpartBic(v string)`

SetCounterpartBic sets CounterpartBic field to given value.

### HasCounterpartBic

`func (o *MoneyTransferOrderParams) HasCounterpartBic() bool`

HasCounterpartBic returns a boolean if a field has been set.

### GetAmount

`func (o *MoneyTransferOrderParams) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MoneyTransferOrderParams) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MoneyTransferOrderParams) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPurpose

`func (o *MoneyTransferOrderParams) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *MoneyTransferOrderParams) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *MoneyTransferOrderParams) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *MoneyTransferOrderParams) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSepaPurposeCode

`func (o *MoneyTransferOrderParams) GetSepaPurposeCode() string`

GetSepaPurposeCode returns the SepaPurposeCode field if non-nil, zero value otherwise.

### GetSepaPurposeCodeOk

`func (o *MoneyTransferOrderParams) GetSepaPurposeCodeOk() (*string, bool)`

GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaPurposeCode

`func (o *MoneyTransferOrderParams) SetSepaPurposeCode(v string)`

SetSepaPurposeCode sets SepaPurposeCode field to given value.

### HasSepaPurposeCode

`func (o *MoneyTransferOrderParams) HasSepaPurposeCode() bool`

HasSepaPurposeCode returns a boolean if a field has been set.

### GetEndToEndId

`func (o *MoneyTransferOrderParams) GetEndToEndId() string`

GetEndToEndId returns the EndToEndId field if non-nil, zero value otherwise.

### GetEndToEndIdOk

`func (o *MoneyTransferOrderParams) GetEndToEndIdOk() (*string, bool)`

GetEndToEndIdOk returns a tuple with the EndToEndId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndToEndId

`func (o *MoneyTransferOrderParams) SetEndToEndId(v string)`

SetEndToEndId sets EndToEndId field to given value.

### HasEndToEndId

`func (o *MoneyTransferOrderParams) HasEndToEndId() bool`

HasEndToEndId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateStandingOrderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int64** | Identifier of the account that should be used for the standing order. If you want to do a standalone standing order (i.e. for an account that is not imported in finAPI) leave this field unset and instead use the field &#39;iban&#39;. | [optional] 
**Iban** | Pointer to **string** | IBAN of the account that should be used for the standing order. Use this field only if you want to do a standalone standing order (i.e. for an account that is not imported in finAPI) otherwise, use the &#39;accountId&#39; field and leave this field unset. | [optional] 
**CounterpartName** | **string** | Name of the counterpart. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the counterpart name. Even if the name does not depict the actual registered account holder of the target account, the order might still be successful. | 
**CounterpartIban** | **string** | IBAN of the counterpart&#39;s account. | 
**Amount** | **float64** | The amount of the standing order. Must be a positive decimal number with at most two decimal places. | 
**Currency** | [**Currency**](Currency.md) | &lt;strong&gt;Type:&lt;/strong&gt; Currency&lt;br/&gt; The currency code of the &#39;amount&#39;. To be provided in the ISO 4217 Alpha 3 format. | 
**Purpose** | Pointer to **string** | The purpose of the standing order. | [optional] 
**SepaPurposeCode** | Pointer to **string** | SEPA purpose code, according to ISO 20022, external codes set.&lt;br/&gt;Please note that the SEPA purpose code may be ignored by some banks. | [optional] 
**EndToEndId** | Pointer to **string** | End-To-End ID of the standing order | [optional] 
**StartDate** | **string** | Start date of the standing order. Date must be in the future (at least tomorrow). To be provided in the format &#39;YYYY-MM-DD&#39;. | 
**EndDate** | Pointer to **string** | Termination date of the standing order. If provided, it must be after the &#39;startDate&#39;. If not provided, then the standing order will have no termination. To be provided in the format &#39;YYYY-MM-DD&#39;. | [optional] 
**Frequency** | [**StandingOrderFrequency**](StandingOrderFrequency.md) | &lt;strong&gt;Type:&lt;/strong&gt; StandingOrderFrequency&lt;br/&gt; The frequency of the standing order | 
**DayOfExecution** | Pointer to **int32** | Desired day of execution. If not provided, it will be derived from the &#39;startDate&#39;. Use &#39;31&#39; for ultimo. Allowed values: 1-31. | [optional] 

## Methods

### NewCreateStandingOrderParams

`func NewCreateStandingOrderParams(counterpartName string, counterpartIban string, amount float64, currency Currency, startDate string, frequency StandingOrderFrequency, ) *CreateStandingOrderParams`

NewCreateStandingOrderParams instantiates a new CreateStandingOrderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandingOrderParamsWithDefaults

`func NewCreateStandingOrderParamsWithDefaults() *CreateStandingOrderParams`

NewCreateStandingOrderParamsWithDefaults instantiates a new CreateStandingOrderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateStandingOrderParams) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateStandingOrderParams) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateStandingOrderParams) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateStandingOrderParams) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetIban

`func (o *CreateStandingOrderParams) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *CreateStandingOrderParams) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *CreateStandingOrderParams) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *CreateStandingOrderParams) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetCounterpartName

`func (o *CreateStandingOrderParams) GetCounterpartName() string`

GetCounterpartName returns the CounterpartName field if non-nil, zero value otherwise.

### GetCounterpartNameOk

`func (o *CreateStandingOrderParams) GetCounterpartNameOk() (*string, bool)`

GetCounterpartNameOk returns a tuple with the CounterpartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartName

`func (o *CreateStandingOrderParams) SetCounterpartName(v string)`

SetCounterpartName sets CounterpartName field to given value.


### GetCounterpartIban

`func (o *CreateStandingOrderParams) GetCounterpartIban() string`

GetCounterpartIban returns the CounterpartIban field if non-nil, zero value otherwise.

### GetCounterpartIbanOk

`func (o *CreateStandingOrderParams) GetCounterpartIbanOk() (*string, bool)`

GetCounterpartIbanOk returns a tuple with the CounterpartIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartIban

`func (o *CreateStandingOrderParams) SetCounterpartIban(v string)`

SetCounterpartIban sets CounterpartIban field to given value.


### GetAmount

`func (o *CreateStandingOrderParams) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateStandingOrderParams) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateStandingOrderParams) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *CreateStandingOrderParams) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateStandingOrderParams) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateStandingOrderParams) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetPurpose

`func (o *CreateStandingOrderParams) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CreateStandingOrderParams) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CreateStandingOrderParams) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *CreateStandingOrderParams) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSepaPurposeCode

`func (o *CreateStandingOrderParams) GetSepaPurposeCode() string`

GetSepaPurposeCode returns the SepaPurposeCode field if non-nil, zero value otherwise.

### GetSepaPurposeCodeOk

`func (o *CreateStandingOrderParams) GetSepaPurposeCodeOk() (*string, bool)`

GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaPurposeCode

`func (o *CreateStandingOrderParams) SetSepaPurposeCode(v string)`

SetSepaPurposeCode sets SepaPurposeCode field to given value.

### HasSepaPurposeCode

`func (o *CreateStandingOrderParams) HasSepaPurposeCode() bool`

HasSepaPurposeCode returns a boolean if a field has been set.

### GetEndToEndId

`func (o *CreateStandingOrderParams) GetEndToEndId() string`

GetEndToEndId returns the EndToEndId field if non-nil, zero value otherwise.

### GetEndToEndIdOk

`func (o *CreateStandingOrderParams) GetEndToEndIdOk() (*string, bool)`

GetEndToEndIdOk returns a tuple with the EndToEndId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndToEndId

`func (o *CreateStandingOrderParams) SetEndToEndId(v string)`

SetEndToEndId sets EndToEndId field to given value.

### HasEndToEndId

`func (o *CreateStandingOrderParams) HasEndToEndId() bool`

HasEndToEndId returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateStandingOrderParams) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateStandingOrderParams) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateStandingOrderParams) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *CreateStandingOrderParams) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateStandingOrderParams) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateStandingOrderParams) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateStandingOrderParams) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFrequency

`func (o *CreateStandingOrderParams) GetFrequency() StandingOrderFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *CreateStandingOrderParams) GetFrequencyOk() (*StandingOrderFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *CreateStandingOrderParams) SetFrequency(v StandingOrderFrequency)`

SetFrequency sets Frequency field to given value.


### GetDayOfExecution

`func (o *CreateStandingOrderParams) GetDayOfExecution() int32`

GetDayOfExecution returns the DayOfExecution field if non-nil, zero value otherwise.

### GetDayOfExecutionOk

`func (o *CreateStandingOrderParams) GetDayOfExecutionOk() (*int32, bool)`

GetDayOfExecutionOk returns a tuple with the DayOfExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfExecution

`func (o *CreateStandingOrderParams) SetDayOfExecution(v int32)`

SetDayOfExecution sets DayOfExecution field to given value.

### HasDayOfExecution

`func (o *CreateStandingOrderParams) HasDayOfExecution() bool`

HasDayOfExecution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



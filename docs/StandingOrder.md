# StandingOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Standing order identifier | 
**AccountId** | **NullableInt64** | Identifier of the account to which this standing order relates. This field is only set if it was specified upon creation of the standing order. | 
**Iban** | **NullableString** | IBAN of the account to which this standing order relates. This field is only set if it was specified upon creation of the standing order. | 
**Amount** | **float64** | Amount of the standing order, as absolute value. | 
**Currency** | [**Currency**](Currency.md) | &lt;strong&gt;Type:&lt;/strong&gt; Currency&lt;br/&gt; The currency code of the &#39;amount&#39;, in the ISO 4217 Alpha 3 format. | 
**StartDate** | **string** | Start date of the standing order, in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**EndDate** | **NullableString** | Termination date of the standing order, in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). If this field is not set, then the standing order has no termination date. | 
**Frequency** | [**StandingOrderFrequency**](StandingOrderFrequency.md) | &lt;strong&gt;Type:&lt;/strong&gt; StandingOrderFrequency&lt;br/&gt; The frequency of the standing order. | 
**DayOfExecution** | **NullableInt32** | Scheduled day of execution. &#39;31&#39; depicts ultimo. | 
**RequestDate** | **NullableString** | Time of when finAPI submitted this standing order to the bank, in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time) | 
**RequestCompletionDate** | **NullableString** | Time of when the submission of this standing order was finalized, in the format ‘YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). Note: When the submission of a standing order is finalized, it does not necessarily mean that the bank accepted the standing order. Please refer to the standing order’s &#39;status&#39; for its final status. | 
**Status** | [**OrderInitiationStatus**](OrderInitiationStatus.md) | &lt;strong&gt;Type:&lt;/strong&gt; OrderInitiationStatus&lt;br/&gt; Current standing order status:&lt;br/&gt; &amp;bull; OPEN: means that this standing order has been created in finAPI, but not yet submitted to the bank.&lt;br/&gt; &amp;bull; PENDING: means that this standing order has been submitted to the bank,  but has not been confirmed yet.&lt;br/&gt; &amp;bull; SUCCESSFUL: means that this standing order has been successfully initiated.&lt;br/&gt; &amp;bull; NOT_SUCCESSFUL: means that this standing order could not be initiated successfully.&lt;br/&gt; &amp;bull; DISCARDED: means that this standing order was discarded, either because another standing order was requested for the same account before this standing order was initiated and the bank does not support this, or because the user has aborted the initiation (when using finAPI&#39;s Web Form). | 
**BankMessage** | **NullableString** | The bank&#39;s response to the most recent request for this standing order. Note that this field may not always (or never) be set. Also, as long as the standing order has not reached its final status, this field can always change. | 

## Methods

### NewStandingOrder

`func NewStandingOrder(id int64, accountId NullableInt64, iban NullableString, amount float64, currency Currency, startDate string, endDate NullableString, frequency StandingOrderFrequency, dayOfExecution NullableInt32, requestDate NullableString, requestCompletionDate NullableString, status OrderInitiationStatus, bankMessage NullableString, ) *StandingOrder`

NewStandingOrder instantiates a new StandingOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandingOrderWithDefaults

`func NewStandingOrderWithDefaults() *StandingOrder`

NewStandingOrderWithDefaults instantiates a new StandingOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandingOrder) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandingOrder) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandingOrder) SetId(v int64)`

SetId sets Id field to given value.


### GetAccountId

`func (o *StandingOrder) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StandingOrder) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StandingOrder) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *StandingOrder) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *StandingOrder) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetIban

`func (o *StandingOrder) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *StandingOrder) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *StandingOrder) SetIban(v string)`

SetIban sets Iban field to given value.


### SetIbanNil

`func (o *StandingOrder) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *StandingOrder) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetAmount

`func (o *StandingOrder) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StandingOrder) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StandingOrder) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *StandingOrder) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *StandingOrder) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *StandingOrder) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetStartDate

`func (o *StandingOrder) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *StandingOrder) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *StandingOrder) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *StandingOrder) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *StandingOrder) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *StandingOrder) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *StandingOrder) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *StandingOrder) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetFrequency

`func (o *StandingOrder) GetFrequency() StandingOrderFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *StandingOrder) GetFrequencyOk() (*StandingOrderFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *StandingOrder) SetFrequency(v StandingOrderFrequency)`

SetFrequency sets Frequency field to given value.


### GetDayOfExecution

`func (o *StandingOrder) GetDayOfExecution() int32`

GetDayOfExecution returns the DayOfExecution field if non-nil, zero value otherwise.

### GetDayOfExecutionOk

`func (o *StandingOrder) GetDayOfExecutionOk() (*int32, bool)`

GetDayOfExecutionOk returns a tuple with the DayOfExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfExecution

`func (o *StandingOrder) SetDayOfExecution(v int32)`

SetDayOfExecution sets DayOfExecution field to given value.


### SetDayOfExecutionNil

`func (o *StandingOrder) SetDayOfExecutionNil(b bool)`

 SetDayOfExecutionNil sets the value for DayOfExecution to be an explicit nil

### UnsetDayOfExecution
`func (o *StandingOrder) UnsetDayOfExecution()`

UnsetDayOfExecution ensures that no value is present for DayOfExecution, not even an explicit nil
### GetRequestDate

`func (o *StandingOrder) GetRequestDate() string`

GetRequestDate returns the RequestDate field if non-nil, zero value otherwise.

### GetRequestDateOk

`func (o *StandingOrder) GetRequestDateOk() (*string, bool)`

GetRequestDateOk returns a tuple with the RequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDate

`func (o *StandingOrder) SetRequestDate(v string)`

SetRequestDate sets RequestDate field to given value.


### SetRequestDateNil

`func (o *StandingOrder) SetRequestDateNil(b bool)`

 SetRequestDateNil sets the value for RequestDate to be an explicit nil

### UnsetRequestDate
`func (o *StandingOrder) UnsetRequestDate()`

UnsetRequestDate ensures that no value is present for RequestDate, not even an explicit nil
### GetRequestCompletionDate

`func (o *StandingOrder) GetRequestCompletionDate() string`

GetRequestCompletionDate returns the RequestCompletionDate field if non-nil, zero value otherwise.

### GetRequestCompletionDateOk

`func (o *StandingOrder) GetRequestCompletionDateOk() (*string, bool)`

GetRequestCompletionDateOk returns a tuple with the RequestCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCompletionDate

`func (o *StandingOrder) SetRequestCompletionDate(v string)`

SetRequestCompletionDate sets RequestCompletionDate field to given value.


### SetRequestCompletionDateNil

`func (o *StandingOrder) SetRequestCompletionDateNil(b bool)`

 SetRequestCompletionDateNil sets the value for RequestCompletionDate to be an explicit nil

### UnsetRequestCompletionDate
`func (o *StandingOrder) UnsetRequestCompletionDate()`

UnsetRequestCompletionDate ensures that no value is present for RequestCompletionDate, not even an explicit nil
### GetStatus

`func (o *StandingOrder) GetStatus() OrderInitiationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StandingOrder) GetStatusOk() (*OrderInitiationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StandingOrder) SetStatus(v OrderInitiationStatus)`

SetStatus sets Status field to given value.


### GetBankMessage

`func (o *StandingOrder) GetBankMessage() string`

GetBankMessage returns the BankMessage field if non-nil, zero value otherwise.

### GetBankMessageOk

`func (o *StandingOrder) GetBankMessageOk() (*string, bool)`

GetBankMessageOk returns a tuple with the BankMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankMessage

`func (o *StandingOrder) SetBankMessage(v string)`

SetBankMessage sets BankMessage field to given value.


### SetBankMessageNil

`func (o *StandingOrder) SetBankMessageNil(b bool)`

 SetBankMessageNil sets the value for BankMessage to be an explicit nil

### UnsetBankMessage
`func (o *StandingOrder) UnsetBankMessage()`

UnsetBankMessage ensures that no value is present for BankMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



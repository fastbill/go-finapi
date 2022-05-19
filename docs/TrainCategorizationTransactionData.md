# TrainCategorizationTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountTypeId** | **int64** | Identifier of account type.&lt;br/&gt;&lt;br/&gt;1 &#x3D; Checking,&lt;br/&gt;2 &#x3D; Savings,&lt;br/&gt;3 &#x3D; CreditCard,&lt;br/&gt;4 &#x3D; Security,&lt;br/&gt;5 &#x3D; Loan,&lt;br/&gt;6 &#x3D; Pocket (DEPRECATED; will not be returned for any account unless this type has explicitly been set via PATCH),&lt;br/&gt;7 &#x3D; Membership,&lt;br/&gt;8 &#x3D; Bausparen&lt;br/&gt;&lt;br/&gt; | 
**Amount** | **float64** | Amount | 
**Purpose** | Pointer to **string** | Purpose. Any symbols are allowed. Maximum length is 2000. Default value: null. | [optional] 
**Counterpart** | Pointer to **string** | Counterpart. Any symbols are allowed. Maximum length is 80. Default value: null. | [optional] 
**CounterpartIban** | Pointer to **string** | Counterpart IBAN. Default value: null. | [optional] 
**CounterpartAccountNumber** | Pointer to **string** | Counterpart account number. Default value: null. | [optional] 
**CounterpartBlz** | Pointer to **string** | Counterpart BLZ. Default value: null. | [optional] 
**CounterpartBic** | Pointer to **string** | Counterpart BIC. Default value: null. | [optional] 
**McCode** | Pointer to **string** | Merchant category code (for credit card transactions only). Default value: null. NOTE: This field is currently not regarded. | [optional] 

## Methods

### NewTrainCategorizationTransactionData

`func NewTrainCategorizationTransactionData(accountTypeId int64, amount float64, ) *TrainCategorizationTransactionData`

NewTrainCategorizationTransactionData instantiates a new TrainCategorizationTransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainCategorizationTransactionDataWithDefaults

`func NewTrainCategorizationTransactionDataWithDefaults() *TrainCategorizationTransactionData`

NewTrainCategorizationTransactionDataWithDefaults instantiates a new TrainCategorizationTransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountTypeId

`func (o *TrainCategorizationTransactionData) GetAccountTypeId() int64`

GetAccountTypeId returns the AccountTypeId field if non-nil, zero value otherwise.

### GetAccountTypeIdOk

`func (o *TrainCategorizationTransactionData) GetAccountTypeIdOk() (*int64, bool)`

GetAccountTypeIdOk returns a tuple with the AccountTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeId

`func (o *TrainCategorizationTransactionData) SetAccountTypeId(v int64)`

SetAccountTypeId sets AccountTypeId field to given value.


### GetAmount

`func (o *TrainCategorizationTransactionData) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TrainCategorizationTransactionData) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TrainCategorizationTransactionData) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPurpose

`func (o *TrainCategorizationTransactionData) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *TrainCategorizationTransactionData) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *TrainCategorizationTransactionData) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *TrainCategorizationTransactionData) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetCounterpart

`func (o *TrainCategorizationTransactionData) GetCounterpart() string`

GetCounterpart returns the Counterpart field if non-nil, zero value otherwise.

### GetCounterpartOk

`func (o *TrainCategorizationTransactionData) GetCounterpartOk() (*string, bool)`

GetCounterpartOk returns a tuple with the Counterpart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpart

`func (o *TrainCategorizationTransactionData) SetCounterpart(v string)`

SetCounterpart sets Counterpart field to given value.

### HasCounterpart

`func (o *TrainCategorizationTransactionData) HasCounterpart() bool`

HasCounterpart returns a boolean if a field has been set.

### GetCounterpartIban

`func (o *TrainCategorizationTransactionData) GetCounterpartIban() string`

GetCounterpartIban returns the CounterpartIban field if non-nil, zero value otherwise.

### GetCounterpartIbanOk

`func (o *TrainCategorizationTransactionData) GetCounterpartIbanOk() (*string, bool)`

GetCounterpartIbanOk returns a tuple with the CounterpartIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartIban

`func (o *TrainCategorizationTransactionData) SetCounterpartIban(v string)`

SetCounterpartIban sets CounterpartIban field to given value.

### HasCounterpartIban

`func (o *TrainCategorizationTransactionData) HasCounterpartIban() bool`

HasCounterpartIban returns a boolean if a field has been set.

### GetCounterpartAccountNumber

`func (o *TrainCategorizationTransactionData) GetCounterpartAccountNumber() string`

GetCounterpartAccountNumber returns the CounterpartAccountNumber field if non-nil, zero value otherwise.

### GetCounterpartAccountNumberOk

`func (o *TrainCategorizationTransactionData) GetCounterpartAccountNumberOk() (*string, bool)`

GetCounterpartAccountNumberOk returns a tuple with the CounterpartAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartAccountNumber

`func (o *TrainCategorizationTransactionData) SetCounterpartAccountNumber(v string)`

SetCounterpartAccountNumber sets CounterpartAccountNumber field to given value.

### HasCounterpartAccountNumber

`func (o *TrainCategorizationTransactionData) HasCounterpartAccountNumber() bool`

HasCounterpartAccountNumber returns a boolean if a field has been set.

### GetCounterpartBlz

`func (o *TrainCategorizationTransactionData) GetCounterpartBlz() string`

GetCounterpartBlz returns the CounterpartBlz field if non-nil, zero value otherwise.

### GetCounterpartBlzOk

`func (o *TrainCategorizationTransactionData) GetCounterpartBlzOk() (*string, bool)`

GetCounterpartBlzOk returns a tuple with the CounterpartBlz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBlz

`func (o *TrainCategorizationTransactionData) SetCounterpartBlz(v string)`

SetCounterpartBlz sets CounterpartBlz field to given value.

### HasCounterpartBlz

`func (o *TrainCategorizationTransactionData) HasCounterpartBlz() bool`

HasCounterpartBlz returns a boolean if a field has been set.

### GetCounterpartBic

`func (o *TrainCategorizationTransactionData) GetCounterpartBic() string`

GetCounterpartBic returns the CounterpartBic field if non-nil, zero value otherwise.

### GetCounterpartBicOk

`func (o *TrainCategorizationTransactionData) GetCounterpartBicOk() (*string, bool)`

GetCounterpartBicOk returns a tuple with the CounterpartBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBic

`func (o *TrainCategorizationTransactionData) SetCounterpartBic(v string)`

SetCounterpartBic sets CounterpartBic field to given value.

### HasCounterpartBic

`func (o *TrainCategorizationTransactionData) HasCounterpartBic() bool`

HasCounterpartBic returns a boolean if a field has been set.

### GetMcCode

`func (o *TrainCategorizationTransactionData) GetMcCode() string`

GetMcCode returns the McCode field if non-nil, zero value otherwise.

### GetMcCodeOk

`func (o *TrainCategorizationTransactionData) GetMcCodeOk() (*string, bool)`

GetMcCodeOk returns a tuple with the McCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcCode

`func (o *TrainCategorizationTransactionData) SetMcCode(v string)`

SetMcCode sets McCode field to given value.

### HasMcCode

`func (o *TrainCategorizationTransactionData) HasMcCode() bool`

HasMcCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



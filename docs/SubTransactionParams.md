# SubTransactionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | Amount | 
**CategoryId** | Pointer to **int64** | Category identifier. If not specified, the original transaction&#39;s category will be applied. If you explicitly want the sub-transaction to have no category, then pass this field with value &#39;0&#39; (zero). | [optional] 
**Purpose** | Pointer to **string** | Purpose. Maximum length is 2000. If not specified, the original transaction&#39;s value will be applied. | [optional] 
**Counterpart** | Pointer to **string** | Counterpart. Maximum length is 80. If not specified, the original transaction&#39;s value will be applied. | [optional] 
**CounterpartAccountNumber** | Pointer to **string** | Counterpart account number. If not specified, the original transaction&#39;s value will be applied. | [optional] 
**CounterpartIban** | Pointer to **string** | Counterpart IBAN. If not specified, the original transaction&#39;s value will be applied. | [optional] 
**CounterpartBic** | Pointer to **string** | Counterpart BIC. If not specified, the original transaction&#39;s value will be applied. | [optional] 
**CounterpartBlz** | Pointer to **string** | Counterpart BLZ. If not specified, the original transaction&#39;s value will be applied. | [optional] 
**LabelIds** | Pointer to **[]int64** | List of connected labels. Note that when this field is not specified, then the labels of the original transaction will NOT get applied to the sub-transaction. Instead, the sub-transaction will have no labels assigned to it. | [optional] 

## Methods

### NewSubTransactionParams

`func NewSubTransactionParams(amount float64, ) *SubTransactionParams`

NewSubTransactionParams instantiates a new SubTransactionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubTransactionParamsWithDefaults

`func NewSubTransactionParamsWithDefaults() *SubTransactionParams`

NewSubTransactionParamsWithDefaults instantiates a new SubTransactionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SubTransactionParams) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SubTransactionParams) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SubTransactionParams) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCategoryId

`func (o *SubTransactionParams) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *SubTransactionParams) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *SubTransactionParams) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *SubTransactionParams) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetPurpose

`func (o *SubTransactionParams) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *SubTransactionParams) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *SubTransactionParams) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *SubTransactionParams) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetCounterpart

`func (o *SubTransactionParams) GetCounterpart() string`

GetCounterpart returns the Counterpart field if non-nil, zero value otherwise.

### GetCounterpartOk

`func (o *SubTransactionParams) GetCounterpartOk() (*string, bool)`

GetCounterpartOk returns a tuple with the Counterpart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpart

`func (o *SubTransactionParams) SetCounterpart(v string)`

SetCounterpart sets Counterpart field to given value.

### HasCounterpart

`func (o *SubTransactionParams) HasCounterpart() bool`

HasCounterpart returns a boolean if a field has been set.

### GetCounterpartAccountNumber

`func (o *SubTransactionParams) GetCounterpartAccountNumber() string`

GetCounterpartAccountNumber returns the CounterpartAccountNumber field if non-nil, zero value otherwise.

### GetCounterpartAccountNumberOk

`func (o *SubTransactionParams) GetCounterpartAccountNumberOk() (*string, bool)`

GetCounterpartAccountNumberOk returns a tuple with the CounterpartAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartAccountNumber

`func (o *SubTransactionParams) SetCounterpartAccountNumber(v string)`

SetCounterpartAccountNumber sets CounterpartAccountNumber field to given value.

### HasCounterpartAccountNumber

`func (o *SubTransactionParams) HasCounterpartAccountNumber() bool`

HasCounterpartAccountNumber returns a boolean if a field has been set.

### GetCounterpartIban

`func (o *SubTransactionParams) GetCounterpartIban() string`

GetCounterpartIban returns the CounterpartIban field if non-nil, zero value otherwise.

### GetCounterpartIbanOk

`func (o *SubTransactionParams) GetCounterpartIbanOk() (*string, bool)`

GetCounterpartIbanOk returns a tuple with the CounterpartIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartIban

`func (o *SubTransactionParams) SetCounterpartIban(v string)`

SetCounterpartIban sets CounterpartIban field to given value.

### HasCounterpartIban

`func (o *SubTransactionParams) HasCounterpartIban() bool`

HasCounterpartIban returns a boolean if a field has been set.

### GetCounterpartBic

`func (o *SubTransactionParams) GetCounterpartBic() string`

GetCounterpartBic returns the CounterpartBic field if non-nil, zero value otherwise.

### GetCounterpartBicOk

`func (o *SubTransactionParams) GetCounterpartBicOk() (*string, bool)`

GetCounterpartBicOk returns a tuple with the CounterpartBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBic

`func (o *SubTransactionParams) SetCounterpartBic(v string)`

SetCounterpartBic sets CounterpartBic field to given value.

### HasCounterpartBic

`func (o *SubTransactionParams) HasCounterpartBic() bool`

HasCounterpartBic returns a boolean if a field has been set.

### GetCounterpartBlz

`func (o *SubTransactionParams) GetCounterpartBlz() string`

GetCounterpartBlz returns the CounterpartBlz field if non-nil, zero value otherwise.

### GetCounterpartBlzOk

`func (o *SubTransactionParams) GetCounterpartBlzOk() (*string, bool)`

GetCounterpartBlzOk returns a tuple with the CounterpartBlz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBlz

`func (o *SubTransactionParams) SetCounterpartBlz(v string)`

SetCounterpartBlz sets CounterpartBlz field to given value.

### HasCounterpartBlz

`func (o *SubTransactionParams) HasCounterpartBlz() bool`

HasCounterpartBlz returns a boolean if a field has been set.

### GetLabelIds

`func (o *SubTransactionParams) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *SubTransactionParams) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *SubTransactionParams) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *SubTransactionParams) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



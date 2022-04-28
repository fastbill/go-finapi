# TriggerCategorizationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankConnectionIds** | Pointer to **[]int64** | List of identifiers of the bank connections that you want to trigger categorization for. Leaving the list unset or empty will trigger categorization for all of the user&#39;s bank connections. Please note that if there are no bank connections, then the service will return with HTTP code 422. | [optional] 

## Methods

### NewTriggerCategorizationParams

`func NewTriggerCategorizationParams() *TriggerCategorizationParams`

NewTriggerCategorizationParams instantiates a new TriggerCategorizationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerCategorizationParamsWithDefaults

`func NewTriggerCategorizationParamsWithDefaults() *TriggerCategorizationParams`

NewTriggerCategorizationParamsWithDefaults instantiates a new TriggerCategorizationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankConnectionIds

`func (o *TriggerCategorizationParams) GetBankConnectionIds() []int64`

GetBankConnectionIds returns the BankConnectionIds field if non-nil, zero value otherwise.

### GetBankConnectionIdsOk

`func (o *TriggerCategorizationParams) GetBankConnectionIdsOk() (*[]int64, bool)`

GetBankConnectionIdsOk returns a tuple with the BankConnectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankConnectionIds

`func (o *TriggerCategorizationParams) SetBankConnectionIds(v []int64)`

SetBankConnectionIds sets BankConnectionIds field to given value.

### HasBankConnectionIds

`func (o *TriggerCategorizationParams) HasBankConnectionIds() bool`

HasBankConnectionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



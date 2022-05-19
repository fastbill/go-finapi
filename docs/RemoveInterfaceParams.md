# RemoveInterfaceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankConnectionId** | **int64** | Bank connection identifier | 
**Interface** | Pointer to [**BankingInterface**](BankingInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankingInterface&lt;br/&gt; The interface which you want to remove. | [optional] 

## Methods

### NewRemoveInterfaceParams

`func NewRemoveInterfaceParams(bankConnectionId int64, ) *RemoveInterfaceParams`

NewRemoveInterfaceParams instantiates a new RemoveInterfaceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveInterfaceParamsWithDefaults

`func NewRemoveInterfaceParamsWithDefaults() *RemoveInterfaceParams`

NewRemoveInterfaceParamsWithDefaults instantiates a new RemoveInterfaceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankConnectionId

`func (o *RemoveInterfaceParams) GetBankConnectionId() int64`

GetBankConnectionId returns the BankConnectionId field if non-nil, zero value otherwise.

### GetBankConnectionIdOk

`func (o *RemoveInterfaceParams) GetBankConnectionIdOk() (*int64, bool)`

GetBankConnectionIdOk returns a tuple with the BankConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankConnectionId

`func (o *RemoveInterfaceParams) SetBankConnectionId(v int64)`

SetBankConnectionId sets BankConnectionId field to given value.


### GetInterface

`func (o *RemoveInterfaceParams) GetInterface() BankingInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *RemoveInterfaceParams) GetInterfaceOk() (*BankingInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *RemoveInterfaceParams) SetInterface(v BankingInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *RemoveInterfaceParams) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



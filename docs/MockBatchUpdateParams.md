# MockBatchUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MockBankConnectionUpdates** | [**[]MockBankConnectionUpdate**](MockBankConnectionUpdate.md) | &lt;strong&gt;Type:&lt;/strong&gt; MockBankConnectionUpdate&lt;br/&gt; List of mock bank connection updates | 
**TriggerNotifications** | Pointer to **bool** | Whether this call should trigger the dispatching of notifications. Default is &#39;false&#39;. | [optional] [default to false]

## Methods

### NewMockBatchUpdateParams

`func NewMockBatchUpdateParams(mockBankConnectionUpdates []MockBankConnectionUpdate, ) *MockBatchUpdateParams`

NewMockBatchUpdateParams instantiates a new MockBatchUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMockBatchUpdateParamsWithDefaults

`func NewMockBatchUpdateParamsWithDefaults() *MockBatchUpdateParams`

NewMockBatchUpdateParamsWithDefaults instantiates a new MockBatchUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMockBankConnectionUpdates

`func (o *MockBatchUpdateParams) GetMockBankConnectionUpdates() []MockBankConnectionUpdate`

GetMockBankConnectionUpdates returns the MockBankConnectionUpdates field if non-nil, zero value otherwise.

### GetMockBankConnectionUpdatesOk

`func (o *MockBatchUpdateParams) GetMockBankConnectionUpdatesOk() (*[]MockBankConnectionUpdate, bool)`

GetMockBankConnectionUpdatesOk returns a tuple with the MockBankConnectionUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMockBankConnectionUpdates

`func (o *MockBatchUpdateParams) SetMockBankConnectionUpdates(v []MockBankConnectionUpdate)`

SetMockBankConnectionUpdates sets MockBankConnectionUpdates field to given value.


### GetTriggerNotifications

`func (o *MockBatchUpdateParams) GetTriggerNotifications() bool`

GetTriggerNotifications returns the TriggerNotifications field if non-nil, zero value otherwise.

### GetTriggerNotificationsOk

`func (o *MockBatchUpdateParams) GetTriggerNotificationsOk() (*bool, bool)`

GetTriggerNotificationsOk returns a tuple with the TriggerNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerNotifications

`func (o *MockBatchUpdateParams) SetTriggerNotifications(v bool)`

SetTriggerNotifications sets TriggerNotifications field to given value.

### HasTriggerNotifications

`func (o *MockBatchUpdateParams) HasTriggerNotifications() bool`

HasTriggerNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



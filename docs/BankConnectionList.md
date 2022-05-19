# BankConnectionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | [**[]BankConnection**](BankConnection.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankConnection&lt;br/&gt; List of bank connections | 

## Methods

### NewBankConnectionList

`func NewBankConnectionList(connections []BankConnection, ) *BankConnectionList`

NewBankConnectionList instantiates a new BankConnectionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankConnectionListWithDefaults

`func NewBankConnectionListWithDefaults() *BankConnectionList`

NewBankConnectionListWithDefaults instantiates a new BankConnectionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *BankConnectionList) GetConnections() []BankConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *BankConnectionList) GetConnectionsOk() (*[]BankConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *BankConnectionList) SetConnections(v []BankConnection)`

SetConnections sets Connections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



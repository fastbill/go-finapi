# MockBankConnectionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankConnectionId** | **int64** | Bank connection identifier | 
**Interface** | Pointer to [**BankingInterface**](BankingInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankingInterface&lt;br/&gt; Banking interface to update. If not specified, then first available interface in bank connection will be used. | [optional] 
**SimulateBankLoginError** | Pointer to **bool** | Whether to simulate the case that the update fails due to incorrect banking credentials. Note that there is no real communication to any bank server involved, so you won&#39;t lock your accounts when enabling this flag. Default value is &#39;false&#39;. | [optional] [default to false]
**MockAccountsData** | Pointer to [**[]MockAccountData**](MockAccountData.md) | &lt;strong&gt;Type:&lt;/strong&gt; MockAccountData&lt;br/&gt; Mock accounts data. Note that for accounts that exist in a bank connection but that you do not specify in this list, the service will act like those accounts are not received by the bank servers. This means that any accounts that you do not specify here will be marked as deprecated. If you do not specify this list at all, all accounts in the bank connection will be marked as deprecated. | [optional] 

## Methods

### NewMockBankConnectionUpdate

`func NewMockBankConnectionUpdate(bankConnectionId int64, ) *MockBankConnectionUpdate`

NewMockBankConnectionUpdate instantiates a new MockBankConnectionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMockBankConnectionUpdateWithDefaults

`func NewMockBankConnectionUpdateWithDefaults() *MockBankConnectionUpdate`

NewMockBankConnectionUpdateWithDefaults instantiates a new MockBankConnectionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankConnectionId

`func (o *MockBankConnectionUpdate) GetBankConnectionId() int64`

GetBankConnectionId returns the BankConnectionId field if non-nil, zero value otherwise.

### GetBankConnectionIdOk

`func (o *MockBankConnectionUpdate) GetBankConnectionIdOk() (*int64, bool)`

GetBankConnectionIdOk returns a tuple with the BankConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankConnectionId

`func (o *MockBankConnectionUpdate) SetBankConnectionId(v int64)`

SetBankConnectionId sets BankConnectionId field to given value.


### GetInterface

`func (o *MockBankConnectionUpdate) GetInterface() BankingInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *MockBankConnectionUpdate) GetInterfaceOk() (*BankingInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *MockBankConnectionUpdate) SetInterface(v BankingInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *MockBankConnectionUpdate) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetSimulateBankLoginError

`func (o *MockBankConnectionUpdate) GetSimulateBankLoginError() bool`

GetSimulateBankLoginError returns the SimulateBankLoginError field if non-nil, zero value otherwise.

### GetSimulateBankLoginErrorOk

`func (o *MockBankConnectionUpdate) GetSimulateBankLoginErrorOk() (*bool, bool)`

GetSimulateBankLoginErrorOk returns a tuple with the SimulateBankLoginError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulateBankLoginError

`func (o *MockBankConnectionUpdate) SetSimulateBankLoginError(v bool)`

SetSimulateBankLoginError sets SimulateBankLoginError field to given value.

### HasSimulateBankLoginError

`func (o *MockBankConnectionUpdate) HasSimulateBankLoginError() bool`

HasSimulateBankLoginError returns a boolean if a field has been set.

### GetMockAccountsData

`func (o *MockBankConnectionUpdate) GetMockAccountsData() []MockAccountData`

GetMockAccountsData returns the MockAccountsData field if non-nil, zero value otherwise.

### GetMockAccountsDataOk

`func (o *MockBankConnectionUpdate) GetMockAccountsDataOk() (*[]MockAccountData, bool)`

GetMockAccountsDataOk returns a tuple with the MockAccountsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMockAccountsData

`func (o *MockBankConnectionUpdate) SetMockAccountsData(v []MockAccountData)`

SetMockAccountsData sets MockAccountsData field to given value.

### HasMockAccountsData

`func (o *MockBankConnectionUpdate) HasMockAccountsData() bool`

HasMockAccountsData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



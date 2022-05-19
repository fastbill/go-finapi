# AccountList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]Account**](Account.md) | &lt;strong&gt;Type:&lt;/strong&gt; Account&lt;br/&gt; List of bank accounts | 

## Methods

### NewAccountList

`func NewAccountList(accounts []Account, ) *AccountList`

NewAccountList instantiates a new AccountList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountListWithDefaults

`func NewAccountListWithDefaults() *AccountList`

NewAccountListWithDefaults instantiates a new AccountList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AccountList) GetAccounts() []Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AccountList) GetAccountsOk() (*[]Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AccountList) SetAccounts(v []Account)`

SetAccounts sets Accounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


